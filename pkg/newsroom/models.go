package newsroom

// Newsroom represents data about a newsroom, including the charter
type Newsroom struct {
	Name    string   `json:"name,omitempty"`
	Charter *Charter `json:"charter,omitempty"`
}

// Charter represents charter data for a newsroom, mirrors structure from the
// FE client store as defined here
// https://github.com/joincivil/Civil/blob/master/packages/core/src/types.ts#L73-L87
type Charter struct {
	Name        string                          `json:"name,omitempty"`
	LogoURL     string                          `json:"logoUrl,omitempty"`
	NewsroomURL string                          `json:"newsroomUrl,omitempty"`
	Tagline     string                          `json:"tagline,omitempty"`
	Roster      []*CharterRosterMember          `json:"roster,omitempty"`
	Signatures  []*CharterConstitutionSignature `json:"signatures,omitempty"`
	Mission     *CharterMission                 `json:"mission,omitempty"`
	SocialURLs  *CharterSocialURLs              `json:"socialUrls,omitempty"`
}

// CharterMission represents mission statements for a charter
type CharterMission struct {
	Purpose       string `json:"purpose,omitempty"`
	Structure     string `json:"structure,omitempty"`
	Revenue       string `json:"revenue,omitempty"`
	Encumbrances  string `json:"encumbrances,omitempty"`
	Miscellaneous string `json:"miscellaneous,omitempty"`
}

// AsMap converts the CharterMission to a map
func (c *CharterMission) AsMap() map[string]interface{} {
	mission := map[string]interface{}{}
	mission["purpose"] = c.Purpose
	mission["structure"] = c.Structure
	mission["revenue"] = c.Revenue
	mission["encumbrances"] = c.Encumbrances
	mission["miscellaneous"] = c.Miscellaneous
	return mission
}

// CharterRosterMember represents a member of a newsroom roster
type CharterRosterMember struct {
	Name       string             `json:"name,omitempty"`
	Role       string             `json:"role,omitempty"`
	Bio        string             `json:"bio,omitempty"`
	EthAddress string             `json:"ethAddress,omitempty"`
	SocialURLs *CharterSocialURLs `json:"socialUrls,omitempty"`
	AvatarURL  string             `json:"avatarUrl,omitempty"`
	Signature  string             `json:"signature,omitempty"`
}

// AsMap converts the CharterRosterMember to a map
func (c *CharterRosterMember) AsMap() map[string]interface{} {
	member := map[string]interface{}{}
	member["name"] = c.Name
	member["role"] = c.Role
	member["bio"] = c.Bio
	member["eth_address"] = c.EthAddress
	member["avatar_url"] = c.AvatarURL
	member["signature"] = c.Signature
	member["social_urls"] = nil

	if c.SocialURLs != nil {
		member["social_urls"] = map[string]string{
			"twitter":   c.SocialURLs.Twitter,
			"facebook":  c.SocialURLs.Facebook,
			"instagram": c.SocialURLs.Instagram,
			"linkedin":  c.SocialURLs.Linkedin,
			"youtube":   c.SocialURLs.Youtube,
		}
	}

	return member
}

// CharterConstitutionSignature represents the signing of the constitution for a
// newsroom
type CharterConstitutionSignature struct {
	Signer    string `json:"signer,omitempty"`
	Signature string `json:"signature,omitempty"`
	Message   string `json:"message,omitempty"`
}

// AsMap converts the CharterConstitutionSignature to a map
func (c *CharterConstitutionSignature) AsMap() map[string]interface{} {
	member := map[string]interface{}{}
	member["signer"] = c.Signer
	member["signature"] = c.Signature
	member["message"] = c.Message
	return member
}

// CharterSocialURLs represents a social URL in the charter
type CharterSocialURLs struct {
	Twitter   string `json:"twitter,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Linkedin  string `json:"linkedin,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Email     string `json:"email,omitempty"`
}

// AsMap converts the CharterSocialURL to a map
func (c *CharterSocialURLs) AsMap() map[string]interface{} {
	social := map[string]interface{}{}
	social["twitter"] = c.Twitter
	social["facebook"] = c.Facebook
	social["instagram"] = c.Instagram
	social["linkedin"] = c.Linkedin
	social["youtube"] = c.Youtube
	social["email"] = c.Email
	return social
}
