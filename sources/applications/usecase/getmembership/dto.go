package getmembership

type MembershipOutput struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}
