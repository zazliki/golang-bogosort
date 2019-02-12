package bogosort

import (
	"math/rand"
	"time"
)

var s int
var t time.Duration

func IsSorted(arr []int) bool {
	l := len(arr) - 1
	for i := 0; i < l; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Shuffle(arr []int) {
	rand.Seed(time.Now().UTC().UnixNano())

	l := len(arr)
	for i := l; i > 0; i-- {
		j := rand.Intn(i)
		arr[i-1], arr[j] = arr[j], arr[i-1]
	}
}

func Bogosort(arr []int) {
	s = 0;
	start := time.Now()

	for IsSorted(arr) == false {
		Shuffle(arr)
		s++
	}

	finish := time.Now()
	t = finish.Sub(start)
}

func Steps() int {
	return s
}

func Time() time.Duration {
	return t
}
