package ad_listing

type AdsResponse struct {
	Variant int  `json:"variant"`
	Total   int  `json:"total"`
	Ads     []Ad `json:"ads"`
}

type Ad struct {
	AdId                  int      `json:"ad_id"`
	ListID                int      `json:"list_id"`
	ListTime              int      `json:"list_time"`
	Date                  string   `json:"date"`
	AccountID             int      `json:"account_id"`
	AccountOID            string   `json:"account_oid"`
	AccountName           string   `json:"account_name"`
	Subject               string   `json:"subject"`
	Body                  string   `json:"body"`
	Category              int      `json:"category"`
	CategoryName          string   `json:"category_name"`
	Area                  int      `json:"area"`
	AreaName              string   `json:"area_name"`
	Region                int      `json:"region"`
	RegionName            string   `json:"region_name"`
	CompanyAD             bool     `json:"company_ad"`
	Type                  string   `json:"type"`
	PriceString           string   `json:"price_string"`
	Image                 string   `json:"image"`
	WebpImage             string   `json:"webp_image"`
	Videos                []string `json:"videos"`
	NumberOfImages        int      `json:"number_of_images"`
	Avatar                string   `json:"avatar"`
	Rooms                 int      `json:"rooms"`
	RegionV2              int      `json:"region_v2"`
	AreaV2                int      `json:"area_v2"`
	Ward                  int      `json:"ward"`
	WardName              string   `json:"ward_name"`
	Direction             int      `json:"direction"`
	PriceMillionPerM2     float32  `json:"price_million_per_m2"`
	HouseType             int      `json:"house_type"`
	Location              string   `json:"location"`
	Longitude             float32  `json:"longitude"`
	Latitude              float32  `json:"latitude"`
	PhoneHidden           bool     `json:"phone_hidden"`
	Owner                 bool     `json:"owner"`
	ProtectionEntitlement bool     `json:"protection_entitlement"`
	EscrowCanDeposit      int      `json:"escrow_can_deposit"`
	ZeroDeposit           bool     `json:"zero_deposit"`
	PtyJupiter            string   `json:"pty_jupiter"`
	LabelCampaigns        *string  `json:"label_campaigns"`
	AdLabels              *string  `json:"ad_labels"`
}
