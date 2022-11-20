package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/KyawKyawThar/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:     util.RandomOwner(),
		Balance:   util.RandomMoney(),
		Currentcy: util.RandomCurrency(),
	}

	account, err := testQuries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currentcy, account.Currentcy)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	getAcc, err := testQuries.GetAccounts(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getAcc)

	require.Equal(t, account1.ID, getAcc.ID)
	require.Equal(t, account1.Owner, getAcc.Owner)
	require.Equal(t, account1.Balance, getAcc.Balance)
	require.Equal(t, account1.Currentcy, getAcc.Currentcy)
	require.WithinDuration(t, account1.CreatedAt, getAcc.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	fmt.Println("account1", account1.Balance)
	args := UpdateAccountsParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	fmt.Println("args", args.Balance)

	updateAcc, err := testQuries.UpdateAccounts(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updateAcc)

	require.Equal(t, account1.ID, updateAcc.ID)
	require.Equal(t, account1.Owner, updateAcc.Owner)
	require.NotEqual(t, account1.Balance, updateAcc.Balance)
	require.Equal(t, account1.Currentcy, updateAcc.Currentcy)
	require.WithinDuration(t, account1.CreatedAt, updateAcc.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQuries.DeleteAccounts(context.Background(), account1.ID)

	require.NoError(t, err)

	getAcc, err := testQuries.GetAccounts(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getAcc)

}

func TestListsAccount(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQuries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		fmt.Println("acc", acc)
		require.NotEmpty(t, acc)
	}
}
