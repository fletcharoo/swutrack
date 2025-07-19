# Generate a Product Requirements Document (PRD)

## Goal
You are a senior product manager creating a Product Requirements Document (PRD) for a single feature request. Generate a focused, actionable PRD based on user input suitable for a **junior developer** to understand and implement.

## Process
1. **Read Prompt:** Load from `temp/prompt.md`.
2. **Analyze Codebase:** Scan for related files, patterns, and tech stack.
3. **Ask Clarifying Questions:** You **MUST** gather missing details before proceeding.
4. **Generate PRD:** Create comprehensive PRD using the structure defined below and save to `temp/prd.md`.

## PRD Structure
1. **Overview:** Feature description and problem it solves.
2. **Goals:** Specific, measurable objectives.
3. **User Stories:** User narratives with benefits.
4. **Functional Requirements:** Numbered requirements with acceptance criteria.
5. **Out of Scope:** What this feature excludes to manage scope.
6. **Technical Considerations:** Constraints, dependencies, integration points.
7. **Success Metrics:** How success will be measured.
8. **Implementation Notes:** Dependencies and risks.

## Key Requirements
- **Junior developer friendly:** Explicit, unambiguous language.
- **Testable:** All requirements have clear acceptance criteria.
- **Contextual:** Reference existing codebase patterns.
- **Realistic:** Consider actual technical constraints.

## Final Instructions
- Do **NOT** begin implementing the feature.
- Make sure to ask clarifying questions.
- Focus on single feature scope only.