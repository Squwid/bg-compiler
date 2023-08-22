package entry

type Submission struct {
	Script string
	Image  string
}

type Result struct {
	StdOut   string
	StdErr   string
	Duration int64
	Memory   int64
}
