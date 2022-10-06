package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"
)

var ctx = context.Background()
var debug *bool = flag.Bool("debug", false, "debug")
var testEndpoint *bool = flag.Bool("test-endpoint", false, "test-endpoint")

func TestExample(t *testing.T) {
	var w os.File
	if *testEndpoint {
		if *debug {
			os.Setenv("LOGLEVEL", "DEBUG")
		}
		fmt.Println("in TestExample")
		if !*debug {
			w.Close()
		}
	}
}

func isEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func TestMain(m *testing.M) {
	flag.Parse()
	if isEnvExist("TESTENDPOINT") {
		fmt.Println("Found environment variable TESTENDPOINT, setting *testEndpoint = true")
		*testEndpoint = true
	}
	fmt.Printf("running with debug mode [%v]\n", *debug)
	fmt.Printf("running with test-endpoint mode [%v]\n", *testEndpoint)
	os.Exit(m.Run())
}
