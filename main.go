package main

import (
	"GitHistory/src/files"
	"GitHistory/src/flags"
	"GitHistory/src/structs"
	"fmt"
	"github.com/go-git/go-git/v5"
	"strings"
	"time"
)

func main() {
	repos := files.SearchForGitRepos()
	flags.Parse()

	var commitMap = make(map[time.Time]structs.GitCommit)

	println("Found git repo:")
	for _, repo := range repos {
		println(repo.LocalPath)
	}

	// loop through all commits to all repos
	// for each commit, check if the author is the similar to flags.Name
	// if it is write it

	for _, repo := range repos {
		repository, err := git.PlainOpen(repo.LocalPath)
		if err != nil {
			panic(err)
		}

		logOptions := git.LogOptions{}
		if flags.ThisMonth {
			now := time.Now()
			firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			logOptions.Since = &firstOfMonth
		} else {
			logOptions.All = true
		}

		commits, err := repository.Log(&logOptions)
		if err != nil {
			println(err.Error())
			continue
		}

		for {
			commit, err := commits.Next()
			if err != nil {
				break // no more commits
			}

			if !strings.Contains(commit.Author.Name, flags.Name) {
				continue
			}

			// set commitmap with key datetime of commit and value commit
			commitMap[commit.Author.When] = structs.GitCommit{
				AuthorName:    commit.Author.Name,
				AuthorEmail:   commit.Author.Email,
				CommitMessage: commit.Message,
				CommitHash:    commit.Hash.String(),
				GitRepo:       repo,
			}
		}

	}

	// sort commitMap by datetime key
	var commits = make([]time.Time, 0, len(commitMap))
	for dateTime := range commitMap {
		commits = append(commits, dateTime)
	}

	// sort commits
	for i := 0; i < len(commits); i++ {
		for j := i + 1; j < len(commits); j++ {
			if commits[i].After(commits[j]) {
				commits[i], commits[j] = commits[j], commits[i]
			}
		}
	}

	for _, dateTime := range commits {
		commit := commitMap[dateTime]
		var templateString = "[%s] %s (%s)"
		println(fmt.Sprintf(templateString, commit.GitRepo.LocalPath, commit.CommitMessage, dateTime.String()))
	}
}
