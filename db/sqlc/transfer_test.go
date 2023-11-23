package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        5,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {

	arg := ListTransfersParams{
		Offset: 5,
		Limit:  5,
	}

	transfer2, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfer2, int(arg.Limit))

	for _, entry := range transfer2 {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)

	arg := UpdateTransferParams{
		Amount: 100,
		ID:     transfer1.ID,
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, transfer2.Amount, arg.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
