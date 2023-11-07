package db

import (
	"context"
	"simplebank/util"
	"testing"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer{
	// must first create two accounts to faciliate the transfer
	// account1 (from) -> account2 (to)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer , err := testQueries.CreateTransfer(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,transfer)

	require.Equal(t,arg.FromAccountID,transfer.FromAccountID)
	require.Equal(t,arg.ToAccountID,transfer.ToAccountID)
	require.Equal(t,arg.Amount,transfer.Amount)

	require.NotZero(t,transfer.ID)
	require.NotZero(t,transfer.CreatedAt.Time)


	return transfer
}


func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}