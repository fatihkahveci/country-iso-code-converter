package converter

import (
	"strings"
)

var countryList2 = map[string]string{
	"AF": "AFG",
	"AX": "ALA",
	"AL": "ALB",
	"DZ": "DZA",
	"AS": "ASM",
	"AD": "AND",
	"AO": "AGO",
	"AI": "AIA",
	"AQ": "ATA",
	"AG": "ATG",
	"AR": "ARG",
	"AM": "ARM",
	"AW": "ABW",
	"AU": "AUS",
	"AT": "AUT",
	"AZ": "AZE",
	"BQ": "BES",
	"BS": "BHS",
	"BH": "BHR",
	"BD": "BGD",
	"BB": "BRB",
	"BY": "BLR",
	"BE": "BEL",
	"BZ": "BLZ",
	"BJ": "BEN",
	"BM": "BMU",
	"BT": "BTN",
	"BO": "BOL",
	"BA": "BIH",
	"BW": "BWA",
	"BV": "BVT",
	"BR": "BRA",
	"VG": "VGB",
	"IO": "IOT",
	"BN": "BRN",
	"BG": "BGR",
	"BF": "BFA",
	"BI": "BDI",
	"KH": "KHM",
	"CM": "CMR",
	"CA": "CAN",
	"CV": "CPV",
	"KY": "CYM",
	"CF": "CAF",
	"TD": "TCD",
	"CL": "CHL",
	"CN": "CHN",
	"CW": "CUW",
	"HK": "HKG",
	"MO": "MAC",
	"CX": "CXR",
	"CC": "CCK",
	"CO": "COL",
	"KM": "COM",
	"CG": "COG",
	"CD": "COD",
	"CK": "COK",
	"CR": "CRI",
	"CI": "CIV",
	"HR": "HRV",
	"CU": "CUB",
	"CY": "CYP",
	"CZ": "CZE",
	"DK": "DNK",
	"DJ": "DJI",
	"DM": "DMA",
	"DO": "DOM",
	"EC": "ECU",
	"EG": "EGY",
	"SV": "SLV",
	"GQ": "GNQ",
	"ER": "ERI",
	"EE": "EST",
	"ET": "ETH",
	"FK": "FLK",
	"FO": "FRO",
	"FJ": "FJI",
	"FI": "FIN",
	"FR": "FRA",
	"GF": "GUF",
	"PF": "PYF",
	"TF": "ATF",
	"GA": "GAB",
	"GM": "GMB",
	"GE": "GEO",
	"DE": "DEU",
	"GH": "GHA",
	"GI": "GIB",
	"GR": "GRC",
	"GL": "GRL",
	"GD": "GRD",
	"GP": "GLP",
	"GU": "GUM",
	"GT": "GTM",
	"GG": "GGY",
	"GN": "GIN",
	"GW": "GNB",
	"GY": "GUY",
	"HT": "HTI",
	"HM": "HMD",
	"VA": "VAT",
	"HN": "HND",
	"HU": "HUN",
	"IS": "ISL",
	"IN": "IND",
	"ID": "IDN",
	"IR": "IRN",
	"IQ": "IRQ",
	"IE": "IRL",
	"IM": "IMN",
	"IL": "ISR",
	"IT": "ITA",
	"JM": "JAM",
	"JP": "JPN",
	"JE": "JEY",
	"JO": "JOR",
	"KZ": "KAZ",
	"KE": "KEN",
	"KI": "KIR",
	"KP": "PRK",
	"KR": "KOR",
	"KW": "KWT",
	"KG": "KGZ",
	"LA": "LAO",
	"LV": "LVA",
	"LB": "LBN",
	"LS": "LSO",
	"LR": "LBR",
	"LY": "LBY",
	"LI": "LIE",
	"LT": "LTU",
	"LU": "LUX",
	"MK": "MKD",
	"MG": "MDG",
	"MW": "MWI",
	"MY": "MYS",
	"MV": "MDV",
	"ML": "MLI",
	"MT": "MLT",
	"MH": "MHL",
	"MQ": "MTQ",
	"MR": "MRT",
	"MU": "MUS",
	"YT": "MYT",
	"MX": "MEX",
	"FM": "FSM",
	"MD": "MDA",
	"MC": "MCO",
	"MN": "MNG",
	"ME": "MNE",
	"MS": "MSR",
	"MA": "MAR",
	"MZ": "MOZ",
	"MM": "MMR",
	"NA": "NAM",
	"NR": "NRU",
	"NP": "NPL",
	"NL": "NLD",
	"AN": "ANT",
	"NC": "NCL",
	"NZ": "NZL",
	"NI": "NIC",
	"NE": "NER",
	"NG": "NGA",
	"NU": "NIU",
	"NF": "NFK",
	"MP": "MNP",
	"NO": "NOR",
	"OM": "OMN",
	"PK": "PAK",
	"PW": "PLW",
	"PS": "PSE",
	"PA": "PAN",
	"PG": "PNG",
	"PY": "PRY",
	"PE": "PER",
	"PH": "PHL",
	"PN": "PCN",
	"PL": "POL",
	"PT": "PRT",
	"PR": "PRI",
	"QA": "QAT",
	"RE": "REU",
	"RO": "ROU",
	"RU": "RUS",
	"RW": "RWA",
	"BL": "BLM",
	"SH": "SHN",
	"KN": "KNA",
	"LC": "LCA",
	"MF": "MAF",
	"PM": "SPM",
	"VC": "VCT",
	"WS": "WSM",
	"SM": "SMR",
	"ST": "STP",
	"SA": "SAU",
	"SN": "SEN",
	"RS": "SRB",
	"SC": "SYC",
	"SL": "SLE",
	"SG": "SGP",
	"SK": "SVK",
	"SI": "SVN",
	"SB": "SLB",
	"SO": "SOM",
	"ZA": "ZAF",
	"GS": "SGS",
	"SS": "SSD",
	"ES": "ESP",
	"LK": "LKA",
	"SD": "SDN",
	"SR": "SUR",
	"SJ": "SJM",
	"SZ": "SWZ",
	"SE": "SWE",
	"CH": "CHE",
	"SY": "SYR",
	"SX": "SXM",
	"TW": "TWN",
	"TJ": "TJK",
	"TZ": "TZA",
	"TH": "THA",
	"TL": "TLS",
	"TG": "TGO",
	"TK": "TKL",
	"TO": "TON",
	"TT": "TTO",
	"TN": "TUN",
	"TR": "TUR",
	"TM": "TKM",
	"TC": "TCA",
	"TV": "TUV",
	"UG": "UGA",
	"UA": "UKR",
	"AE": "ARE",
	"GB": "GBR",
	"US": "USA",
	"UM": "UMI",
	"UY": "URY",
	"UZ": "UZB",
	"VU": "VUT",
	"VE": "VEN",
	"VN": "VNM",
	"VI": "VIR",
	"WF": "WLF",
	"EH": "ESH",
	"YE": "YEM",
	"ZM": "ZMB",
	"ZW": "ZWE",
}

