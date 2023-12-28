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
	Playing  *SeqItem
	Len      time.Duration
	Shift    time.Duration
	Time     time.Duration
	Cursor   uint32
	Loop     bool
}

func (s *Sequencer) Init(rate beep.SampleRate) {
	s.Module.Init(rate, s)
	s.Sequence = make([]SeqItem, 0)
}

func (s *Sequencer) SetLoop(loop bool) {
	s.Loop = loop
}

func (s *Sequencer) Update(time time.Duration) {
	s.Module.Update(time)

	s.Time = time
	time -= s.Shift

	if s.Playing != nil && s.Playing.At+s.Playing.Len < time {
		s.Playing = nil
		s.ConnectionWrite(PortOutGate, 0)
	}

	for i := s.Cursor; i < uint32(len(s.Sequence)); i++ {
		item := s.Sequence[i]

		if item.At > time {
			break
		}

		if s.Playing != nil {
			s.ConnectionWrite(PortOutGate, 0)
		}

		if item.Freq > 0 {
			s.ConnectionWrite(PortOutFreq, item.Freq)
			s.ConnectionWrite(PortOutGate, 1)
		}
		s.Playing = &item

		s.Cursor++
	}

	if s.Loop && s.Cursor >= uint32(len(s.Sequence)) && s.Playing == nil {
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
