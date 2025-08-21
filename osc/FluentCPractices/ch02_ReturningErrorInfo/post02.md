# Chapter 2: Returning Error Information

This chapter continues the discussion of Error Handling, but this time we primarily focus on how to exchange error information across functions correctly so that the callers can react to according to the errors. The following patterns are covered:

- **Return Status Codes:** Introduce a specific error code that both you and the caller mutually understand (using enums for example).
- **Return Relevant Errors:** Return errors only which are relevant to the caller. The more error is returned the more code to handle them and the harder to maintain the codebase.
- **Special Return Values:** 
- **Log Errors:** Sometimes just logging the error is enough to developer to find the root cause of the problem, instead of returning to the caller.
