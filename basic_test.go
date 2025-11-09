package capture_test

import (
	"bytes"
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/capture"
	"github.com/yupsh/echo"
)

func ExampleCapture_basic() {
	// echo "Hello, World!" | capture
	var stdout, stderr bytes.Buffer

	gloo.MustRun(Pipeline(
		echo.Echo("Hello, World!"),
		capture.Capture(&stdout, &stderr),
	))

	fmt.Print(stdout.String())
	// Output:
	// Hello, World!
}
