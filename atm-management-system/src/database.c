#include "header.h"


sqlite3 *db = NULL;

void createDatabase()
{
    int rc;

    rc = sqlite3_open("data/database.db", &db);
    if (rc != SQLITE_OK)
    {
        fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
        exit(1);
    }
}

void initDatabase()
{
    const char *user_sql, *record_sql;
    char *errmsg = NULL;
    int rc;

    user_sql =   "CREATE TABLE IF NOT EXISTS users ("
                "id INTEGER PRIMARY KEY AUTOINCREMENT, "
                "name TEXT NOT NULL UNIQUE, "
                "pass TEXT);";
    record_sql = "CREATE TABLE IF NOT EXISTS records ("
                "id INTEGER PRIMARY KEY AUTOINCREMENT, "
                "user_id INTEGER, "
                "user_name TEXT, "
                "account_id INTEGER UNIQUE, "
                "date_of_creation TEXT, "
                "country TEXT, "
                "phone_number TEXT, "
                "balance REAL, "
                "type_of_account TEXT, "
                "FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE);";

    rc = sqlite3_exec(db, user_sql, 0, 0, &errmsg);
    handleDbError(rc, db, errmsg);
    rc = sqlite3_exec(db, record_sql, 0, 0, &errmsg);
    handleDbError(rc, db, errmsg);
}

void insertRecord(struct Record r, struct User u)
{
    char *sql;
    int rc;
    sqlite3_stmt *stmt;

    sql =   "INSERT INTO records (user_id, user_name, account_id, date_of_creation, country, phone_number, balance, type_of_account)"
            "SELECT id, name, ?, ?, ?, ?, ?, ? FROM users WHERE name = ? LIMIT 1;";
    rc = sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
    handleDbError(rc, db, sqlite3_errmsg(db));
    rc = sqlite3_bind_int(stmt, 1, r.accountNbr);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_text(stmt, 2, r.deposit, -1, SQLITE_STATIC);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_text(stmt, 3, r.country, -1, SQLITE_STATIC);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_text(stmt, 4, r.phone, -1, SQLITE_STATIC);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_double(stmt, 5, r.amount);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_text(stmt, 6, r.accountType, -1, SQLITE_STATIC);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_bind_text(stmt, 7, u.name, -1, SQLITE_STATIC);
    handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    rc = sqlite3_step(stmt);
    if (rc != SQLITE_DONE)
        handleStatementError(rc, db, sqlite3_errmsg(db), stmt);
    sqlite3_finalize(stmt);
}