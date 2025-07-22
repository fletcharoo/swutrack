# Update Documentation

## Goal
To guide an AI assistant through updating relevant Markdown and inline comment documentation.

## Process
1. **Read PRD:** Read the Product Requirements Document (PRD) from `temp/prd.md` to understand the purpose and intent of the current changes.
2. **Read Changes:** Read through the uncommitted changes in this repo to fully understand the new or modified implementation of features in this repository.
3. **Read Markdown Documentation:** Read and understand all relevant Markdown documentation in this repository.
4. **Read Inline Comments:** Read and understand all relevant inline comments in this repository.
5. **Update:** Use your best judgement to identify any outdated documentation and update them.

## Exclude
You **MUST** exclude the following files and directories when identifying outdated documentation:
- `./docs/`
- `./.claude/`
- `./temp/`
- `./CLAUDE.md`
- `./README.md`
- `/Makefile`

## Final Instructions
- All changes **MUST** be meaningful, do **NOT** make changes that are simply synonyms.