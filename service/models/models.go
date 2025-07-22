package models

type Card struct {
	ID       string
	Name     string
	Subtitle string
	Legal    LegalFormats
}

type LegalFormats struct {
	Premier bool
}
