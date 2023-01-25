package cli

type ArgumentInformation struct {
	Args        []string
	StringFlags map[string]string
	BoolFlags   map[string]bool
}

func (a *ArgumentInformation) VerifySetup() bool {
	if (a.StringFlags["moduleName"] == "" && a.Args[0] == "") && a.StringFlags["dirName"] == "" && a.StringFlags["binaryDirName"] == "" {
		return false
	}

	return true
}
