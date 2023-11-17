package internal

import (
	"errors"
	"regexp"
	"strings"
)

var conventionalCommitPrefixes = []string{"fix", "chore", "feat", "build", "ci", "docs", "style", "refactor", "perf", "test"}

var InvalidTitleError = errors.New("invalid title given ")

type Commit struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	CreatedAt   string `json:"created_at"`
	Message     string `json:"message"`
}

func (commit *Commit) IsValidCommit() error {
	if len(commit.Message) == 0 {
		return errors.New("no message given")
	}
	return commit.IsValidTitle()
}

func (commit *Commit) IsValidTitle() error {
	found := false
	for _, prefix := range conventionalCommitPrefixes {
		found = strings.HasPrefix(commit.Title, prefix)
		if found {
			break
		}
	}
	if !found {
		return InvalidTitleError
	}
	matched, err := regexp.MatchString("\\w*\\(SEVDESK-[0-9]{4,6}\\).*", commit.Title)
	if err != nil {
		return err
	}
	if !matched {
		return InvalidTitleError
	}
	return nil
}
