# COVID&ndash;19 Global / Greece API

> last updated at 2020-11-18

COVID&ndash;19 API is open-source and developed by the team at [Civic Information Office](https://cvcio.org/) in collaboration with [iMedD](https://imedd.org/) to help academics, data scientists, journalists, goverment agencies, health professionals and the public, understand the COVID-19 outbreak in Greece and worldwide.

Data is sourced from [Johns Hopkins CSSE](https://github.com/CSSEGISandData/COVID-19) and [WorlOMeter](https://www.worldometers.info/coronavirus/) for global data, and from [EODY](https://eody.gov.gr/) and [iMedD](https://imedd.org/) for data related to Greece. Data updates run approximately every two hours. For Greece, we collect data from [EODY](https://eody.gov.gr/epidimiologika-statistika-dedomena/ektheseis-covid-19/)'s daily reports, **when and if published**, and direct updates sent to [iMedD](https://imedd.org/)'s dedicated team by [EODY](https://eody.gov.gr/). Unfortunatelly, the Greek administration has no official data sharing mechanism in place, thus there might be some delays between updates. Additionaly to COVID&ndash;19 related data we collect information regarding policy responses, such as area specific lockdowns, and information about educational structures, such as schools suspension, as published by [GSCP](https://www.civilprotection.gr/) and [SCH](https://www.sch.gr/anastoli/web/) respectively.

You can read more about the data in [iMedD](https://imedd.org/)'s [open-data relevant repository](https://github.com/iMEdD-Lab/open-data) or see in action the [COVID&ndash;19 dashboard](https://lab.imedd.org/covid19/). If you are interested in the data collection mechanism you can refer to [COVID&ndash;19 automation](https://github.com/cvcio/covid-19-automation) service.

## Data Format

## Endpoints

## Rate Limiting

We introduced rate limiting from the begining as it is a critical aspect of the API's performance, and to prevent abuse by automated system and humans. The global rate limit is set to **300 requests per minute**, but this may change without direct notice. We plan to introduce a token based authentication to bypass the limiting in the near future.

## License and Attribution

In general, we are making this data publicly available for broad, noncommercial public use, including by medical and public health researchers, policymakers, analysts and local news media.

If you use this API, you must attribute it to “Civic Information Office / Incubator for Media Education and Development” or "CVCIO/iMedD" in any publication.

If you use this data, please let us know at [info@cvcio.org](mailto:info@cvcio.org).

See our LICENSE for the full terms of use for this data.

This license is co-extensive with the Creative Commons Attribution-NonCommercial 4.0 International license, and licensees should refer to that license (CC BY-NC) if they have questions about the scope of the license.

## Related Repositories

- [COVID&ndash;19 Automation](https://github.com/cvcio/covid-19-automation)
- [COVID&ndash;19 API](https://github.com/cvcio/covid-19-api)
- [COVID&ndash;19 Application](https://github.com/cvcio/covid-19)
- [COVID&ndash;19 Open Data](https://github.com/iMEdD-Lab/open-data)

## Contributors

- Ilias Dimos ([@dosko64](https://github.com/dosko64))
- Dimitris Papaevagelou ([@andefined](https://github.com/andefined)).
