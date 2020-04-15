package plaid

type Environment struct {
	url string
}

var Sandbox = Environment{
	url: "https://sandbox.plaid.com",
}

var Development = Environment{
	url: "https://development.plaid.com",
}

var Production = Environment{
	url: "https://production.plaid.com",
}

var validEnvironments = [...]Environment{
	Sandbox,
	Development,
	Production,
}

func (c Environment) Valid() bool {
	for _, v := range validEnvironments {
		if v == c {
			return true
		}
	}
	return false
}
