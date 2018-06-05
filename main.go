package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/plaid/plaid-go/plaid"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Creates a new Plaid Client
	clientOptions := plaid.ClientOptions{
		os.Getenv("PLAID_CLIENT_ID"),
		os.Getenv("PLAID_SECRET"),
		os.Getenv("PLAID_PUBLIC_KEY"),
		plaid.Sandbox,
		&http.Client{},
	}
	client, err := plaid.NewClient(clientOptions)
	handleError(err)

	// POST /institutions/get
	instsResp, err := client.GetInstitutions(5, 0)
	handleError(err)
	fmt.Println(instsResp.Institutions[0].Name, "has products:", instsResp.Institutions[0].Products)

	// POST /institutions/get_by_id
	instResp, err := client.GetInstitutionByID(instsResp.Institutions[0].ID)
	handleError(err)
	fmt.Println(instResp.Institution.Name, "has MFA:", instResp.Institution.MFA)

	// POST /institutions/search
	instSearchResp, err := client.SearchInstitutions("Ally", []string{"transactions"})
	handleError(err)
	fmt.Println(instSearchResp.Institutions[0].Name, "has ID:", instSearchResp.Institutions[0].ID)

	// POST /categories/get
	categoriesResp, err := client.GetCategories()
	handleError(err)
	fmt.Println("Category group", categoriesResp.Categories[0].Group, "has items:", categoriesResp.Categories[0].Hierarchy)

	// POST /sandbox/public_token/create
	publicTokenResp, err := client.CreateSandboxPublicToken(instResp.Institution.ID, []string{"auth", "transactions"})
	handleError(err)
	fmt.Println("Created sandbox public token:", publicTokenResp.PublicToken)

	// POST /item/public_token/exchange
	accessTokenResp, err := client.ExchangePublicToken(publicTokenResp.PublicToken)
	handleError(err)
	fmt.Println("Public token -> Access Token", accessTokenResp.AccessToken, "for item:", accessTokenResp.ItemID)

	// POST /accounts/balance/get
	balanceResp, err := client.GetBalances(accessTokenResp.AccessToken)
	handleError(err)
	fmt.Println("Account with name", balanceResp.Accounts[0].Name, "has available balance", balanceResp.Accounts[0].Balances.Available)

	// POST /accounts/get
	accountsResp, err := client.GetAccounts(accessTokenResp.AccessToken)
	handleError(err)
	fmt.Println("Access token is associated with", len(accountsResp.Accounts), "accounts")

	// POST /auth/get
	authResp, err := client.GetAuth(accessTokenResp.AccessToken)
	handleError(err)
	fmt.Println("Account has number:", authResp.Numbers.ACH[0].Account)

	// POST /transactions/get
	transactionsResp, err := client.GetTransactions(accessTokenResp.AccessToken, "2010-01-01", "2018-01-01")
	if plaidErr, ok := err.(plaid.Error); ok {
		// Poll until transactions are ready
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			transactionsResp, err = client.GetTransactions(accessTokenResp.AccessToken, "2010-01-01", "2018-01-01")
			plaidErr, ok = err.(plaid.Error)
		}
		handleError(err)
	}
	fmt.Println("Number of transactions:", len(transactionsResp.Transactions))
}
