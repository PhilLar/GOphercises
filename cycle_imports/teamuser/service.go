package teamuser

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/lokalise/go-lokalise-api/handlers"
	"github.com/lokalise/go-lokalise-api/pagination"
)

const (
	pathTeams = "teams"
)

type Client interface {
	Get(ctx context.Context, path string, res interface{}) (*resty.Response, error)
	GetList(ctx context.Context, path string, res interface{}, options pagination.OptionsApplier) (*resty.Response, error)
}

type Service struct {
	Client Client
}

func pathTeamUsers(teamID int64) string {
	return fmt.Sprintf("%s/%d/users", pathTeams, teamID)
}

func (c *Service) List(ctx context.Context, teamID int64, pageOptions pagination.PageOptions) (ResponseMultiple, error) {
	var res ResponseMultiple
	resp, err := c.Client.GetList(ctx, pathTeamUsers(teamID), &res, &pageOptions)
	if err != nil {
		return res, err
	}
	handlers.ApplyPaged(resp, &res.Paged)
	return res, handlers.ApiError(resp)
}

func (c *Service) Retrieve(ctx context.Context, teamID, userID int64) (Response, error) {
	var res Response
	resp, err := c.Client.Get(ctx, fmt.Sprintf("%s/%d", pathTeamUsers(teamID), userID), &res)
	if err != nil {
		return Response{}, err
	}
	return res, handlers.ApiError(resp)
}
