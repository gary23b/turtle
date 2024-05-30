package turtleutil

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBroker(t *testing.T) {
	b := NewBroker[int]()
	wg := &sync.WaitGroup{}
	readers := 100
	quitters := 100
	wg.Add(readers + quitters)
	for i := 0; i < readers; i++ {
		go func() {
			c := b.Subscribe()
			defer b.Unsubscribe(c)

			got := []int{}
			for x := range c {
				got = append(got, x)
			}

			expectedValue := 0
			for _, x := range got {
				// require.Equal(t, expectedValue, x)
				if expectedValue != x {
					t.Error("not equal", expectedValue, x)
				}
				expectedValue++
			}
			wg.Done()
		}()
	}

	for i := 0; i < quitters; i++ {
		go func() {
			c := b.Subscribe()
			defer b.Unsubscribe(c)
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
			wg.Done()
		}()
	}

	time.Sleep(time.Second * 2)
	for i := 0; i < 200; i++ {
		b.Publish(i)
	}

	b.Stop()
	wg.Wait()
}

func TestBroker_noSubscribers(t *testing.T) {
	b := NewBroker[int]()

	for i := 0; i < 2000; i++ {
		b.Publish(i)
	}

	time.Sleep(time.Millisecond * 250)
	b.Stop()
}

func TestBroker_subscriberDoesReadMessages(t *testing.T) {
	b := NewBroker[int]()

	c := b.Subscribe()
	time.Sleep(time.Millisecond * 200)

	for i := 0; i < 200; i++ {
		b.Publish(i)
	}

	time.Sleep(time.Millisecond * 200)
	b.Stop()

	got := []int{}
	for x := range c {
		got = append(got, x)
	}
	require.Equal(t, 100, len(got))

	expectedValue := 0
	for _, x := range got {
		// require.Equal(t, expectedValue, x)
		if expectedValue != x {
			t.Error("not equal", expectedValue, x)
		}
		expectedValue++
	}

	cBad := make(chan int, 100)
	b.Unsubscribe(cBad)

	b.Unsubscribe(c)
	time.Sleep(time.Millisecond * 200)
}
