package models

type Card struct {
	ID    string
	Name  string
	Legal LegalFormats
}

type Leader struct {
	Card

	Subtitle string
}

type LegalFormats struct {
	Premier bool
}
