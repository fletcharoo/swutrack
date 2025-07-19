# Commit Changes

## Goal
Analyze uncommitted changes in the repository and generate an appropriate commit message following conventional commit standards, then commit all changes after user approval.

## Process
1. Run `git status` to see all uncommitted files.
2. Run `git diff` to analyze all staged and unstaged changes.
3. Run `git log -5 --oneline` to understand recent commit message patterns.
4. Summarize the nature of the changes (e.g., new feature, bug fix, refactoring, tests, docs).
5. Identify the main purpose and impact of the changes.
6. Based on your analysis, generate a concise one-line commit message that:
    - Write in plain English.
    - Is written in imperative mood (e.g., "Add", "Fix", "Update").
    - Clearly describes what the change does, not how it does it.
    - Is no longer than 72 characters.
7. Present the commit message to the user in this format:
    ```
    Proposed commit message: "[your message here]"

    Type 'Go' to proceed with this commit message, or suggest changes
    ```
8. Wait for user confirmation before proceeding.
9. Once the user approves with `Go`, stage all changes with `git add -A`.
10. Create the commit using the approved message.
11. Run `git status` to confirm the commit was successful.
12. Display the commit hash and confirmation to the user.