var countryList3 = map[string]string{
	"AFG": "AF",
	"ALA": "AX",
	"ALB": "AL",
	"DZA": "DZ",
	"ASM": "AS",
	"AND": "AD",
	"AGO": "AO",
	"AIA": "AI",
	"ATA": "AQ",
	"ATG": "AG",
	"ARG": "AR",
	"ARM": "AM",
	"ABW": "AW",
	"AUS": "AU",
	"AUT": "AT",
	"AZE": "AZ",
	"BEQ": "BQ",
	"BHS": "BS",
	"BHR": "BH",
	"BGD": "BD",
	"BRB": "BB",
	"BLR": "BY",
	"BEL": "BE",
	"BLZ": "BZ",
	"BEN": "BJ",
	"BMU": "BM",
	"BTN": "BT",
	"BOL": "BO",
	"BIH": "BA",
	"BWA": "BW",
	"BVT": "BV",
	"BRA": "BR",
	"VGB": "VG",
	"IOT": "IO",
	"BRN": "BN",
	"BGR": "BG",
	"BFA": "BF",
	"BDI": "BI",
	"KHM": "KH",
	"CMR": "CM",
	"CAN": "CA",
	"CPV": "CV",
	"CYM": "KY",
	"CAF": "CF",
	"CUW": "CW",
	"TCD": "TD",
	"CHL": "CL",
	"CHN": "CN",
	"HKG": "HK",
	"MAC": "MO",
	"CXR": "CX",
	"CCK": "CC",
	"COL": "CO",
	"COM": "KM",
	"COG": "CG",
	"COD": "CD",
	"COK": "CK",
	"CRI": "CR",
	"CIV": "CI",
	"HRV": "HR",
	"CUB": "CU",
	"CYP": "CY",
	"CZE": "CZ",
	"DNK": "DK",
	"DJI": "DJ",
	"DMA": "DM",
	"DOM": "DO",
	"ECU": "EC",
	"EGY": "EG",
	"SLV": "SV",
	"GNQ": "GQ",
	"ERI": "ER",
	"EST": "EE",
	"ETH": "ET",
	"FLK": "FK",
	"FRO": "FO",
	"FJI": "FJ",
	"FIN": "FI",
	"FRA": "FR",
	"GUF": "GF",
	"PYF": "PF",
	"ATF": "TF",
	"GAB": "GA",
	"GMB": "GM",
	"GEO": "GE",
	"DEU": "DE",
	"GHA": "GH",
	"GIB": "GI",
	"GRC": "GR",
	"GRL": "GL",
	"GRD": "GD",
	"GLP": "GP",
	"GUM": "GU",
	"GTM": "GT",
	"GGY": "GG",
	"GIN": "GN",
	"GNB": "GW",
	"GUY": "GY",
	"HTI": "HT",
	"HMD": "HM",
	"VAT": "VA",
	"HND": "HN",
	"HUN": "HU",
	"ISL": "IS",
	"IND": "IN",
	"IDN": "ID",
	"IRN": "IR",
	"IRQ": "IQ",
	"IRL": "IE",
	"IMN": "IM",
	"ISR": "IL",
	"ITA": "IT",
	"JAM": "JM",
	"JPN": "JP",
	"JEY": "JE",
	"JOR": "JO",
	"KAZ": "KZ",
	"KEN": "KE",
	"KIR": "KI",
	"PRK": "KP",
	"KOR": "KR",
	"KWT": "KW",
	"KGZ": "KG",
	"LAO": "LA",
	"LVA": "LV",
	"LBN": "LB",
	"LSO": "LS",
	"LBR": "LR",
	"LBY": "LY",
	"LIE": "LI",
	"LTU": "LT",
	"LUX": "LU",
	"MKD": "MK",
	"MDG": "MG",
	"MWI": "MW",
	"MYS": "MY",
	"MDV": "MV",
	"MLI": "ML",
	"MLT": "MT",
	"MHL": "MH",
	"MTQ": "MQ",
	"MRT": "MR",
	"MUS": "MU",
	"MYT": "YT",
	"MEX": "MX",
	"FSM": "FM",
	"MDA": "MD",
	"MCO": "MC",
	"MNG": "MN",
	"MNE": "ME",
	"MSR": "MS",
	"MAR": "MA",
	"MOZ": "MZ",
	"MMR": "MM",
	"NAM": "NA",
	"NRU": "NR",
	"NPL": "NP",
	"NLD": "NL",
	"ANT": "AN",
	"NCL": "NC",
	"NZL": "NZ",
	"NIC": "NI",
	"NER": "NE",
	"NGA": "NG",
	"NIU": "NU",
	"NFK": "NF",
	"MNP": "MP",
	"NOR": "NO",
	"OMN": "OM",
	"PAK": "PK",
	"PLW": "PW",
	"PSE": "PS",
	"PAN": "PA",
	"PNG": "PG",
	"PRY": "PY",
	"PER": "PE",
	"PHL": "PH",
	"PCN": "PN",
	"POL": "PL",
	"PRT": "PT",
	"PRI": "PR",
	"QAT": "QA",
	"REU": "RE",
	"ROU": "RO",
	"RUS": "RU",
	"RWA": "RW",
	"BLM": "BL",
	"SHN": "SH",
	"KNA": "KN",
	"LCA": "LC",
	"MAF": "MF",
	"SPM": "PM",
	"VCT": "VC",
	"WSM": "WS",
	"SMR": "SM",
	"STP": "ST",
	"SAU": "SA",
	"SEN": "SN",
	"SRB": "RS",
	"SYC": "SC",
	"SLE": "SL",
	"SGP": "SG",
	"SVK": "SK",
	"SVN": "SI",
	"SLB": "SB",
	"SOM": "SO",
	"ZAF": "ZA",
	"SGS": "GS",
	"SSD": "SS",
	"ESP": "ES",
	"LKA": "LK",
	"SDN": "SD",
	"SUR": "SR",
	"SJM": "SJ",
	"SWZ": "SZ",
	"SXM": "SX",
	"SWE": "SE",
	"CHE": "CH",
	"SYR": "SY",
	"TWN": "TW",
	"TJK": "TJ",
	"TZA": "TZ",
	"THA": "TH",
	"TLS": "TL",
	"TGO": "TG",
	"TKL": "TK",
	"TON": "TO",
	"TTO": "TT",
	"TUN": "TN",
	"TUR": "TR",
	"TKM": "TM",
	"TCA": "TC",
	"TUV": "TV",
	"UGA": "UG",
	"UKR": "UA",
	"ARE": "AE",
	"GBR": "GB",
	"USA": "US",
	"UMI": "UM",
	"URY": "UY",
	"UZB": "UZ",
	"VUT": "VU",
	"VEN": "VE",
	"VNM": "VN",
	"VIR": "VI",
	"WLF": "WF",
	"ESH": "EH",
	"YEM": "YE",
	"ZMB": "ZM",
	"ZWE": "ZW",
}

func ConvertIso2To3(isoCode string) string {
	return countryList2[strings.ToUpper(isoCode)]
}

func ConvertIso3To2(isoCode string) string {
	return countryList3[strings.ToUpper(isoCode)]
}
