package flags

import "flag"

var (
	RemoteUrl = "git@github.com"
	Name      = "Thijmen"
	ThisMonth = false
)

func Parse() {
	flag.StringVar(&RemoteUrl, "remote-url", RemoteUrl, "The remote url of the git repository")
	flag.StringVar(&Name, "name", Name, "The name of the user")
	flag.BoolVar(&ThisMonth, "this-month", ThisMonth, "Only show commits from this month")

	flag.Parse()
}
