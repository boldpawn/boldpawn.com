#!/usr/bin/env node

import { promises as fs } from "node:fs";
import path from "node:path";

const defaultSourceDir = "/Users/marcelschroer/Desktop/AI-ICS2";
const defaultOutputDir = "site/knowledge/ics2";
const maxChunkChars = 14_000;

const sourceDir = path.resolve(process.argv[2] || defaultSourceDir);
const outputDir = path.resolve(process.argv[3] || defaultOutputDir);

const markdownExtensions = new Set([".md", ".markdown"]);

async function main() {
  const sourceFiles = await listMarkdownFiles(sourceDir);
  if (sourceFiles.length === 0) {
    throw new Error(`No Markdown files found in ${sourceDir}`);
  }

  await fs.rm(outputDir, { force: true, recursive: true });
  await fs.mkdir(outputDir, { recursive: true });

  const manifest = [];
  let chunkTotal = 0;

  for (const filePath of sourceFiles) {
    const relativePath = normalizePath(path.relative(sourceDir, filePath));
    const raw = await fs.readFile(filePath, "utf8");
    const normalized = normalizeMarkdown(raw);
    const title = titleFromMarkdown(normalized) || titleFromFilename(filePath);
    const sourceSlug = slugFromPath(relativePath);
    const chunks = chunkMarkdown(normalized, title);

    for (let index = 0; index < chunks.length; index += 1) {
      const chunk = chunks[index];
      const chunkNumber = index + 1;
      const chunkSlug = chunks.length === 1
        ? `${sourceSlug}.md`
        : `${sourceSlug}--part-${String(chunkNumber).padStart(2, "0")}.md`;
      const outputPath = path.join(outputDir, chunkSlug);
      const chunkTitle = chunks.length === 1
        ? title
        : `${title} - part ${chunkNumber}`;
      const content = renderChunk({
        title: chunkTitle,
        sourceTitle: title,
        sourcePath: relativePath,
        section: chunk.sectionTitle,
        chunkNumber,
        chunkCount: chunks.length,
        body: chunk.body,
      });

      await fs.writeFile(outputPath, content);
      const size = Buffer.byteLength(content, "utf8");
      manifest.push({
        key: normalizePath(path.relative(path.resolve("site"), outputPath)),
        title: chunkTitle,
        sourceTitle: title,
        sourcePath: relativePath,
        section: chunk.sectionTitle,
        chunk: chunkNumber,
        chunks: chunks.length,
        bytes: size,
      });
      chunkTotal += 1;
    }
  }

  await writeIndex(manifest);
  await fs.writeFile(
    path.join(outputDir, "manifest.json"),
    `${JSON.stringify({ generatedAt: new Date().toISOString(), sourceDir, files: manifest }, null, 2)}\n`,
  );

  const largest = manifest.reduce((max, item) => Math.max(max, item.bytes), 0);
  console.log(`Processed ${sourceFiles.length} Markdown source file(s).`);
  console.log(`Generated ${chunkTotal} Markdown chunk file(s).`);
  console.log(`Largest generated chunk is ${largest} bytes.`);
  console.log(`Output: ${normalizePath(path.relative(process.cwd(), outputDir))}`);
}

async function listMarkdownFiles(root) {
  const entries = await fs.readdir(root, { withFileTypes: true });
  const files = [];

  for (const entry of entries) {
    const fullPath = path.join(root, entry.name);
    if (entry.isDirectory()) {
      files.push(...await listMarkdownFiles(fullPath));
    } else if (entry.isFile() && markdownExtensions.has(path.extname(entry.name).toLowerCase())) {
      files.push(fullPath);
    }
  }

  return files.sort((a, b) => normalizePath(a).localeCompare(normalizePath(b)));
}

function normalizeMarkdown(raw) {
  return raw
    .replace(/\r\n?/g, "\n")
    .split("\n")
    .filter((line) => !/^\s*!\[[^\]]*]\([^)]*\)\s*$/.test(line))
    .join("\n")
    .replace(/\n{4,}/g, "\n\n\n")
    .trim();
}

