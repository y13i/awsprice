package awsprice

import "github.com/urfave/cli"

func listTermTypes() cli.Command {
	return cli.Command{
		Name:  "listTermTypes",
		Usage: "list term types",

		Flags: []cli.Flag{
			flags["offerCode"],
			flags["region"],
			flags["offerVersion"],
		},

		Action: func(c *cli.Context) error {
			ov, e := getOfferVersion(
				OfferCode(c.String("offerCode")),
				RegionCode(c.String("region")),
				Version(c.String("offerVersion")),
			)

			if e != nil {
				return e
			}

			tts := []TermType{}

			for k := range ov.Terms {
				tts = append(tts, k)
			}

			print(tts)

			return nil
		},
	}
}
