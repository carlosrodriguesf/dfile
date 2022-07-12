package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
)

func (a *appImpl) List(ctx context.Context) []string {
	return ctx.DBFile().GetPathKeys()
}
