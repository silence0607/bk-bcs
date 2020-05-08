/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

const (
	// ENVPREFIX is prefix for env variables.
	ENVPREFIX = "BSCP_DM"
)

// config for local module.
type config struct {
	viper *viper.Viper
}

// init initialize and check the module configs.
func (c *config) init(localConfigFile string) (*viper.Viper, error) {
	c.viper = viper.New()
	c.viper.SetConfigFile(localConfigFile)

	if err := c.viper.ReadInConfig(); err != nil {
		return nil, err
	}
	c.viper.WatchConfig()

	c.viper.SetEnvPrefix(ENVPREFIX)
	c.viper.AutomaticEnv()

	if err := c.check(); err != nil {
		return nil, err
	}
	return c.viper, nil
}

func (c *config) envName(bindKey string) string {
	return ENVPREFIX + "_" + bindKey
}

// check bind the env vars and checks base config content.
func (c *config) check() error {
	c.viper.BindEnv("server.servicename", c.envName("SERVICE_NAME"))
	if !c.viper.IsSet("server.servicename") {
		return errors.New("config check, missing 'server.servicename'")
	}

	c.viper.BindEnv("server.endpoint.ip", c.envName("ENDPOINT_IP"))
	if !c.viper.IsSet("server.endpoint.ip") {
		return errors.New("config check, missing 'server.endpoint.ip'")
	}

	c.viper.BindEnv("server.endpoint.port", c.envName("ENDPOINT_PORT"))
	if !c.viper.IsSet("server.endpoint.port") {
		return errors.New("config check, missing 'server.endpoint.port'")
	}

	c.viper.BindEnv("server.discoveryttl", c.envName("DISCOVERY_TTL"))
	c.viper.SetDefault("server.discoveryttl", 60)

	c.viper.BindEnv("server.commitCacheSize", c.envName("COMMIT_CACHE_SIZE"))
	c.viper.SetDefault("server.commitCacheSize", 100)

	c.viper.BindEnv("server.releaseCacheSize", c.envName("RELEASE_CACHE_SIZE"))
	c.viper.SetDefault("server.releaseCacheSize", 100)

	c.viper.BindEnv("server.configsCacheSize", c.envName("CONFIGS_CACHE_SIZE"))
	c.viper.SetDefault("server.configsCacheSize", 200)

	c.viper.BindEnv("metrics.endpoint", c.envName("METRICS_ENDPOINT"))
	c.viper.SetDefault("metrics.endpoint", ":9100")

	c.viper.BindEnv("etcdCluster.endpoints", c.envName("ETCD_ENDPOINTS"))
	if !c.viper.IsSet("etcdCluster.endpoints") {
		return errors.New("config check, missing 'etcdCluster.endpoints'")
	}

	c.viper.BindEnv("etcdCluster.dialtimeout", c.envName("ETCD_DIAL_TIMEOUT"))
	if !c.viper.IsSet("etcdCluster.dialtimeout") {
		c.viper.SetDefault("etcdCluster.dialtimeout", 3*time.Second)
	}

	c.viper.BindEnv("database.dialect", c.envName("DB_DIALECT"))
	c.viper.SetDefault("database.dialect", "mysql")

	c.viper.BindEnv("database.host", c.envName("DB_HOST"))
	if !c.viper.IsSet("database.host") {
		return errors.New("config check, missing 'database.host'")
	}

	c.viper.BindEnv("database.port", c.envName("DB_PORT"))
	c.viper.SetDefault("database.port", 3306)

	c.viper.BindEnv("database.user", c.envName("DB_USER"))
	if !c.viper.IsSet("database.user") {
		return errors.New("config check, missing 'database.user'")
	}

	c.viper.BindEnv("database.conntimeout", c.envName("CONN_TIMEOUT"))
	c.viper.SetDefault("database.conntimeout", 3*time.Second)

	c.viper.BindEnv("database.readtimeout", c.envName("DB_READ_TIMEOUT"))
	c.viper.SetDefault("database.readtimeout", 3*time.Second)

	c.viper.BindEnv("database.writetimeout", c.envName("DB_WRITE_TIMEOUT"))
	c.viper.SetDefault("database.writetimeout", 3*time.Second)

	c.viper.BindEnv("database.maxopenconns", c.envName("DB_MAX_OPEN_CONNS"))
	c.viper.SetDefault("database.maxopenconns", 10)

	c.viper.BindEnv("database.maxidleconns", c.envName("DB_MAX_IDLE_CONNS"))
	c.viper.SetDefault("database.maxidleconns", 5)

	c.viper.BindEnv("database.keepalive", c.envName("DB_KEEPALIVE"))
	c.viper.SetDefault("database.keepalive", 10*time.Minute)

	c.viper.BindEnv("database.cachesize", c.envName("DB_CACHE_SIZE"))
	c.viper.SetDefault("database.cachesize", 1000)

	c.viper.BindEnv("database.purgeInterval", c.envName("DB_PURGE_INTERVAL"))
	c.viper.SetDefault("database.purgeInterval", time.Minute)

	c.viper.BindEnv("logger.directory", c.envName("LOG_DIR"))
	c.viper.SetDefault("logger.directory", "./log")

	c.viper.BindEnv("logger.maxsize", c.envName("LOG_MAXSIZE"))
	c.viper.SetDefault("logger.maxsize", 200)

	c.viper.BindEnv("logger.maxnum", c.envName("LOG_MAXNUM"))
	c.viper.SetDefault("logger.maxnum", 5)

	c.viper.BindEnv("logger.stderr", c.envName("LOG_STDERR"))
	c.viper.SetDefault("logger.stderr", false)

	c.viper.BindEnv("logger.alsoStderr", c.envName("LOG_ALSOSTDERR"))
	c.viper.SetDefault("logger.alsoStderr", false)

	c.viper.BindEnv("logger.level", c.envName("LOG_LEVEL"))
	c.viper.SetDefault("logger.level", 3)

	c.viper.BindEnv("logger.stderrThreshold", c.envName("LOG_STDERR_THRESHOLD"))
	c.viper.SetDefault("logger.stderrThreshold", 2)

	return nil
}