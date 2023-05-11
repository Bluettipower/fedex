package fedex

import "fmt"

const shipmentBasePath = "ship/v1/shipments"

type ShipmentServiceOp struct {
	client *FedExClient
}

type ShipmentService interface {
	Create(Shipment) (*ShipmentResponse, error)
}

type Shipment struct {

	// requestedShipment
	RequestedShipment *RequestedShipment `json:"requestedShipment,omitempty"`
	// labelResponseOptions  "URL_ONLY" "LABEL"
	LabelResponseOptions string `json:"labelResponseOptions,omitempty"`
	// accountNumber
	AccountNumber *AccountNumber `json:"accountNumber,omitempty"`
	// shipAction
	ShipAction string `json:"shipAction,omitempty"`
}

type RequestedShipment struct {
	// shipTimestamp
	ShipTimestamp string `json:"shipTimestamp,omitempty"`
	// totalDeclaredValue
	TotalDeclaredValue *Amount `json:"totalDeclaredValue,omitempty"`
	// shipper
	Shipper *ContactInfo `json:"shipper,omitempty"`
	// recipient
	// Recipient *ContactInfo `json:"recipient,omitempty"`
	// recipients
	Recipients []*ContactInfo `json:"recipients,omitempty"`
	// soldTo
	SoldTo *ContactInfo `json:"soldTo,omitempty"`
	// recipientLocationNumber
	RecipientLocationNumber string `json:"recipientLocationNumber,omitempty"`
	// pickupType
	PickupType string `json:"pickupType,omitempty"`
	// serviceType
	ServiceType string `json:"serviceType,omitempty"`
	// packagingType
	PackagingType string `json:"packagingType,omitempty"`
	// totalWeight
	TotalWeight float64 `json:"totalWeight,omitempty"`
	// origin
	Origin *ContactInfo `json:"origin,omitempty"`
	// shippingChargesPayment
	ShippingChargesPayment *Payment `json:"shippingChargesPayment,omitempty"`
	// labelSpecification
	LabelSpecification *LabelSpecification `json:"labelSpecification,omitempty"`
	// rateRequestType
	RateRequestType []string `json:"rateRequestType,omitempty"`
	// preferredCurrency
	PreferredCurrency string `json:"preferredCurrency,omitempty"`
	// totalPackageCount
	TotalPackageCount int32 `json:"totalPackageCount,omitempty"`
	// masterTrackingId
	MasterTrackingId *TrackingId `json:"masterTrackingId,omitempty"`
	// requestedPackageLineItems
	RequestedPackageLineItems []*RequestedPackageLineItem `json:"requestedPackageLineItems,omitempty"`
}

type RequestedPackageLineItem struct {
	// sequenceNumber
	SequenceNumber int `json:"sequenceNumber,omitempty"`
	// subPackagingType
	SubPackagingType string `json:"subPackagingType,omitempty"`
	// customerReferences
	CustomerReferences []*CustomerReference `json:"customerReferences,omitempty"`
	// declaredValue
	DeclaredValue *Amount `json:"declaredValue,omitempty"`
	// weight
	Weight *Weight `json:"weight,omitempty"`
	// dimensions
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// groupPackageCount
	GroupPackageCount int32 `json:"groupPackageCount,omitempty"`
	// itemDescriptionForClearance
	ItemDescriptionForClearance string `json:"itemDescriptionForClearance,omitempty"`
	// trackingNumber
	TrackingNumber string `json:"trackingNumber,omitempty"`
}

type Dimensions struct {
	// length
	Length int32 `json:"length,omitempty"`
	// width
	Width int32 `json:"width,omitempty"`
	// height
	Height int32 `json:"height,omitempty"`
	// units
	Units string `json:"units,omitempty"`
}

type Weight struct {
	// units
	Units string `json:"units,omitempty"`
	// value
	Value float64 `json:"value,omitempty"`
}

type CustomerReference struct {
	// customerReferenceType
	CustomerReferenceType string `json:"customerReferenceType,omitempty"`
	// value
	Value string `json:"value,omitempty"`
}

type TrackingId struct {
	// trackingIdType
	TrackingIdType string `json:"trackingIdType,omitempty"`
	// trackingNumber
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// formId
	FormId string `json:"formId,omitempty"`
	// uspsApplicationId
	UspsApplicationId string `json:"uspsApplicationId,omitempty"`
}

type ContactInfo struct {
	// address
	Address *Address `json:"address,omitempty"`
	// contact
	Contact *Contact `json:"contact,omitempty"`
	// tins
	Tins []*Tin `json:"tins,omitempty"`
	// deliveryInstructions
	DeliveryInstructions string `json:"deliveryInstructions,omitempty"`
	// accountNumber
	// AccountNumber *AccountNumber `json:"accountNumber,omitempty"`
}

type AccountNumber struct {
	// value
	Value string `json:"value,omitempty"`
}

type Contact struct {
	// personName
	PersonName string `json:"personName,omitempty"`
	// companyName
	CompanyName string `json:"companyName,omitempty"`
	// phoneNumber
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// emailAddress
	EmailAddress string `json:"emailAddress,omitempty"`
	// phoneExtension
	PhoneExtension string `json:"phoneExtension,omitempty"`
	// faxNumber
	FaxNumber string `json:"faxNumber,omitempty"`
}

