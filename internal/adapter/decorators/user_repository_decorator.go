package decorators

import (
	"encoding/json"
	"file_manager/configs"
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/common/caching"
	"file_manager/internal/common/log"
	"file_manager/internal/enums"
	"file_manager/internal/models"
	"fmt"
	"time"
)

type UserRepositoryDecorator struct {
	cache          caching.CacheStrategy
	userRepository *repositories.UserQueryRepository
	expCacheTime   time.Duration
	setKeyEnv      string
}

func NewUserRepositoryDecorator(
	cache caching.CacheStrategy,
	userRepository *repositories.UserQueryRepository,
	env string,
) *UserRepositoryDecorator {
	return &UserRepositoryDecorator{
		cache:          cache,
		userRepository: userRepository,
		expCacheTime:   configs.Get().ExpCacheTimeDb,
		setKeyEnv:      fmt.Sprintf("%s_%s", enums.DefaultSetKeyDB, env),
	}
}

func (d *UserRepositoryDecorator) FindByUsername(username string) (*models.User, error) {
	key := d.generateCachingKeyDB(enums.DefaultNameSpace, d.setKeyEnv, username)
	data, found := d.cache.Get(key)
	if found {
		user, err := d.getUserModelFromRemoteCachedData(data)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	userModel, err := d.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	cachingData, err := d.parseMapFromUser(userModel)
	if err != nil {
		log.Errorf("Can not parseMapFrom newsModels for cache with error: [%s]", err)
		return userModel, nil
	}
	d.cache.Set(key, cachingData, d.expCacheTime)
	return userModel, nil
}

func (d *UserRepositoryDecorator) generateCachingKeyDB(namespace, setKey, id string) string {
	return fmt.Sprintf("%s:%s:%s", namespace, setKey, id)
}

func (d *UserRepositoryDecorator) getUserModelFromRemoteCachedData(cachedData interface{}) (*models.User, error) {
	data, _ := cachedData.(map[string]interface{})
	newsDataBytes := (data[enums.DefaultKeyMapBin]).([]byte)
	var user *models.User
	err := json.Unmarshal(newsDataBytes, &user)
	if err != nil {
		log.Errorf("json unmarshal user fail, err:[%s]", err)
		return nil, err
	}
	return user, nil
}

func (d *UserRepositoryDecorator) parseMapFromUser(user *models.User) (interface{}, error) {
	cachingData := make(map[string]interface{})
	newsBytes, err := json.Marshal(user)
	if err != nil {
		log.Errorf("json marshal user fail, err:[%s]", err)
		return nil, err
	}
	cachingData[enums.DefaultKeyMapBin] = newsBytes
	return cachingData, nil
}
