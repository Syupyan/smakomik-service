package helpers

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenU(c *fiber.Ctx) (string, string, error) {
	// Ambil token dari variabel lokal
	user := c.Locals("user").(*jwt.Token)

	// Verifikasi token
	if user == nil {
		return "", "", fiber.ErrUnauthorized
	}

	// Ambil klaim-klaim dari token
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fiber.ErrUnauthorized
	}

	// Ambil name dan role dari klaim-klaim
	name, nameOk := claims["name"].(string)
	role, roleOk := claims["role"].(string)

	// Periksa apakah klaim name dan role ditemukan
	if !nameOk || !roleOk {
		return "", "", fiber.ErrUnauthorized
	}

	return name, role, nil
}

func VerifyToken(c *fiber.Ctx) (string, error) {
	// Ambil token dari variabel lokal
	user := c.Locals("user").(*jwt.Token)

	// Verifikasi token
	if user == nil {
		return "", fiber.ErrUnauthorized
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", fiber.ErrUnauthorized
	}

	// Ambil name dan role dari klaim-klaim
	role, roleOk := claims["role"].(string)

	// Periksa apakah klaim name dan role ditemukan
	if !roleOk {
		return "", fiber.ErrUnauthorized
	}

	return role, nil
}

func DeleteFileWithRetry(filePath string) error {
	// Cek apakah file ada sebelum mencoba menghapusnya
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File tidak ada, tidak perlu dihapus
	}

	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := os.Remove(filePath)
		if err == nil || os.IsNotExist(err) {
			return nil // Berhasil menghapus file atau file tidak ada
		}
		// Jika gagal menghapus file, tunggu sebentar sebelum mencoba lagi
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	return fmt.Errorf("failed to delete file after %d attempts", maxAttempts)
}

type RateLimiter struct {
    mu           sync.Mutex
    comicAccess  sync.Map // comicID -> IP -> accessed
}

func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        comicAccess: sync.Map{},
    }
}

func (r *RateLimiter) HasAccessedComic(comicID, ip string) bool {
    comicIDIPKey := comicID + "_" + ip
    _, exists := r.comicAccess.Load(comicIDIPKey)
    return exists
}

func (r *RateLimiter) MarkComicAccess(comicID, ip string) {
    comicIDIPKey := comicID + "_" + ip
    r.comicAccess.Store(comicIDIPKey, true)
}