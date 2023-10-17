package define

import "github.com/bwmarrin/snowflake"

type Server struct {
	Domain string `json:"domain" mapstructure:"domain" ini:"domain"`
	Host   string `json:"host" mapstructure:"host" ini:"host"`
	Port   int    `json:"port" mapstructure:"port" ini:"port"`
	Debug  bool   `json:"debug" mapstructure:"debug" ini:"debug"`
	Node   int64  `json:"node" mapstructure:"node" ini:"node"`
}

type Redis struct {
	Host     string `json:"redis_host" mapstructure:"redis_host" ini:"redis_host"`
	Port     int    `json:"redis_port" mapstructure:"redis_port" ini:"redis_port"`
	Db       int    `json:"redis_db" mapstructure:"redis_db" ini:"redis_db"`
	Password string `json:"redis_password" mapstructure:"redis_password" ini:"redis_password"`
}

type config struct {
	Server Server `json:"server" ini:"server"`
	Redis  Redis  `json:"redis" ini:"redis"`
}

var Config = &config{}

var Snow *snowflake.Node
