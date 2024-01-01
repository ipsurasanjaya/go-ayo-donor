<!-- markdownlint-disable MD014 MD024 MD026 MD033 MD036 MD041 -->

<div align='center'>

![go-ayo-donor](./header.png)

</div>

go-ayo-donor is an app that provides an API related to blood donation. Service data from [Palang Merah Indonesia](https://ayodonor.pmi.or.id) official site as JSON API.

> I use this app as a personal project and I thought maybe this API could be useful for someone else, so I decided to make this app open source.

---

<details>
<summary><b>View table of contents</b></summary><br/>

- [Routes](#routes)
- [Current Limitations](#current-limitations)
- [UDD (Unit Donor Darah) Code List](#udd-unit-donor-darah-code-list)
- [Province List](#provinces-list)
- [Contributing](#contributing)
- [License](#license)
</details>

---

## Routes
- /api/v1/bloods/supplies [GET] retrieve list of blood supplies
- /api/v1/bloods/supplies/:udd [GET] retrieve list of blood supplies by UDD (Unit Donor Darah)
- /api/v1/mobiles [GET] retrieve list of PMI mobile unit
- /api/v1/mobiles/:province [GET] retrieve list of PMI mobile unit by province. this is the list of the [provinces](#provinces-list) 

## Current Limitations
1. The performance of all APIs is slow because it is very dependent on the performance of PMI's official site.
2. Currently get bood supplies by UDD endpoint does not support inputting the name of the Kabupaten or City (should be inputting the UDD). 
3. limited UDD data that has been collected.
4. Currently get bood mobile by province endpoint should input the exact [province](#provinces-list) name. e.g. this endpoint will return error when receiving input Sumbar (an acronym for West Sumatra)

## UDD (Unit Donor Darah) Code List
| UDD                             | Name                                |
| ------------------------------- | ----------------------------------- |
| 3171                            | PMI DKI Jakarta                     |
| 1103                            | PMI Aceh Timur                      |
| 1173                            | PMI Aceh Utara                      |
| 1171                            | PMI Kota Banda Aceh                 |
| 6301                            | PMI Kabupaten Tanah Laut            |
| 7571                            | PMI Gorontalo                       |
| 3300                            | PMI Jawa Tengah                     |
| 6309                            | PMI Tabalong                        |
| 1206                            | PMI Kabupaten Balige                |
| 3204                            | PMI Kabupaten Bandung               |
| 7207                            | PMI Kabupaten Banggai               |
| 3526                            | PMI Kabupaten Bangkalan             |
| 5106                            | PMI Kabupaten Bangli                |
| 3279                            | PMI Kabupaten Banjar                |

## Provinces List

This tables shows the list of the provinces and its PMI mobile data readiness.

:white_check_mark: : ready

:x: : not ready

| Province                        | Status                              |
| ------------------------------- | ----------------------------------- |
| Nanggroe Aceh Darussalam        | :x:                                 |
| Sumatera Utara                  | :white_check_mark:                  |
| Sumatera Selatan                | :x:                                 |
| Sumatera Barat                  | :x:                                 |
| Bengkulu                        | :x:                                 |
| Riau                            | :x:                                 |
| Kepulauan Riau                  | :x:                                 |
| Jambi                           | :x:                                 |
| Lampung                         | :x:                                 |
| Bangka Belitung                 | :x:                                 |
| Kalimantan Barat                | :x:                                 |
| Kalimantan Timur                | :white_check_mark:                  |
| Kalimantan Selatan              | :white_check_mark:                  |
| Kalimantan Tengah               | :x:                                 |
| Kalimantan Utara                | :x:                                 |
| Banten                          | :white_check_mark:                  |
| DKI Jakarta                     | :x:                                 |
| Jawa Barat                      | :white_check_mark:                  |
| Jawa Tengah                     | :white_check_mark:                  |
| Daerah Istimewa Yogyakarta      | :white_check_mark:                  |
| Jawa Timur                      | :white_check_mark:                  |
| Bali                            | :white_check_mark:                  |
| Nusa Tenggara Timur             | :x:                                 |
| Nusa Tenggara Barat             | :x:                                 |
| Gorontalo                       | :x:                                 |
| Sulawesi Barat                  | :x:                                 |
| Sulawesi Tengah                 | :x:                                 |
| Sulawesi Utara                  | :x:                                 |
| Sulawesi Tenggara               | :x:                                 |
| Sulawesi Selatan                | :x:                                 |
| Maluku Utara                    | :x:                                 |
| Maluku                          | :x:                                 |
| Papua Barat                     | :x:                                 |
| Papua                           | :x:                                 |
| Papua Tengah                    | :x:                                 |
| Papua Pegunungan                | :x:                                 |
| Papua Selatan                   | :x:                                 |
| Papua Barat Daya                | :x:                                 |

## Contributing

No rules for now. Feel free to add issue first and optionally submit a PR. Cheers

## License

MIT. Copyright 2022 [Sura Sanjaya](./LICENSE)