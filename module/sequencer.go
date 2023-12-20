package module

import (
	"github.com/gopxl/beep"
	"slices"
	"time"
)

type SeqItem struct {
	Freq float64
	At   time.Duration
	Len  time.Duration
}

type Sequencer struct {
	Module
	Sequence []SeqItem
	Playing  []SeqItem
	Len      time.Duration
	Shift    time.Duration
	Time     time.Duration
	Cursor   uint32
	Loop     bool
}

func (s *Sequencer) Init(rate beep.SampleRate) {
	s.Module.Init(rate)
	s.Sequence = make([]SeqItem, 0)
	s.Playing = make([]SeqItem, 0)

	s.AddOutput("freq", PortOutFreq)
	s.AddOutput("gate", PortOutGate)
}

func (s *Sequencer) GetName() string {
	return "Sequencer"
}

func (s *Sequencer) SetLoop(loop bool) {
	s.Loop = loop
}

func (s *Sequencer) Update(time time.Duration) {
	s.Time = time
	time -= s.Shift

	for i := 0; i < len(s.Playing); i++ {
		item := s.Playing[i]

		if item.At+item.Len < time {
			s.ConnectionWrite(PortOutGate, 0)
			s.Playing = append(s.Playing[:i], s.Playing[i+1:]...)
			i--
		}
	}

	for i := s.Cursor; i < uint32(len(s.Sequence)); i++ {
		item := s.Sequence[i]

		if item.At > time {
			break
		}

		s.ConnectionWrite(PortOutFreq, item.Freq)
		s.ConnectionWrite(PortOutGate, 1)

		s.Playing = append(s.Playing, item)

		s.Cursor++
	}

	if s.Loop && s.Cursor >= uint32(len(s.Sequence)) && len(s.Playing) == 0 {
		s.Seek(0)
	}
}

func (s *Sequencer) Add(freq float64, at time.Duration, len time.Duration) {
	if s.Len < at+len {
		s.Len = at + len
	}

	s.Sequence = append(s.Sequence, SeqItem{Freq: freq, At: at, Len: len})

	slices.SortFunc(s.Sequence, func(a, b SeqItem) int {
		if a.At < b.At {
			return -1
		} else if a.At > b.At {
			return 1
		} else {
			return 0
		}
	})
}

func (s *Sequencer) Append(freq float64, len time.Duration) {
	s.Add(freq, s.Len, len)
}

func (s *Sequencer) AppendAfter(freq float64, after time.Duration, len time.Duration) {
	s.Add(freq, after+s.Len, len)
}

func (s *Sequencer) Seek(time time.Duration) {
	cursor := uint32(0)
	s.Shift = s.Time - time

	for _, note := range s.Sequence {
		if note.At >= time {
			break
		}
		s.Cursor++
	}

	s.Cursor = cursor
}
