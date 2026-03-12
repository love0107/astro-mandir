package kundali

type KundaliRequest struct {
	Name  string `json:"name"`
	DOB   string `json:"dob"`
	TOB   string `json:"tob"`
	Place string `json:"place"`
}

type KundaliData struct {
	Name      string `json:"name"`
	Rashi     string `json:"rashi"`
	Swami     string `json:"swami"`
	Tatva     string `json:"tatva"`
	BhagyaAnk string `json:"bhagya_ank"`
	ShubhRang string `json:"shubh_rang"`
	Mantra    string `json:"mantra"`
	Guidance  string `json:"guidance"`
}

// Rashi details map
var rashiDetails = map[string]map[string]string{
	"mesh": {
		"swami":    "Mangal",
		"tatva":    "Agni",
		"bhagya":   "1, 8",
		"rang":     "Laal",
		"mantra":   "Om Aim Hreem Shreem",
		"guidance": "Aaj mehnat rang laayegi. Parivar ka saath milega.",
	},
	"vrishabh": {
		"swami":    "Shukra",
		"tatva":    "Prithvi",
		"bhagya":   "2, 6",
		"rang":     "Safed",
		"mantra":   "Om Shum Shukraya Namah",
		"guidance": "Aarthik sthiti mazboot rahegi. Naye kaam ki shuruat shubh.",
	},
	"mithun": {
		"swami":    "Budh",
		"tatva":    "Vayu",
		"bhagya":   "3, 5",
		"rang":     "Hara",
		"mantra":   "Om Bum Budhaya Namah",
		"guidance": "Vyapar mein labh milega. Nayi mulakat shubh rahegi.",
	},
	"kark": {
		"swami":    "Chandra",
		"tatva":    "Jal",
		"bhagya":   "2, 7",
		"rang":     "Safed",
		"mantra":   "Om Som Somaya Namah",
		"guidance": "Ghar mein sukh shanti rahegi. Mata ka ashirwad milega.",
	},
	"simha": {
		"swami":    "Surya",
		"tatva":    "Agni",
		"bhagya":   "1, 4",
		"rang":     "Sona",
		"mantra":   "Om Hram Hreem Hraum Sah Suryaya Namah",
		"guidance": "Maan samman badhega. Sarkar se labh milne ki sambhavna.",
	},
	"kanya": {
		"swami":    "Budh",
		"tatva":    "Prithvi",
		"bhagya":   "3, 6",
		"rang":     "Hara",
		"mantra":   "Om Bum Budhaya Namah",
		"guidance": "Kaam mein safalta milegi. Swasthya ka dhyan rakhen.",
	},
	"tula": {
		"swami":    "Shukra",
		"tatva":    "Vayu",
		"bhagya":   "2, 7",
		"rang":     "Neela",
		"mantra":   "Om Shum Shukraya Namah",
		"guidance": "Rishton mein madhurta aayegi. Naye avsar milenge.",
	},
	"vrishchik": {
		"swami":    "Mangal",
		"tatva":    "Jal",
		"bhagya":   "1, 9",
		"rang":     "Laal",
		"mantra":   "Om Kram Kreem Kraum Sah Bhaumaya Namah",
		"guidance": "Aaj ka din sahas ke liye shubh. Dushmano par vijay milegi.",
	},
	"dhanu": {
		"swami":    "Guru",
		"tatva":    "Agni",
		"bhagya":   "3, 9",
		"rang":     "Peela",
		"mantra":   "Om Gram Greem Graum Sah Gurave Namah",
		"guidance": "Dharmik karyon mein man lagega. Vidya mein safalta milegi.",
	},
	"makar": {
		"swami":    "Shani",
		"tatva":    "Prithvi",
		"bhagya":   "6, 8",
		"rang":     "Neela",
		"mantra":   "Om Pram Preem Praum Sah Shanaischaraya Namah",
		"guidance": "Mehnat ka phal zaroor milega. Sabr rakhen.",
	},
	"kumbh": {
		"swami":    "Shani",
		"tatva":    "Vayu",
		"bhagya":   "4, 8",
		"rang":     "Neela",
		"mantra":   "Om Pram Preem Praum Sah Shanaischaraya Namah",
		"guidance": "Naye dost banenge. Samajik karyon mein safalta milegi.",
	},
	"meen": {
		"swami":    "Guru",
		"tatva":    "Jal",
		"bhagya":   "3, 7",
		"rang":     "Peela",
		"mantra":   "Om Gram Greem Graum Sah Gurave Namah",
		"guidance": "Aaj bhavnaatmak din hai. Pooja path se man ko shanti milegi.",
	},
}

// CalculateRashi — basic sun sign calculation from DOB
func CalculateRashi(dob string) string {
	if len(dob) < 10 {
		return "mesh"
	}

	month := dob[5:7]
	day := dob[8:10]

	switch {
	case (month == "03" && day >= "21") || (month == "04" && day <= "19"):
		return "mesh"
	case (month == "04" && day >= "20") || (month == "05" && day <= "20"):
		return "vrishabh"
	case (month == "05" && day >= "21") || (month == "06" && day <= "20"):
		return "mithun"
	case (month == "06" && day >= "21") || (month == "07" && day <= "22"):
		return "kark"
	case (month == "07" && day >= "23") || (month == "08" && day <= "22"):
		return "simha"
	case (month == "08" && day >= "23") || (month == "09" && day <= "22"):
		return "kanya"
	case (month == "09" && day >= "23") || (month == "10" && day <= "22"):
		return "tula"
	case (month == "10" && day >= "23") || (month == "11" && day <= "21"):
		return "vrishchik"
	case (month == "11" && day >= "22") || (month == "12" && day <= "21"):
		return "dhanu"
	case (month == "12" && day >= "22") || (month == "01" && day <= "19"):
		return "makar"
	case (month == "01" && day >= "20") || (month == "02" && day <= "18"):
		return "kumbh"
	default:
		return "meen"
	}
}

// GenerateKundali — main function
func GenerateKundali(req KundaliRequest) *KundaliData {
	rashi := CalculateRashi(req.DOB)

	details, ok := rashiDetails[rashi]
	if !ok {
		details = rashiDetails["mesh"]
	}

	return &KundaliData{
		Name:      req.Name,
		Rashi:     rashi,
		Swami:     details["swami"],
		Tatva:     details["tatva"],
		BhagyaAnk: details["bhagya"],
		ShubhRang: details["rang"],
		Mantra:    details["mantra"],
		Guidance:  details["guidance"],
	}
}
