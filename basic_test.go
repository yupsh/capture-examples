package capture_test

import (
	"bytes"
	"fmt"

	capture "github.com/yupsh/capture"
	echo "github.com/yupsh/echo"
	yup "github.com/gloo-foo/framework"
	. "github.com/gloo-foo/pipe"
)

func ExampleCapture_basic() {
	// echo "Hello, World!" | capture
	var stdout, stderr bytes.Buffer

	yup.MustRun(Pipeline(
		echo.Echo("Hello, World!"),
		capture.Capture(&stdout, &stderr),
	))

	fmt.Print(stdout.String())
	// Output:
	// Hello, World!
}

