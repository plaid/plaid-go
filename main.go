package main

import (
	"fmt"
	"strings"

	"github.com/plaid/plaid-go/plaid"
)

// main contains example usage of all functions
func main() {
	// GET /institutions
	res, err := plaid.GetInstitutions(plaid.Tartan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res[0].Name, "has type:", res[0].Type)

	// GET /institutions/5301a93ac140de84910000e0
	inst, err := plaid.GetInstitution(plaid.Tartan, "5301a93ac140de84910000e0")
	if err != nil {
		fmt.Println("Institution Error:", err)
	}
	fmt.Println(inst.Name, "has mfa:", strings.Join(inst.MFA, ", "))

	// GET /categories
	categories, err := plaid.GetCategories(plaid.Tartan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First category:", categories[0])

	// GET /categories/13001001
	category, err := plaid.GetCategory(plaid.Tartan, "13001001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("category", category.ID, "is", strings.Join(category.Hierarchy, ", "))

	client := plaid.NewClient("test_id", "test_secret", plaid.Tartan)

	// POST /auth
	postRes, mfaRes, err :=
		client.AuthAddUser("plaid_test", "plaid_good", "", "citi", nil)
	if err != nil {
		fmt.Println(err)
	} else if mfaRes != nil {
		switch mfaRes.Type {
		case "device":
			fmt.Println("--Device MFA--")
			fmt.Println("Message:", mfaRes.Device.Message)
		case "list":
			fmt.Println("--List MFA--")
			fmt.Println("Mask:", mfaRes.List[0].Mask, "\nType:", mfaRes.List[0].Type)
		case "questions":
			fmt.Println("--Questions MFA--")
			fmt.Println("Question:", mfaRes.Questions[0].Question)
		case "selections":
			fmt.Println("--Selections MFA--")
			fmt.Println("Question:", mfaRes.Selections[1].Question)
			fmt.Println("Answers:", mfaRes.Selections[1].Answers)
		}

		postRes2, mfaRes2, err := client.AuthStepSendMethod(mfaRes.AccessToken, "type", "email")
		if err != nil {
			fmt.Println("Error submitting send_method", err)
		}
		fmt.Println(mfaRes2, postRes2)

		postRes2, mfaRes2, err = client.AuthStep(mfaRes.AccessToken, "tomato")
		if err != nil {
			fmt.Println("Error submitting mfa", err)
		}
		fmt.Println(mfaRes2, postRes2)
	} else {
		fmt.Println(postRes.Accounts)
		fmt.Println("Auth Get")
		fmt.Println(client.AuthGet("test"))

		fmt.Println("Auth DELETE")
		fmt.Println(client.AuthDelete("test"))
	}
}
