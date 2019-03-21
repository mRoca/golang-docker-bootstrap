package cli

import "github.com/mRoca/golang-docker-bootstrap/internal/common"

// BuildCliString concatenates strings. This is an example function
func BuildCliString(input string, suffix string) string {
	return common.BuildString("Cli text: ", input, suffix)
}
