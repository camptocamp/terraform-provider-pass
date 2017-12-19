/*
 * goprogressbar
 *     Copyright (c) 2016-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package main

import (
	"fmt"
	"time"

	"github.com/muesli/goprogressbar"
)

func main() {
	pb := &goprogressbar.ProgressBar{
		Text:    "Current Progress",
		Total:   1000,
		Current: 0,
		Width:   60,
	}

	for i := 1; i <= 1000; i++ {
		pb.PrependText = fmt.Sprintf("%d of %d", i, pb.Total)
		pb.Current = int64(i)

		time.Sleep(23 * time.Millisecond)
		pb.LazyPrint()
	}

	fmt.Println()
}
