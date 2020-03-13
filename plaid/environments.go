package plaid

type Environment string

const (
	Sandbox     Environment = "https://sandbox.plaid.com"
	Development Environment = "https://development.plaid.com"
	Production  Environment = "https://production.plaid.com"
)

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
