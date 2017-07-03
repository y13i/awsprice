package awsprice

import "github.com/urfave/cli"

// ProductTerm conbines Product and OfferTerms
type ProductTerm struct {
	Product    Product
	OfferTerms []OfferTerm
}

func listProductTerms() cli.Command {
	return cli.Command{
		Name:  "listProductTerms",
		Usage: "list offer terms of products match given filters",

		Flags: []cli.Flag{
			flags["offerCode"],
			flags["region"],
			flags["offerVersion"],
			flags["productFamily"],
			flags["attribute"],
			flags["termType"],
			flags["priceUnit"],
			flags["termAttribute"],
		},

		Action: func(c *cli.Context) error {
			as, e := stringsToAttributes(c.StringSlice("attribute"))
			tam, e := stringsToTermAttributes(c.StringSlice("termAttribute"))

			if e != nil {
				return e
			}

			ps, e := getProductSlice(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
				stringsToProductFamilies(c.StringSlice("productFamily")),
				as,
			)

			if e != nil {
				return e
			}

			ov, e := getOfferVersion(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
			)

			if e != nil {
				return e
			}

			ptm := ov.Terms[TermType(c.String("termType"))]

			pts := []ProductTerm{}

			for _, p := range ps {
				ots := filterOfferTerms(ptm[p.SKU], tam)

				if len(ots) == 0 {
					continue
				}

				productPrice := ProductTerm{
					Product:    p,
					OfferTerms: ots,
				}

				pts = append(pts, productPrice)
			}

			print(pts)

			return nil
		},
	}
}
