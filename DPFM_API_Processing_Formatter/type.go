package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Station							int		`json:"Station"`
	StationType						string	`json:"StationType"`
	StationOwner					int		`json:"StationOwner"`
	StationOwnerBusinessPartnerRole	string	`json:"StationOwnerBusinessPartnerRole"`
	PrimaryLine						int		`json:"PrimaryLine"`
	PersonResponsible				*string	`json:"PersonResponsible"`
	URL								*string	`json:"URL"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	StationSymbol					*string	`json:"StationSymbol"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	AvailabilityOfParking			*bool	`json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			*int	`json:"NumberOfParkingSpaces"`
	Site							int		`json:"Site"`
	SiteType						string	`json:"SiteType"`
	AirPort							*int	`json:"AirPort"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	LastChangeUser					int		`json:"LastChangeUser"`
}

type PartnerUpdates struct {
	Station                 int     `json:"Station"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type AddressUpdates struct {
	Station     	int     	`json:"Station"`
	AddressID   	int     	`json:"AddressID"`
	PostalCode  	string 		`json:"PostalCode"`
	LocalSubRegion 	string 		`json:"LocalSubRegion"`
	LocalRegion 	string 		`json:"LocalRegion"`
	Country     	string 		`json:"Country"`
	GlobalRegion   	string 		`json:"GlobalRegion"`
	TimeZone   		string 		`json:"TimeZone"`
	District    	*string 	`json:"District"`
	StreetName  	*string 	`json:"StreetName"`
	CityName    	*string 	`json:"CityName"`
	Building    	*string 	`json:"Building"`
	Floor       	*int		`json:"Floor"`
	Room        	*int		`json:"Room"`
	XCoordinate 	*float32	`json:"XCoordinate"`
	YCoordinate 	*float32	`json:"YCoordinate"`
	ZCoordinate 	*float32	`json:"ZCoordinate"`
	Site			*int		`json:"Site"`
}

type PlatformUpdates struct {
	Station				int		`json:"Station"`
	PlatformNumber		string	`json:"PlatformNumber"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	Description			string	`json:"Description"`
	LongText			string	`json:"LongText"`
	OperationRemarks	*string	`json:"OperationRemarks"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	LastChangeUser		int		`json:"LastChangeUser"`
}

type RailwayLineUpdates struct {
	Station				int		`json:"Station"`
	RailwayLine			int		`json:"RailwayLine"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	LastChangeUser		int		`json:"LastChangeUser"`
}
