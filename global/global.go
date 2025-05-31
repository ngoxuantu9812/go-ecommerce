package global

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go-ecomm/pkg/logger"
	"go-ecomm/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Rds           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
