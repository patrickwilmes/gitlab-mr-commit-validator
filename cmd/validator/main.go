/*
 * Copyright (c) 2023, Patrick Wilmes <patrick.wilmes@bit-lake.com>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */
package main

import (
	"fmt"
	"gitlab-mr-commit-validator/internal"
	"os"
)

func main() {
	args := os.Args
	gitlabToken := args[1]
	fmt.Println(gitlabToken)
	data, err := internal.FetchMRs(gitlabToken)
	if err != nil {
		panic(err)
	}
	for _, mr := range data {
		commits, err := internal.FetchCommitsForMergeRequest(mr.IID, gitlabToken)
		if err != nil {
			panic(err)
		}
		fmt.Println("----")
		for _, commit := range commits {
			if commit.IsValidCommit() != nil {
				fmt.Printf("Invalid commit detected! %d\n", mr.IID)
				if !mr.IsDraft {
					if err := internal.PostNote(mr.IID, gitlabToken); err != nil {
						panic(err)
					}
					if err := mr.DenyMR(gitlabToken); err != nil {
						panic(err)
					}
				}
			}
		}
		fmt.Println("----")
	}
}
