package main

import (
	"fmt"

	"github.com/fandreuz/godaft"
)

func stdentAccomodation(payload *godaft.PayloadType) {
	payload.SetSearchType(godaft.ST_STUDENT_ACCOMMODATION)

	payload.SetAddedSinceRange(godaft.AS_LAST_MONTH)

	// See locations.go for more locations
	payload.AddGeoFilter(godaft.LOC_DUNDALK_INSTITUTE_OF_TECHNOLOGY_LOUTH, godaft.DISTANCE_KM10)

	payload.SetPriceRange(400, 1000)
}

func rentInDublin4(payload *godaft.PayloadType) {
	payload.SetSearchType(godaft.ST_RESIDENTIAL_RENT)

	payload.AddPropertyType(godaft.PT_APARTMENT)
	payload.AddPropertyType(godaft.PT_STUDIO_APARTMENT)

	// Older apartments are alreadt taken ;)
	payload.SetAddedSinceRange(godaft.AS_LAST_3_DAYS)

	// See locations.go for more locations
	payload.AddGeoFilter(godaft.LOC_DUBLIN_4_DUBLIN, godaft.DISTANCE_KM5)

	payload.SetPriceRange(1200, 2000)
}

func main() {
	payload := godaft.NewPayload()

	// Pagination: fetch 10 elements starting from element 0
	payload.SetFetchCount(10)

	rentInDublin4(&payload)

	parsed, err := godaft.DoRequest(payload)
	if err != nil {
		panic(err)
	}
	fmt.Print(parsed)
}
