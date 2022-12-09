package instructions

import (
	"github.com/pkg/errors"
)

var nocacheKey = "dockerfile/run/nocache"

func init() {
	parseRunPreHooks = append(parseRunPreHooks, runNoCachePreHook)
	parseRunPostHooks = append(parseRunPostHooks, runNoCachePostHook)
}

func runNoCachePreHook(cmd *RunCommand, req parseRequest) error {
	st := &noCacheState{}
	st.flag = req.flags.AddBool("no-cache", false)
	cmd.setExternalValue(nocacheKey, st)
	return nil
}

func runNoCachePostHook(cmd *RunCommand, req parseRequest) error {
	st := cmd.getExternalValue(nocacheKey).(*noCacheState)
	if st == nil {
		return errors.Errorf("no no-cache state")
	}

	value := st.flag.Value
	if value != "true" && value != "false" {
		return errors.Errorf("no-cache %q is not valid", value)
	}

	return nil
}

func GetNoCache(cmd *RunCommand) bool {
	return cmd.getExternalValue(nocacheKey).(*noCacheState).flag.Value == "true"
}

type noCacheState struct {
	flag     *Flag
}
