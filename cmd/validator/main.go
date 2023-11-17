package main

import (
	"encoding/json"
	"fmt"
	"gitlab-mr-commit-validator/internal"
	"net/http"
	"os"
	"strconv"
)

type MR struct {
	ID  int `json:"id"`
	IID int `json:"iid"`
}

func main() {
	args := os.Args
	gitlabToken := args[1]
	fmt.Println(gitlabToken)
	var data []MR
	err := getJson("https://gitlab.com/api/v4/projects/26177235/merge_requests?state=opened&access_token="+gitlabToken, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	for _, id := range data {
		var commits []internal.Commit
		err := getJson("https://gitlab.com/api/v4/projects/26177235/merge_requests/"+strconv.Itoa(id.IID)+"/commits?state=opened&access_token="+gitlabToken, &commits)
		if err != nil {
			panic(err)
		}
		fmt.Println("----")
		for _, commit := range commits {
			if commit.IsValidCommit() != nil {
				fmt.Printf("Invalid commit detected! %d\n", id.IID)
			}
		}
		fmt.Println("----")
	}
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
