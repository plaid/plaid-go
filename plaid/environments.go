package plaid

type Environment struct {
	URL string
}

var Sandbox = Environment{
	URL: "https://sandbox.plaid.com",
}

var Development = Environment{
	URL: "https://development.plaid.com",
}

var Production = Environment{
	URL: "https://production.plaid.com",
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
