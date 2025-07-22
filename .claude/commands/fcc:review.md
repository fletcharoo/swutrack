# Conduct a Comprehensive Code Review

## Goal
To guide an AI assistant in performing a complete code and documentation review of the repository. The goal is to identify any potential bugs, security vulnerabilities, poor coding practices, outdated documentation, and to verify that all tests pass. The output should be a clearly written report appropriate for a developer or engineering team to act upon.

## Process
1. **Run Tests:** Run all tests. If any tests fail, notify the user and stop.
2. **Read PRD:** Load the feature requirements from `temp/prd.md`.
3. **Analyze PRD:** Use deep thinking to fully understand the purpose and intent of the PRD.
4. **Feature Alignment:** Analyze the uncommitted changes on the branch. Use deep thinking to fully understand the purpose and intent of the changes. Use your best judgement to determine whether the uncommitted changes align with the feature requirements described in the PRD.
5. **Test Review:** Use your best judgement to determine whether tests adequately test the feature defined in the PRD. Including possible bugs and edge cases.
6. **Code Coverage:** Run all tests with code coverage. If the code coverage is below 85%, add this finding to the report.
7. **Code Quality Review:** Assess naming conventions, code organization, and adherence to established code style & conventions. Look for code duplication, complex functionality, and lack of modularity. Identify areas that would benefit from refactoring or improved readability.
8. **Static Analysis:** Check for bugs, bad coding style, and potential logic errors. Identify dead code or unused variables/functions.
9. **Security Analysis:** Identify hardcoded secrets, insecure coding patterns, or outdated dependencies. If the code uses external services or APIs, ensure proper handling of secrets and credentials.
10. **Generate Report:** Using all your findings from above, generate the report using the template described below and save it to `temp/report.md`.

## Report Structure
1. **Overview:** Briefly describe the findings in this report.
2. **Critical:** Severe issues requiring immediate action. These pose significant threats to system security, functionality, or data integrity. Examples include SQL injection vulnerabilities, authentication bypasses, or critical business logic flaws.
3. **High:** Serious issues that should be addressed promptly. While not immediately exploitable, they represent substantial risks. Examples include weak encryption, insufficient access controls, or sensitive data exposure.
4. **Medium:** Moderate issues that should be remediated in normal development cycles. These present real risks but with limited impact or exploitability. Examples include missing security headers, verbose error messages, or outdated dependencies.
5. **Low:** Minor issues representing minimal risk. These should be fixed when convenient but don't pose immediate threats. Examples include missing best practices, minor information disclosure, or code quality issues.

## Final Instructions
- Do **NOT** begin implementing fixes for the findings.
- Ensure all findings include specific file paths and line numbers where applicable.
- Include code snippets for critical and high priority issues.
- Provide clear, actionable recommendations for each finding.