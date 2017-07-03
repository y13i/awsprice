package awsprice

import (
	"encoding/json"
	"errors"
	"fmt"

	"strings"

	"github.com/k0kubun/pp"
)

func print(v ...interface{}) (int, error) {
	var vv interface{}

	if len(v) > 1 {
		vv = v
	} else {
		vv = v[0]
	}

	switch outputFormat {
	case "json":
		bytes, e := json.Marshal(vv)

		if e != nil {
			return 0, e
		}

		return fmt.Println(string(bytes))
	case "pp":
		return pp.Println(vv)
	}

	return 0, errors.New("unknown format")
}

func getOfferIndex() (OfferIndex, error) {
	var o OfferIndex

	e := backend.getData(OfferIndexPath, &o)
	return o, e
}

func getOffer(offerCode OfferCode) (Offer, error) {
	var o Offer
	oi, e := getOfferIndex()

	if e != nil {
		return o, e
	}

	o, ok := oi.Offers[offerCode]

	if !ok {
		e = errors.New("no such offer")
	}

	return o, e
}

func getOfferRegionIndex(offerCode OfferCode) (OfferRegionIndex, error) {
	var ori OfferRegionIndex

	o, e := getOffer(offerCode)

	if e != nil {
		return ori, e
	}

	if o.CurrentRegionIndexURL == "" {
		return ori, errors.New("no region index")
	}

	backend.getData(o.CurrentRegionIndexURL, &ori)

	return ori, e
}

func getOfferVersionIndex(offerCode OfferCode) (OfferVersionIndex, error) {
	var ovi OfferVersionIndex

	o, e := getOffer(offerCode)

	if e == nil {
		backend.getData(o.VersionIndexURL, &ovi)
	}

	return ovi, e
}

func getOfferVersion(offerCode OfferCode, regionCode RegionCode, version Version) (OfferVersion, error) {
	var ov OfferVersion

	o, e := getOffer(offerCode)

	if e != nil {
		return ov, e
	}

	if version != Version("") {
		ovi, e := getOfferVersionIndex(offerCode)

		if e != nil {
			return ov, e
		}

		offerVersionPath := ovi.Versions[version].OfferVersionURL

		if offerVersionPath == APIPath("") {
			return ov, errors.New("no such offer version")
		}

		backend.getData(offerVersionPath, &ov)
	} else if regionCode != RegionCode("") {
		ori, e := getOfferRegionIndex(offerCode)

		if e != nil {
			return ov, e
		}

		regionCurrentVersionPath := ori.Regions[regionCode].CurrentVersionURL

		if regionCurrentVersionPath == APIPath("") {
			return ov, errors.New("no such region")
		}

		backend.getData(regionCurrentVersionPath, &ov)

		if e != nil {
			return ov, e
		}
	} else {
		backend.getData(o.CurrentVersionURL, &ov)
	}

	return ov, e
}

func getProductSlice(offerCode OfferCode, regionCode RegionCode, version Version, productFamilies []ProductFamily, attributes Attributes) ([]Product, error) {
	ps := []Product{}

	ov, e := getOfferVersion(
		offerCode,
		regionCode,
		version,
	)

	if e != nil {
		return ps, e
	}

	pfm := map[ProductFamily]bool{}

	for _, pf := range productFamilies {
		pfm[pf] = true
	}

EachProduct:
	for _, p := range ov.Products {
		if _, ok := pfm[p.ProductFamily]; len(productFamilies) > 0 && !ok {
			continue EachProduct
		}

		if len(attributes) > 0 {
			for ak, av := range attributes {
				pav, ok := p.Attributes[ak]

				if !ok {
					continue EachProduct
				}

				if av != "" && av != pav {
					continue EachProduct
				}
			}
		}

		ps = append(ps, p)
	}

	return ps, e
}

func stringsToProductFamilies(ss []string) []ProductFamily {
	pfs := []ProductFamily{}

	for _, s := range ss {
		pfs = append(pfs, ProductFamily(s))
	}

	return pfs
}

func stringsToAttributes(ss []string) Attributes {
	as := Attributes{}

	for _, s := range ss {
		ps := strings.Split(s, "=")

		k := ps[0]
		v := ps[1]

		as[AttributeKey(k)] = AttributeValue(v)
	}

	return as
}
