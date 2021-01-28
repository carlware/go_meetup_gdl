package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const MaxGorutines = 10000
const maxSleep = 10

func main() {
	// gN := runtime.NumGoroutine()

	for i := 0; i < MaxGorutines; i++ {
		go func() {
			s := rand.Int() % maxSleep
			time.Sleep(time.Duration(s) * time.Second)
		}()
	}

	monitor()
	// fmt.Println("gorutines:", gN)
}

// 2. Crear una gorutina para monitorear las gorutinas
func monitor() {
	for range time.Tick(1 * time.Second) {
		gN := runtime.NumGoroutine()
		fmt.Println("gorutines:", gN)
	}
}
