#include <termios.h>
#include "header.h"
#include <unistd.h>
#include <stdio.h>

void loginMenu(char a[50], char pass[50])
{
    struct termios oflags, nflags;

    system("clear");
    printf("\n\n\n\t\t\t\t   Bank Management System\n\t\t\t\t\t User Login:");
    scanf("%s", a);

    // disabling echo
    tcgetattr(fileno(stdin), &oflags);
    nflags = oflags;
    nflags.c_lflag &= ~ECHO;
    nflags.c_lflag |= ECHONL;

    if (tcsetattr(fileno(stdin), TCSANOW, &nflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
    printf("\n\n\n\n\n\t\t\t\tEnter the password to login:");
    scanf("%s", pass);
    // restore terminal
    if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
}

void removeNewline(char **str)
{
    int i;

    i = 0;
    while((*str)[i] != '\0')
    {
        if ((*str)[i] == '\n')
        {
            (*str)[i] = '\0';
        }
        i++;
    }
}
const char *getPassword(struct User u)
{
    FILE *fp;
    struct User userChecker;
    char *line;
    char *user[4];
    size_t len;
    int i;

    line = NULL;
    if ((fp = fopen(USERS, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    while (getline(&line, &len, fp) != -1)
    {
        i = 0;
        user[i] = strtok(line, " ");
        removeNewline(&user[i++]);
        while (line != NULL)
        {
            if (i == 3)
            {
                user[i] = NULL;
                break;
            }
            user[i] = strtok(NULL, " ");
            removeNewline(&user[i++]);
        }
        if (_arrlen(user) == 3)
        {
            strcpy(userChecker.name, user[1]);
            strcpy(userChecker.password, user[2]);
            if (strcmp(userChecker.name, u.name) == 0)
            {
                free(line);
                fclose(fp);
                char *buff = userChecker.password;
                return buff;
            }
        }
        free(line);
        line = NULL;
    }
    free(line);
    fclose(fp);
    return "no user found";
}