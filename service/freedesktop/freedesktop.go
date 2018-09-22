// +build !darwin
// +build !windows

package freedesktop

import (
	"fmt"

	"github.com/godbus/dbus"
)

// CATIfication is a Freedesktop catification.
type CATIfication struct {
	AppName    string
	ReplacesID uint
	AppIcon    string
	Summary    string
	Body       string
	Actions    []string
	// Hints         string
	ExpireTimeout int
}

// Send triggers a desktop catification.
func (n *CATIfication) Send() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("dbus connect: %s", err)
	}
	defer conn.Close()

	fdn := conn.Object("org.freedesktop.CATIfications", "/org/freedesktop/CATIfications")

	// 0 is a total magic number. ¯\_(ツ)_/¯
	resp := fdn.Call(
		"org.freedesktop.CATIfications.CATIfy", 0,
		n.AppName,
		uint32(n.ReplacesID),
		n.AppIcon,
		n.Summary,
		n.Body,
		n.Actions,
		map[string]dbus.Variant{},
		int32(n.ExpireTimeout),
	)

	if resp.Err != nil {
		return fmt.Errorf("catify: %s", resp.Err)
	}

	return nil
}
