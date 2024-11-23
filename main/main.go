package main

import (
	"fmt"

	"github.com/fandreuz/godaft"
)

func main() {
	payload := make(map[string]any)
	godaft.SetFetchCount(payload, 10)
	godaft.AddPropertyType(payload, godaft.PT_APARTMENT)
	godaft.AddPropertyType(payload, godaft.PT_STUDIO_APARTMENT)
	godaft.SetSearchType(payload, godaft.ST_RESIDENTIAL_RENT)
	godaft.SetAddedSinceRange(payload, godaft.AS_LAST_3_DAYS)
	godaft.SetPriceRange(payload, 1000, 1300)

	parsed, err := godaft.DoRequest(payload)
	if err != nil {
		panic(err)
	}
	fmt.Print(parsed.GetArray("listings"))
}
