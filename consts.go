package godaft

type PropertyType string

const (
	PT_HOUSE                   PropertyType = "houses"
	PT_DETACHED_HOUSE          PropertyType = "detached-houses"
	PT_SEMI_DETACHED_HOUSE     PropertyType = "semi-detached-houses"
	PT_TERRACED_HOUSE          PropertyType = "terraced-houses"
	PT_END_OF_TERRACE_HOUSE    PropertyType = "end-of-terrace-houses"
	PT_TOWNHOUSE               PropertyType = "townhouses"
	PT_DUPLEX                  PropertyType = "duplexes"
	PT_BUNGALOW                PropertyType = "bungalows"
	PT_APARTMENT               PropertyType = "apartments"
	PT_STUDIO_APARTMENT        PropertyType = "studio-apartments"
	PT_SITE                    PropertyType = "sites"
	PT_OFFICE_SPACE            PropertyType = "office-spaces"
	PT_RETAIL_UNIT             PropertyType = "retail-units"
	PT_INDUSTRIAL_UNIT         PropertyType = "industrial-units"
	PT_INDUSTRIAL_SITES        PropertyType = "industrial-sites"
	PT_RESTAURANTS_BARS_HOTELS PropertyType = "restaurants-bars-hotels"
	PT_COMMERCIAL_SITES        PropertyType = "commercial-sites"
	PT_AGRICULTURAL_LAND       PropertyType = "agricultural-land"
	PT_DEVELOPMENT_LAND        PropertyType = "development-land"
	PT_INVESTMENT_PROPERTY     PropertyType = "investment-properties"
)

type SearchType string

const (
	ST_RESIDENTIAL_SALE      SearchType = "residential-for-sale"
	ST_RESIDENTIAL_RENT      SearchType = "residential-to-rent"
	ST_COMMERCIAL_SALE       SearchType = "commercial-for-sale"
	ST_COMMERCIAL_RENT       SearchType = "commercial-for-rent"
	ST_SHARING               SearchType = "sharing"
	ST_STUDENT_ACCOMMODATION SearchType = "student-accommodation-to-share"
	ST_NEW_HOMES             SearchType = "new-homes"
)

type Gender string

const (
	GENDER_MALE   Gender = "male"
	GENDER_FEMALE Gender = "female"
)

type SortCriteria string

const (
	PUBLISH_DATE_DESC SortCriteria = "publishDateDesc"
	PRICE_ASC         SortCriteria = "priceAsc"
	PRICE_DESC        SortCriteria = "priceDesc"
)

type AddedSince string

const (
	AS_LAST_3_DAYS    = "now-3d/d"
	AS_LAST_WEEK      = "now-7d/d"
	AS_LAST_TWO_WEEKS = "now-14d/d"
	AS_LAST_MONTH     = "now-30d/d"
)

// Building Energy Rating
type Ber int

const (
	A1 Ber = 0
	A2 Ber = 1
	A3 Ber = 2
	B1 Ber = 3
	B2 Ber = 4
	B3 Ber = 5
	C1 Ber = 6
	C2 Ber = 7
	C3 Ber = 8
	D1 Ber = 9
	D2 Ber = 10
	E1 Ber = 11
	E2 Ber = 12
	F  Ber = 13
	G  Ber = 14
)

type Distance string

const (
	DISTANCE_KM0  Distance = ""
	DISTANCE_KM1  Distance = "_1000"
	DISTANCE_KM3  Distance = "_3000"
	DISTANCE_KM5  Distance = "_5000"
	DISTANCE_KM10 Distance = "_10000"
	DISTANCE_KM20 Distance = "_20000"
)
