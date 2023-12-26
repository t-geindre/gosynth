package node

import (
	"gosynth/output"
)

type Streamer struct {
	Node
	Streamer *output.Streamer
}

func NewStreamer(width, height int, streamer *output.Streamer) *Streamer {
	s := &Streamer{}
	s.Node = *NewNode(width, height, s)
	s.Streamer = streamer
	s.Resize(width, height)

	return s
}
