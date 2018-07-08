package pget

import "fmt"

type ranges struct {
	high, low, worker uint
}

func (o *Object) makeRanges(i, split uint) *ranges {
	low := split * i
	high := low + split - 1
	if i == uint(o.Procs)-1 {
		high = o.filesize
	}
	return &ranges{
		low:    low,
		high:   high,
		worker: i,
	}
}

func (r *ranges) filesize() uint {
	return r.high - r.low
}

func (r *ranges) String() string {
	return fmt.Sprintf("bytes=%d-%d", r.low, r.high)
}
