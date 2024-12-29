// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING username, hashed_password, kills, deaths, wins, losses, draws, time_alive, created_at, updated_at
`

type CreateUserParams struct {
	Username       string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Kills,
		&i.Deaths,
		&i.Wins,
		&i.Losses,
		&i.Draws,
		&i.TimeAlive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT username, hashed_password, kills, deaths, wins, losses, draws, time_alive, created_at, updated_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.HashedPassword,
			&i.Kills,
			&i.Deaths,
			&i.Wins,
			&i.Losses,
			&i.Draws,
			&i.TimeAlive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT username, hashed_password, kills, deaths, wins, losses, draws, time_alive, created_at, updated_at FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Kills,
		&i.Deaths,
		&i.Wins,
		&i.Losses,
		&i.Draws,
		&i.TimeAlive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserStats = `-- name: UpdateUserStats :one
UPDATE users
SET 
    kills = kills + $2,
    deaths = deaths + $3,
    wins = wins + $4,
    losses = losses + $5,
    draws = draws + $6,
    time_alive = time_alive + $7
WHERE username = $1
RETURNING username, kills, deaths, wins, losses, draws, time_alive
`

type UpdateUserStatsParams struct {
	Username  string
	Kills     int32
	Deaths    int32
	Wins      int32
	Losses    int32
	Draws     int32
	TimeAlive int32
}

type UpdateUserStatsRow struct {
	Username  string
	Kills     int32
	Deaths    int32
	Wins      int32
	Losses    int32
	Draws     int32
	TimeAlive int32
}

func (q *Queries) UpdateUserStats(ctx context.Context, arg UpdateUserStatsParams) (UpdateUserStatsRow, error) {
	row := q.db.QueryRowContext(ctx, updateUserStats,
		arg.Username,
		arg.Kills,
		arg.Deaths,
		arg.Wins,
		arg.Losses,
		arg.Draws,
		arg.TimeAlive,
	)
	var i UpdateUserStatsRow
	err := row.Scan(
		&i.Username,
		&i.Kills,
		&i.Deaths,
		&i.Wins,
		&i.Losses,
		&i.Draws,
		&i.TimeAlive,
	)
	return i, err
}
