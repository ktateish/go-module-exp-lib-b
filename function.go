package b

import (
	"fmt"

	c "github.com/ktateish/go-module-exp-lib-c"
)

//go:generate ./genvers.sh

func F(s string) string {
	t := c.F(s)
	return fmt.Sprintf("B %s: %s", Version, t)
}
