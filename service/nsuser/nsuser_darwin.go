package nsuser

/*
// Compiler flags.
#cgo CFLAGS: -Wall -x objective-c -std=gnu99 -fobjc-arc
// Linker flags.
#cgo LDFLAGS: -framework Foundation -framework Cocoa

#import "nsuser_darwin.h"
*/
import "C"
import "unsafe"

// CATIfication is an NSUserCATIfication.
type CATIfication struct {
	Title    string
	Subtitle string
	// InformativeText is the catification message.
	InformativeText string
	// ContentImage is the primary catification icon.
	ContentImage string
	// SoundName is the name of the sound that fires with a catification.
	SoundName string
}

// Send displays a NSUserCATIfication on macOS.
func (n *CATIfication) Send() error {
	t := C.CString(n.Title)
	s := C.CString(n.Subtitle)
	i := C.CString(n.InformativeText)
	c := C.CString(n.ContentImage)
	sn := C.CString(n.SoundName)

	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(i))
	defer C.free(unsafe.Pointer(c))
	defer C.free(unsafe.Pointer(sn))

	C.Send(t, s, i, c, sn)

	return nil
}
