package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"log"
)

var (
	ballastSize = flag.Int("s", 0, "ballast size MB")
	n           = flag.Int("n", 10000, "")
)

const (
	KB = 1 << 10
	MB = 1 << 20
	GB = 1 << 30
)

var ballast []byte

func main() {
	flag.Parse()
	ballast = make([]byte, *ballastSize<<20)

	go printMemstat()

	for j := 0; j < *n; j++ {
		b := make([]byte, 1*MB)
		for k := 0; k < 1<<20; k++ {
			b[k] = '0'
		}
	}
}

var format = "NumGC: %d %d/s, STW: max %s, avg %s; Heap: %s %s/s, live objects: %s, ReadMemStats: %s"

func printMemstat() {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	var lastNumGC uint32
	var lastTotalAlloc uint64
	memstat := &runtime.MemStats{}

	for range t.C {
		t := time.Now()
		runtime.ReadMemStats(memstat)
		cost := time.Since(t)

		gcCnt := memstat.NumGC - lastNumGC
		lastNumGC = memstat.NumGC

		alloc := memstat.TotalAlloc - uint64(lastTotalAlloc)
		lastTotalAlloc = memstat.TotalAlloc

		maxGCTimes := time.Duration(max(memstat.PauseNs))
		avgGCTimes := time.Duration(avg(memstat.PauseNs))

		info := fmt.Sprintf(format, memstat.NumGC, gcCnt,
			maxGCTimes, avgGCTimes, formatUint64(memstat.HeapAlloc),
			formatUint64(alloc), formatUint64(memstat.Mallocs-memstat.Frees),
			cost,
		)

		log.Println(info)
	}
}

func avg(a [256]uint64) uint64 {
	sum := uint64(0)
	for _, v := range a {
		sum += v
	}
	return sum / 256
}

func max(a [256]uint64) uint64 {
	v := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > v {
			v = a[i]
		}
	}
	return v
}

func formatUint64(v uint64) string {
	var buf [16]byte
	w := len(buf)
	w--
	buf[w] = 'B'

	var b byte
	if v < KB {
	} else if v < MB {
		v >>= 10
		b = 'K'
	} else if v < GB {
		v >>= 20
		b = 'M'
	} else {
		b = 'G'
		v >>= 30
	}

	w--
	buf[w] = b

	for v > 0 {
		digit := v % 10
		w--
		buf[w] = byte(digit) + '0'
		v /= 10
	}
	return string(buf[w:])
}
