# golang-bank-backend

## Current Schema
![Alt text](/static/image.png)


## [Before you begin]
1. ```docker start postrgres12```,  ```make createdb```, ```make migrateup``` if u did not done so previously

## [3]
 
1. ```history | grep "docker run"``` is extremely useful to view docker history, apparently it does show that some of my docker images were removed somehow. Not sure what happened exactly
<br>

2. MakeFile is useful to run short hand commands like ```make migrateup```. it is actually short for 
   ```migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up```

        
## [4]
1. SQLC is recommended for GOlang CRUD operations.

## [5]
1. Convention for unit testing in GoLang : testfile is in the same folder as the actual code.
<br>

2. Convention for unit testing in Golang : main_test is the entry point of all unit tests inside a specific golang package

## [5]

1. where to put the dependencies? consider the following snippets of code.
example 1 is by me, example 2 is by instructor

example 1: dependencies are created within the createRandomTransfer function
```
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
```


example 2: dependencies are created within the TestCreateTransfer function
```
func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}
```

heres what gpt has to say about the differences
```
Example 1:

Pros:

    Self-contained: Example 1 creates random accounts within the createRandomTransfer function, which means you don't need to manage the creation of accounts explicitly in your test cases. This makes it convenient for tests where you don't want to worry about the setup.

    Simple to use: The function is straightforward and easy to use. You call it, and it handles everything from creating random accounts to performing the transfer.

Cons:

    Lack of control: The random account creation may not be suitable for all testing scenarios. If you need precise control over the accounts involved in the transfer, this approach may not be the best choice.

    Limited reusability: This approach may not be suitable for cases where you want to reuse specific accounts across multiple test cases.

Example 2:

Pros:

    Control: Example 2 allows you to provide existing accounts as arguments to the createRandomTransfer function, giving you more control over the test scenario. This is beneficial when you want to create specific test cases or when you want to reuse existing accounts.

    Reusability: This approach is more versatile and can be used in a wider range of test scenarios because you can provide any accounts you want.

Cons:

    Requires more setup: This approach requires you to set up the accounts explicitly in your test cases, which can be a bit more work. However, it gives you greater control.

    Increased complexity: The test cases using this approach may become more complex due to the need to create accounts separately before creating a transfer.

Best Practice:

The best approach depends on the specific requirements of your test cases:

    If you need self-contained tests with random accounts and want to avoid managing account creation in each test case, Example 1 may be more suitable for you. It simplifies the testing process and is easier to use.

    If you require more control over the test scenario, want to use existing accounts, or need to create specific test scenarios, Example 2 is a better choice. It gives you the flexibility to set up the test environment according to your needs.

In general, both approaches have their use cases, and the choice should be driven by the specific requirements of your testing strategy and the level of control and reusability you need in your test scenarios.
```

[6]
database transactions must be ACID, this is to handle and deal with potential race conditions
```
A - atmoic. either all operations complete successfuly, or the transaction files and db is unchaged
C - consistent. db must be valid after transaction
I - isolation. concurrent db transactions must not affect each other. There are diff levels of isolation
D - Durability. data written must be in persistent storage
```