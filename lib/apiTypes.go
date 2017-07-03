// Reference - http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html

package awsprice

// FormatVersion - The version of the file format
// An attribute that tracks which format version the offer file is in. The formatVersion of the file is updated when the structure is changed. For example, the version will change from v1 to v2.
// example: "v1.0"
type FormatVersion string

// Desclaimer - The disclaimers for the offer file
// Any disclaimers that apply to the offer file.
// example: "This pricing list is for informational purposes only."
type Desclaimer string

// PublicationDate - The publication date of the offer file
// The date and time (UTC) when an offer file was published. For example, 2015-04-09T02:22:05Z, 2015-09-10T18:21:05Z.
// example: "2015-12-09T10:13:42Z"
type PublicationDate string

// APIPath - example "/offers/v1.0/aws/AmazonS3/index.json"
type APIPath string

// DateString - example "2017-02-01T00:00:00Z"
type DateString string

// SKU - example "GFSDHARN4P5AGZQ7"
type SKU string

// TermType - "OnDemand" or "Reserved"
type TermType string

// OfferTermID - SKU.OfferTermCode, example "GFSDHARN4P5AGZQ7.JRTCKXETXF"
type OfferTermID string

// OfferTermCode - example "JRTCKXETXF"
type OfferTermCode string

// ProductFamily - The product family of the product
// example: "Compute Instance"
type ProductFamily string

// RateCode - SKU.OfferTermCode.XXXXXXXXXX, example "GFSDHARN4P5AGZQ7.JRTCKXETXF.6YS6EN2CT7"
type RateCode string

// Description - example "Queries to LBR Alias records are free of charge"
type Description string

// TermAttributeKey - example "LeaseContractLength"
type TermAttributeKey string

// TermAttributeValue - example "3yr"
type TermAttributeValue string

// Range - integer string or "Inf", example "0", "Inf"
type Range string

// Unit - example "Hrs"
type Unit string

// OfferCode - The code for the service
// A unique code for the product of an AWS service. For example, AmazonEC2 for Amazon EC2 or AmazonS3 for Amazon S3.
// example: "AmazonEC2"
type OfferCode string

// Version - The version of the offer file
// An attribute that tracks the version of the offer file. Each time a new file is published, it contains a new version number. For example, 20150409T022205 and 20150910T182105.
// example: "20170605233259"
type Version string

// RegionCode - example "us-east-1"
type RegionCode string

// AttributeKey - example "instanceType"
type AttributeKey string

// AttributeValue - example "t2.small"
type AttributeValue string

// PriceUnit - example "USD"
type PriceUnit string

// PriceValue - example "0.5"
type PriceValue string

// Offers - map, key: OfferCode, value: Offer
type Offers map[OfferCode]Offer

// Versions - map, key: Version, value: OfferVersionInfo
type Versions map[Version]OfferVersionInfo

// Attributes - map, key: AttributeKey, value: AttributeValue
type Attributes map[AttributeKey]AttributeValue

// Regions - map, key: RegionCode, value: OfferRegionInfo
type Regions map[RegionCode]OfferRegionInfo

// Products - map, key: SKU, value: Product
type Products map[SKU]Product

// Terms - map, key: TermType, value: ProductTermMap
type Terms map[TermType]ProductTermMap

// ProductTermMap - map, key: SKU, value: OfferTerms
type ProductTermMap map[SKU]OfferTerms

// OfferTerms - map, key: OfferTermID, value: OfferTerm
type OfferTerms map[OfferTermID]OfferTerm

// PriceDimensions - map, key: RateCode, value: PriceDimension
type PriceDimensions map[RateCode]PriceDimension

// TermAttributes - map, key: TermAttributeKey, value: TermAttributeValue
type TermAttributes map[TermAttributeKey]TermAttributeValue

// PricePerUnit - map, key: PriceUnit, value: PriceValue
type PricePerUnit map[PriceUnit]PriceValue

// AppliesTo - slice of SKU
type AppliesTo []SKU

// OfferIndex - [TODO: write description]
type OfferIndex struct {
	FormatVersion   FormatVersion
	Disclaimer      Desclaimer
	PublicationDate PublicationDate
	Offers          Offers
}

// Offer - [TODO: write description]
type Offer struct {
	OfferCode             OfferCode
	VersionIndexURL       APIPath
	CurrentVersionURL     APIPath
	CurrentRegionIndexURL APIPath
}

// OfferVersionIndex - [TODO: write description]
type OfferVersionIndex struct {
	FormatVersion  FormatVersion
	Disclaimer     Desclaimer
	OfferCode      OfferCode
	CurrentVersion Version
	Versions       Versions
}

// OfferVersionInfo - [TODO: write description]
type OfferVersionInfo struct {
	VersionEffectiveBeginDate DateString
	VersionEffectiveEndDate   DateString
	OfferVersionURL           APIPath
}

// OfferRegionIndex - [TODO: write description]
type OfferRegionIndex struct {
	FormatVersion   FormatVersion
	Disclaimer      Desclaimer
	PublicationDate PublicationDate
	Regions         Regions
}

// OfferRegionInfo - [TODO: write description]
type OfferRegionInfo struct {
	RegionCode        RegionCode
	CurrentVersionURL APIPath
}

// OfferVersion - [TODO: write description]
type OfferVersion struct {
	FormatVersion   FormatVersion
	Disclaimer      Desclaimer
	OfferCode       OfferCode
	Version         Version
	PublicationDate PublicationDate
	Products        Products
	Terms           Terms
}

// Product - [TODO: write description]
type Product struct {
	SKU           SKU
	ProductFamily ProductFamily
	Attributes    Attributes
}

// OfferTerm - [TODO: write description]
type OfferTerm struct {
	OfferTermCode   OfferTermCode
	SKU             SKU
	EffectiveDate   DateString
	PriceDimensions PriceDimensions
	TermAttributes  TermAttributes
}

// PriceDimension - [TODO: write description]
type PriceDimension struct {
	RateCode     RateCode
	Description  Description
	BeginRange   Range
	EndRange     Range
	Unit         Unit
	PricePerUnit PricePerUnit
	AppliesTo    AppliesTo
}
