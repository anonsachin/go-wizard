package cli

import (
	"errors"
	"fmt"
)

var (
	ERR_WRONG_MODULE_CREATION_FORMAT = errors.New(fmt.Sprintf(`
The inputs are not in the right format they have to look like:
%s`, New_Example))
)
