package capture_test

import (
	"bytes"
	"fmt"

	capture "github.com/yupsh/capture"
	echo "github.com/yupsh/echo"
	yup "github.com/gloo-foo/framework"
	. "github.com/gloo-foo/pipe"
)

func ExampleCapture_merged() {
	// Merge stdout and stderr to the same writer (like 2>&1)
	var combined bytes.Buffer

	yup.MustRun(Pipeline(
		echo.Echo("combined output"),
		capture.Capture(&combined, &combined),
	))

	fmt.Print(combined.String())
	// Output:
	// combined output
}

