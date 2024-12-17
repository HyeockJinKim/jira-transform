package destination

import (
	"context"

	"github.com/HyeockJinKim/jira-transform/data"
)

type Destination interface {
	Mirror(ctx context.Context, issue data.Issue) error
}
