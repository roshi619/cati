package espeak

import (
	"os"
	"os/exec"
)

// CATIfication is an espeak catification.
type CATIfication struct {
	// -g
	WordGap int
	// -p
	PitchAdjustment int
	// -s
	Rate int
	// -v
	VoiceName string

	Text string
}

// Send triggers a spoken catification.
func (n *CATIfication) Send() error {
	cmd := exec.Command("espeak", "-v", n.VoiceName, "--", n.Text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
