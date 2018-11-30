package nr

import (
	"fmt"
)

func nerror(s string, a ...interface{}) error {
	return fmt.Errorf(s, a...)
}
