package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct{
	*Queries
	db *sql.DB
}

func NewStore( db *sql.DB) *Store{
	return &Store{
		db:db,
		Queries:New(db), // recall that New(db) initilisae a queries struct based on the *sql.db connection
	}
}


func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
	tx,err := store.db.BeginTx(context.Background(),nil)

	if err != nil{
		return err
	}

	q := New(tx)
	err = fn(q)

	if err!=nil{
		if rbErr := tx.Rollback() ; rbErr !=nil{
			return fmt.Errorf("tx err: %v , rb err:%v",err,rbErr)
		}
		return err
	}

	return tx.Commit()
}
