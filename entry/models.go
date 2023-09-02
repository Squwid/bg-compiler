package entry

type Submission struct {
	Script string  `json:"script"`
	Image  string  `json:"image"`
	Count  int     `json:"count"`
	Input  *string `json:"input"` // Text input to write to a file.

	Cmd string `json:"cmd"` // Command to run the script.
}

type Result struct {
	StdOut   string `json:"std_out"`
	StdErr   string `json:"std_err"`
	TimedOut bool   `json:"timed_out"`
}
