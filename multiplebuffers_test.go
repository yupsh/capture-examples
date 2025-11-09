package capture_test

import (
	"bytes"
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/capture"
	"github.com/yupsh/echo"
)

func ExampleCapture_multipleBuffers() {
	// echo "standard output data" | capture (separate stdout and stderr)
	var stdout, stderr bytes.Buffer

	gloo.MustRun(Pipeline(
		echo.Echo("standard output data"),
		capture.Capture(&stdout, &stderr),
	))

	fmt.Printf("Stdout: %s", stdout.String())
	fmt.Printf("Stderr length: %d\n", stderr.Len())

	// Output:
	// Stdout: standard output data
	// Stderr length: 0
}
