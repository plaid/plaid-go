package plaid

type Environment string

const (
	Sandbox     Environment = "sandbox.plaid.com"
	Development Environment = "development.plaid.com"
	Production  Environment = "production.plaid.com"
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
