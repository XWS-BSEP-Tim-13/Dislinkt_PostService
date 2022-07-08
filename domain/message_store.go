package domain

type MessageStore interface {
	Insert(messages *MessageUsers) error
	DeleteAll()
	GetByUsers(firstUsername, secondUsername string) (*MessageUsers, error)
	SendMessage(message *Message) error
	GetByUser(firstUsername string) ([]*MessageUsers, error)
}
