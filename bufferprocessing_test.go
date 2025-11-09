package capture_test

import (
	"bytes"
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/capture"
	"github.com/yupsh/cat"
	"github.com/yupsh/echo"
)

func ExampleCapture_bufferProcessing() {
	// echo "one\ntwo\nthree" | cat | capture (then process captured buffer)
	var stdout, stderr bytes.Buffer

	gloo.MustRun(Pipeline(
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
