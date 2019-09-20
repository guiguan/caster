/*
 * @Author: guiguan
 * @Date:   2019-09-19T23:52:54+10:00
 * @Last modified by:   guiguan
 * @Last modified time: 2019-09-20T10:26:20+10:00
 */

package caster_test

import (
	"fmt"
	"sync"

	"github.com/guiguan/caster"
)

func Example() {
	wg := new(sync.WaitGroup)

	c := caster.New(nil)

	// register subscribers
	for i := 0; i < 5; i++ {
		ch, _ := c.Sub(nil, 1)
		go receiveFromChannel(ch, wg)
	}

	// broadcast
	for i := 0; i < 100; i++ {
		c.Pub(i)
	}

	// close caster and all subscriber channels
	c.Close()

	wg.Wait()

	// Output:
	// received 100 messages
	// received 100 messages
	// received 100 messages
	// received 100 messages
	// received 100 messages
}

func receiveFromChannel(ch chan interface{}, wg *sync.WaitGroup) {
	wg.Add(1)

	counter := 0

	for range ch {
		counter++
	}

	fmt.Printf("received %v messages\n", counter)

	wg.Done()
}
