// Copyright ©2018 Peter Paolucci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package custplotter_test

import (
	"log"
	"testing"

	"github.com/pplcc/plotext/custplotter"
	"github.com/pplcc/plotext/custplotter/internal"
	"gonum.org/v1/plot"
)

func TestNewVBars(t *testing.T) {
	testTOHLCVs := internal.CreateTOHLCVTestData()

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	bars, err := custplotter.NewVBars(testTOHLCVs)
	if err != nil {
		log.Panic(err)
	}

	p.Add(bars)

	testFile := "testdata/vbars.png"
	err = p.Save(180, 100, testFile)
	if err != nil {
		log.Panic(err)
	}

	internal.TestImage(t, testFile)
}
