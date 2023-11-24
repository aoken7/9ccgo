package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		sourceFile       string
		expectedExitCode int
	}{
		{"test001.c", 123},
		{"test002.c", 3},
		{"test003.c", 10},
		{"test004.c", 3},
		{"test005.c", 126},
	}

	for _, tc := range testCases {
		t.Run(tc.sourceFile, func(t *testing.T) {
			var stderr bytes.Buffer

			defer os.Remove("tmp.s")

			outputFile, err := os.Create("tmp.s")
			if err != nil {
				log.Fatalf("Failed to create/open file: %v", err)
			}
			defer outputFile.Close() // 適切にファイルを閉じることを確保

			// C言語ソースファイルをコンパイルして実行
			compile := exec.Command("go", "run", "main.go", "test_c/"+tc.sourceFile)

			compile.Stdout = outputFile
			compile.Stderr = &stderr

			err = compile.Run()
			if err != nil {
				exitError, ok := err.(*exec.ExitError)
				if !ok {
					t.Fatalf("Error executing the command: %v", exitError)
				}
				if exitError.ExitCode() != 0 {
					t.Fatalf("Compile error: %v", stderr.String())
				}
			}

			asmbly := exec.Command("cc", "-no-pie", "-o", "tmp", "tmp.s")
			asmbly.Stderr = &stderr

			err = asmbly.Run()
			if err != nil {
				exitError, ok := err.(*exec.ExitError)
				if !ok {
					t.Fatalf("Error executing the command: %v", exitError)
				}
				if exitError.ExitCode() != 0 {
					t.Fatalf("Asemble error: %v", stderr.String())
				}
			}

			run := exec.Command("./tmp")
			err = run.Run()

			if err != nil {
				exitError, ok := err.(*exec.ExitError)
				if !ok {
					t.Fatalf("Error executing the command: %v", err)
				}
				actualExitCode := exitError.ExitCode()
				if actualExitCode != tc.expectedExitCode {
					t.Fatalf("Exit code mismatch. Expected: %d, Actual: %d", tc.expectedExitCode, actualExitCode)
				}
			}
		})
	}
}
