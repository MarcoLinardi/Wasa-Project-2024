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
	GetUserById(userID int) (User, error)
	CreateNewUser(name string) (int, error)
	UpdateUserPhoto(photo string, userId int) error
	UpdateUsername(newName string, userId int) error
	GetUsers(userToIgnore int) ([]User, error)
	GetChats(userId int) ([]Chat, error)
	CreateChat(name string, users []int, isGroup bool, authenticatedUserID int) (int, error)
	FindPrivateChat(user1ID, user2ID int) (int, error)
	GetChatDetails(chatID int, userID int) (*Chat, error)
	DeleteChat(chatID int) error
	SaveMessage(chatID int, msg Message) (int, error)
	GetAllMessages(chatID int) ([]Message, error)
	UpdateMessagesStatus(chatID int, messageIDs []int, newStatus string) error
	ForwardMessage(originalChatID, originalMsgID, destinationChatID, senderID int) error
	AddReaction(messageID, userID int, reaction string) error
	CheckMessageInChat(chatID, messageID int) (bool, error)
	DeleteReaction(messageID, userID int) error
	DeleteMessage(messageID int) error
	AddMemberToGroup(chatID, userID int) error
	RemoveMemberFromGroup(chatID int, userID int) error
	SetGroupName(chatID int, newName string) error
	SetGroupPhoto(chatID int, newPhoto string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

type User struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type Chat struct {
	ChatID      int          `json:"chatId"`
	Name        string       `json:"name"`
	Photo       string       `json:"photo"`
	IsGroup     bool         `json:"isGroup"`
	Members     []User       `json:"participants"`
	LastMessage *LastMessage `json:"lastMessage,omitempty"`
}

type LastMessage struct {
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	SenderID  int    `json:"senderID,omitempty"`
}

type Message struct {
	MessageID   int        `json:"messageId"`
	SenderID    int        `json:"senderId"`
	Content     string     `json:"content"`
	Photo       *string    `json:"photo,omitempty"`
	Timestamp   string     `json:"timestamp"`
	Status      string     `json:"status"`
	IsForwarded bool       `json:"isForwarded"`
	ReplyToID   *int       `json:"replyToId,omitempty"`
	Reactions   []Reaction `json:"reactions"`
}

type Reaction struct {
	MessageID int    `json:"messageId"`
	UserID    int    `json:"userId"`
	Reaction  string `json:"reaction"`
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
		sqlStmt := `CREATE TABLE chats_table (chatId INTEGER NOT NULL PRIMARY KEY, 
												name TEXT,
												isGroup BOOLEAN DEFAULT 0,
												photo TEXT NOT NULL DEFAULT ''
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// TABELLA MEMBRI CHAT
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='chat_members';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE chat_members (chatId INTEGER NOT NULL,
    											userId INTEGER NOT NULL,
												PRIMARY KEY (chatId, userId),
												FOREIGN KEY (chatId) REFERENCES chats_table(chatId) ON DELETE CASCADE,
												FOREIGN KEY (userId) REFERENCES users_table(userId) ON DELETE CASCADE
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
												name TEXT,
												photo TEXT NOT NULL DEFAULT ''
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// TABELLA MESSAGES
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='messages';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE messages (messageId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
										chatId INTEGER NOT NULL,
										senderId INTEGER NOT NULL,
										content TEXT DEFAULT NULL,
										photo TEXT DEFAULT NULL,
										timestamp TEXT NOT NULL,
										status TEXT DEFAULT 'sent',
										isForwarded BOOLEAN DEFAULT 0,
										replyToId INTEGER DEFAULT NULL,
										FOREIGN KEY (chatId) REFERENCES chats_table(chatId) ON DELETE CASCADE,
										FOREIGN KEY (senderId) REFERENCES users_table(userId) ON DELETE CASCADE,
										FOREIGN KEY (replyToId) REFERENCES messages(messageId) ON DELETE SET NULL
										);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating messages table: %w", err)
		}
	}

	// TABELLA REACTION
	tableName = ""
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='message_reactions';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE message_reactions (messageId INTEGER NOT NULL,
  													userId INTEGER NOT NULL,
  													reaction TEXT NOT NULL,
													PRIMARY KEY (messageId, userId),
													FOREIGN KEY (messageId) REFERENCES messages(messageId) ON DELETE CASCADE,
													FOREIGN KEY (userId) REFERENCES users_table(userId) ON DELETE CASCADE
												);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
