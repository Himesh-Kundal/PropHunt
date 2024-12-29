/*
  # Create players table and security policies

  1. New Tables
    - `players`
      - `username` (text, primary key)
      - `hashed_password` (text)
      - `kills` (integer)
      - `deaths` (integer)
      - `wins` (integer)
      - `losses` (integer)
      - `draws` (integer)
      - `time_alive` (integer)
      - `created_at` (timestamp)

  2. Security
    - Enable RLS on `players` table
    - Add policies for:
      - Anyone can read player stats
      - Authenticated users can only update their own stats
*/

CREATE TABLE players (
    username TEXT PRIMARY KEY,
    hashed_password TEXT NOT NULL,
    kills INTEGER NOT NULL DEFAULT 0,
    deaths INTEGER NOT NULL DEFAULT 0,
    wins INTEGER NOT NULL DEFAULT 0,
    losses INTEGER NOT NULL DEFAULT 0,
    draws INTEGER NOT NULL DEFAULT 0,
    time_alive INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE players ENABLE ROW LEVEL SECURITY;

-- Allow anyone to read player stats
CREATE POLICY "Anyone can read player stats"
ON players
FOR SELECT
TO public
USING (true);

-- Allow users to update only their own stats
CREATE POLICY "Users can update own stats"
ON players
FOR UPDATE
TO authenticated
USING (auth.uid()::text = username);