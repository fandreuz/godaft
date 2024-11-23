package godaft

import (
	"math"
	"strconv"
)

type payloadType map[string]any
type pagingPayloadType map[string]int

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

func initPagingPayload() pagingPayloadType {
	tmp := make(pagingPayloadType)
	tmp["from"] = 0
	tmp["pagesize"] = math.MaxInt
	return tmp
}

func SetFetchCount(payload payloadType, count int) {
	pagingPayload, ok := payload["paging"]
	if !ok {
		tmp := initPagingPayload()
		payload["paging"] = tmp
		pagingPayload = tmp
	}

	pagingPayload.(pagingPayloadType)["pagesize"] = count
}

func SetFetchStart(payload payloadType, start int) {
	pagingPayload, ok := payload["paging"]
	if !ok {
		tmp := initPagingPayload()
		payload["paging"] = tmp
		pagingPayload = tmp
	}

	pagingPayload.(pagingPayloadType)["from"] = start
}

func ensureArrayField[T any](payload payloadType, fieldName string) []T {
	_, ok := payload[fieldName]
	if !ok {
		payload[fieldName] = make([]T, 0)
	}
	return payload[fieldName].([]T)
}

func ensureFilters(payload payloadType) []filterType {
	return ensureArrayField[filterType](payload, "filters")
}

func addFilter(payload payloadType, filterName string, value string) {
	filterPayloads := ensureFilters(payload)
	for idx, filter := range filterPayloads {
		if filter.Name == filterName {
			filterPayloads[idx].Value = append(filter.Value, value)
			return
		}
	}
	payload["filters"] = append(filterPayloads, filterType{Name: filterName, Value: []string{value}})
}

func AddPropertyType(payload payloadType, propertyType PropertyType) {
	addFilter(payload, "propertyType", string(propertyType))
}

func SetSearchType(payload payloadType, searchType SearchType) {
	payload["section"] = searchType
}

func AddSuitableGender(payload payloadType, gender Gender) {
	addFilter(payload, "suitableFor", string(gender))
}

func ensureRanges(payload payloadType) []rangeType {
	return ensureArrayField[rangeType](payload, "ranges")
}

func addRange(payload payloadType, rangeName string, min string, max string) {
	rangePayloads := ensureRanges(payload)
	for idx, rangeItem := range rangePayloads {
		if rangeItem.Name == rangeName {
			rangePayloads[idx].From = min
			rangePayloads[idx].To = max
			return
		}
	}
	payload["ranges"] = append(rangePayloads, rangeType{Name: rangeName, From: min, To: max})
}

func SetPriceRange(payload payloadType, min int, max int) {
	var rangeName string
	switch payload["section"] {
	case ST_RESIDENTIAL_RENT, ST_COMMERCIAL_RENT, ST_SHARING, ST_STUDENT_ACCOMMODATION:
		rangeName = "rentalPrice"
	default:
		rangeName = "salePrice"
	}
	addRange(payload, rangeName, strconv.Itoa(min), strconv.Itoa(max))
}

func SetBedsRange(payload payloadType, min int, max int) {
	addRange(payload, "numBeds", strconv.Itoa(min), strconv.Itoa(max))
}

func SetBathsRange(payload payloadType, min int, max int) {
	addRange(payload, "numBaths", strconv.Itoa(min), strconv.Itoa(max))
}

func SetTenantsRange(payload payloadType, min int, max int) {
	addRange(payload, "numTenants", strconv.Itoa(min), strconv.Itoa(max))
}

func SetLeaseRange(payload payloadType, min int, max int) {
	addRange(payload, "leaseLength", strconv.Itoa(min), strconv.Itoa(max))
}

func SetFloorSizeRange(payload payloadType, min int, max int) {
	addRange(payload, "floorSize", strconv.Itoa(min), strconv.Itoa(max))
}

func SetAddedSinceRange(payload payloadType, addedSince AddedSince) {
	addRange(payload, "firstPublishDate", string(addedSince), "")
}

func SetBerRange(payload payloadType, min Ber, max Ber) {
	addRange(payload, "ber", strconv.Itoa(int(min)), strconv.Itoa(int(max)))
}

func SetSortCriteria(payload payloadType, criteria SortCriteria) {
	payload["sort"] = criteria
}
