package cli

import (
	"fmt"
)

type Handler struct {
	Debug bool
}

func (h *Handler) Router(funcName string, args *ArgumentInformation) {
	switch funcName {

	case "new":
		{
			if args.VerifySetup() {
				// setting up the file structure
				err := h.SetupFiles(args.StringFlags["dirName"], args.StringFlags["binaryDirName"], args.BoolFlags["library"])
				if err != nil {
					fmt.Printf("Unable to setup project %s\n", err)
					break
				}

			} else {
				// malformed input error
				fmt.Printf("Error:\n%s", ERR_WRONG_MODULE_CREATION_FORMAT.Error())
			}

		}

	case "help":
		fmt.Printf(fmt.Sprintf(`This is a tool for managing go code-bases 
Use: go-wizard [function] [optional input] [flags]
%s
`, New_Example))

	default:
		fmt.Printf("Unknown function %s\n", funcName)
	}
}
