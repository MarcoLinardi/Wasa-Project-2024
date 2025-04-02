/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserIdByName(name string) (int, error)
	SetName(name string) error
	UserIdExist(userId int) (bool, error)
	CreateUser(name string) (int, error)
	UpdateUserPhoto(photo string, userId int) error
	UpdateUsername(newName string, userId int) error
	GetUsers() ([]User, error)
	GetChats(userId int) ([]Chat, error)
	CreateChat(name string, users []int) (int, error)
	GetChatDetails(chatID int, userID int) (*Chat, error)
	DeleteChat(chatID int) error
	SaveMessage(msg Message) (int, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// User rappresenta un utente nel database
type User struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type Chat struct {
	ChatID int    `json:"chatId"`
	Name   string `json:"name"`
	// Photo       Photo  `json:"photo"`
	IsGroup  bool      `json:"isGroup"`
	Messages []Message `json:"messages"`
}

type Message struct {
	MessageID   int    `json:"messageId"`
	SenderID    int    `json:"senderId"`
	ChatID      int    `json:"chatId"`
	Content     string `json:"content"`
	Timestamp   string `json:"timestamp"`
	Status      string `json:"status"`
	IsForwarded bool   `json:"isForwarded"`
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	var tableName string

	// TABELLA CHAT
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='chats_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE chats_table (chat_id INTEGER NOT NULL PRIMARY KEY, 
												name TEXT
												is_group BOOLEAN DEFAULT 0
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// TABELLA MEMBRI CHAT
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='chat_members';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE chat_members (chat_id INTEGER NOT NULL,
    											user_id INTEGER NOT NULL,
												PRIMARY KEY (chat_id, user_id),
												FOREIGN KEY (chat_id) REFERENCES chats_table(id) ON DELETE CASCADE,
												FOREIGN KEY (user_id) REFERENCES users_table(userId) ON DELETE CASCADE
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// TABELLA UTENTI
	tableName = ""
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users_table (userId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
												name TEXT
												photo TEXT
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// TABELLA MESSAGES
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='messages';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE messages (message_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
										chat_id INTEGER NOT NULL,
										sender_id INTEGER NOT NULL,
										content TEXT NOT NULL,
										timestamp TEXT NOT NULL,
										FOREIGN KEY (chat_id) REFERENCES chats_table(id) ON DELETE CASCADE,
										FOREIGN KEY (sender_id) REFERENCES users_table(userId) ON DELETE CASCADE
										);`

		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating messages table: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
