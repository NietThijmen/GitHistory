package structs

type GitCommit struct {
	AuthorName    string
	AuthorEmail   string
	CommitMessage string
	CommitHash    string

	GitRepo GitRepo
}
