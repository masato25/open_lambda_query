package g

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	log "github.com/Sirupsen/logrus"

	"github.com/Cepave/open-falcon-backend/modules/query/logger"
	"github.com/toolkits/file"
)

type GinHttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type GraphConfig struct {
	ConnTimeout int32             `json:"connTimeout"`
	CallTimeout int32             `json:"callTimeout"`
	MaxConns    int32             `json:"maxConns"`
	MaxIdle     int32             `json:"maxIdle"`
	Replicas    int32             `json:"replicas"`
	Cluster     map[string]string `json:"cluster"`
}

type DbConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

type GraphDB struct {
	Addr  string `json:"addr"`
	Idle  int    `json:"idle"`
	Max   int    `json:"max"`
	Limit int    `json:"limit"`
}

type GlobalConfig struct {
	Debug   bool           `json:"debug"`
	RootDir string         `json:"root_dir"`
	Graph   *GraphConfig   `json:"graph"`
	Db      *DbConfig      `json:"db"`
	GinHttp *GinHttpConfig `json:"gin_http"`
	GraphDB *GraphDB       `json:"graphdb"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

// Gets the configuration
func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// Sets the config directly
func SetConfig(newConfig *GlobalConfig) {
	configLock.RLock()
	defer configLock.RUnlock()
	config = newConfig
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("config file not specified: use -c $filename")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file specified not found:", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file", cfg, "error:", err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file", cfg, "error:", err.Error())
	}

	//support develop mode
	if c.RootDir == "" {
		c.RootDir = filepath.Dir(os.Args[0])
	}

	SetConfig(&c)

	logger.InitLogger(c.Debug)
	log.Println("g.ParseConfig ok, file", cfg)
}
