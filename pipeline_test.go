package capture_test

import (
	"bytes"
	"fmt"

	capture "github.com/yupsh/capture"
	echo "github.com/yupsh/echo"
	yup "github.com/gloo-foo/framework"
	grep "github.com/yupsh/grep"
	. "github.com/gloo-foo/pipe"
	sort "github.com/yupsh/sort"
)

func ExampleCapture_pipeline() {
	// echo "apple\nbanana\napple\ncherry" | grep apple | sort | capture
	var stdout, stderr bytes.Buffer

	yup.MustRun(Pipeline(
		echo.Echo("apple\nbanana\napple\ncherry"),
		grep.Grep("apple"),
		sort.Sort(),
		capture.Capture(&stdout, &stderr),
	))

	fmt.Print(stdout.String())
	// Output:
	// apple
	// apple
}

