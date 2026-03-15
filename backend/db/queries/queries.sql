-- name: GetPanchang :one
SELECT * FROM panchang WHERE date = ?;

-- name: GetBhajanByRashi :one
SELECT * FROM bhajans
WHERE (rashi = ? OR rashi = 'all')
ORDER BY RANDOM()
LIMIT 1;

-- name: CreateKundaliRequest :one
INSERT INTO kundali_requests (name, dob, tob, place, rashi)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (phone, rashi, name)
VALUES (?, ?, ?)
RETURNING *;

-- name: InsertPanchang :exec
INSERT OR REPLACE INTO panchang
(date, vrat, tithi, nakshatra, sunrise, sunset, muhurat, festival)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);