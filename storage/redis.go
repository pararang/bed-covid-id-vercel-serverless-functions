package storage

import (
	"api-bed-covid/model"
	"api-bed-covid/utils"
	"context"
	"fmt"
	"os"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

type Redis interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	SetScrapedAvailableHospitals(hashURL, hashHTML string, hospitals []model.HospitalSummary) error
	GetScrapedAvailableHospitals(key string) (string, error)
}

type redis struct {
	client  *goredis.Client
	context context.Context
}

func NewRedis() redis {
	var host = os.Getenv("REDIS_HOST")
	var password = os.Getenv("REDIS_PASSWORD")

	var url = fmt.Sprintf("redis://:%s@%s:30820", password, host) // TODO get port from env var

	opt, _ := goredis.ParseURL(url)
	newClient := goredis.NewClient(opt)

	return redis{
		client:  newClient,
		context: context.Background(),
	}
}

// Get returns the value of the key
func (r redis) Get(key string) (string, error) {
	return r.client.Get(r.context, key).Result()
}

// Set sets the value of the key
func (r redis) Set(key string, value string) error {
	return r.client.Set(r.context, key, value, 0).Err()
}

// SetEx sets the value of the key with an expire time
func (r redis) SetEx(key string, value string, expire time.Duration) error {
	return r.client.Set(r.context, key, value, expire).Err()
}

// SetScrapedAvailableHospitals sets the scraped available hospitals
func (r redis) SetScrapedAvailableHospitals(url string, hospitals []model.HospitalSummary) error {
	var key = buildKeyAvailableHospital(url)
	var value = utils.JSONString(hospitals)
	var expireTime = time.Duration(5 * 60 * time.Second) // TODO: set to env var

	return r.SetEx(key, value, expireTime)
}

// GetScrapedAvailableHospitals sets the scraped available hospitals
func (r redis) GetScrapedAvailableHospitals(url string) (string, error) {
	return r.Get(buildKeyAvailableHospital(url))
}

// buildKeyAvailableHospital returns the key for the scraped available hospitals
func buildKeyAvailableHospital(url string) string {
	const prefix = "available_hospitals"
	return prefix + "." + utils.GetMD5String(url)
}
