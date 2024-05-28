package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-station-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-station-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-station-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Station:                          	*input.Header.Station,
		StationType:						input.Header.StationType,
		StationOwner:						input.Header.StationOwner,
		StationOwnerBusinessPartnerRole:	input.Header.StationOwnerBusinessPartnerRole,
		PrimaryLine:						input.Header.PrimaryLine,
		PersonResponsible:					input.Header.PersonResponsible,
		URL:								input.Header.URL,
		ValidityStartDate:					input.Header.ValidityStartDate,
		ValidityEndDate:					input.Header.ValidityEndDate,
		DailyOperationStartTime:			input.Header.DailyOperationStartTime,
		DailyOperationEndTime:				input.Header.DailyOperationEndTime,
		Description:						input.Header.Description,
		LongText:							input.Header.LongText,
		Introduction:						input.Header.Introduction,
		StationSymbol:						input.Header.StationSymbol,
		OperationRemarks:					input.Header.OperationRemarks,
		PhoneNumber:						input.Header.PhoneNumber,
		AvailabilityOfParking:				input.Header.AvailabilityOfParking,
		NumberOfParkingSpaces:				input.Header.NumberOfParkingSpaces,
		Site:								input.Header.Site,
		SiteType:							input.Header.SiteType,
		Airport:							input.Header.Airport,
		Project:							input.Header.Project,
		WBSElement:							input.Header.WBSElement,
		Tag1:								input.Header.Tag1,
		Tag2:								input.Header.Tag2,
		Tag3:								input.Header.Tag3,
		Tag4:								input.Header.Tag4,
		CreationDate:						input.Header.CreationDate,
		CreationTime:						input.Header.CreationTime,
		LastChangeDate:						input.Header.LastChangeDate,
		LastChangeTime:						input.Header.LastChangeTime,
		CreateUser:							input.Header.CreateUser,
		LastChangeUser:						input.Header.LastChangeUser,
		IsReleased:							input.Header.IsReleased,
		IsMarkedForDeletion:				input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToAddress(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var addresses []sub_func_complementer.Address

	addresses = append(
		addresses,
		sub_func_complementer.Address{
			Station:        *input.Header.Station,
			AddressID:      input.Header.Address[0].AddressID,
			PostalCode:     input.Header.Address[0].PostalCode,
			LocalSubRegion: input.Header.Address[0].LocalSubRegion,
			LocalRegion:    input.Header.Address[0].LocalRegion,
			Country:        input.Header.Address[0].Country,
			GlobalRegion:   input.Header.Address[0].GlobalRegion,
			TimeZone:       input.Header.Address[0].TimeZone,
			District:       input.Header.Address[0].District,
			StreetName:     input.Header.Address[0].StreetName,
			CityName:       input.Header.Address[0].CityName,
			Building:       input.Header.Address[0].Building,
			Floor:          input.Header.Address[0].Floor,
			Room:           input.Header.Address[0].Room,
			XCoordinate:    input.Header.Address[0].XCoordinate,
			YCoordinate:    input.Header.Address[0].YCoordinate,
			ZCoordinate:    input.Header.Address[0].ZCoordinate,
			Site:			input.Header.Address[0].Site,
		},
	)

	subfuncSDC.Message.Address = &addresses

	return subfuncSDC
}

func ConvertToPartner(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var partners []sub_func_complementer.Partner

	partners = append(
		partners,
		sub_func_complementer.Partner{
			Station:          		 *input.Header.Station,
			PartnerFunction:         input.Header.Partner[0].PartnerFunction,
			BusinessPartner:         input.Header.Partner[0].BusinessPartner,
			BusinessPartnerFullName: input.Header.Partner[0].BusinessPartnerFullName,
			BusinessPartnerName:     input.Header.Partner[0].BusinessPartnerName,
			Organization:            input.Header.Partner[0].Organization,
			Country:                 input.Header.Partner[0].Country,
			Language:                input.Header.Partner[0].Language,
			Currency:                input.Header.Partner[0].Currency,
			ExternalDocumentID:      input.Header.Partner[0].ExternalDocumentID,
			AddressID:               input.Header.Partner[0].AddressID,
			EmailAddress:            input.Header.Partner[0].EmailAddress,
		},
	)

	subfuncSDC.Message.Partner = &partners

	return subfuncSDC
}

func ConvertToPlatform(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var platforms []sub_func_complementer.Platform

	platforms = append(
		platforms,
		sub_func_complementer.Platform{
			Station:          		 	*input.Header.Station,
			Platform:					input.Header.Platform[0].Platform,
			ValidityStartDate:			input.Header.Platform[0].ValidityStartDate,
			ValidityEndDate:			input.Header.Platform[0].ValidityEndDate,
			Description:				input.Header.Platform[0].Description,
			LongText:					input.Header.Platform[0].LongText,
			OperationRemarks:			input.Header.Platform[0].OperationRemarks,
			CreationDate:				input.Header.Platform[0].CreationDate,
			CreationTime:				input.Header.Platform[0].CreationTime,
			LastChangeDate:				input.Header.Platform[0].LastChangeDate,
			LastChangeTime:				input.Header.Platform[0].LastChangeTime,
			CreateUser:					input.Header.Platform[0].CreateUser,
			LastChangeUser:				input.Header.Platform[0].LastChangeUser,
			IsReleased:					input.Header.Platform[0].IsReleased,
			IsMarkedForDeletion:		input.Header.Platform[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Platform = &platforms

	return subfuncSDC
}

func ConvertToRailwayLine(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var railwayLines []sub_func_complementer.RailwayLine

	railwayLines = append(
		railwayLines,
		sub_func_complementer.RailwayLine{
			Station:          		 	*input.Header.Station,
			RailwayLine:				input.Header.RailwayLine[0].RailwayLine,
			ValidityStartDate:			input.Header.RailwayLine[0].ValidityStartDate,
			ValidityEndDate:			input.Header.RailwayLine[0].ValidityEndDate,
			CreationDate:				input.Header.RailwayLine[0].CreationDate,
			CreationTime:				input.Header.RailwayLine[0].CreationTime,
			LastChangeDate:				input.Header.RailwayLine[0].LastChangeDate,
			LastChangeTime:				input.Header.RailwayLine[0].LastChangeTime,
			CreateUser:					input.Header.RailwayLine[0].CreateUser,
			LastChangeUser:				input.Header.RailwayLine[0].LastChangeUser,
			IsReleased:					input.Header.RailwayLine[0].IsReleased,
			IsMarkedForDeletion:		input.Header.RailwayLine[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.RailwayLine = &railwayLines

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
