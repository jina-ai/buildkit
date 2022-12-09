package dockerfile2llb

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunNoCache(c *instructions.RunCommand) (bool, error) {
	noCache := instructions.GetNoCache(c)

	return noCache, nil
}
