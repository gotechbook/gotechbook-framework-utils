package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func isFunction(f interface{}) bool {
	actual := reflect.TypeOf(f)
	return actual.Kind() == reflect.Func && actual.NumIn() == 0 && actual.NumOut() > 0
}
func isChan(a interface{}) bool {
	if isNil(a) {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Chan
}
func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}
	return false
}
func getMapKeys(t *testing.T, m interface{}) []string {
	if reflect.ValueOf(m).Kind() != reflect.Map {
		t.Fatal(errors.New("GetMapKeys should receive a map"))
	}
	if reflect.TypeOf(m).Key() != reflect.TypeOf("bla") {
		t.Fatal(errors.New("GetMapKeys should receive a map with string keys"))
	}
	t.Helper()
	res := make([]string, 0)
	for _, k := range reflect.ValueOf(m).MapKeys() {
		res = append(res, k.String())
	}
	return res
}
func shouldEventuallyReturn(t testing.TB, f interface{}, v interface{}, timeouts ...time.Duration) {
	t.Helper()
	interval := 10 * time.Millisecond
	timeout := time.After(500 * time.Millisecond)
	switch len(timeouts) {
	case 1:
		interval = timeouts[0]
		break
	case 2:
		interval = timeouts[0]
		timeout = time.After(timeouts[1])
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if isFunction(f) {
		for {
			select {
			case <-timeout:
				t.Fatalf("function f never returned value %s", v)
			case <-ticker.C:
				val, err := pollFuncReturn(f)
				if err != nil {
					t.Fatal(err)
				}
				if v == val {
					return
				}
			}
		}
	} else {
		t.Fatal("ShouldEventuallyEqual should receive a function with no args and more than 0 outs")
		return
	}
}
func shouldAlwaysReturn(t testing.TB, f interface{}, v interface{}, timeouts ...time.Duration) {
	t.Helper()
	interval := 10 * time.Millisecond
	timeout := time.After(50 * time.Millisecond)
	switch len(timeouts) {
	case 1:
		interval = timeouts[0]
		break
	case 2:
		interval = timeouts[0]
		timeout = time.After(timeouts[1])
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if isFunction(f) {
		for {
			select {
			case <-timeout:
				return
			case <-ticker.C:
				val, err := pollFuncReturn(f)
				if err != nil {
					t.Fatal(err)
				}
				if v != val {
					t.Fatalf("function f returned wrong value %s", val)
				}
			}
		}
	} else {
		t.Fatal("ShouldAlwaysReturn should receive a function with no args and more than 0 outs")
		return
	}
}
func shouldEventuallyReceive(t testing.TB, c interface{}, timeouts ...time.Duration) interface{} {
	t.Helper()
	if !isChan(c) {
		t.Fatal("ShouldEventuallyReceive c argument should be a channel")
	}
	v := reflect.ValueOf(c)

	timeout := time.After(500 * time.Millisecond)

	if len(timeouts) > 0 {
		timeout = time.After(timeouts[0])
	}

	recvChan := make(chan reflect.Value)

	go func() {
		v, ok := v.Recv()
		if ok {
			recvChan <- v
		}
	}()

	select {
	case <-timeout:
		t.Fatal(errors.New("timed out waiting for channel to receive"))
	case a := <-recvChan:
		return a.Interface()
	}

	return nil
}
func pollFuncReturn(f interface{}) (interface{}, error) {
	values := reflect.ValueOf(f).Call([]reflect.Value{})

	extras := []interface{}{}
	for _, value := range values[1:] {
		extras = append(extras, value.Interface())
	}

	success, message := vetExtras(extras)

	if !success {
		return nil, errors.New(message)
	}

	return values[0].Interface(), nil
}
func vetExtras(extras []interface{}) (bool, string) {
	for i, extra := range extras {
		if extra != nil {
			zeroValue := reflect.Zero(reflect.TypeOf(extra)).Interface()
			if !reflect.DeepEqual(zeroValue, extra) {
				message := fmt.Sprintf("unexpected non-nil/non-zero extra argument at index %d:\n\t<%T>: %#v", i+1, extra, extra)
				return false, message
			}
		}
	}
	return true, ""
}
func fixtureGoldenFileName(t *testing.T, name string) string {
	t.Helper()
	return filepath.Join("fixtures", name+".golden")
}
func waitForServerToBeReady(t testing.TB, out *bufio.Reader) {
	t.Helper()
	shouldEventuallyReturn(t, func() bool {
		line, _, err := out.ReadLine()
		if err != nil {
			t.Fatal(err)
		}
		return strings.Contains(string(line), "all modules started!")
	}, true, 100*time.Millisecond, 30*time.Second)
}
func startProcess(t testing.TB, program string, args ...string) *exec.Cmd {
	t.Helper()
	return exec.Command(program, args...)
}
