package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-station-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
			Station:							data.Station,
			StationType:						data.StationType,
			StationOwner:						data.StationOwner,
			StationOwnerBusinessPartnerRole:	data.StationOwnerBusinessPartnerRole,
			PrimaryLine:						data.PrimaryLine,
			PersonResponsible:					data.PersonResponsible,
			URL:								data.URL,
			ValidityStartDate:					data.ValidityStartDate,
			ValidityEndDate:					data.ValidityEndDate,
			DailyOperationStartTime:			data.DailyOperationStartTime,
			DailyOperationEndTime:				data.DailyOperationEndTime,
			Description:						data.Description,
			LongText:							data.LongText,
			Introduction:						data.Introduction,
			StationSymbol:						data.StationSymbol,
			OperationRemarks:					data.OperationRemarks,
			PhoneNumber:						data.PhoneNumber,
			AvailabilityOfParking:				data.AvailabilityOfParking,
			NumberOfParkingSpaces:				data.NumberOfParkingSpaces,
			Site:								data.Site,
			SiteType:							data.SiteType,
			Airport:							data.Airport,
			Project:							data.Project,
			WBSElement:							data.WBSElement,
			Tag1:								data.Tag1,
			Tag2:								data.Tag2,
			Tag3:								data.Tag3,
			Tag4:								data.Tag4,
			LastChangeDate:						data.LastChangeDate,
			LastChangeTime:						data.LastChangeTime,
			LastChangeUser:						data.LastChangeUser,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Station:                 dataHeader.Station,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		Station:       	dataHeader.Station,
		AddressID:   	data.AddressID,
		PostalCode:  	data.PostalCode,
		LocalSubRegion: data.LocalSubRegion,
		LocalRegion: 	data.LocalRegion,
		Country:     	data.Country,
		GlobalRegion: 	data.GlobalRegion,
		TimeZone:	 	data.TimeZone,
		District:    	data.District,
		StreetName:  	data.StreetName,
		CityName:    	data.CityName,
		Building:    	data.Building,
		Floor:       	data.Floor,
		Room:        	data.Room,
		XCoordinate: 	data.XCoordinate,
		YCoordinate: 	data.YCoordinate,
		ZCoordinate: 	data.ZCoordinate,
		Site:			data.Site,
	}
}

func ConvertToPlatformUpdates(header dpfm_api_input_reader.Header, platform dpfm_api_input_reader.Platform) *PlatformUpdates {
	dataHeader := header
	data := platform

	return &PlatformUpdates{
		Station:       			dataHeader.Station,
		Platform:				data.Platform,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		Description:			data.Description,
		LongText:				data.LongText,
		OperationRemarks:		data.OperationRemarks,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}

func ConvertToRailwayLineUpdates(header dpfm_api_input_reader.Header, railwayLine dpfm_api_input_reader.RailwayLine) *RailwayLineUpdates {
	dataHeader := header
	data := railwayLine

	return &RailwayLineUpdates{
		Station:       			dataHeader.Station,
		RailwayLine:			data.RailwayLine,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}
