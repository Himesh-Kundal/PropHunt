-- name: CreateUser :one
INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUserStats :one
UPDATE users
SET 
    kills = kills + $2,
    deaths = deaths + $3,
    wins = wins + $4,
    losses = losses + $5,
    draws = draws + $6,
    time_alive = time_alive + $7
WHERE username = $1
RETURNING username, kills, deaths, wins, losses, draws, time_alive;