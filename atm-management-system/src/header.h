#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define USERS "./data/users.txt"

struct Date
{
    int month, day, year;
};

// checkers
struct checker
{
    int EOF_checker;
};

// all fields for each record of an account
struct Record
{
    int id;
    int userId;
    char name[100];
    char country[100];
    int phone;
    char accountType[10];
    int accountNbr;
    double amount;
    struct Date deposit;
    struct Date withdraw;
};

struct User
{
    int id;
    char name[50];
    char password[50];
};

// authentication functions
void loginMenu(char a[50], char pass[50]);
void registerMenu(char a[50], char pass[50]);
void addUser(char a[50], char pass[50]);
int checkUser(struct User u);
const char *getPassword(struct User u);

// system function
void createNewAcc(struct User u);
void mainMenu(struct User u);
void checkAllAccounts(struct User u);

// utils
int _getline(FILE *fp, char **line, int *newLine);
int _strlen(char *str);
char **_split(char *str, char delim);
int _arrlen(char **arr);
void freeArr(char **arr);