# Chapter 9: Escaping `#ifdef` Hell

This chapter gives detailed guidance on how to implement variants, like operating system variants or hardware variants, in C code. It discusses five patterns on how to cope with code variants as well as how to organize or even get rid of #ifdef statements. The patterns can be viewed as an introduction to organizing such code or as a guide on how to refactor unstructured #ifdef code.

The following patterns are covered:
- **Avoid Variants:** 
- **Isolated Primitives:** 
- **Atomic Primitives:** 
- **Abstraction Layer:** 
- **Split Variant Implementations:** 