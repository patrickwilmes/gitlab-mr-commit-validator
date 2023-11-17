package internal

import "testing"

func TestCommit_isValidTitleCheckAllConventionalCommitPrefixes(t *testing.T) {
	table := []struct {
		commit  Commit
		isValid bool
	}{
		{Commit{Title: "chore(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "fix(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "feat(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "ci(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "build(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "docs(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "style(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "refactor(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "perf(SEVDESK-222222) some fancy test"}, true},
		{Commit{Title: "test(SEVDESK-222222) some fancy test"}, true},

		{Commit{Title: "Chore(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Fix(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Feat(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Ci(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Build(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Docs(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Style(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Refactor(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Perf(SEVDESK-222222) some fancy test"}, false},
		{Commit{Title: "Test(SEVDESK-222222) some fancy test"}, false},

		{Commit{Title: "test(SEVDESK-222222)"}, false},
		{Commit{Title: "test(-222222)"}, false},
		{Commit{Title: "test(Something else)"}, false},
		{Commit{Title: "test(Something else) with text"}, false},
		{Commit{Title: "totally different"}, false},
	}

	for _, entry := range table {
		commit := &entry.commit
		if err := commit.IsValidTitle(); err != nil && entry.isValid {
			t.Errorf("Invalid title detected but should be valid! title %s\n", commit.Title)
		}
	}
}

func TestCommit_IsValidCommitCheckAllConventionalCommitPrefixes(t *testing.T) {
	table := []struct {
		commit  Commit
		isValid bool
	}{
		{
			Commit{
				Title:   "chore(SEVDESK-222222) some fancy test",
				Message: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut l",
			},
			true,
		},
		{
			Commit{
				Title: "chore(SEVDESK-222222) some fancy test",
			},
			false,
		},
		{
			Commit{
				Title:   "chore(SEVDESK-222222) some fancy test",
				Message: "",
			},
			false,
		},
	}

	for _, entry := range table {
		commit := &entry.commit
		if err := commit.IsValidCommit(); err != nil && entry.isValid {
			t.Errorf("Invalid message detected but should be valid! message %s\n", commit.Message)
		}
	}
}
