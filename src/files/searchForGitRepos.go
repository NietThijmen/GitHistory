package files

import (
	"GitHistory/src/flags"
	"GitHistory/src/structs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func SearchForGitRepos() []structs.GitRepo {
	// start treewalking the current directory
	// for each directory, check if it is a git repo
	// if it is, add it to the list of git repos
	repos := []structs.GitRepo{}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if strings.Contains(path, "vendor") {
			return nil
		}

		if strings.Contains(path, "node_modules") {
			return nil
		}

		_, err = os.Stat(filepath.Join(path, ".git"))

		if err == nil {
			var repo structs.GitRepo

			repository, err := git.PlainOpen(path)

			if err != nil {
				return nil
			}

			remotes, err := repository.Remotes()

			if err != nil {
				return nil
			}

			for _, remote := range remotes {
				url := remote.Config().URLs[0]
				if !strings.Contains(url, flags.RemoteUrl) && flags.RemoteUrl != "" {
					continue
				}

				repo.RemoteName = remote.Config().Name
				repo.RemoteURL = remote.Config().URLs[0]
				repo.LocalPath = path

				repos = append(repos, repo)
			}
		}

		return nil
	})

	return repos
}
