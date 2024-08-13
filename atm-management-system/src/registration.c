#include "header.h"
#include <termios.h>
#include <unistd.h>
#include <fcntl.h>

void registerMenu(char a[50], char pass[50])
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
};

int checkUser(struct User u)
{
    FILE *fp;
    struct User userChecker;
    char *line;
    char **user;

    if ((fp = fopen(USERS, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    while (_getline(fp, &line, NULL) != 1)
    {
        user = _split(line);
        if (_arrlen(user) == 3)
        {
            if (strcmp(user[1], u.name) == 0)
            {
                fclose(fp);
                return 1;
            }
        }
        free(line);
    }
    fclose(fp);
    return 0;
}

void addUser(char a[50], char pass[50])
{
    FILE *fp;
    int id;
    int newLine;
    char *line;
    char **user;

    newLine = 0;
    if ((fp = fopen(USERS, "a+")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    while (_getline(fp, &line, &newLine) != 1)
    {
        user = _split(line);
        if (_arrlen(user) == 3)
        {
            id = atoi(user[0]);
        }
        free(line);
    }
    if (newLine == 1)
    {
        if (fprintf(fp, "%d %s %s\n", id+1, a, pass) < 0)
        {
            printf("Can't write to file!!");
            exit(1);
        }
    }
    else
    {
        if (fprintf(fp, "\n%d %s %s\n", id+1, a, pass) < 0)
        {
            printf("Can't write to file!!");
            exit(1);
        }
    }
    fclose(fp);
}