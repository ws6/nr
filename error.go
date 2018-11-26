package nr

import (
	"fmt"
)

func nerror(s string) error {
	return fmt.Errorf(s)
}
