//go:build windows
// +build windows

package events_stream

import (
	"io"

	"github.com/dest4590/evilcap/v2/session"
)

func (mod *EventsStream) viewBLEEvent(output io.Writer, e session.Event) {

}
