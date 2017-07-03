# awsprice

CLI for [AWS Price List API](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html).

[日本語 README](README.ja.md)

## Installation

Download binary from [releases](https://github.com/y13i/awsprice/releases).

Put it into your `$PATH`.

### Homebrew

```
$ brew install https://raw.githubusercontent.com/y13i/awsprice/master/homebrew/awsprice.rb
```

## Usage

```
$ awsprice [global options] command [command options] [arguments...]
```

### Commands

#### `listOffers`

List all offer codes.

```
$ awsprice listOffers | head -10
[]awsprice.OfferCode{
  "AWSBudgets",
  "AWSCloudTrail",
  "AWSCodeCommit",
  "AWSCodeDeploy",
  "AWSCodePipeline",
  "AWSConfig",
  "AWSDatabaseMigrationSvc",
  "AWSDeveloperSupport",
  "AWSDeviceFarm",
```

#### `listOfferRegions`

List all regions of specified offer. (note: `--offerCode` defaults to `AmazonEC2`)

```
$ awsprice listOfferRegions | head -5
[]awsprice.RegionCode{
  "ap-northeast-1",
  "ap-northeast-2",
  "ap-south-1",
  "ap-southeast-1",
```

```
$ awsprice listOfferRegions --offerCode AmazonLex
[]awsprice.RegionCode{
  "us-east-1",
}
```

##### Command Options

###### `--offerCode value, -o value`

(default: `AmazonEC2`)

A unique code for the product of an AWS service.

#### `listOfferVersions`

List all versions of specified offer.

```
$ awsprice listOfferVersions | head -10
[]awsprice.Version{
  "20151209144527",
  "20160126001708",
  "20160628000628",
  "20160901005907",
  "20161026205455",
  "20161213014831",
  "20170210223144",
  "20170224022054",
  "20170302183221",
```

##### Command Options

###### `--offerCode value, -o value`

#### `listProductFamilies`

List product families of the offer products.

```
$ awsprice listProductFamilies
[]awsprice.ProductFamily{
  "Compute Instance",
  "Data Transfer",
  "Dedicated Host",
  "Fee",
  "IP Address",
  "Load Balancer",
  "Load Balancer-Application",
  "NAT Gateway",
  "Storage",
  "Storage Snapshot",
  "System Operation",
}
```

##### Command Options

###### `--offerCode value, -o value`

###### `--region value, -r value`

AWS region code. If nvironment variable `$AWS_REGION` is set, it will be used as this option value

###### `--offerVersion value`

(default: current version)

The version of the offer file.

Example: `20170302183221`

###### `--productFamily value, -p value`

Filter products by the name of product family.

Accept multiple values.

```
-p familyName1 -p familyName2
```

###### `--attribute value, -a value`

Filter products by the attribute in `KEY=[VALUE]` format.

Accept multiple values.

```
-a attribute1=value1 -a attribute2= -a ...
```

#### `listAttributes`

List attributes of the offer products.

```
$ awsprice listAttributes | head -10
{
  "instanceCapacity2xlarge": []awsprice.AttributeValue{
    "4",
    "5",
    "8",
  },
  "physicalCores": []awsprice.AttributeValue{
    "20",
    "24",
    "32",
```

```
$ awsprice listAttributes -o AmazonRoute53 -r "" | head -10
{
  "group": []awsprice.AttributeValue{
    "Route53-Basic",
    "Route53-Optional",
  },
  "routingType": []awsprice.AttributeValue{
    "Geo DNS",
    "Latency Based Routing",
    "Standard",
  },
```

##### Command Options

###### `--offerCode value, -o value`

###### `--region value, -r value`

###### `--offerVersion value`

###### `--productFamily value, -p value`

###### `--attribute value, -a value`

#### `listProducts`

List products match given filters.

```
$ awsprice listProducts -p "Compute Instance" -a "instanceType=t2.micro" -a "operatingSystem=Linux" -a "tenancy=Shared"
[]awsprice.Product{
  awsprice.Product{
    SKU:           "N4D3MGNKSH7Q9KT3",
    ProductFamily: "Compute Instance",
    Attributes:    {
      "preInstalledSw":        "NA",
      "location":              "US West (Oregon)",
      "locationType":          "AWS Region",
      "operatingSystem":       "Linux",
      "usagetype":             "USW2-BoxUsage:t2.micro",
      "operation":             "RunInstances",
      "ecu":                   "Variable",
      "currentGeneration":     "Yes",
      "physicalProcessor":     "Intel Xeon Family",
      "memory":                "1 GiB",
      "processorArchitecture": "32-bit or 64-bit",
      "servicecode":           "AmazonEC2",
      "instanceFamily":        "General purpose",
      "storage":               "EBS only",
      "tenancy":               "Shared",
      "licenseModel":          "No License required",
      "processorFeatures":     "Intel AVX; Intel Turbo",
      "instanceType":          "t2.micro",
      "vcpu":                  "1",
      "clockSpeed":            "Up to 3.3 GHz",
      "networkPerformance":    "Low to Moderate",
    },
  },
}
```

##### Command Options

###### `--offerCode value, -o value`

###### `--region value, -r value`

###### `--offerVersion value`

###### `--productFamily value, -p value`

###### `--attribute value, -a value`

#### `listTermTypes`

List term types.

```
$ awsprice listTermTypes
[]awsprice.TermType{
  "OnDemand",
  "Reserved",
}
```

```
$ awsprice listTermTypes -o AmazonRDS
[]awsprice.TermType{
  "OnDemand",
  "Reserved",
}
```

```
$ awsprice listTermTypes -o AmazonS3
[]awsprice.TermType{
  "OnDemand",
}
```

##### Command Options

###### `--offerCode value, -o value`

###### `--region value, -r value`

###### `--offerVersion value`

#### `listProductTerms`

List offer terms of products match given filters.

```
$ awsprice listProductTerms -a "instanceType=c4.2xlarge" -a "operatingSystem=Linux" -a "tenancy=Shared"
[]awsprice.ProductTerm{
  awsprice.ProductTerm{
    Product: awsprice.Product{
      SKU:           "YBN8Q7AQJD9ZT57S",
      ProductFamily: "Compute Instance",
      Attributes:    {
        "vcpu":                        "8",
        "storage":                     "EBS only",
        "dedicatedEbsThroughput":      "1000 Mbps",
        "location":                    "US West (Oregon)",
        "instanceType":                "c4.2xlarge",
        "instanceFamily":              "Compute optimized",
        "locationType":                "AWS Region",
        "ecu":                         "31",
        "licenseModel":                "No License required",
        "enhancedNetworkingSupported": "Yes",
        "preInstalledSw":              "NA",
        "processorFeatures":           "Intel AVX; Intel AVX2; Intel Turbo",
        "processorArchitecture":       "64-bit",
        "tenancy":                     "Shared",
        "operatingSystem":             "Linux",
        "clockSpeed":                  "2.9 GHz",
        "memory":                      "15 GiB",
        "networkPerformance":          "High",
        "usagetype":                   "USW2-BoxUsage:c4.2xlarge",
        "operation":                   "RunInstances",
        "servicecode":                 "AmazonEC2",
        "currentGeneration":           "Yes",
        "physicalProcessor":           "Intel Xeon E5-2666 v3 (Haswell)",
      },
    },
    OfferTerms: {
      "YBN8Q7AQJD9ZT57S.JRTCKXETXF": awsprice.OfferTerm{
        OfferTermCode:   "JRTCKXETXF",
        SKU:             "YBN8Q7AQJD9ZT57S",
        EffectiveDate:   "2017-06-01T00:00:00Z",
        PriceDimensions: {
          "YBN8Q7AQJD9ZT57S.JRTCKXETXF.6YS6EN2CT7": awsprice.PriceDimension{
            RateCode:     "YBN8Q7AQJD9ZT57S.JRTCKXETXF.6YS6EN2CT7",
            Description:  "$0.398 per On Demand Linux c4.2xlarge Instance Hour",
            BeginRange:   "0",
            EndRange:     "Inf",
            Unit:         "Hrs",
            PricePerUnit: {
              "USD": "0.3980000000",
            },
            AppliesTo: awsprice.AppliesTo{},
          },
        },
        TermAttributes: awsprice.TermAttributes{},
      },
    },
  },
}
```

##### Command Options

###### `--offerCode value, -o value`

###### `--region value, -r value`

###### `--offerVersion value`

###### `--productFamily value, -p value`

###### `--attribute value, -a value`


###### `--termType value`

(default: `OnDemand`)

Type of the terms. "OnDemand" or "Reserved".

###### `--priceUnit value`

(default: `USD`)

Unit of prices.

Not used. Currently.

###### `--termAttribute value, -t value`

Filter terms by the term attribute in `KEY=[VALUE]` format.

Accept multiple values.

```
-t attribute1=value1 -t attribute2= -t ...
```

#### `help, h`

Shows a list of commands or help for one command.

To view subcommand usage...

```
$ awsprice help [command name]
```

### Global Options

#### `--format value, -f value`

(default: `pp`)

Output format. ("pp" or "json")

#### `--cacheTTL value`

(default: `24h`)

Max age of cache. If existing cache file is older than this value, cache will be recreated. using [`time.ParseDuration(s string)`](https://golang.org/pkg/time/#ParseDuration).

#### `--clearCache`

Remove all cache files before exiting.

#### `--help, -h`

Show help.

#### `--version, -v`

Print the version.
