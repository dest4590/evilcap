package net_sniff

import (
	"time"

	"github.com/dest4590/evilcap/v2/log"
)

type SnifferStats struct {
	NumLocal     uint64
	NumMatched   uint64
	NumDumped    uint64
	NumWrote     uint64
	BytesSeen    uint64
	BytesMatched uint64
	Started      time.Time
	FirstPacket  time.Time
	LastPacket   time.Time
}

func NewSnifferStats() *SnifferStats {
	return &SnifferStats{
		NumLocal:     0,
		NumMatched:   0,
		NumDumped:    0,
		NumWrote:     0,
		BytesSeen:    0,
		BytesMatched: 0,
		Started:      time.Now(),
		FirstPacket:  time.Time{},
		LastPacket:   time.Time{},
	}
}

// PPS returns the average packets per second over the whole capture duration.
func (s *SnifferStats) PPS() float64 {
	if s.FirstPacket.IsZero() {
		return 0
	}
	elapsed := s.LastPacket.Sub(s.FirstPacket).Seconds()
	if elapsed <= 0 {
		return 0
	}
	return float64(s.NumMatched) / elapsed
}

func (s *SnifferStats) Print() error {
	first := "never"
	last := "never"

	if !s.FirstPacket.IsZero() {
		first = s.FirstPacket.String()
	}
	if !s.LastPacket.IsZero() {
		last = s.LastPacket.String()
	}

	log.Info("Sniffer Started    : %s", s.Started)
	log.Info("First Packet Seen  : %s", first)
	log.Info("Last Packet Seen   : %s", last)
	log.Info("Local Packets      : %d", s.NumLocal)
	log.Info("Matched Packets    : %d", s.NumMatched)
	log.Info("Dumped Packets     : %d", s.NumDumped)
	log.Info("Wrote Packets      : %d", s.NumWrote)
	log.Info("Bytes Seen         : %d", s.BytesSeen)
	log.Info("Bytes Matched      : %d", s.BytesMatched)
	log.Info("Avg PPS            : %.2f", s.PPS())

	return nil
}
