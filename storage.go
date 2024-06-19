package main

import "database/sql"

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccount(int) (*Account, error)
}

type PostGreStore struct {
	db *sql.DB
}

func NewPostGreStore() (*PostGreStore, error) {

}
