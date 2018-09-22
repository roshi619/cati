package say

import (
	"fmt"
	"os"
	"os/exec"
)

// CATIfication is a say catification.
type CATIfication struct {
	Text string
	// Voice is the name of the voice to speak the catification.
	Voice string
	// Rate controls how fast the voice will speak. It's measured in words per
	// minute.
	Rate int
}

// Send triggers a spoken catification.
func (n *CATIfication) Send() error {
	r := fmt.Sprint(n.Rate)
	cmd := exec.Command("say", "-v", n.Voice, "-r", r, n.Text)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
