package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    account1.Balance - 1,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {

	arg := ListEntriesParams{
		Offset: 5,
		Limit:  5,
	}

	entry2, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entry2, int(arg.Limit))

	for _, entry := range entry2 {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		Amount: 100,
		ID:     entry1.ID,
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry2.Amount, arg.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
