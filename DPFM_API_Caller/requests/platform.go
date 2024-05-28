package requests

type Platform struct {
	Station				int		`json:"Station"`
	PlatformNumber		string	`json:"PlatformNumber"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	Description			string	`json:"Description"`
	LongText			string	`json:"LongText"`
	OperationRemarks	*string	`json:"OperationRemarks"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	CreateUser			int		`json:"CreateUser"`
	LastChangeUser		int		`json:"LastChangeUser"`
	IsReleased			*bool	`json:"IsReleased"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
