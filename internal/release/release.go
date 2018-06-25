package release

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr, "[release] ", log.LUTC)
}

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

// Main is the entry point into the release script
func Main(args []string) error {
	pkgName := "github.com/plaid/plaid-go"
	pkg, err := build.Import(pkgName, "", build.FindOnly)
	if err != nil {
		return err
	}

	// Open the git repository
	logger.Println("Opening", pkgName)
	r, err := git.PlainOpen(pkg.Dir)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Check the git status to make sure the working tree is clean
	status, err := w.Status()
	if err != nil {
		return err
	}

	if !status.IsClean() {
		return fmt.Errorf("working tree contains uncommitted changes, cannot proceed")
	}

	release, err := validateArguments(args)
	if err != nil {
		return err
	}

	newVersion, err := incrementVersion(pkg, release)
	if err != nil {
		return err
	}

	if _, err := w.Add("internal"); err != nil {
		return err
	}

	logger.Println("added internal/... to working tree")

	if err := commitAndTagRepo(r, w, newVersion); err != nil {
		return err
	}

	return nil
}

func commitAndTagRepo(r *git.Repository, w *git.Worktree, newVersion string) error {
	now := time.Now()
	signature := object.Signature{
		Name:  "Plaid",
		Email: "api@plaid.com",
		When:  now,
	}

	commit, err := w.Commit(newVersion, &git.CommitOptions{
		Author: &signature,
	})
	if err != nil {
		return err
	}

	logger.Println("committing changes")

	if _, err := r.CommitObject(commit); err != nil {
		return err
	}

	tag := object.Tag{
		Name:       newVersion,
		Message:    "Release of " + newVersion,
		Tagger:     signature,
		Target:     commit,
		TargetType: plumbing.CommitObject,
	}

	logger.Println("tagging repo with", tag.Name)

	e := r.Storer.NewEncodedObject()
	tag.Encode(e)
	hash, err := r.Storer.SetEncodedObject(e)
	if err != nil {
		return err
	}

	if err := r.Storer.SetReference(plumbing.NewReferenceFromStrings("refs/tags/"+newVersion, hash.String())); err != nil {
		return err
	}

	return nil
}
