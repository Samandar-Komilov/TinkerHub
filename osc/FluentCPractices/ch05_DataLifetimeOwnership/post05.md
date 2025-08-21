# Chapter 5: Data Lifetime and Ownership

This chapter discusses patterns for how to structure your C program with object-like elements. For these object-like elements, the patterns put special focus on who is responsible for creating and destroying them—in other words, they put special focus on lifetime and ownership. This topic is especially important for C because C has no automatic destructor and no garbage collection mechanism, and thus special attention has to be paid to cleanup of resources.

The following patterns are covered:
- **Stateless Software Module:** 
- **Software-Module With Global State:** 
- **Caller-Owned Instance:** 
- **Shared Instance:** 

