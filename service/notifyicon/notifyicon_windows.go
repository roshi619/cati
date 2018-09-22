package catifyicon

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

// Balloon icons.
const (
	BalloonTipIconError   = "Error"
	BalloonTipIconWarning = "Warning"
	BalloonTipIconInfo    = "Info"
	BalloonTipIconNone    = "None"

	// DefaultIcon is the default icon.
	DefaultIcon = "[System.Drawing.Icon]::ExtractAssociatedIcon([System.Windows.Forms.Application]::ExecutablePath)"
)

const script = `
[void] [System.Reflection.Assembly]::LoadWithPartialName("System.Windows.Forms")

$n = New-Object System.Windows.Forms.CATIfyIcon
$n.Icon = {{.Icon}}
$n.BalloonTipIcon = "{{.BalloonTipIcon}}"
$n.BalloonTipText = "{{.BalloonTipText}}"
$n.BalloonTipTitle = "{{.BalloonTipTitle}}"
$n.Text = "{{.Text}}"

$n.Visible = $True
$n.ShowBalloonTip({{.Duration}})
`

// CATIfication is a Windows catification.
type CATIfication struct {
	// BalloonTipIcon is the catification icon.
	BalloonTipIcon string
	// BalloonTipText is the catification message.
	BalloonTipText string
	// BalloonTipTitle is the catification title.
	BalloonTipTitle string
	// Icon is the path to an .ico file.
	// Icon sets the icon that will appear in the systray for this application.
	// Icon must be 16 pixels high by 16 pixels wide
	// This is required to show the catification.
	Icon string

	// Text is the text shown when you hover over the app icon.
	Text string

	Duration int
}

// Send sends a Windows catification.
func (n *CATIfication) Send() error {
	if n.Icon == "" {
		n.Icon = DefaultIcon
	} else {
		n.Icon = fmt.Sprintf("%q", n.Icon)
	}

	tmpl, err := template.New("").Parse(script)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, n); err != nil {
		return err
	}

	cmd := exec.Command("PowerShell", "-Command", buf.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
