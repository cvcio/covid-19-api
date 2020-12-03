# COVID&ndash;19 API

> last updated at Dec 12, 2020

COVID&ndash;19 API is open-source and developed by [Civic Information Office](https://cvcio.org/) in collaboration with [iMEdD](https://imedd.org/) to help academics, data scientists, journalists, the public, understand the COVID&ndash;19 outbreak in Greece (and worldwide), as the Greek government lacks transparency on this issue. The core development team consists of two people, working very hard to keep this up and running. This is an open-source project, so please contribute to make it better.

In general, we collect data from [Johns Hopkins CSSE]((https://github.com/CSSEGISandData/COVID-19)) and [WoldOMeter](https://www.worldometers.info/coronavirus/) for the global collection, and from [iMEdD](https://github.com/iMEdD-Lab) for data related to Greece. Data updates run approximately every two hours, and every 15 minutes during 17:00-21:00 +2. You can read more about the data in [iMEdD](https://imedd.org/)'s [open-data relevant repository](https://github.com/iMEdD-Lab/open-data) or see in action the [COVID&ndash;19 Data Visualization Dashboard](https://lab.imedd.org/covid19/). If you are interested in the data collection mechanism you can refer to [COVID&ndash;19 automation](https://github.com/cvcio/covid-19-automation) service.

## Public Endpoints

We provide a public url for the API only for testing purposes only. You should use [iMEdD](https://github.com/iMEdD-Lab)'s public url in production, as they are responsible for the data.

- https://covid.cvcio.org (Testing)

## Data Format

Data format may vary accross documents as we enrich data related to Greece. In general we serve 3 different endpoints -raw, total and aggregared- for 2 different levels -global, greece. We are working to introducing even more.

###### Raw Global Data

```json
// GET /global/all/all/2020-12-01
[
    {
        "active": 62,
        "case_fatality_ratio": 1.6486,
        "cases": 5156,
        "country": "Equatorial Guinea",
        "date": "2020-12-01T02:00:00+02:00",
        "deaths": 85,
        "incidence_rate": 367.5021,
        "iso2": "GQ",
        "iso3": "GNQ",
        "last_updated_at": "2020-12-02T19:55:37.507+02:00",
        "loc": {
            "coordinates": [
                10.2679,
                1.6508
            ],
            "type": "Point"
        },
        "new_cases": 3,
        "new_deaths": 0,
        "new_recovered": 0,
        "population": 1402985,
        "recovered": 5009,
        "source": "jhu",
        "uid": 226
    },
    {
        "active": 45229,
        "cases": 74205,
        "country": "Greece",
        "critical": 392,
        "cumulative_rapid_tests_raw": 112416,
        "cumulative_rtpcr_tests_raw": 2092324,
        "date": "2020-11-15T02:00:00+02:00",
        "deaths": 1106,
        "esitmated_new_rapid_tests": 1635,
        "estimated_new_rtpcr_tests": 14262,
        "intensive_care": 0,
        "iso2": "GR",
        "iso3": "GRC",
        "last_updated_at": "2020-12-02T19:56:23.868+02:00",
        "loc": {
            "coordinates": [
                21.8243,
                39.0742
            ],
            "type": "Point"
        },
        "new_cases": 1695,
        "new_deaths": 71,
        "new_hospitalized": 0,
        "new_recovered": 0,
        "new_tests": 15897,
        "population": 10423056,
        "recovered": 27870,
        "source": "imedd",
        "tests": 2204743,
        "uid": 300
    },
    ...
]
```

###### Raw Greece Local Data

```json
// GET /greece/all/all/2020-12-01
[
    {
        "case_fatality_ratio": 0,
        "cases": 549,
        "date": "2020-12-01T02:00:00+02:00",
        "deaths": 0,
        "geo_unit": "-",
        "incidence_rate": 0,
        "last_updated_at": "2020-12-02T19:56:17.232+02:00",
        "new_cases": 2,
        "new_deaths": 0,
        "population": 0,
        "region": "Imported (They asked to be tested)",
        "source": "imedd",
        "state": "-",
        "uid": "EL002"
    },
    {
        "case_fatality_ratio": 1.6015,
        "cases": 1561,
        "date": "2020-12-01T02:00:00+02:00",
        "deaths": 25,
        "geo_unit": "Thrace",
        "incidence_rate": 1055.1076,
        "last_updated_at": "2020-12-02T19:56:17.232+02:00",
        "loc": {
            "coordinates": [
                26.135943100000002,
                41.2443761
            ],
            "type": "Point"
        },
        "new_cases": 33,
        "new_deaths": 0,
        "population": 147947,
        "region": "Evros",
        "source": "imedd",
        "state": "East Macedonia-Thrace",
        "uid": "EL111"
    },
    (...)
]
```

###### Global Total Data (Beta)

The `total` endpoint is still in active development and may change without further notice.

```json
// GET /total/global
[
    {
        "cases": 0,
        "country": "Equatorial Guinea",
        "deaths": 0,
        "iso2": "GQ",
        "iso3": "GNQ",
        "last_updated_at": "2020-12-03T14:34:35.696+02:00",
        "loc": {
            "coordinates": [
                10.2679,
                1.6508
            ],
            "type": "Point"
        },
        "population": 1402985,
        "recovered": 0,
        "sources": [
            "worldometer"
        ],
        "total_active": 62,
        "total_cases": 5156,
        "total_critical": 0,
        "total_deaths": 85,
        "total_recovered": 5009,
        "total_tests": 70718,
        "uid": 226
    },
    {
        "cases": 0,
        "country": "Greece",
        "deaths": 0,
        "iso2": "GR",
        "iso3": "GRC",
        "last_updated_at": "2020-12-03T14:35:09.394+02:00",
        "loc": {
            "coordinates": [
                21.8243,
                39.0742
            ],
            "type": "Point"
        },
        "population": 10423056,
        "recovered": 0,
        "sources": [
            "imedd"
        ],
        "total_active": 29015,
        "total_cases": 109655,
        "total_critical": 0,
        "total_deaths": 2606,
        "total_recovered": 78034,
        "total_tests": 0,
        "uid": 300
    },
    (...)
]
```

###### Greece Total Data (Beta)

```json
// GET /total/greece
[
    {
        "cases": 65,
        "deaths": 0,
        "geo_unit": "Macedonia",
        "last_updated_at": "2020-12-03T14:35:03.245+02:00",
        "loc": {
            "coordinates": [
                24.412866,
                40.937607
            ],
            "type": "Point"
        },
        "population": 138687,
        "recovered": 0,
        "region": "Kavala",
        "sources": [
            "imedd"
        ],
        "state": "East Macedonia-Thrace",
        "total_active": null,
        "total_cases": 1637,
        "total_critical": null,
        "total_deaths": 8,
        "total_recovered": null,
        "uid": "EL115"
    },
    (...)
]
```

###### Global Aggragated Data (Beta)

The `agg` endpoint is still in active development and may change without further notice.

```json
// GET /agg/global/all/all/2020-11-22
[
    {
        "active": [62119, 63422, 65452, 67516, 69435, 21151, 22777, 23872,  24831, 26919, 29015, 29015],
        "cases": [91619, 93006, 95137, 97288, 99306, 101287, 103034, 104227, 105271, 107470, 109655, 109655],
        "country": "Greece",
        "critical": [540, 549, 562, 597, 608, 607, 606, 607, 600, 596, 613, 0],
        "deaths": [1630, 1714, 1815, 1902, 2001, 2102, 2223, 2321, 2406, 2517, 2606, 2606],
        "from": "2020-11-22T02:00:00+02:00",
        "iso2": "GR",
        "iso3": "GRC",
        "last_updated_at": "2020-12-03T14:35:09.394+02:00",
        "loc": {
            "coordinates": [
                21.8243,
                39.0742
            ],
            "type": "Point"
        },
        "new_cases": [1498, 1387, 2131, 2151, 2018, 1981, 1747, 1193, 1044, 2199, 2185, 0],
        "new_deaths": [103, 84, 101, 87, 99, 101, 121, 98, 85, 111, 89, 0],
        "population": 10423056,
        "recovered": [27870, 27870, 27870, 27870, 27870, 78034, 78034, 78034, 78034, 78034, 78034, 78034],
        "sources": [
            "imedd"
        ],
        "to": "2020-12-03T02:00:00+02:00",
        "uid": 300
    },
    (...)
]
```

###### Greece Aggregated Data (Beta)

```json
// GET /agg/greece/all/all/2020-11-22
[
    {
        "active": [],
        "cases": [162, 167, 170, 182, 189, 201, 204, 205, 211, 215, 220, 220],
        "critical": [],
        "deaths": [2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2],
        "from": "2020-11-22T02:00:00+02:00",
        "geo_unit": "Epirus",
        "last_updated_at": "2020-12-03T14:35:03.245+02:00",
        "loc": {
            "coordinates": [
                20.987683899999997,
                39.1582421
            ],
            "type": "Point"
        },
        "new_cases": [8, 5, 3, 12, 7, 12, 3, 1, 6, 4, 5, 5],
        "new_deaths": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        "population": 67877,
        "recovered": [],
        "region": "Arta",
        "sources": [
            "imedd"
        ],
        "state": "Epirus",
        "to": "2020-12-03T02:00:00+02:00",
        "uid": "EL211"
    },
    (...)
]
```

## Endpoints

All endpoints accept various parameters. We store country level information on `global` collection and region level information on `greece` collection. 

Country and region parameter refer to country specific iso3 code or nuts codes, for global and greece collections accordingly. Accepted `:country` and `:region` params are either `all` to get data for all countries / regions, or iso3 country codes / nuts codes to get specific country or regional data. You can refer to this file [countries-mapping-jhu-wom.csv](https://github.com/cvcio/covid-19-automation/blob/main/data/countries-mapping-jhu-wom.csv), that will help you understand how we map [JHU](https://github.com/CSSEGISandData/COVID-19) and [WoldOMeter](https://www.worldometers.info/coronavirus/) data (column `iso3`) and to this file [region-mapping-imedd.csv](https://github.com/cvcio/covid-19-automation/blob/main/data/region-mapping-imedd.csv) for mapping data from [iMedD](https://github.com/iMEdD-Lab/open-data) (column `uid`).

Accepted `:keys` are either `all`, which will retrieve all the defaults, or document specific keys, single or comma seperated. This parameter is not included in the aggregated data endpoint.

Finally, `:from` and `:to` parameters will retrieve data in the specified date range in `YYYY-MM-DD` format. If no `:from` parameter  provided, we will only return tha last date. If no `:to` parameter  provided, we will return up-to current date.

Please, if you think that we are missing some countries, or any other data related issues, open an issue on [COVID&ndash;19 automation](https://github.com/cvcio/covid-19-automation) repository.

Accepted keys, available for all countries are:

- **all**: default, will return all available
- **cases**: cumulative cases
- **deaths**: cumulative deaths
- **recovered**: cumulative recovered cases 
- **active**: daily active cases (cases - deaths - recovered)
- **critical**: daily critical cases available only for current date
- **tests**: daily tests available only for current date
- **new_cases**: daily cases
- **new_deaths**: daily deaths
- **new_recovered**: daily recovered cases
- **case_fatality_ratio**: daily case fatality ratio ((death / cases) * 100)
- **incidence_rate**: daily incidence ratio per 100K population ((cases * 100000) / population)

and for Greece (country:grc) only:

- **cumulative_rtpcr_tests_raw**: cumulative rt-pcr tests
- **estimated_new_rtpcr_tests**: daily rt-pcr tests
- **cumulative_rapid_tests_raw**: cumulative rapid tests
- **esitmated_new_rapid_tests**: daily rapid tests
- **estimated_new_total_tests**: daily total tests

For regioanal data (Greece) available keys are:

- **all**: default, will return all available
- **cases**: cumulative cases
- **deaths**: cumulative deaths
- **new_cases**: daily cases
- **new_deaths**: daily deaths
- **case_fatality_ratio**: daily case fatality ratio ((death / cases) * 100)
- **incidence_rate**: daily incidence ratio per 100K population ((cases * 100000) / population)

###### Raw Global Data

```bash
GET /global/:country/:keys/:from/:to

# ex. get all data, for all countries, from the begining
# of the pandemic.
curl -XGET https://covid.cvcio.org/global/all/all/2020-01-01

# ex. get all data, for all countries just for today
# in this scenario :country and :keys are optional
# the following requests will return the same response
curl -XGET https://covid.cvcio.org/global
curl -XGET https://covid.cvcio.org/global/all
curl -XGET https://covid.cvcio.org/global/all/all

# ex. get only the new_cases key for Greece
curl -XGET https://covid.cvcio.org/global/grc/new_cases

# ex. get only the new_cases and cases keys for all countries
# just for today
curl -XGET https://covid.cvcio.org/global/all/new_cases,cases
```

###### Raw Greece Data

```bash
GET /greece/:region/:keys/:from/:to

# ex. get all data, for all regions in greece, from the begining 
# of the pandemic.
curl -XGET https://covid.cvcio.org/greece/all/all/2020-01-01

# get all cases for Thessaloniki region between Oct and Nov
curl -XGET https://covid.cvcio.org/greece/EL122/cases/2020-10-01/2020-11-31
```

###### Global Total Data (Beta)

```bash
GET /total/global/:country/:keys/:from/:to

# ex. get total recovered cases for all countries,
# from the begining of the pandemic.
curl -XGET https://covid.cvcio.org/total/global/all/cases/2020-10-01
```

###### Greece Total Data (Beta)

```bash
GET /total/greece/:region/:keys/:from/:to

# ex. get total imported (detected at the entry points) cases just for today
curl -XGET https://covid.cvcio.org/total/greece/EL001/cases
```

###### Global Aggregated Data (Beta)

```bash
GET /agg/global/:country/:from/:to

# ex. get all aggregated data for Greece, from
# the begining of the pandemic
curl -XGET https://covid.cvcio.org/agg/global/GRC/2020-01-01
```

###### Greece Aggregated Data (Beta)

```bash
GET /agg/greece/:region/:from/:to

# ex. get all aggregated data for Attica region, from
# the begining of the pandemic
curl -XGET https://covid.cvcio.org/agg/greece/EL300/2020-01-01
```

*Note: the aggregated data endpoint doesn't include the `:keys` parameter*

## Rate Limiting

We introduced rate limiting from the begining as it is a critical aspect of the API's performance, and/or prevent abuse by automated system and humans. The global rate limit is set to **300 requests per minute**, but this may change without direct notice. We plan to introduce a token based authentication to bypass the limiting in the near future.

## Getting started

You will need to run [golang](https://golang.org/) (>= version 1.14) to build the api, [mongodb](https://www.mongodb.com/) to store the documents and [redis](https://redis.io/) for caching and rate limiting. We suggest to use docker during development.

##### Development

```bash
make db-start
make run-api
```

## Contribution

If you're new to contributing to Open Source on Github, [this guide](https://opensource.guide/how-to-contribute/) can help you get started. Please check out the contribution guide for more details on how issues and pull requests work. Before contributing be sure to review the [code of conduct](https://github.com/cvcio/covid-19-api/blob/main/CODE_OF_CONDUCT.md).

## License and Attribution

In general, we are making this software publicly available for broad, noncommercial public use, including by medical and public health researchers, policymakers, analysts and local news media.

If you use this API, please let us know at [info@cvcio.org](mailto:info@cvcio.org).

See our [LICENSE](https://github.com/cvcio/covid-19-api/blob/main/LICENSE.md) for the full terms of use for this software.

## Related Repositories

- [COVID&ndash;19 Automation](https://github.com/cvcio/covid-19-automation)
- [COVID&ndash;19 API](https://github.com/cvcio/covid-19-api)
- [COVID&ndash;19 Application](https://github.com/cvcio/covid-19)
- [COVID&ndash;19 Open Data](https://github.com/iMEdD-Lab/open-data) (iMEdD)

## Contributors

- Ilias Dimos ([@dosko64](https://github.com/dosko64))
- Dimitris Papaevagelou ([@andefined](https://github.com/andefined)).
