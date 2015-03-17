# plaid-go

plaid-go is a Go client implementation of the [Plaid API](https://plaid.com/docs).

Install via `go get github.com/plaid/plaid-go`.

TODO:
- Complete README
- Complete testing
- Add CI


## Examples

### Adding an Auth user

```go
client := plaid.NewClient("test_id", "test_secret", plaid.Tartan)

// POST /auth
postRes, mfaRes, err := client.AuthAddUser("plaid_test", "plaid_good", "", "bofa", nil)
if err != nil {
    fmt.Println(err)
} else if mfaRes != nil {
    // Need to switch on different MFA types. See https://plaid.com/docs/#mfa-auth.
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
    } else {
        fmt.Println(mfaRes2, postRes2)
    }
} else {
    fmt.Println(postRes.Accounts)
    fmt.Println("Auth Get")
    fmt.Println(client.AuthGet("test_bofa"))

    fmt.Println("Auth DELETE")
    fmt.Println(client.AuthDelete("test_bofa"))
}
```

### Querying a category
```go
// GET /categories/13001001
category, err := plaid.GetCategory(plaid.Tartan, "13001001")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("category", category.ID, "is", strings.Join(category.Hierarchy, ", "))
}
```
