# Star Wars Unlimited - Domain Knowledge

This document provides essential context about the Star Wars Unlimited trading card game to help with application development.

## Game Overview

Star Wars Unlimited is a competitive trading card game where players build decks and battle against each other using characters, vehicles, and events from the Star Wars universe. Each player needs a deck to play, and tracking performance across different decks and matchups is crucial for competitive improvement.

## Core Game Components

### Leaders
- **Definition**: Powerful character cards that define your deck's strategy and provide ongoing abilities
- **Examples**: Luke Skywalker, Darth Vader, Princess Leia, Emperor Palpatine
- **Deck Building**: Each deck must have exactly one leader
- **Gameplay Impact**: Leaders start in play and provide consistent effects throughout the game
- **Identity**: The leader fundamentally defines what type of deck you're playing

### Bases
- **Definition**: Location cards that provide resources and strategic advantages
- **Examples**: Echo Base, Death Star, Mos Eisley Cantina, Jabba's Palace
- **Deck Building**: Each deck must have exactly one base
- **Gameplay Impact**: Bases provide ongoing effects and resource generation
- **Identity**: Combined with leader, the base helps define your deck's strategic approach

### Main Deck Cards
- **Definition**: The primary cards used during gameplay (typically 50 cards)
- **Types**:
  - **Units**: Characters, vehicles, and creatures that can attack and defend
  - **Events**: One-time effects that provide immediate impact
  - **Upgrades**: Attachments that enhance units with additional abilities
- **Deck Building Rules**: Usually limited quantities per card (typically 3 copies max)

### Sideboard
- **Definition**: Additional cards that can be swapped in between games in competitive play
- **Purpose**: Allows players to adapt their deck against specific opponents or strategies
- **Typical Size**: 10 cards
- **Usage**: In best-of-three matches, players can modify their deck using sideboard cards

## Deck Building Concepts

### Deck Identity
- **Leader + Base Combination**: This pairing defines the core strategy and theme of the deck
- **Archetype Examples**: 
  - Luke Skywalker + Echo Base = Rebel aggro deck
  - Darth Vader + Death Star = Imperial control deck
- **Why Immutable**: Changing the leader or base fundamentally changes what deck you're playing

### Deck Versioning
- **Purpose**: Track iterations and improvements to the same core deck concept
- **When to Version**: 
  - Meta shifts requiring card changes
  - Performance analysis leading to optimization
  - New card releases enabling new strategies
- **Version Continuity**: All versions share the same leader/base combination

### Deck Names
- **Player Choice**: Players create custom names for their decks
- **Common Patterns**: 
  - Character-based: "Luke Aggro", "Vader Control"
  - Strategy-based: "Rebel Rush", "Imperial Midrange"
  - Meta-based: "Anti-Aggro Build", "Tournament Special"

## Competitive Play Context

### Game Results
- **Win/Loss/Draw**: Standard competitive outcomes
- **Best of Three**: Most competitive play uses this format
- **Tournament Structure**: Players face multiple opponents in events

### Matchup Analysis
- **Leader vs Leader**: How your leader performs against different opposing leaders
- **Base Interactions**: How base combinations affect resource races and tempo
- **Archetype Matchups**: Aggro vs Control, Midrange vs Combo, etc.
- **Meta Considerations**: Popular deck choices in the current competitive environment

### Performance Tracking Importance
- **Deck Optimization**: Understanding which cards perform well
- **Matchup Preparation**: Knowing your weak matchups for sideboard planning
- **Meta Analysis**: Tracking what opponents are playing
- **Tournament Preparation**: Historical data helps with deck selection

## Application Context

### Why Players Need Tracking
- **Deck Performance**: Which versions of decks work best
- **Matchup Data**: Win rates against specific leader/base combinations
- **Meta Insights**: What decks are being played most
- **Tournament Prep**: Historical performance for event planning
- **Improvement**: Data-driven deck building decisions

### User Types
- **Competitive Players**: Need detailed analytics for tournament preparation
- **Casual Players**: Want to track improvement and deck performance
- **Deck Builders**: Experiment with different iterations and need version tracking
- **Local Game Store Players**: Track performance in weekly events

### Key Metrics Players Care About
- **Overall Win Rate**: General performance indicator
- **Deck-Specific Win Rates**: How each deck performs
- **Matchup Win Rates**: Performance against specific leaders/bases
- **Recent Performance**: Trends and current form
- **Games Played**: Activity levels and data reliability
- **Version Comparison**: How deck iterations compare

## Data Considerations

### Card Identification
- **Card IDs**: Use unique identifiers rather than names for data consistency
- **Future Proofing**: Card databases may be integrated later
- **Localization**: Card IDs work across different language versions

### Deck List Storage
- **JSON Format**: Flexible structure for card lists and sideboards
- **Card Quantities**: Track how many copies of each card
- **Card Metadata**: Type, cost, and other relevant properties for analysis

### Historical Data Preservation
- **Game Records**: Must maintain integrity even if decks are "deleted"
- **Version History**: Complete audit trail of deck changes
- **Performance Trends**: Long-term data for meaningful analysis

## Business Rules

### Deck Creation
- **Required Fields**: Name, leader, base, initial decklist
- **Automatic Versioning**: Start at version 1.0
- **Unique Names**: Per user, prevent confusion

### Deck Updates
- **Immutable Identity**: Leader and base cannot change
- **Mutable Metadata**: Name and description can be updated
- **Version Creation**: New decklists create new versions

### Game Recording
- **Required Data**: Deck version used, opponent leader/base, result
- **Optional Data**: Notes, custom game date
- **Data Integrity**: Cannot delete versions with recorded games

### Version Management
- **Current Version**: One active version per deck
- **Historical Versions**: Preserved for analysis
- **Automatic Numbering**: System controls version increments