# Chapter 1: Error Handling Patterns

This initial chapter of Fluent C book provides good error handling techniques in C with a running example. The following patterns are covered:
- **Function Split:** when a function has several responsibilities, split it into smaller functions.
- **Guard Clause:** The function mixes pre-condition checks with main logic, therefore is hard to maintain. Hence, check mandatory pre-conditions first and if they are not met, return error info immediately.
- **Samurai Principle:** Caller can sometimes omit the error checks of your function. If you know that the error can't be handled and there is no point to return to the caller, simply abort the program.
- **Goto Handling:** Collect all resource cleanup code in one place and use goto to jump to that code when needed, instead of writing multiple cleanup code in each block.
- **Cleanup Record:** 
- **Object-Based Error Handling:** One function doing both resource acquision, cleanup and logic is hard to maintain. Therefore, make separate constructor and destructor functions just like in OOP.

---

We have given a file parser source code as an example:
```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 1024

enum {
    ERROR = -1,
    NO_KEYWORD_FOUND = 10,
    KEYWORD_ONE_FOUND_FIRST = 11,
    KEYWORD_TWO_FOUND_FIRST = 12
};

int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    if (filename != NULL){
        if (fp = fopen(filename, "r") != NULL){
            if (buffer = malloc(BUFFER_SIZE) != NULL){
                // parse file content
                retval = NO_KEYWORD_FOUND;
                while (fgets(buffer, BUFFER_SIZE, fp) != NULL){
                    if (strcmp("KEYWORD_ONE\n", buffer) == 0){
                        retval = KEYWORD_ONE_FOUND_FIRST;
                        break;
                    }
                    if (strcmp("KEYWORD_TWO\n", buffer) == 0){
                        retval = KEYWORD_TWO_FOUND_FIRST;
                        break;
                    }
                }
                free(buffer);
            }
            fclose(fp);
        }
    }

    return retval;
}
```

The function has several drawbacks:
- the function is doing many things at once (resource acquision and cleanup, error handling and logic);
- nested if conditions are hard to maintain and test;
- cleanup functions are not easy to find, may cause memory leaks if forgotten somewhere.

**Function Split.** To solve the issues, we first need to split the function into 2 parts: `searchKeywordsFromFile()` and `parseFile()`:

```c
int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    if (filename != NULL){
        if (fp = fopen(filename, "r") != NULL){
            if (buffer = malloc(BUFFER_SIZE) != NULL){
                retval = searchFileForKeywords(buffer, fp);
                free(buffer);
            }
            fclose(fp);
        }
    }

    return retval;
}


int searchFileForKeywords(char* buffer, FILE *fp){
    while (fgets(buffer, BUFFER_SIZE, fp) != NULL){
        if (strcmp("KEYWORD_ONE\n", buffer) == 0){
            return KEYWORD_ONE_FOUND_FIRST;
        }
        if (strcmp("KEYWORD_TWO\n", buffer) == 0){
            return KEYWORD_TWO_FOUND_FIRST;
        }
    }

    return NO_KEYWORD_FOUND;
}
```
This is much better now, but still not perfect. The main function is still doing many `if` checks which is making function long. 

**Guard Clause.** To solve this, we need to check mandatory pre-conditions first and if they are not met, return error info immediately:
```c
int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    if (filename == NULL){
        return ERROR;
    }

    if (fp = fopen(filename, "r") == NULL){
        return ERROR;
    }

    if (buffer = malloc(BUFFER_SIZE) != NULL){
        retval = searchFileForKeywords(buffer, fp);
        free(buffer);
    }
    fclose(fp);

    return retval;
}
```

**Samurai Principle.** Caller can sometimes omit the error checks of your function. If you know that the error can't be handled and there is no point to return to the caller, simply abort the program:
```c
int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    assert(filename != NULL && "Invalid filename");

    if (fp = fopen(filename, "r") == NULL){
        return ERROR;
    }

    if (buffer = malloc(BUFFER_SIZE) != NULL){
        retval = searchFileForKeywords(buffer, fp);
        free(buffer);
    }
    fclose(fp);

    return retval;
}
```

**Goto Error Handling.** You might still have nested ifs, which is complicating resource cleanup as in every `if` condition, you must cleanup the resource if the result is not successful. In that case, even mostly discouraged, we can use `goto` to stop repeating the same code, instead jumping to a specific code block that cleans up the resources:
```c
int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    assert(filename != NULL && "Invalid filename");

    if (fp = fopen(filename, "r") == NULL){
        goto error_fopen;
    }

    if (buffer = malloc(BUFFER_SIZE) != NULL){
        goto error_malloc;
    }
    retval = searchFileForKeywords(buffer, fp);
    free(buffer);

error_malloc:
    fclose(fp);
error fopen:
    return retval;
}
```

**Cleanup Record.** Even though `goto` gets the job done, it is highly discouraged still. And if you don't want to use `goto` in your code, you may choose this approach: label the results of each functions used in the error handling code, run the main logic only if they all succeed, and check for each failure case using ifs:
```c
int parseFile(char* filename){
    int retval = ERROR;
    FILE *fp = 0;
    char *buffer = 0;

    assert(filename != NULL && "Invalid filename");

    if ((fp = fopen(filename, "r") != NULL) && (buffer = malloc(BUFFER_SIZE) != NULL)){
        retval = searchFileForKeywords(buffer, fp);
        free(buffer);
    }

    // Cleanup Record without goto
    if (fp)
        fclose(fp);
    if (buffer)
        free(buffer);

    return retval;
}
```

**Object-Based Error Handling.** Still, the function is doing 3 things: resource initialization, usage of that resource and resource cleanup, which makes hard to maintain this function. What if we split these 3, just like we do in object-based languages?
```c
int parseFile(char* filename){
    int retval = ERROR;
    FileParser* parser = createParser(filename);
    retval = searchFileForKeywords(parser);
    cleanupParser(parser);

    return retval;
}

FileParser* createParser(filename){
    assert(filename!=NULL && "Invalid filename");
    FileParser* parser = malloc(sizeof(FileParser));

    if (parser){
        parser->fp=fopen(filename, "r");
        parser->buffer = malloc(BUFFER_SIZE);
        if (!parser->fp || !parser->buffer){
            cleanupParser(parser);
            return NULL;
        }
    }

    return parser;
}

void cleanupParser(FileParser *parser){
    assert(parser!=NULL && "Invalid parser");

    if (parser->buffer)
        free(parser->buffer);
    if (parser->fp)
        fclose(parser->fp);
    free(parser);
}

int searchFileForKeywords(FileParser *parser){
    assert(parser!=NULL && "Invalid parser");

    while (fgets(parser->buffer, BUFFER_SIZE, parser->fp) != NULL){
        if (strcmp("KEYWORD_ONE\n", parser->buffer) == 0){
            return KEYWORD_ONE_FOUND_FIRST;
        }
        if (strcmp("KEYWORD_TWO\n", parser->buffer) == 0){
            return KEYWORD_TWO_FOUND_FIRST;
        }
    }

    return NO_KEYWORD_FOUND;
}
```

Even though the lines of code is increased, we successfully applied all of the patterns and made our code readable, maintainable and testable. Yes, less code is better, but maintainable code is much better than less but later painful code.