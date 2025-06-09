package main

import (
	"fmt"
	"io"
	// "iter"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

//TODO: i need read this again

// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := range countdownFrom(3) {
// 		fmt.Fprintln(out, i)
// 		sleeper.Sleep()
// 	}

// 	fmt.Fprint(out, finalWord)
// }

// func countdownFrom(from int) iter.Seq[int] {
// 	return func(yield func(int) bool) {
// 		for i := from; i > 0; i-- {
// 			if !yield(i) {
// 				return
// 			}
// 		}
// 	}
// }
