package qwik

import (
	"time"

	"github.com/google/uuid"
)

type EventTransformer interface {
	FromFSEvent(FSEvent) QwikEvent
}

func NewEventTransformer(srcIP string) EventTransformer {
	return eventTransformer{
		srcIP: srcIP,
	}
}

type eventTransformer struct {
	srcIP string
}

func (et eventTransformer) FromFSEvent(evt FSEvent) QwikEvent {
	var modTime uint64
	var size uint64

	if evt.Info != nil {
		modTime = uint64(evt.Info.ModTime().UnixNano())
		size = uint64(evt.Info.Size())
	}

	return QwikEvent{
		Id:        newEventID(),
		Type:      QwikEvent_FILESYSTEM,
		Timestamp: uint64(time.Now().UnixNano()),
		SrcIp:     et.srcIP,
		Event: &QwikEvent_Fsevent{
			Fsevent: &FileSystemEvent{
				Path:         evt.Path,
				Op:           evt.Op,
				ModifiedTime: modTime,
				Size:         size,
			},
		},
	}
}

func newEventID() string {
	return uuid.New().String()
}
