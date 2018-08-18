package testmain_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	func() {
		fmt.Println("Prepare test")
	}()
	ret := m.Run()
	func() {
		fmt.Println("Teardown test")
	}()
	os.Exit(ret)
}
