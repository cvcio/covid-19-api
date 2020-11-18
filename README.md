# COVID-19 Global / Greece API

> last updated at 2020-11-06 10:08 +2000

Covid&mdash;19 API is open-source and developed by the team at [Civic Information Office](https://cvcio.org/) in collaboration with [iMedD](https://imedd.org/) to help academics, data scientists, newsrooms & journalists, goverment agencies, health professionals and the public understand the COVID-19 outbreak in Greece.

Data is sourced from [Johns Hopkins CSSE]() and [WorlOMeter]() for global data, from [EODY]() and [iMedD](https://imedd.org/) for Greece. Data updates run approximately every two hours for global data. For Greece, we collect data from [EODY]()s daily reports, **when and if published**, and direct updates sent to [iMedD](https://imedd.org/)s dedicated team, by [EODY]()s officials. Unfortunatelly, the Greek administration has no official data sharing mechanism in place, thus our work in not trivial.

Additionaly to COVID-19 related data we collect information regarding policy responses, such as area specific lockdowns, and information about educational structures, such as schools suspension, as published by [GSCP](https://www.civilprotection.gr/) and [SCH](https://www.sch.gr/anastoli/web/) respectively.

## Data Format

## Endpoints

## Rate Limiting

We introduced rate limiting from the begining as it is a critical aspect of the API's scalability and performance, and prevent abuse by automated system and humans. The global rate limit is set to **300 requests per minute**, but this may change without direct notice. We plan to introduce a token based authentication to bypass the limiting in the near future.

## License and Attribution

In general, we are making this data publicly available for broad, noncommercial public use including by medical and public health researchers, policymakers, analysts and local news media.

If you use this API, you must attribute it to “Civic Information Office / Incubator for Media Education and Development” or "CVCIO/iMedD" in any publication. If you would like a more expanded description of the data, you could say “Data from The New York Times, based on reports from state and local health agencies.”

If you use it in an online presentation, we would appreciate it if you would link to our U.S. tracking page at https://www.nytimes.com/interactive/2020/us/coronavirus-us-cases.html.

If you use this data, please let us know at info@cvcio.org.

See our LICENSE for the full terms of use for this data.

This license is co-extensive with the Creative Commons Attribution-NonCommercial 4.0 International license, and licensees should refer to that license (CC BY-NC) if they have questions about the scope of the license.

## Contributors

Kelly Kikki (@KellyKiki_), Thanasis Troboukis (@Troboukis), Ilias Dimos (@dosko64), Dimitris Papaevagelou (@andefined).
