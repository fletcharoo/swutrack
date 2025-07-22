package leaders

import (
	"fmt"
	"swutrack/service/models"
)

var leaderCards = map[string]models.Card{
	"SOR_001": {
		ID:       "SOR_001",
		Name:     "Director Krennic",
		Subtitle: "Aspiring to Authority",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_002": {
		ID:       "SOR_002",
		Name:     "Iden Versio",
		Subtitle: "Inferno Squad Commander",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_003": {
		ID:       "SOR_003",
		Name:     "Chewbacca",
		Subtitle: "Walking Carpet",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_004": {
		ID:       "SOR_004",
		Name:     "Chirrut Imwe",
		Subtitle: "One with the Force",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_005": {
		ID:       "SOR_005",
		Name:     "Luke Skywalker",
		Subtitle: "Faithful Friend",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_006": {
		ID:       "SOR_006",
		Name:     "Emperor Palpatine",
		Subtitle: "Galactic Ruler",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_007": {
		ID:       "SOR_007",
		Name:     "Grand Moff Tarkin",
		Subtitle: "Oversector Governor",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_008": {
		ID:       "SOR_008",
		Name:     "Hera Syndulla",
		Subtitle: "Spectre Two",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_009": {
		ID:       "SOR_009",
		Name:     "Leia Organa",
		Subtitle: "Alliance General",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_010": {
		ID:       "SOR_010",
		Name:     "Darth Vader",
		Subtitle: "Dark Lord of the Sith",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_011": {
		ID:       "SOR_011",
		Name:     "Grand Inquisitor",
		Subtitle: "Hunting the Jedi",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_012": {
		ID:       "SOR_012",
		Name:     "IG-88",
		Subtitle: "Ruthless Bounty Hunter",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_013": {
		ID:       "SOR_013",
		Name:     "Cassian Andor",
		Subtitle: "Dedicated to the Rebellion",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_014": {
		ID:       "SOR_014",
		Name:     "Sabine Wren",
		Subtitle: "Galvanized Revolutionary",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_015": {
		ID:       "SOR_015",
		Name:     "Boba Fett",
		Subtitle: "Collecting the Bounty",
		Legal: models.LegalFormats{
			Premier: false,
		},
	},
	"SOR_016": {
		ID:       "SOR_016",
		Name:     "Grand Admiral Thrawn",
		Subtitle: "Patient and Insightful",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_017": {
		ID:       "SOR_017",
		Name:     "Han Solo",
		Subtitle: "Audacious Smuggler",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SOR_018": {
		ID:       "SOR_018",
		Name:     "Jyn Erso",
		Subtitle: "Resisting Oppression",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_001": {
		ID:       "SHD_001",
		Name:     "Gar Saxon",
		Subtitle: "Viceroy of Mandalore",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_002": {
		ID:       "SHD_002",
		Name:     "Qi'ra",
		Subtitle: "I Alone Survived",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_003": {
		ID:       "SHD_003",
		Name:     "Finn",
		Subtitle: "This Is a Rescue",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_004": {
		ID:       "SHD_004",
		Name:     "Rey",
		Subtitle: "More than a Scavenger",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_005": {
		ID:       "SHD_005",
		Name:     "Hondo Ohnaka",
		Subtitle: "That's Good Business",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_006": {
		ID:       "SHD_006",
		Name:     "Jabba the Hutt",
		Subtitle: "His High Exaltedness",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_007": {
		ID:       "SHD_007",
		Name:     "Moff Gideon",
		Subtitle: "Formidable Commander",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_008": {
		ID:       "SHD_008",
		Name:     "Boba Fett",
		Subtitle: "Daimyo",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_009": {
		ID:       "SHD_009",
		Name:     "Hunter",
		Subtitle: "Outcast Sergeant",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_010": {
		ID:       "SHD_010",
		Name:     "Bossk",
		Subtitle: "Hunting His Prey",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_011": {
		ID:       "SHD_011",
		Name:     "Kylo Ren",
		Subtitle: "Rash and Deadly",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_012": {
		ID:       "SHD_012",
		Name:     "Bo-Katan Kryze",
		Subtitle: "Princess in Exile",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_013": {
		ID:       "SHD_013",
		Name:     "Han Solo",
		Subtitle: "Worth the Risk",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_014": {
		ID:       "SHD_014",
		Name:     "Cad Bane",
		Subtitle: "He Who Needs No Introduction",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_015": {
		ID:       "SHD_015",
		Name:     "Doctor Aphra",
		Subtitle: "Rapacious Archaeologist",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_016": {
		ID:       "SHD_016",
		Name:     "Fennec Shand",
		Subtitle: "Honoring the Deal",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_017": {
		ID:       "SHD_017",
		Name:     "Lando Calrissian",
		Subtitle: "With Impeccable Taste",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"SHD_018": {
		ID:       "SHD_018",
		Name:     "The Mandalorian",
		Subtitle: "Sworn to the Creed",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_001": {
		ID:       "TWI_001",
		Name:     "Nala Se",
		Subtitle: "Clone Engineer",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_002": {
		ID:       "TWI_002",
		Name:     "Nute Gunray",
		Subtitle: "Vindictive Viceroy",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_003": {
		ID:       "TWI_003",
		Name:     "Obi-Wan Kenobi",
		Subtitle: "Patient Mentor",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_004": {
		ID:       "TWI_004",
		Name:     "Yoda",
		Subtitle: "Sensing Darkness",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_005": {
		ID:       "TWI_005",
		Name:     "Count Dooku",
		Subtitle: "Face of the Confederacy",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_006": {
		ID:       "TWI_006",
		Name:     "Wat Tambor",
		Subtitle: "Techno Union Foreman",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_007": {
		ID:       "TWI_007",
		Name:     "Captain Rex",
		Subtitle: "Fighting for His Brothers",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_008": {
		ID:       "TWI_008",
		Name:     "Padme Amidala",
		Subtitle: "Serving the Republic",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_009": {
		ID:       "TWI_009",
		Name:     "Maul",
		Subtitle: "A Rival in Darkness",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_010": {
		ID:       "TWI_010",
		Name:     "Pre Vizsla",
		Subtitle: "Pursuing the Throne",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_011": {
		ID:       "TWI_011",
		Name:     "Ahsoka Tano",
		Subtitle: "Snips",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_012": {
		ID:       "TWI_012",
		Name:     "Anakin Skywalker",
		Subtitle: "What It Takes to Win",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_013": {
		ID:       "TWI_013",
		Name:     "Mace Windu",
		Subtitle: "Vaapad Form Master",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_014": {
		ID:       "TWI_014",
		Name:     "Asajj Ventress",
		Subtitle: "Unparalleled Adversary",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_015": {
		ID:       "TWI_015",
		Name:     "General Grievous",
		Subtitle: "General of the Droid Armies",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_016": {
		ID:       "TWI_016",
		Name:     "Jango Fett",
		Subtitle: "Concealing the Conspiracy",
		Legal: models.LegalFormats{
			Premier: false,
		},
	},
	"TWI_017": {
		ID:       "TWI_017",
		Name:     "Chancellor Palpatine / Darth Sidious",
		Subtitle: "Playing Both Sides",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"TWI_018": {
		ID:       "TWI_018",
		Name:     "Quinlan Vos",
		Subtitle: "Sticking the Landing",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_001": {
		ID:       "JTL_001",
		Name:     "Asajj Ventress",
		Subtitle: "I Work Alone",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_002": {
		ID:       "JTL_002",
		Name:     "Grand Admiral Thrawn",
		Subtitle: "...How Unfortunate",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_003": {
		ID:       "JTL_003",
		Name:     "Lando Calrissian",
		Subtitle: "Buying Time",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_004": {
		ID:       "JTL_004",
		Name:     "Rose Tico",
		Subtitle: "Saving What We Love",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_005": {
		ID:       "JTL_005",
		Name:     "Admiral Piett",
		Subtitle: "Commanding the Armada",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_006": {
		ID:       "JTL_006",
		Name:     "Darth Vader",
		Subtitle: "Victor Squadron Leader",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_007": {
		ID:       "JTL_007",
		Name:     "Admiral Holdo",
		Subtitle: "We're Not Alone",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_008": {
		ID:       "JTL_008",
		Name:     "Wedge Antilles",
		Subtitle: "Leader of Red Squadron",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_009": {
		ID:       "JTL_009",
		Name:     "Boba Fett",
		Subtitle: "Any Methods Necessary",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_010": {
		ID:       "JTL_010",
		Name:     "Captain Phasma",
		Subtitle: "Chrome Dome",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_011": {
		ID:       "JTL_011",
		Name:     "Major Vonreg",
		Subtitle: "Red Baron",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_012": {
		ID:       "JTL_012",
		Name:     "Luke Skywalker",
		Subtitle: "Hero of Yavin",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_013": {
		ID:       "JTL_013",
		Name:     "Poe Dameron",
		Subtitle: "I Can Fly Anything",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_014": {
		ID:       "JTL_014",
		Name:     "Admiral Trench",
		Subtitle: "Chk-Chk-Chk-Chk",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_015": {
		ID:       "JTL_015",
		Name:     "Rio Durant",
		Subtitle: "Wisecracking Wheelman",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_016": {
		ID:       "JTL_016",
		Name:     "Admiral Ackbar",
		Subtitle: "It's a Trap!",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_017": {
		ID:       "JTL_017",
		Name:     "Han Solo",
		Subtitle: "Never Tell Me the Odds",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"JTL_018": {
		ID:       "JTL_018",
		Name:     "Kazuda Xiono",
		Subtitle: "Best Pilot in the Galaxy",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_001": {
		ID:       "LOF_001",
		Name:     "Kylo Ren",
		Subtitle: "We're Not Done Yet",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_002": {
		ID:       "LOF_002",
		Name:     "Mother Talzin",
		Subtitle: "Power Through Magick",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_003": {
		ID:       "LOF_003",
		Name:     "Ahsoka Tano",
		Subtitle: "Fighting for Peace",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_004": {
		ID:       "LOF_004",
		Name:     "Kanan Jarrus",
		Subtitle: "Help Us Survive",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_005": {
		ID:       "LOF_005",
		Name:     "Morgan Elsbeth",
		Subtitle: "Following the Call",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_006": {
		ID:       "LOF_006",
		Name:     "Supreme Leader Snoke",
		Subtitle: "In the Seat of Power",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_007": {
		ID:       "LOF_007",
		Name:     "Avar Kriss",
		Subtitle: "Marshal of Starlight",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_008": {
		ID:       "LOF_008",
		Name:     "Obi-Wan Kenobi",
		Subtitle: "Courage Makes Heroes",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_009": {
		ID:       "LOF_009",
		Name:     "Darth Maul",
		Subtitle: "Sith Revealed",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_010": {
		ID:       "LOF_010",
		Name:     "Third Sister",
		Subtitle: "Seething with Ambition",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_011": {
		ID:       "LOF_011",
		Name:     "Kit Fisto",
		Subtitle: "Focused Jedi Master",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_012": {
		ID:       "LOF_012",
		Name:     "Rey",
		Subtitle: "Nobody",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_013": {
		ID:       "LOF_013",
		Name:     "Barriss Offee",
		Subtitle: "We Have Become Villains",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_014": {
		ID:       "LOF_014",
		Name:     "Grand Inquisitor",
		Subtitle: "Stories Travel Quickly",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_015": {
		ID:       "LOF_015",
		Name:     "Cal Kestis",
		Subtitle: "I Can't Keep Hiding",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_016": {
		ID:       "LOF_016",
		Name:     "Qui-Gon Jinn",
		Subtitle: "Student of the Living Force",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_017": {
		ID:       "LOF_017",
		Name:     "Darth Revan",
		Subtitle: "Scourge of the Old Republic",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
	"LOF_018": {
		ID:       "LOF_018",
		Name:     "Anakin Skywalker",
		Subtitle: "Tempted by the Dark Side",
		Legal: models.LegalFormats{
			Premier: true,
		},
	},
}

// GetByID retrieves a leader card by its unique identifier.
func GetByID(id string) (card models.Card, err error) {
	if id == "" {
		err = fmt.Errorf("card ID cannot be empty")
		return
	}

	var exists bool
	card, exists = leaderCards[id]
	if !exists {
		err = fmt.Errorf("card with ID %q not found", id)
		return
	}

	return card, nil
}

// GetAll returns all leader cards as a slice.
func GetAll() (cards []models.Card) {
	cards = make([]models.Card, 0, len(leaderCards))

	for _, card := range leaderCards {
		cards = append(cards, card)
	}

	return cards
}
