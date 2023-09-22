package processor

type Submission struct {
	Script string  `json:"script"`
	Image  string  `json:"image"`
	Count  int     `json:"count"`
	Input  *string `json:"input"` // Text input to write to a file.

	Cmd string `json:"cmd"` // Command to run the script.

	// File extension of the script. Defaults to .ext if not present
	Extension string `json:"extension"`
}
