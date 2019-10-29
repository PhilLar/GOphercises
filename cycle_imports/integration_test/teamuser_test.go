package integration_test

import (
	"context"
	"testing"

	"github.com/lokalise/go-lokalise-api"
	"github.com/lokalise/go-lokalise-api/pagination"
)

const (
	token = "c4fab2a40efb0256a57f92f05fd975f6eb6e0866"
)

func TestGetTeamUser(t *testing.T) {

	client, err := lokalise.NewClient(token)
	if err != nil {
		t.Fatalf("client instantiation: %v", err)
	}
	resp, err := client.TeamUsers.Retrieve(context.Background(), 170090, 5715)
	if err != nil {
		t.Fatalf("request err: %v", err)
	}
	t.Logf("team id %d", resp.TeamID)
	t.Logf("user email %s", resp.TeamUser.Email)
}

func TestGetTeamUsers(t *testing.T) {

	client, err := lokalise.NewClient(token)
	if err != nil {
		t.Fatalf("client instantiation: %v", err)
	}

	resp, err := client.TeamUsers.List(context.Background(), 193277, pagination.PageOptions{
		Limit: 10,
		Page:  1,
	})

	if err != nil {
		t.Fatalf("request err: %v", err)
	}
	t.Logf("team id %d", resp.TeamID)
	t.Logf("users %v", resp.TeamUsers)
	t.Logf("paged %+v", resp.Paged)
}
