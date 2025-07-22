package leaders

import (
	"fmt"
	"testing"

	"swutrack/service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetByID_ValidCardReturnsCard(t *testing.T) {
	// Setup expectations
	expectedID := "SOR_001"
	expectedCard := models.Card{
		ID:       "SOR_001",
		Name:     "Director Krennic",
		Subtitle: "Aspiring to Authority",
		Legal: models.LegalFormats{
			Premier: true,
		},
	}

	// Act
	card, err := GetByID(expectedID)

	// Assert
	require.NoError(t, err, "GetByID should not return an error for valid card ID")
	assert.Equal(t, expectedCard, card, "returned card should match expected")
}

func Test_GetByID_DifferentValidCardReturnsCorrectCard(t *testing.T) {
	// Setup expectations
	expectedID := "LOF_018"
	expectedCard := models.Card{
		ID:       "LOF_018",
		Name:     "Anakin Skywalker",
		Subtitle: "Tempted by the Dark Side",
		Legal: models.LegalFormats{
			Premier: true,
		},
	}

	// Act
	card, err := GetByID(expectedID)

	// Assert
	require.NoError(t, err, "GetByID should not return an error for valid card ID")
	assert.Equal(t, expectedCard, card, "returned card should match expected")
}

func Test_GetByID_EmptyIDReturnsError(t *testing.T) {
	// Setup expectations
	emptyID := ""
	expectedErrorMsg := "card ID cannot be empty"

	// Act
	card, err := GetByID(emptyID)

	// Assert
	require.Error(t, err, "GetByID should return an error for empty ID")
	assert.Equal(t, expectedErrorMsg, err.Error(), "error message should match expected")
	assert.Empty(t, card, "card should be empty when error is returned")
}

func Test_GetByID_NonExistentIDReturnsError(t *testing.T) {
	// Setup expectations
	nonExistentID := "INVALID_ID"
	expectedErrorMsg := `card with ID "INVALID_ID" not found`

	// Act
	card, err := GetByID(nonExistentID)

	// Assert
	require.Error(t, err, "GetByID should return an error for non-existent ID")
	assert.Equal(t, expectedErrorMsg, err.Error(), "error message should match expected")
	assert.Empty(t, card, "card should be empty when error is returned")
}

func Test_GetByID_CaseSensitiveCheck(t *testing.T) {
	// Setup expectations
	lowercaseID := "sor_001"
	expectedErrorMsg := `card with ID "sor_001" not found`

	// Act
	card, err := GetByID(lowercaseID)

	// Assert
	require.Error(t, err, "GetByID should return an error for wrong case ID")
	assert.Equal(t, expectedErrorMsg, err.Error(), "error message should match expected")
	assert.Empty(t, card, "card should be empty when error is returned")
}

func Test_GetAll(t *testing.T) {
	// Act
	cards := GetAll()

	// Assert basic properties
	assert.NotEmpty(t, cards, "GetAll should return cards")
	assert.Len(t, cards, len(leaderCards), "GetAll should return all cards from the map")

	// Create a map to track unique IDs
	seenIDs := make(map[string]bool)

	// Verify each card
	for _, card := range cards {
		// Check that the card has required fields
		assert.NotEmpty(t, card.ID, "card ID should not be empty")
		assert.NotEmpty(t, card.Name, "card name should not be empty")
		assert.NotEmpty(t, card.Subtitle, "card subtitle should not be empty")

		// Check for duplicate IDs in the result
		assert.False(t, seenIDs[card.ID], fmt.Sprintf("card ID %q should not be duplicated in results", card.ID))
		seenIDs[card.ID] = true

		// Verify the card exists in the original map
		originalCard, exists := leaderCards[card.ID]
		assert.True(t, exists, fmt.Sprintf("card with ID %q should exist in leaderCards map", card.ID))
		assert.Equal(t, originalCard, card, fmt.Sprintf("card with ID %q should match the original", card.ID))
	}

	// Verify we got all cards from the map
	assert.Equal(t, len(leaderCards), len(seenIDs), "should have retrieved all unique cards from the map")
}

func Test_GetAll_ReturnsNewSlice(t *testing.T) {
	// Act - get cards twice
	cards1 := GetAll()
	cards2 := GetAll()

	// Assert - verify they are different slices
	assert.NotSame(t, &cards1, &cards2, "GetAll should return a new slice each time")

	// Verify content is the same (convert to maps for order-independent comparison)
	cards1Map := make(map[string]models.Card)
	cards2Map := make(map[string]models.Card)
	for _, card := range cards1 {
		cards1Map[card.ID] = card
	}
	for _, card := range cards2 {
		cards2Map[card.ID] = card
	}
	assert.Equal(t, cards1Map, cards2Map, "GetAll should return the same content each time")

	// Verify modifying one doesn't affect the other
	if len(cards1) > 0 {
		// Store original first card
		originalID := cards1[0].ID
		originalCard, exists := leaderCards[originalID]
		require.True(t, exists, "original card should exist in leaderCards")

		// Modify the returned slice
		cards1[0] = models.Card{ID: "MODIFIED"}

		// Get fresh data and verify it wasn't affected
		cards3 := GetAll()
		cards3Map := make(map[string]models.Card)
		for _, card := range cards3 {
			cards3Map[card.ID] = card
		}

		// Verify the original card is still unchanged in new results
		assert.Equal(t, originalCard, cards3Map[originalID], "modifying returned slice should not affect source data")
	}
}

func Test_LeaderCardsData(t *testing.T) {
	// Verify the leader cards data structure
	assert.NotEmpty(t, leaderCards, "leaderCards map should not be empty")

	// Track sets and check for required sets
	setCounts := make(map[string]int)

	for id, card := range leaderCards {
		// Verify ID matches the map key
		assert.Equal(t, id, card.ID, fmt.Sprintf("card ID should match map key for %q", id))

		// Extract set prefix (first 3 characters)
		if len(id) >= 3 {
			setPrefix := id[:3]
			setCounts[setPrefix]++
		}

		// Verify all required fields are populated
		assert.NotEmpty(t, card.ID, fmt.Sprintf("card %q should have an ID", id))
		assert.NotEmpty(t, card.Name, fmt.Sprintf("card %q should have a name", id))
		assert.NotEmpty(t, card.Subtitle, fmt.Sprintf("card %q should have a subtitle", id))
	}

	// Verify we have cards from expected sets
	expectedSets := []string{"SOR", "SHD", "TWI", "JTL", "LOF"}
	for _, set := range expectedSets {
		assert.Greater(t, setCounts[set], 0, fmt.Sprintf("should have cards from set %q", set))
		assert.Equal(t, 18, setCounts[set], fmt.Sprintf("should have exactly 18 cards from set %q", set))
	}
}

func Test_PremierLegality(t *testing.T) {
	// Known banned cards based on the data
	bannedCards := map[string]bool{
		"SOR_015": true, // Boba Fett - Collecting the Bounty
		"TWI_016": true, // Jango Fett - Concealing the Conspiracy
	}

	premierLegalCount := 0
	for id, card := range leaderCards {
		if bannedCards[id] {
			assert.False(t, card.Legal.Premier, fmt.Sprintf("card %q (%s) should not be Premier legal", id, card.Name))
		} else {
			// Most cards should be Premier legal
			if card.Legal.Premier {
				premierLegalCount++
			}
		}
	}

	// Verify most cards are Premier legal
	totalCards := len(leaderCards)
	bannedCount := len(bannedCards)
	expectedLegalCount := totalCards - bannedCount
	assert.Equal(t, expectedLegalCount, premierLegalCount, "most cards should be Premier legal except for known banned cards")
}
