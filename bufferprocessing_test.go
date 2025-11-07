package capture_test

import (
	"bytes"
	"fmt"

	capture "github.com/yupsh/capture"
	cat "github.com/yupsh/cat"
	echo "github.com/yupsh/echo"
	yup "github.com/gloo-foo/framework"
	. "github.com/gloo-foo/pipe"
)

func ExampleCapture_bufferProcessing() {
	// echo "one\ntwo\nthree" | cat | capture (then process captured buffer)
	var stdout, stderr bytes.Buffer

	yup.MustRun(Pipeline(
		echo.Echo("one\ntwo\nthree"),
		cat.Cat(),
		capture.Capture(&stdout, &stderr),
	))

	// Process the captured output
	lines := bytes.Count(stdout.Bytes(), []byte("\n"))
	fmt.Printf("Captured %d bytes in %d lines\n", stdout.Len(), lines)
	fmt.Printf("Content:\n%s", stdout.String())

	// Output:
	// Captured 14 bytes in 3 lines
	// Content:
	// one
	// two
	// three
}

