package concurrent

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConcurrentMap(t *testing.T) {

}

func TestGo(t *testing.T) {
	fs := []func() error{
		func() error {
			time.Sleep(time.Second)
			fmt.Println("func 1 finished")
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("func 2 finished")
			panic("func 2 panic")
			return errors.New("func 2 error")
		},
		func() error {
			time.Sleep(3 * time.Second)
			fmt.Println("func 3 finished")
			return nil
		},
	}

	if err := goAndWait(fs...); err != nil {
		fmt.Println("goAndWait err", err)
	}
	if es := goAndWaitErrors(fs...); len(es) > 0 {
		fmt.Println("goAndWait err", es)
	}

}

func goAndWait(fs ...func() error) error {
	var e error
	var wg sync.WaitGroup
	for _, f := range fs {
		wg.Add(1)
		f := f
		go func() {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					e = errors.New(err.(string))
				}
			}()
			if err := f(); err != nil {
				e = err
			}
		}()
	}
	wg.Wait()
	return e
}
func goAndWaitErrors(fs ...func() error) []error {
	var es []error = make([]error, len(fs))
	var wg sync.WaitGroup
	for i, f := range fs {
		wg.Add(1)
		f := f
		i := i
		go func() {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					es[i] = errors.New(err.(string))
				}
			}()
			if err := f(); err != nil {
				es[i] = err
			}
		}()
	}
	wg.Wait()
	return es
}

func TestChan(t *testing.T) {
	//c := make(chan)
}
