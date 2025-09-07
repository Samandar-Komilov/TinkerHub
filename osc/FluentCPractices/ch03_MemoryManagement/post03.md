# Working with Data Efficiently

This is a giant multichapter paper, gathering all related chapters into one place. 

## Chapter 3: Memory Management

This chapter primarily focuses on memory management patterns in C, which are crucial for developers where we don't have garbage collection or Rust-style borrow checkers, therefore putting all responsibility to them. The following patterns are covered:

!!! abstract
    - **Stack First:** if you know the size of the data and it is not so huge, you may store it in stack.
    - **Eternal Memory:** If the data size is big, but you know its size, its better to put in static memory which lives during the whole lifetime of the program.
    - **Lazy Cleanup:** If the data size is not known beforehand, and maybe its size changes at runtime, then you should use heap memory. But you may let the OS cleanup if your program is shortlived and does not allocate too much space (no `free()`).
    - **Dedicated Ownership:** The one who allocated the memory, is responsible to clean it up. Or it may transfer the ownership, for example to the caller function. But in any case, you have to document who allocated the memory and who will free it.
    - **Allocation Wrapper:** Checking for `malloc()` potential errors every time is cumbersome. Instead, we should create a wrapper function for that.
    - **Pointer Check:** Always nullify the uninitialized and freed pointers to prevent use-after-free and dangling pointer issues.
    - **Memory Pool:** Frequently allocating and deallocating memory is expensive and can cause memory fragmentation. Thus, it is better to allocate a pool of memory at the start of the program and use it at runtime by getting by chunks from it.



## Chapter 4: Returning Data from Functions

This chapter primarily focuses on how we should exchange data across functions well enough, so that our code remains maintainable. The following patterns are covered:

!!! abstract
    - **Return Value:** Return values are the most common way to exchange data between functions. The returning value of the callee is copied to the caller function and become accessible.
    - **Out-Parameters:** Use reference arguments when needed. Instead of copying the values to the callee, give pointers to the data and allow the callee function to change them, without returning anything.
    
    !!! danger
        You have to consider multi-threaded cases here. Use mutexes or semaphores while changing "remote" data by reference.
    - **Aggregate Instance:** As C functions can return a single instance of data, you may wrap them into a struct as a single unit, which is a common practice. (Even official nginx uses it.)
    - **Immutable Instance:** 
    - **Caller-Owned Buffer:** 
    - **Callee Allocates:** 

## Chapter 5: Data Lifetime and Ownership

This chapter discusses patterns for how to structure your C program with object-like elements. For these object-like elements, the patterns put special focus on who is responsible for creating and destroying them—in other words, they put special focus on lifetime and ownership. This topic is especially important for C because C has no automatic destructor and no garbage collection mechanism, and thus special attention has to be paid to cleanup of resources.

The following patterns are covered:

!!! abstract
    - **Stateless Software Module:** 
    - **Software-Module With Global State:** 
    - **Caller-Owned Instance:** 
    - **Shared Instance:** 
