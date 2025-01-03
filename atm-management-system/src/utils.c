#include "header.h"

void getPrompt(char **str)
{
    *str =  NULL;
    size_t size = 0;
    ssize_t chars_read;

    chars_read = getline(str, &size, stdin);
    if (chars_read < 0)
    {
        fprintf(stderr, "getline error\n");
        exit(1);
    }
    if (chars_read > 0 && (*str)[chars_read - 1] == '\n')
        (*str)[chars_read - 1] = '\0';
}

static int countChar(char *str, int index, char delim)
{
    while (str[index] && str[index] != delim)
        index++;
    return index;
}
static int countWord(char *str, char delim)
{
    int i;
    int count;

    i = 0;
    count = 0;
    while (str[i])
    {
        if (str[i] == delim || !str[i+1])
            count++;
        i++;
    }
    return count;
}

static void arrMemSet(char **arr, int len)
{
    while (len >= 0)
        arr[len--] = NULL;
}

static void fillStr(char *toFill, char *filler, char delim, int index)
{
    int i;

    i = 0;
    while(filler[index] && filler[index] != delim)
    {
        toFill[i] = filler[index];
        index++;
        i++;
    }
}
char **split(char *str, char delim)
{
    char **result;
    int i, j;
    int delim_index;
    int len;

    i = 0;
    j = 0;
    delim_index = 0;
    if ((len = countWord(str, delim)) == 0)
        return NULL;
    result = malloc(sizeof(char *) * (len + 1));
    arrMemSet(result, len);
    while (j < len)
    {
        delim_index = countChar(str, i, delim);
        result[j] = malloc(delim_index + 1);
        memset(result[j], '\0', delim_index +1);
        fillStr(result[j++], str, delim, i);
        i = delim_index+1;
    }
    return result;
}

void getDateFromStrs(char **str, struct Date *date)
{
    date->month = atoi(str[0]);
    date->day = atoi(str[1]);
    date->year = atoi(str[2]);
}

int countArr(char **arr)
{
    int i;

    i = -1;
    while(arr[++i] != NULL);
    return i;
}

void handleDbError(int rc, sqlite3 *db, const char *errmsg)
{
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", errmsg);
        if (errmsg)
            sqlite3_free((void *)errmsg);
        sqlite3_close(db);
        exit(1);
    }
}

void handleStatementError(int rc, sqlite3 *db, const char *errmsg, sqlite3_stmt *stmt)
{
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "SQL error: %s\n", errmsg);
        if (errmsg)
            sqlite3_free((void *)errmsg);
        sqlite3_finalize(stmt);
        sqlite3_close(db);
        exit(1);
    }
}

int checkExistingAccByUser(char *user, int id)
{
    char *sql;
    int rc;
    sqlite3_stmt *stmt;

    sql = "SELECT * FROM records WHERE user_name = ? AND account_id = ?";
    rc = sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_bind_text(stmt, 1, user, -1, SQLITE_STATIC);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_bind_int(stmt, 2, id);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_step(stmt);
    if (rc == SQLITE_ROW)
        return 0;
    else if (rc == SQLITE_DONE)
        return 1;
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    return 1;
}

int checkAmount(char *str)
{
    char *tmp;
    double res;

    res = strtod(str, &tmp);
    if (res == 0 || strlen(tmp) != 0)
        return 1;
    return 0;
}

int checkExistingUser(char *user, int *user_id)
{
    char *sql;
    int rc;
    sqlite3_stmt *stmt;

    sql = "SELECT id FROM users WHERE name = ?;";
    rc = sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_bind_text(stmt, 1, user, -1, SQLITE_STATIC);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_step(stmt);
    if (rc == SQLITE_ROW)
    {
        *user_id = sqlite3_column_int(stmt, 0);
        return 0;
    }
    else if (rc == SQLITE_DONE)
        return 1;
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    return 1;
}

int strIsInt(char *str)
{
    int i;

    i = 0;
    if (!strlen(str))
        return 1;
    while(str[i])
    {
        if (!isdigit(str[i]))
            return 1;
        i++;
    }
    return 0;
}

int checkExistingAcc(int id)
{
    char *sql;
    int rc;
    sqlite3_stmt *stmt;

    sql = "SELECT * FROM records WHERE account_id = ?;";
    rc = sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_bind_int(stmt, 1, id);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_step(stmt);
    if (rc == SQLITE_ROW)
        return 1;
    else if (rc == SQLITE_DONE)
        return 0;
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    return 1;
}

void success(struct User u)
{
    int option;
    char *input;

    sleep(1);
invalid:
    printf("\n\nEnter 1 to go to the main menu and 0 to exit!\n");
    getPrompt(&input);
    option = atoi(input);
    if ((option = atoi(input)) <= 0 && strcmp(input, "0") != 0)
        option = 2;
    switch (option)
    {
    case 1:
        mainMenu(u);
        break;
    case 0:
        sqlite3_close(db);
        exit(0);
    default:
        printf("Insert a valid operation!\n");
        fflush(stdout);
        sleep(1);
        system("clear");
        goto invalid;
    }
}