package main

import (
	"fmt"
	"strings"

	"github.com/plaid/plaid-go/plaid"
)

// main contains example usage of all functions
func main() {
	res, err := plaid.GetInstitutions(plaid.Tartan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res[0].Name, "has type:", res[0].Type)

	inst, err := plaid.GetInstitution(plaid.Tartan, "5301a93ac140de84910000e0")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inst.Name, "has mfa:", strings.Join(inst.MFA, ", "))

	categories, err := plaid.GetCategories(plaid.Tartan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("First category:", categories[0])

	category, err := plaid.GetCategory(plaid.Tartan, "13001001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("category", category.ID, "is", strings.Join(category.Hierarchy, ", "))

}
