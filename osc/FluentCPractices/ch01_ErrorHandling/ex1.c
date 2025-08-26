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

typedef struct {
    FILE *fp;
    char *buffer;
} FileParser;


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