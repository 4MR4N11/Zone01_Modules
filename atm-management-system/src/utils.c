#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define BUFF_SIZE 5

int _arrlen(char **arr)
{
    int i = 0;
    for (; arr[i] != NULL; i++);
    return i;
}

int _strlen(char *str)
{
    int i = 0;
    for (; str[i] != '\0'; i++);
    return i;
}

int _getline(FILE *fp, char **line, int *newLine)
{
    char c;
    int len;

    len = 0;
    *line = malloc(sizeof(char) * BUFF_SIZE);
    if (!(*line))
    {
        perror("malloc error!\n");
        exit(1);
    }
    (*line)[0] = '\0';
    while((c = fgetc(fp)) != '\n' && c != EOF)
    {
        if (newLine != NULL)
        {
            *newLine = 0;
        }
        if (len+1 > BUFF_SIZE)
        {
            *line = realloc(*line, len+2);
            (*line)[len+1] = '\0';
        }
        (*line)[len++] = c;
    }
    if (newLine != NULL && c == '\n')
    {
        *newLine = 1;
    }
    if (c == EOF && _strlen(*line) == 0)
    {
        return 1;
    }
    return 0;
}

char **_split(char *str, char delim)
{
    char **strs;
    char *tmp;
    int tmpIndex = 0;
    int arrIndex = 0;

    tmp = malloc(1);
    strs = malloc(sizeof(char *));
    if (!tmp || !strs)
    {
        perror("malloc error!\n");
        exit(1);
    }
    strs[0] = NULL;
    tmp[0] = '\0';
    for (int i = 0; str[i] != '\0';i++)
    {
        if (str[i] == delim)
        {
            if (_strlen(tmp) != 0)
            {
                tmpIndex = 0;
                if (strs[arrIndex] == NULL)
                {
                    strs = realloc(strs, arrIndex+2);
                    strs[arrIndex+1] = NULL;
                }
                strs[arrIndex++] = strdup(tmp);
                free(tmp);
                tmp = malloc(1);
                if (!tmp)
                {
                    perror("malloc error!\n");
                    exit(1);
                }
                tmp[0] = '\0';
            }
        }
        else
        {
            if (tmp[tmpIndex] == '\0')
            {
                tmp = realloc(tmp, tmpIndex+2);
                tmp[tmpIndex+1] = '\0';
            }
            tmp[tmpIndex++] = str[i];
        }
    }

    if (_strlen(tmp) != 0)
    {
        if (strs[arrIndex] == NULL)
        {
            strs = realloc(strs, arrIndex+2);
            strs[arrIndex+1] = NULL;
        }
        strs[arrIndex++] = strdup(tmp);
        free(tmp);
    }
    return strs;
}

void freeArr(char **arr)
{
    int	i;

	i = 0;
	while (arr && arr[i])
	{
        write(1, arr[i], _strlen(arr[i]));
		free(arr[i]);
		arr[i] = NULL;
		i++;
	}
	free(arr);
	arr = NULL;
}