package main

import (
	"fmt"
	"strings"

	"github.com/fandreuz/godaft"
)

func main() {
	payload := godaft.NewPayload()
	payload.SetFetchCount(10)
	payload.AddPropertyType(godaft.PT_APARTMENT)
	payload.AddPropertyType(godaft.PT_STUDIO_APARTMENT)
	payload.SetSearchType(godaft.ST_RESIDENTIAL_RENT)
	payload.SetAddedSinceRange(godaft.AS_LAST_3_DAYS)
	payload.SetPriceRange(1000, 1300)

	for _, location := range godaft.ALL_LOCATIONS {
		if strings.Contains(location.GetDisplayName(), "Dublin") {
			payload.AddGeoFilter(location.GetId(), godaft.DISTANCE_KM5)
		}
	}

	parsed, err := godaft.DoRequest(payload)
	if err != nil {
		panic(err)
	}
	fmt.Print(parsed.GetArray("listings")[0])
}
