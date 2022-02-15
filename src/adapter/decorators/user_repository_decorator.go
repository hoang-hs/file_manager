package decorators

import (
	"encoding/json"
	"file_manager/src/adapter/database/repositories"
	"file_manager/src/common/caching"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"file_manager/src/core/entities"
	"file_manager/src/core/enums"
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
) *UserRepositoryDecorator {
	cf := configs.Get()
	return &UserRepositoryDecorator{
		cache:          cache,
		userRepository: userRepository,
		expCacheTime:   cf.ExpCacheTimeDb,
		setKeyEnv:      fmt.Sprintf("%s_%s", enums.DefaultSetKeyDB, cf.AppEnv),
	}
}

func (d *UserRepositoryDecorator) FindByUsername(username string) (*entities.User, error) {
	key := d.generateCachingKeyDB(enums.DefaultNameSpace, d.setKeyEnv, username)
	data, found := d.cache.Get(key)
	if found {
		user, err := d.getUserModelFromRemoteCachedData(data)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := d.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	cachingData, err := d.parseMapFromUser(user)
	if err != nil {
		return user, nil
	}
	d.cache.Set(key, cachingData, d.expCacheTime)
	return user, nil
}

func (d *UserRepositoryDecorator) generateCachingKeyDB(namespace, setKey, id string) string {
	return fmt.Sprintf("%s:%s:%s", namespace, setKey, id)
}

func (d *UserRepositoryDecorator) getUserModelFromRemoteCachedData(cachedData interface{}) (*entities.User, error) {
	userBytes := cachedData.([]byte)
	var user *entities.User
	err := json.Unmarshal(userBytes, &user)
	if err != nil {
		log.Errorf("json unmarshal user fail, err:[%s]", err)
		return nil, err
	}
	return user, nil
}

func (d *UserRepositoryDecorator) parseMapFromUser(user *entities.User) (interface{}, error) {
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Errorf("json marshal user fail, err:[%s]", err)
		return nil, err
	}
	return userBytes, nil
}
