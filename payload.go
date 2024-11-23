package godaft

import (
	"math"
	"strconv"
)

type PayloadType struct {
	Filters   []filterType      `json:"filters"`
	Ranges    []rangeType       `json:"ranges"`
	GeoFilter geoFilterType     `json:"geoFilter"`
	Paging    pagingPayloadType `json:"paging"`
	Section   SearchType        `json:"section"`
	Sort      SortCriteria      `json:"sort"`
}

type pagingPayloadType struct {
	From     int `json:"from"`
	Pagesize int `json:"pagesize"`
}

type filterType struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

type rangeType struct {
	Name string `json:"name"`
	// Might be empty, thus shall be represented as a string
	From string `json:"from"`
	To   string `json:"to"`
}

type geoFilterType struct {
	StoredShapeIds []string `json:"storedShapeIds"`
	GeoSearchType  string   `json:"geoSearchType"`
}

func newGeoFilter() geoFilterType {
	return geoFilterType{GeoSearchType: "STORED_SHAPES", StoredShapeIds: []string{}}
}

func newPagingPayload() pagingPayloadType {
	return pagingPayloadType{From: 0, Pagesize: math.MaxInt}
}

func NewPayload() PayloadType {
	return PayloadType{Filters: []filterType{}, Ranges: []rangeType{}, GeoFilter: newGeoFilter(), Paging: newPagingPayload()}
}

func (payload *PayloadType) SetFetchCount(count int) {
	payload.Paging.Pagesize = count
}

func (payload *PayloadType) SetFetchStart(start int) {
	payload.Paging.From = start
}

func (payload *PayloadType) addFilter(filterName string, value string) {
	for idx, filter := range payload.Filters {
		if filter.Name == filterName {
			payload.Filters[idx].Value = append(filter.Value, value)
			return
		}
	}
	payload.Filters = append(payload.Filters, filterType{Name: filterName, Value: []string{value}})
}

func (payload *PayloadType) AddPropertyType(propertyType PropertyType) {
	payload.addFilter("propertyType", string(propertyType))
}

func (payload *PayloadType) AddSuitableGender(gender Gender) {
	payload.addFilter("suitableFor", string(gender))
}

func (payload *PayloadType) addRange(rangeName string, min string, max string) {
	for idx, rangeItem := range payload.Ranges {
		if rangeItem.Name == rangeName {
			payload.Ranges[idx].From = min
			payload.Ranges[idx].To = max
			return
		}
	}
	payload.Ranges = append(payload.Ranges, rangeType{Name: rangeName, From: min, To: max})
}

func (payload *PayloadType) SetPriceRange(min int, max int) {
	var rangeName string
	switch payload.Section {
	case ST_RESIDENTIAL_RENT, ST_COMMERCIAL_RENT, ST_SHARING, ST_STUDENT_ACCOMMODATION:
		rangeName = "rentalPrice"
	default:
		rangeName = "salePrice"
	}
	payload.addRange(rangeName, strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetBedsRange(min int, max int) {
	payload.addRange("numBeds", strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetBathsRange(min int, max int) {
	payload.addRange("numBaths", strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetTenantsRange(min int, max int) {
	payload.addRange("numTenants", strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetLeaseRange(min int, max int) {
	payload.addRange("leaseLength", strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetFloorSizeRange(min int, max int) {
	payload.addRange("floorSize", strconv.Itoa(min), strconv.Itoa(max))
}

func (payload *PayloadType) SetAddedSinceRange(addedSince AddedSince) {
	payload.addRange("firstPublishDate", string(addedSince), "")
}

func (payload *PayloadType) SetBerRange(min Ber, max Ber) {
	payload.addRange("ber", strconv.Itoa(int(min)), strconv.Itoa(int(max)))
}

func (payload *PayloadType) AddGeoFilter(location Location, distance Distance) {
	idWithDistance := strconv.Itoa(location.GetId()) + string(distance)
	payload.GeoFilter.StoredShapeIds = append(payload.GeoFilter.StoredShapeIds, idWithDistance)
}

func (payload *PayloadType) SetSortCriteria(criteria SortCriteria) {
	payload.Sort = criteria
}

func (payload *PayloadType) SetSearchType(searchType SearchType) {
	payload.Section = searchType
}
