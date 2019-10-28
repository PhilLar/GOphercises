package teamuser

import (
	"context"
	"fmt"
	"github.com/lokalise/go-lokalise-api/handlers"
)

const (
	pathTeams = "teams"
)

type Service struct {
	Client *Client
}

func pathTeamUsers(teamID int64) string {
	return fmt.Sprintf("%s/%d/users", pathTeams, teamID)
}

func (c *Service) List(ctx context.Context, teamID int64, pageOptions handlers.PageOptions) (ResponseMultiple, error) {
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
