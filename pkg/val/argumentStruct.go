package val

type ArgumentsStruct struct {
	LaunchShellPath   string
	SrcTsvPaths       []string
	ArgsCon           string
	ArgsMap           map[string]string
	ImportPaths       []string
	IsSaveRepbashLine bool
}
