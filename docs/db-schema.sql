-- swutrack Database Schema

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table for OAuth authentication
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    oauth_provider VARCHAR(50) NOT NULL CHECK (oauth_provider IN ('google', 'discord')),
    oauth_id VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(oauth_provider, oauth_id)
);

-- Create index on email for quick lookups
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_oauth ON users(oauth_provider, oauth_id);

-- Decks table for deck management
CREATE TABLE decks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    leader VARCHAR(255) NOT NULL, -- card ID
    base VARCHAR(255) NOT NULL, -- card ID
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Ensure unique deck name per user
    UNIQUE(user_id, name)
);

-- Indexes for deck queries
CREATE INDEX idx_decks_user_id ON decks(user_id);
CREATE INDEX idx_decks_active ON decks(user_id, is_active);
CREATE INDEX idx_decks_name ON decks(user_id, name);

-- Deck Versions table for deck version management
CREATE TABLE deck_versions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    deck_id UUID NOT NULL REFERENCES decks(id) ON DELETE CASCADE,
    version VARCHAR(20) NOT NULL,
    decklist JSON,
    is_current BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Ensure unique version per deck
    UNIQUE(deck_id, version)
);

-- Indexes for deck version queries
CREATE INDEX idx_deck_versions_deck_id ON deck_versions(deck_id);
CREATE INDEX idx_deck_versions_current ON deck_versions(deck_id, is_current);
CREATE INDEX idx_deck_versions_created ON deck_versions(deck_id, created_at DESC);

-- Games table for tracking game results
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    deck_version_id UUID NOT NULL REFERENCES deck_versions(id) ON DELETE RESTRICT, -- Track specific version played
    opponent_leader VARCHAR(255) NOT NULL,
    opponent_base VARCHAR(255) NOT NULL,
    result VARCHAR(20) NOT NULL CHECK (result IN ('win', 'loss', 'draw')),
    notes TEXT,
    game_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for game queries and analytics
CREATE INDEX idx_games_user_id ON games(user_id);
CREATE INDEX idx_games_deck_version_id ON games(deck_version_id);
CREATE INDEX idx_games_date ON games(user_id, game_date);
CREATE INDEX idx_games_result ON games(user_id, result);
CREATE INDEX idx_games_opponent ON games(user_id, opponent_leader, opponent_base);

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Function to ensure only one current version per deck
CREATE OR REPLACE FUNCTION ensure_single_current_version()
RETURNS TRIGGER AS $$
BEGIN
    -- If setting a version to current, unset all other versions for this deck
    IF NEW.is_current = true THEN
        UPDATE deck_versions 
        SET is_current = false 
        WHERE deck_id = NEW.deck_id AND id != NEW.id;
    END IF;
    
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers to automatically update updated_at
CREATE TRIGGER update_users_updated_at 
    BEFORE UPDATE ON users 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_decks_updated_at 
    BEFORE UPDATE ON decks 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_games_updated_at 
    BEFORE UPDATE ON games 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Trigger to ensure only one current version per deck
CREATE TRIGGER ensure_single_current_version_trigger
    BEFORE INSERT OR UPDATE ON deck_versions
    FOR EACH ROW
    EXECUTE FUNCTION ensure_single_current_version();

-- Comments explaining design decisions:
-- 1. Uses UUIDs for all primary keys to avoid enumeration attacks
-- 2. Separate deck_versions table for better version management
-- 3. JSON decklist storage for flexibility and simplicity
-- 4. Games reference specific deck versions for accurate historical tracking
-- 5. RESTRICT on deck version deletion from games to preserve historical data
-- 6. is_current flag with trigger to ensure only one current version per deck
-- 7. Comprehensive indexing for common query patterns
-- 8. Check constraints for data validation
-- 9. Proper foreign key relationships with appropriate cascade rules
-- 10. Timestamp tracking with automatic updates
-- 11. Support for multiple OAuth providers