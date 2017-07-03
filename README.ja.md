# awsprice

[AWS Price List API](http://docs.aws.amazon.com/ja_jp/awsaccountbilling/latest/aboutv2/price-changes.html) 用のコマンドラインインターフェイスです。

## インストール

[releases](https://github.com/y13i/awsprice/releases) から実行ファイルがダウンロードできます。

展開して `$PATH` の通ったディレクトリーに置いてください。

### Homebrew

macOS であれば Homebrew でもインストールできます。

```
$ brew install https://raw.githubusercontent.com/y13i/awsprice/master/homebrew/awsprice.rb
```

## 使用方法

```
$ awsprice [global options] command [command options] [arguments...]
```

### Commands

#### `listOffers`

オファーコード一覧を表示します。

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

指定されたオファーのリージョン一覧を表示します。

`--offerCode` のデフォルトが `AmazonEC2` である点に注意してください。

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

(デフォルト: `AmazonEC2`)

AWS サービスを特定するコードです。

#### `listOfferVersions`

オファーのバージョン一覧を表示します。

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

オファーに含まれるプロダクトの、プロダクトファミリー一覧を表示します。

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

AWS リージョンコードを指定します。 `$AWS_REGION` 環境変数がセットされており、オプションが未指定の場合は環境変数の値が使用されます。

###### `--offerVersion value`

(デフォルト: 最新バージョン)

オファーのバージョンを指定します。

例: `20170302183221`

###### `--productFamily value, -p value`

プロダクトをプロダクトファミリーの名前で絞り込みます。

複数指定可。

```
-p familyName1 -p familyName2
```

###### `--attribute value, -a value`

プロダクトをプロダクトの属性で絞り込みます。

`KEY[=VALUE]` の形式で指定してください。

複数指定可。

```
-a attribute1=value1 -a attribute2= -a ...
```

#### `listAttributes`

オファーに含まれるプロダクトの、属性一覧を表示します。

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

オプションで絞り込まれたプロダクト一覧を表示します。

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

オファーに含まれる term type (価格タイプ？) 一覧を表示します。

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

オプションで絞り込まれたプロダクトと、それに紐付く価格情報の一覧を表示します。

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

(デフォルト: `OnDemand`)

価格タイプを指定します。 "OnDemand" または "Reserved"

###### `--priceUnit value`

(デフォルト: `USD`)

価格の単位を指定します。現状特に意味はないです。

#### `help, h`

ヘルプを表示します。

各サブコマンドの使用方法を見るには以下のように指定します。

```
$ awsprice help [command name]
```

### Global Options

#### `--format value, -f value`

(デフォルト: `pp`)

出力形式を指定します。 "pp" または "json"

#### `--cacheTTL value`

(デフォルト: `24h`)

キャッシュの有効期限を指定します。もし存在しているキャッシュがこの値より古い場合、キャッシュは再作成されます。 [`time.ParseDuration(s string)`](https://golang.org/pkg/time/#ParseDuration) でオプションの値がパースされます。

#### `--clearCache`

終了時にすべてのキャッシュファイルを削除します。

#### `--help, -h`

ヘルプを表示します。

#### `--version, -v`

バージョンを表示します。
