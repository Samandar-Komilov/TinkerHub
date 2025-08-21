# Chapter 1: Error Handling Patterns

This initial chapter of Fluent C book provides good error handling techniques in C with a running example. The following patterns are covered:
- **Function Split:** when a function has several responsibilities, split it into smaller functions.
- **Guard Clause:** The function mixes pre-condition checks with main logic, therefore is hard to maintain. Hence, check mandatory pre-conditions first and if they are not met, return error info immediately.
- **Samurai Principle:** Caller can sometimes omit the error checks of your function. If you know that the error can't be handled and there is no point to return to the caller, simply abort the program.
- **Goto Handling:** Collect all resource cleanup code in one place and use goto to jump to that code when needed, instead of writing multiple cleanup code in each block.
- **Cleanup Record:** 
- **Object-Based Error Handling:** One function doing both resource acquision, cleanup and logic is hard to maintain. Therefore, make separate constructor and destructor functions just like in OOP.
