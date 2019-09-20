/*
 * @Author: guiguan
 * @Date:   2019-09-19T00:53:54+10:00
 * @Last modified by:   guiguan
 * @Last modified time: 2019-09-20T00:48:19+10:00
 */

package caster

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCaster_PubSub(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := New(ctx)

	go func() {
		ch, _ := c.Sub(nil, 0)

		var lastM interface{}

		for m := range ch {
			fmt.Println("c1", m)
			lastM = m
			if m == 5 {
				c.Unsub(ch)
			}
		}

		fmt.Println("done")

		if lastM != 5 {
			t.Fatalf("c1 didn't close at 5")
		}
	}()

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch, _ := c.Sub(ctx, 0)

		var lastM interface{}

		for m := range ch {
			fmt.Println("c2", m)
			lastM = m
			if m == 3 {
				cancel()
			}
		}

		fmt.Println("done")

		if lastM != 3 {
			t.Fatalf("c1 didn't close at 5")
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		c.Pub(i)
	}

	cancel()

	time.Sleep(10 * time.Millisecond)

	ok := c.Pub("test")
	if ok {
		t.Fatalf("`Pub` should not be ok after `Close`")
	}
}
