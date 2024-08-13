#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main()
{
    char *line;
    line = NULL;
    line = malloc(10);
    size_t c = strlen(line);
    printf("|%ld|\n", c);
}