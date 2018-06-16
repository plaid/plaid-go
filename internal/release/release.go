package release

import (
	"fmt"
)

type releaseType string

var (
	major = releaseType("major")
	minor = releaseType("minor")
	patch = releaseType("patch")
)

func validateArguments(args []string) (releaseType, error) {
	if len(args) != 1 {
		return releaseType(""), fmt.Errorf("exactly one argument must be supplied")
	}

	spec := args[0]
	if spec != "major" && spec != "minor" && spec != "patch" {
		return releaseType(""), fmt.Errorf("invalid release-type, must be one of (major|minor|patch)")
	}

	return releaseType(spec), nil
}

// Main is the main entry point into the release script
func Main(args []string) error {
	release, err := validateArguments(args)
	if err != nil {
		return err
	}

	if err := incrementVersion(release); err != nil {
		return err
	}

	return nil
}
