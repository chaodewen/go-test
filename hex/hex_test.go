package hex

import (
"testing"
"strconv"
)

func TestHexConvert(t *testing.T) {
	t.Logf("hex: %+v", strconv.FormatInt(6635586986474340356, 36))
}
