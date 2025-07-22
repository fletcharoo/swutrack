package leaders

import (
	"fmt"
	"swutrack/service/models"
)

var leaderCards = map[string]models.Leader{
	"SOR_001": {
		Card: models.Card{
			ID:   "SOR_001",
			Name: "Director Krennic",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Aspiring to Authority",
	},
	"SOR_002": {
		Card: models.Card{
			ID:   "SOR_002",
			Name: "Iden Versio",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Inferno Squad Commander",
	},
	"SOR_003": {
		Card: models.Card{
			ID:   "SOR_003",
			Name: "Chewbacca",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Walking Carpet",
	},
	"SOR_004": {
		Card: models.Card{
			ID:   "SOR_004",
			Name: "Chirrut Imwe",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "One with the Force",
	},
	"SOR_005": {
		Card: models.Card{
			ID:   "SOR_005",
			Name: "Luke Skywalker",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Faithful Friend",
	},
	"SOR_006": {
		Card: models.Card{
			ID:   "SOR_006",
			Name: "Emperor Palpatine",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Galactic Ruler",
	},
	"SOR_007": {
		Card: models.Card{
			ID:   "SOR_007",
			Name: "Grand Moff Tarkin",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Oversector Governor",
	},
	"SOR_008": {
		Card: models.Card{
			ID:   "SOR_008",
			Name: "Hera Syndulla",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Spectre Two",
	},
	"SOR_009": {
		Card: models.Card{
			ID:   "SOR_009",
			Name: "Leia Organa",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Alliance General",
	},
	"SOR_010": {
		Card: models.Card{
			ID:   "SOR_010",
			Name: "Darth Vader",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Dark Lord of the Sith",
	},
	"SOR_011": {
		Card: models.Card{
			ID:   "SOR_011",
			Name: "Grand Inquisitor",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Hunting the Jedi",
	},
	"SOR_012": {
		Card: models.Card{
			ID:   "SOR_012",
			Name: "IG-88",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Ruthless Bounty Hunter",
	},
	"SOR_013": {
		Card: models.Card{
			ID:   "SOR_013",
			Name: "Cassian Andor",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Dedicated to the Rebellion",
	},
	"SOR_014": {
		Card: models.Card{
			ID:   "SOR_014",
			Name: "Sabine Wren",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Galvanized Revolutionary",
	},
	"SOR_015": {
		Card: models.Card{
			ID:   "SOR_015",
			Name: "Boba Fett",
			Legal: models.LegalFormats{
				Premier: false,
			},
		},
		Subtitle: "Collecting the Bounty",
	},
	"SOR_016": {
		Card: models.Card{
			ID:   "SOR_016",
			Name: "Grand Admiral Thrawn",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Patient and Insightful",
	},
	"SOR_017": {
		Card: models.Card{
			ID:   "SOR_017",
			Name: "Han Solo",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Audacious Smuggler",
	},
	"SOR_018": {
		Card: models.Card{
			ID:   "SOR_018",
			Name: "Jyn Erso",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Resisting Oppression",
	},
	"SHD_001": {
		Card: models.Card{
			ID:   "SHD_001",
			Name: "Gar Saxon",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Viceroy of Mandalore",
	},
	"SHD_002": {
		Card: models.Card{
			ID:   "SHD_002",
			Name: "Qi'ra",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "I Alone Survived",
	},
	"SHD_003": {
		Card: models.Card{
			ID:   "SHD_003",
			Name: "Finn",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "This Is a Rescue",
	},
	"SHD_004": {
		Card: models.Card{
			ID:   "SHD_004",
			Name: "Rey",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "More than a Scavenger",
	},
	"SHD_005": {
		Card: models.Card{
			ID:   "SHD_005",
			Name: "Hondo Ohnaka",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "That's Good Business",
	},
	"SHD_006": {
		Card: models.Card{
			ID:   "SHD_006",
			Name: "Jabba the Hutt",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "His High Exaltedness",
	},
	"SHD_007": {
		Card: models.Card{
			ID:   "SHD_007",
			Name: "Moff Gideon",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Formidable Commander",
	},
	"SHD_008": {
		Card: models.Card{
			ID:   "SHD_008",
			Name: "Boba Fett",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Daimyo",
	},
	"SHD_009": {
		Card: models.Card{
			ID:   "SHD_009",
			Name: "Hunter",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Outcast Sergeant",
	},
	"SHD_010": {
		Card: models.Card{
			ID:   "SHD_010",
			Name: "Bossk",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Hunting His Prey",
	},
	"SHD_011": {
		Card: models.Card{
			ID:   "SHD_011",
			Name: "Kylo Ren",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Rash and Deadly",
	},
	"SHD_012": {
		Card: models.Card{
			ID:   "SHD_012",
			Name: "Bo-Katan Kryze",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Princess in Exile",
	},
	"SHD_013": {
		Card: models.Card{
			ID:   "SHD_013",
			Name: "Han Solo",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Worth the Risk",
	},
	"SHD_014": {
		Card: models.Card{
			ID:   "SHD_014",
			Name: "Cad Bane",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "He Who Needs No Introduction",
	},
	"SHD_015": {
		Card: models.Card{
			ID:   "SHD_015",
			Name: "Doctor Aphra",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Rapacious Archaeologist",
	},
	"SHD_016": {
		Card: models.Card{
			ID:   "SHD_016",
			Name: "Fennec Shand",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Honoring the Deal",
	},
	"SHD_017": {
		Card: models.Card{
			ID:   "SHD_017",
			Name: "Lando Calrissian",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "With Impeccable Taste",
	},
	"SHD_018": {
		Card: models.Card{
			ID:   "SHD_018",
			Name: "The Mandalorian",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Sworn to the Creed",
	},
	"TWI_001": {
		Card: models.Card{
			ID:   "TWI_001",
			Name: "Nala Se",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Clone Engineer",
	},
	"TWI_002": {
		Card: models.Card{
			ID:   "TWI_002",
			Name: "Nute Gunray",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Vindictive Viceroy",
	},
	"TWI_003": {
		Card: models.Card{
			ID:   "TWI_003",
			Name: "Obi-Wan Kenobi",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Patient Mentor",
	},
	"TWI_004": {
		Card: models.Card{
			ID:   "TWI_004",
			Name: "Yoda",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Sensing Darkness",
	},
	"TWI_005": {
		Card: models.Card{
			ID:   "TWI_005",
			Name: "Count Dooku",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Face of the Confederacy",
	},
	"TWI_006": {
		Card: models.Card{
			ID:   "TWI_006",
			Name: "Wat Tambor",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Techno Union Foreman",
	},
	"TWI_007": {
		Card: models.Card{
			ID:   "TWI_007",
			Name: "Captain Rex",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Fighting for His Brothers",
	},
	"TWI_008": {
		Card: models.Card{
			ID:   "TWI_008",
			Name: "Padme Amidala",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Serving the Republic",
	},
	"TWI_009": {
		Card: models.Card{
			ID:   "TWI_009",
			Name: "Maul",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "A Rival in Darkness",
	},
	"TWI_010": {
		Card: models.Card{
			ID:   "TWI_010",
			Name: "Pre Vizsla",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Pursuing the Throne",
	},
	"TWI_011": {
		Card: models.Card{
			ID:   "TWI_011",
			Name: "Ahsoka Tano",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Snips",
	},
	"TWI_012": {
		Card: models.Card{
			ID:   "TWI_012",
			Name: "Anakin Skywalker",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "What It Takes to Win",
	},
	"TWI_013": {
		Card: models.Card{
			ID:   "TWI_013",
			Name: "Mace Windu",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Vaapad Form Master",
	},
	"TWI_014": {
		Card: models.Card{
			ID:   "TWI_014",
			Name: "Asajj Ventress",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Unparalleled Adversary",
	},
	"TWI_015": {
		Card: models.Card{
			ID:   "TWI_015",
			Name: "General Grievous",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "General of the Droid Armies",
	},
	"TWI_016": {
		Card: models.Card{
			ID:   "TWI_016",
			Name: "Jango Fett",
			Legal: models.LegalFormats{
				Premier: false,
			},
		},
		Subtitle: "Concealing the Conspiracy",
	},
	"TWI_017": {
		Card: models.Card{
			ID:   "TWI_017",
			Name: "Chancellor Palpatine / Darth Sidious",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Playing Both Sides",
	},
	"TWI_018": {
		Card: models.Card{
			ID:   "TWI_018",
			Name: "Quinlan Vos",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Sticking the Landing",
	},
	"JTL_001": {
		Card: models.Card{
			ID:   "JTL_001",
			Name: "Asajj Ventress",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "I Work Alone",
	},
	"JTL_002": {
		Card: models.Card{
			ID:   "JTL_002",
			Name: "Grand Admiral Thrawn",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "...How Unfortunate",
	},
	"JTL_003": {
		Card: models.Card{
			ID:   "JTL_003",
			Name: "Lando Calrissian",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Buying Time",
	},
	"JTL_004": {
		Card: models.Card{
			ID:   "JTL_004",
			Name: "Rose Tico",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Saving What We Love",
	},
	"JTL_005": {
		Card: models.Card{
			ID:   "JTL_005",
			Name: "Admiral Piett",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Commanding the Armada",
	},
	"JTL_006": {
		Card: models.Card{
			ID:   "JTL_006",
			Name: "Darth Vader",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Victor Squadron Leader",
	},
	"JTL_007": {
		Card: models.Card{
			ID:   "JTL_007",
			Name: "Admiral Holdo",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "We're Not Alone",
	},
	"JTL_008": {
		Card: models.Card{
			ID:   "JTL_008",
			Name: "Wedge Antilles",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Leader of Red Squadron",
	},
	"JTL_009": {
		Card: models.Card{
			ID:   "JTL_009",
			Name: "Boba Fett",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Any Methods Necessary",
	},
	"JTL_010": {
		Card: models.Card{
			ID:   "JTL_010",
			Name: "Captain Phasma",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Chrome Dome",
	},
	"JTL_011": {
		Card: models.Card{
			ID:   "JTL_011",
			Name: "Major Vonreg",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Red Baron",
	},
	"JTL_012": {
		Card: models.Card{
			ID:   "JTL_012",
			Name: "Luke Skywalker",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Hero of Yavin",
	},
	"JTL_013": {
		Card: models.Card{
			ID:   "JTL_013",
			Name: "Poe Dameron",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "I Can Fly Anything",
	},
	"JTL_014": {
		Card: models.Card{
			ID:   "JTL_014",
			Name: "Admiral Trench",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Chk-Chk-Chk-Chk",
	},
	"JTL_015": {
		Card: models.Card{
			ID:   "JTL_015",
			Name: "Rio Durant",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Wisecracking Wheelman",
	},
	"JTL_016": {
		Card: models.Card{
			ID:   "JTL_016",
			Name: "Admiral Ackbar",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "It's a Trap!",
	},
	"JTL_017": {
		Card: models.Card{
			ID:   "JTL_017",
			Name: "Han Solo",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Never Tell Me the Odds",
	},
	"JTL_018": {
		Card: models.Card{
			ID:   "JTL_018",
			Name: "Kazuda Xiono",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Best Pilot in the Galaxy",
	},
	"LOF_001": {
		Card: models.Card{
			ID:   "LOF_001",
			Name: "Kylo Ren",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "We're Not Done Yet",
	},
	"LOF_002": {
		Card: models.Card{
			ID:   "LOF_002",
			Name: "Mother Talzin",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Power Through Magick",
	},
	"LOF_003": {
		Card: models.Card{
			ID:   "LOF_003",
			Name: "Ahsoka Tano",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Fighting for Peace",
	},
	"LOF_004": {
		Card: models.Card{
			ID:   "LOF_004",
			Name: "Kanan Jarrus",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Help Us Survive",
	},
	"LOF_005": {
		Card: models.Card{
			ID:   "LOF_005",
			Name: "Morgan Elsbeth",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Following the Call",
	},
	"LOF_006": {
		Card: models.Card{
			ID:   "LOF_006",
			Name: "Supreme Leader Snoke",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "In the Seat of Power",
	},
	"LOF_007": {
		Card: models.Card{
			ID:   "LOF_007",
			Name: "Avar Kriss",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Marshal of Starlight",
	},
	"LOF_008": {
		Card: models.Card{
			ID:   "LOF_008",
			Name: "Obi-Wan Kenobi",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Courage Makes Heroes",
	},
	"LOF_009": {
		Card: models.Card{
			ID:   "LOF_009",
			Name: "Darth Maul",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Sith Revealed",
	},
	"LOF_010": {
		Card: models.Card{
			ID:   "LOF_010",
			Name: "Third Sister",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Seething with Ambition",
	},
	"LOF_011": {
		Card: models.Card{
			ID:   "LOF_011",
			Name: "Kit Fisto",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Focused Jedi Master",
	},
	"LOF_012": {
		Card: models.Card{
			ID:   "LOF_012",
			Name: "Rey",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Nobody",
	},
	"LOF_013": {
		Card: models.Card{
			ID:   "LOF_013",
			Name: "Barriss Offee",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "We Have Become Villains",
	},
	"LOF_014": {
		Card: models.Card{
			ID:   "LOF_014",
			Name: "Grand Inquisitor",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Stories Travel Quickly",
	},
	"LOF_015": {
		Card: models.Card{
			ID:   "LOF_015",
			Name: "Cal Kestis",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "I Can't Keep Hiding",
	},
	"LOF_016": {
		Card: models.Card{
			ID:   "LOF_016",
			Name: "Qui-Gon Jinn",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Student of the Living Force",
	},
	"LOF_017": {
		Card: models.Card{
			ID:   "LOF_017",
			Name: "Darth Revan",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Scourge of the Old Republic",
	},
	"LOF_018": {
		Card: models.Card{
			ID:   "LOF_018",
			Name: "Anakin Skywalker",
			Legal: models.LegalFormats{
				Premier: true,
			},
		},
		Subtitle: "Tempted by the Dark Side",
	},
}

// GetByID retrieves a leader card by its unique identifier.
func GetByID(id string) (leader models.Leader, err error) {
	if id == "" {
		err = fmt.Errorf("card ID cannot be empty")
		return
	}
	var exists bool
	leader, exists = leaderCards[id]
	if !exists {
		err = fmt.Errorf("card with ID %q not found", id)
		return
	}
	return leader, nil
}

// GetAll returns all leader cards as a slice.
func GetAll() (leaders []models.Leader) {
	leaders = make([]models.Leader, 0, len(leaderCards))
	for _, leader := range leaderCards {
		leaders = append(leaders, leader)
	}
	return leaders
}
