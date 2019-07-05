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
	Name       string      `json:"name"`
	Role       string      `json:"role"`
	Bio        string      `json:"bio"`
	EthAddress string      `json:"ethAddress"`
	SocialUrls []SocialURL `json:"socialUrls"`
	AvatarURL  string      `json:"avatarUrl"`
	Signature  string      `json:"signature"`
}

// SocialURL contains social urls
type SocialURL struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
}

// Signature contains the charter signature
type Signature struct {
	Signer    string `json:"signer"`
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

// MissionFields contains fields for the newsroom's mission
type MissionFields struct {
	Purpose       string `json:"purpose"`
	Structure     string `json:"structure"`
	Revenue       string `json:"revenue"`
	Encumbrances  string `json:"encumbrances"`
	Miscellaneous string `json:"miscellaneous"`
}
