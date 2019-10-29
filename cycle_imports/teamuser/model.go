package teamuser

import "github.com/lokalise/go-lokalise-api/pagination"

type TeamUser struct {
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	CreatedAt string `json:"created_at,omitempty"`
	Role      Role   `json:"role"`
}

type Response struct {
	TeamID   int64    `json:"team_id"`
	TeamUser TeamUser `json:"team_user,omitempty"`
}

type ResponseMultiple struct {
	pagination.Paged
	TeamID    int64      `json:"team_id"`
	TeamUsers []TeamUser `json:"team_users,omitempty"`
}

type Role string

//noinspection GoUnusedConst
const (
	Owner  Role = "owner"
	Admin  Role = "admin"
	Member Role = "member"
)

type DeleteResponse struct {
	TeamID  int64 `json:"team_id,omitempty"`
	Deleted bool  `json:"team_user_deleted"`
}
