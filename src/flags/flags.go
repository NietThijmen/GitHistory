package flags

import "flag"

var (
	RemoteUrl = ""
	Name      = "Thijmen"
	ThisMonth = false
	Month     = 0
)

func Parse() {
	flag.StringVar(&RemoteUrl, "remote-url", RemoteUrl, "The remote url of the git repository")
	flag.StringVar(&Name, "name", Name, "The name of the user")
	flag.BoolVar(&ThisMonth, "this-month", ThisMonth, "Only show commits from this month")
	flag.IntVar(&Month, "month", Month, "Only show commits from the specified month (1-12)")

	flag.Parse()
}