function titleFromMarkdown(markdown) {
  const match = markdown.match(/^#\s+(.+)$/m);
  return match?.[1]?.trim();
}

function titleFromFilename(filePath) {
  return path.basename(filePath, path.extname(filePath)).replace(/[-_]+/g, " ").trim();
}

function chunkMarkdown(markdown, fallbackTitle) {
  const sections = splitIntoSections(markdown, fallbackTitle);
  const chunks = [];

  for (const section of sections) {
    if (section.body.length <= maxChunkChars) {
      chunks.push(section);
      continue;
    }

    const subChunks = splitLongSection(section);
    chunks.push(...subChunks);
  }

  return chunks.filter((chunk) => chunk.body.trim() !== "");
}

function splitIntoSections(markdown, fallbackTitle) {
  const lines = markdown.split("\n");
  const sections = [];
  let currentTitle = fallbackTitle;
  let current = [];

  for (const line of lines) {
    const heading = line.match(/^(#{1,4})\s+(.+)$/);
    if (heading && current.join("\n").trim().length > 0) {
      sections.push({
        sectionTitle: currentTitle,
        body: current.join("\n").trim(),
      });
      current = [];
    }
    if (heading) {
      currentTitle = heading[2].trim();
    }
    current.push(line);
  }

  if (current.join("\n").trim().length > 0) {
    sections.push({
      sectionTitle: currentTitle,
      body: current.join("\n").trim(),
    });
  }

  return sections;
}

function splitLongSection(section) {
  const lines = section.body.split("\n");
  const chunks = [];
  let current = [];
  let currentLength = 0;
  let currentTableHeader = [];

  for (const line of lines) {
    const tableHeader = detectTableHeader(current, line);
    if (tableHeader.length > 0) {
      currentTableHeader = tableHeader;
    }

    if (currentLength > 0 && currentLength + line.length + 1 > maxChunkChars) {
      chunks.push({
        sectionTitle: section.sectionTitle,
        body: current.join("\n").trim(),
      });
      current = currentTableHeader.length > 0 ? [...currentTableHeader] : [];
      currentLength = current.join("\n").length;
    }

    current.push(line);
    currentLength += line.length + 1;
  }

  if (current.join("\n").trim().length > 0) {
    chunks.push({
      sectionTitle: section.sectionTitle,
      body: current.join("\n").trim(),
    });
  }

  return chunks;
}

function detectTableHeader(currentLines, nextLine) {
  if (!/^\s*\|.*\|\s*$/.test(nextLine)) {
    return [];
  }
  if (currentLines.length === 0) {
    return [];
  }
  const previous = currentLines[currentLines.length - 1];
  if (!/^\s*\|.*\|\s*$/.test(previous)) {
    return [];
  }
  if (!/^\s*\|[\s:-]+\|\s*$/.test(nextLine.replace(/[^|:-]/g, ""))) {
    return [];
  }
  return [previous, nextLine];
}

function renderChunk({ title, sourceTitle, sourcePath, section, chunkNumber, chunkCount, body }) {
  return `---\n${yamlField("title", title)}${yamlField("source_title", sourceTitle)}${yamlField("source_path", sourcePath)}${yamlField("section", section)}chunk: ${chunkNumber}\nchunk_count: ${chunkCount}\ngenerated: true\n---\n\n# ${title}\n\nSource document: ${sourceTitle}\nSource path: \`${sourcePath}\`\nChunk: ${chunkNumber} of ${chunkCount}\n\n${body.trim()}\n`;
}

function yamlField(name, value) {
  return `${name}: ${JSON.stringify(value || "")}\n`;
}

async function writeIndex(manifest) {
  const bySource = new Map();
  for (const item of manifest) {
    if (!bySource.has(item.sourcePath)) {
      bySource.set(item.sourcePath, []);
    }
    bySource.get(item.sourcePath).push(item);
  }

  const lines = [
    "---",
    'title: "ICS2 Knowledge Base"',
    'description: "Generated Markdown chunks from the ICS2 HTI source documents."',
    "generated: true",
    "---",
    "",
    "# ICS2 Knowledge Base",
    "",
    "This directory contains generated, MCP-friendly chunks from the ICS2 HTI Markdown source documents.",
    "",
    `Generated chunks: ${manifest.length}`,
    "",
    "## Curated Guides",
    "",
    "- [Entry Summary Declaration (ENS)](../ics2-guides/entry-summary-declaration.md): Start here for broad questions about ENS meaning, lifecycle, filings, messages, amendment, invalidation, consultation, and answer strategy.",
    "",
    "## Source Documents",
    "",
  ];

  for (const [sourcePath, chunks] of [...bySource.entries()].sort()) {
    lines.push(`- ${sourcePath} (${chunks.length} chunk${chunks.length === 1 ? "" : "s"})`);
  }

  await fs.writeFile(path.join(outputDir, "index.md"), `${lines.join("\n")}\n`);
}

function slugFromPath(relativePath) {
  const withoutExtension = relativePath.replace(/\.[^.]+$/, "");
  return normalizePath(withoutExtension)
    .split("/")
    .map(slugify)
    .filter(Boolean)
    .join("--");
}

function slugify(value) {
  return value
    .normalize("NFKD")
    .replace(/[\u0300-\u036f]/g, "")
    .toLowerCase()
    .replace(/&/g, " and ")
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "")
    .slice(0, 120);
}

function normalizePath(value) {
  return value.split(path.sep).join("/");
}

main().catch((error) => {
  console.error(error instanceof Error ? error.message : error);
  process.exitCode = 1;
});
