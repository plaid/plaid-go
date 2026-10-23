// modulezip uses Go's module ZIP implementation, including its exclusion rules.
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	modzip "golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: modulezip <sdk-directory> <module-path> <v-version> <output.zip>")
		os.Exit(2)
	}
	out, err := os.OpenFile(os.Args[4], os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = modzip.CreateFromDir(out, module.Version{Path: os.Args[2], Version: os.Args[3]}, os.Args[1])
	closeErr := out.Close()
	if err == nil {
		err = closeErr
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
