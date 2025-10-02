package implementations

import (
	"database/sql"
	"fmt"
	"strings"

	"yaus/services/model"
	"yaus/services/repositories"
)

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base        = uint64(len(base62Chars))
)

// Max value for a 7-character Base62 string is 62^7 - 1.
// This is 3,521,614,606,207.
const maxVal = uint64(3521614606207)

type urlMapRepository struct {
	db *sql.DB
}

func NewUrlMapRepository(db *sql.DB) repositories.UrlMapRepository {
	return &urlMapRepository{db: db}
}

func (r *urlMapRepository) Create(payload model.UrlMapPayload) (string, error) {
	var err error
	var id int64
	var shortedUrl string
	shortedUrl, exists := r.getByLongUrl(payload.LongUrl)
	if exists {
		return shortedUrl, nil
	}
	str := "INSERT INTO url_maps (short_url, long_url) VALUES ($1, $2) RETURNING id"

	err = r.db.QueryRow(str, "", payload.LongUrl).Scan(&id)
	if err != nil {
		return "", err
	}

	shortedUrl, err = toBase62(uint64(id))
	if err != nil {
		return "", err
	}
	_, err = r.db.Exec("UPDATE url_maps SET short_url = $1 WHERE id = $2", shortedUrl, id)
	if err != nil {
		return "", err
	}
	return shortedUrl, err
}

func (r *urlMapRepository) GetByShortUrl(shortUrl string) (*model.UrlMap, error) {
	var longUrl string
	err := r.db.QueryRow("SELECT long_url FROM url_maps WHERE short_url = $1", shortUrl).Scan(&longUrl)
	if err != nil {
		return nil, err
	}
	return &model.UrlMap{
		LongUrl: longUrl,
	}, nil
}

func (r *urlMapRepository) getByLongUrl(longUrl string) (string, bool) {
	var shortUrl string
	err := r.db.QueryRow("SELECT short_url FROM url_maps WHERE long_url = $1", longUrl).Scan(&shortUrl)
	if err != nil {
		return "", false
	}
	return shortUrl, true
}

// ToBase62 converts a uint64 to a Base62 string.
// It returns an error if the input number is too large for 7 characters.
func toBase62(n uint64) (string, error) {
	if n > maxVal {
		return "", fmt.Errorf("input %d is too large for a 7-character Base62 string", n)
	}

	if n == 0 {
		return string(base62Chars[0]), nil
	}

	var sb strings.Builder
	// Reserve space to avoid reallocations. 7 is our max length.
	sb.Grow(7)

	for n > 0 {
		// Get the remainder to find the character index.
		remainder := n % base
		// Append the corresponding character to the string builder.
		sb.WriteByte(base62Chars[remainder])
		// Integer division to prepare for the next digit.
		n /= base
	}

	// The result is built in reverse order, so we need to reverse it.
	return reverse(sb.String()), nil
}

// reverse is a helper function to reverse a string.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