type Address struct {
	// streetLines
	StreetLines []string `json:"streetLines,omitempty"`
	// city
	City string `json:"city,omitempty"`
	// stateOrProvinceCode
	StateOrProvinceCode string `json:"stateOrProvinceCode,omitempty"`
	// postalCode
	PostalCode string `json:"postalCode,omitempty"`
	// countryCode
	CountryCode string `json:"countryCode,omitempty"`
	// residential
	Residential bool `json:"residential,omitempty"`
}

type Tin struct {
	// tinType
	TinType string `json:"tinType,omitempty"`
	// number
	Number string `json:"number,omitempty"`
	// usage
	Usage string `json:"usage,omitempty"`
	// effectiveDate
	EffectiveDate string `json:"effectiveDate,omitempty"`
	// expirationDate
	ExpirationDate string `json:"expirationDate,omitempty"`
}

type Payment struct {
	// paymentType
	PaymentType string `json:"paymentType,omitempty"`
	// payor
	Payor *Payor `json:"payor,omitempty"`
}

type Payor struct {
	// responsibleParty
	ResponsibleParty *ContactInfo `json:"responsibleParty,omitempty"`
}

type Amount struct {
	// currency
	Currency string `json:"currency,omitempty"`
	// amount
	Amount string `json:"amount,omitempty"`
}

type LabelSpecification struct {
	// labelFormatType
	LabelFormatType string `json:"labelFormatType,omitempty"`
	// labelOrder
	LabelOrder string `json:"labelOrder,omitempty"`
	// labelStockType
	LabelStockType string `json:"labelStockType,omitempty"`
	// labelRotation
	LabelRotation string `json:"labelRotation,omitempty"`
	// imageType
	ImageType string `json:"imageType,omitempty"`
	// labelPrintingOrientation
	LabelPrintingOrientation string `json:"labelPrintingOrientation,omitempty"`
	// returnedDispositionDetail
	ReturnedDispositionDetail bool `json:"returnedDispositionDetail,omitempty"`
}

type ShipmentResponse struct {

	// output
	Output *Output `json:"output,omitempty"`
	// alerts
	Alerts []*Alert `json:"alerts,omitempty"`
	// errors
	Errors []*Error `json:"errors,omitempty"`
}

func (s *ShipmentResponse) Error() string {
	return fmt.Sprintf("%+v", s.Errors)
}

type Output struct {
	// transactionShipments
	TransactionShipments []*TransactionShipment `json:"transactionShipments,omitempty"`
	// alerts
	Alerts []*Alert `json:"alerts,omitempty"`
	// jobId
	JobId string `json:"jobId,omitempty"`
}

type TransactionShipment struct {
	// serviceType
	ServiceType string `json:"serviceType,omitempty"`
	// shipDatestamp
	ShipDatestamp string `json:"shipDatestamp,omitempty"`
	// serviceCategory
	ServiceCategory string `json:"serviceCategory,omitempty"`
	// serviceName
	ServiceName string `json:"serviceName,omitempty"`
	//masterTrackingNumber
	MasterTrackingNumber string `json:"masterTrackingNumber,omitempty"`
	// shipmentDocuments
	ShipmentDocuments []*ShipmentDocument `json:"shipmentDocuments,omitempty"`
	// pieceResponses
	PieceResponses []*PieceResponse `json:"pieceResponses,omitempty"`
}

type PieceResponse struct {
	// trackingNumber
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// packageSequenceNumber
	PackageSequenceNumber string `json:"packageSequenceNumber,omitempty"`
	// packageDocuments
	PackageDocuments []*PackageDocument `json:"packageDocuments,omitempty"`
}

type PackageDocument struct {
	// contentKey
	ContentKey string `json:"contentKey,omitempty"`
	// copiesToPrint
	CopiesToPrint int32 `json:"copiesToPrint,omitempty"`
	// contentType
	ContentType string `json:"contentType,omitempty"`
	// trackingNumber
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// docType
	DocType string `json:"docType,omitempty"`
	// encodedLabel
	EncodedLabel string `json:"encodedLabel,omitempty"`
	// url
	URL string `json:"url,omitempty"`
}

type ShipmentDocument struct {
	// contentKey
	ContentKey string `json:"contentKey,omitempty"`
	// copiesToPrint
	CopiesToPrint int64 `json:"copiesToPrint,omitempty"`
	// contentType
	ContentType string `json:"contentType,omitempty"`
	// trackingNumber
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// docType
	DocType string `json:"docType,omitempty"`
	// encodedLabel
	EncodedLabel string `json:"encodedLabel,omitempty"`
	// url
	URL string `json:"url,omitempty"`
	// alerts
	Alerts []*Alert `json:"alerts,omitempty"`
}

// alerts
type Alert struct {
	// code
	Code string `json:"code,omitempty"`
	// message
	Message string `json:"message,omitempty"`
	// alertType
	AlertType string `json:"alertType,omitempty"`
}

type Error struct {
	// code
	Code string `json:"code,omitempty"`
	// message
	Message string `json:"message,omitempty"`
	// parameterList
	ParameterList []*ParameterList `json:"parameterList,omitempty"`
}

type ParameterList struct {
	// key
	Key string `json:"key,omitempty"`
	// value
	Value string `json:"value,omitempty"`
}

func (s *ShipmentServiceOp) Create(shipment Shipment) (*ShipmentResponse, error) {
	resource := new(ShipmentResponse)
	err := s.client.Post(shipmentBasePath, shipment, resource)
	if resource.Errors != nil {
		return nil, resource
	}
	return resource, err
}
