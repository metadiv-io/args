package args

import (
	"os"
	"strings"
)

type Args struct {
	Args  []string
	Flags map[string]string
}

func New() *Args {
	args := &Args{
		Args:  make([]string, 0),
		Flags: make(map[string]string),
	}
	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") {
			flagKey := strings.TrimPrefix(os.Args[i], "--")
			flagKey = strings.TrimPrefix(flagKey, "-")

			if len(os.Args) >= i+2 && !strings.HasPrefix(os.Args[i+1], "-") {
				args.Flags[flagKey] = os.Args[i+1]
			} else {
				args.Flags[flagKey] = ""
			}
		} else if !strings.HasPrefix(os.Args[i-1], "-") {
			args.Args = append(args.Args, os.Args[i])
		}
	}
	return args
}
