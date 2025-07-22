// Package models defines the core data structures used throughout the swutrack application.
// It includes types for representing game cards, leaders, and format legality information.
package models

// Card represents a basic game card with identification, name, and legality information.
type Card struct {
	ID    string
	Name  string
	Legal LegalFormats
}

// Leader represents a leader card which extends Card with additional subtitle information.
// It embeds the Card type to inherit its fields and adds leader-specific attributes.
type Leader struct {
	Card

	Subtitle string
}

// LegalFormats indicates which game formats a card is legal to play in.
// Currently tracks Premier format legality status.
type LegalFormats struct {
	Premier bool
}
