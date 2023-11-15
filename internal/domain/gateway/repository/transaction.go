package repository

type ITx interface {
	Commit() error
	Rollback() error
}
