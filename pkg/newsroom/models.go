package newsroom

// Charter is a newsroom charter
type Charter struct {
	Name        string        `json:"name"`
	Tagline     string        `json:"tagline"`
	LogoURL     string        `json:"logoUrl"`
	NewsroomURL string        `json:"newsroomUrl"`
	Roster      []RosterItem  `json:"roster"`
	Signatures  []Signature   `json:"signatures"`
	Mission     MissionFields `json:"mission"`
	SocialUrls  SocialURL     `json:"socialUrls"`
}

// RosterItem contains a member of the newsroom roster
type RosterItem struct {
	Name       string    `json:"name"`
	Role       string    `json:"role"`
	Bio        string    `json:"bio"`
	EthAddress string    `json:"ethAddress,omitempty"`
	SocialUrls SocialURL `json:"socialUrls"`
	AvatarURL  string    `json:"avatarUrl,omitempty"`
	Signature  string    `json:"signature"`
}

// SocialURL contains social urls
type SocialURL struct {
	Twitter  string `json:"twitter,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Signature contains the charter signature
type Signature struct {
	Signer    string `json:"signer"`
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

// MissionFields contains fields for the newsroom's mission
type MissionFields struct {
	Purpose       string `json:"purpose,omitempty"`
	Structure     string `json:"structure,omitempty"`
	Revenue       string `json:"revenue,omitempty"`
	Encumbrances  string `json:"encumbrances,omitempty"`
	Miscellaneous string `json:"miscellaneous,omitempty"`
}
