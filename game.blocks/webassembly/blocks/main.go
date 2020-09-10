// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example jsgo

package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"syscall/js"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/blocks/blocks"
)

// set the value to js -> output
// func setResult(val js.Value) {
// 	js.Global().Get("document").Call("getElementById", "result").Set("value", val)
// }

func add(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].String())
	blocks.NameUserGame = result.String()
	//setResult(result)
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
}

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	//blocks.NameUserGame = "--"
	registerCallbacks()

	blocks.KeyUUid = uuid.New().String()
	ebiten.SetWindowSize(blocks.ScreenWidth*2, blocks.ScreenHeight*2)
	ebiten.SetWindowTitle("Blocks (Ebiten Demo POC)")
	if err := ebiten.RunGame(&blocks.Game{}); err != nil {
		log.Fatal(err)
	}
}
