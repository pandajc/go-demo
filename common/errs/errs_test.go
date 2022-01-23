package errs

import (
	"fmt"
	pe "github.com/pkg/errors"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Printf("err: %+v", produceRuntimeErr())
	fmt.Println("\n==================")
	err := produceBizErr()
	fmt.Printf("%+v", err)
	fmt.Println("\n", IsBizErr(err))
}

func produceRuntimeErr() error {
	_, err := strconv.Atoi("")
	if err != nil {
		return pe.WithStack(err)
	}
	return nil
}

func produceRuntimeErr2() error {
	_, err := strconv.Atoi("")
	if err != nil {
		return err
	}
	return nil
}

func produceBizErr() error {
	return NewBizErrWithStack(1, "i am biz err")
}

const times = 100000

func TestPressure(t *testing.T) {
	startTime := time.Now()
	for i := 0; i < times; i++ {
		produceRuntimeErr2()
	}
	printBenchmarkInfo("err", startTime)
}

func printBenchmarkInfo(info string, startTime time.Time) {
	var memStats runtime.MemStats
	//var rusage syscall.Rusage
	var bToKb = func(b uint64) uint64 {
		return b / 1024
	}
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Alloc = %v \n TotalAlloc = %v \n Sys = %v \n gc = %v \n cost = %v",
		bToKb(memStats.Alloc), bToKb(memStats.TotalAlloc), bToKb(memStats.Sys), memStats.NumGC, time.Since(startTime))
}

//=== RUN   TestPressure
//Alloc = 3604
//TotalAlloc = 34561
//Sys = 13078
//gc = 8
//cost = 74.341ms--- PASS: TestPressure (0.07s)
//PASS

//Alloc = 1294
//TotalAlloc = 4868
//Sys = 11352
//gc = 1
//cost = 7.0592ms--- PASS: TestPressure (0.01s)
//PASS
