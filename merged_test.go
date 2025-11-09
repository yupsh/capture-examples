package capture_test

import (
	"bytes"
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/capture"
	"github.com/yupsh/echo"
)

func ExampleCapture_merged() {
	// Merge stdout and stderr to the same writer (like 2>&1)
	var combined bytes.Buffer

	gloo.MustRun(Pipeline(
		echo.Echo("combined output"),
		capture.Capture(&combined, &combined),
	))

	fmt.Print(combined.String())
	// Output:
	// combined output
}
