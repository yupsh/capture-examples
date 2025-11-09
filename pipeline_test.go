package capture_test

import (
	"bytes"
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/yupsh/capture"
	"github.com/yupsh/echo"
	"github.com/yupsh/grep"
	"github.com/yupsh/sort"
)

func ExampleCapture_pipeline() {
	// echo "apple\nbanana\napple\ncherry" | grep apple | sort | capture
	var stdout, stderr bytes.Buffer

	gloo.MustRun(Pipeline(
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
