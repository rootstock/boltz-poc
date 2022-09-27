package storage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Key   string `db:"Key"`
	Value string `db:"Value"`
}

type ConfigRequest struct {
	Key   string `json:"key"`
	Value string `json:"Value"`
}

const createConfigTable = `
CREATE TABLE IF NOT EXISTS config (
	Key TEXT PRIMARY KEY,
	Value TEXT
)
`

const insertConfig = `
INSERT INTO config (
    Key,
	Value
)
VALUES (
    :Key,
	:Value
)
`

const updateConfig = `
UPDATE config SET  Value = :Value
Where Key = :Key`

const getConfigs = `
SELECT *
FROM config`

const getConfig = `
SELECT *
FROM config
WHERE Key = ?
LIMIT 1`

func (db *DB) SetDefaultConfig() error {
	defaultConf := &Config{
		Key:   "swaptype",
		Value: "mint", //"liquidity",
	}
	if res, _ := db.GetConfig(defaultConf.Key); res.Key == "" {
		return db.SaveConfig(defaultConf)
	}
	return nil
}

func (db *DB) SaveConfig(conf *Config) error {
	exist, err := db.GetConfig(conf.Key)
	if err != nil {
		log.Error("Error saving config", err)
		return err
	}
	var queryString string
	if (exist != Config{}) {
		log.Debug("updating config{", conf.Key, "}", ": ", conf)
		queryString = updateConfig
	} else {
		log.Debug("inserting config{", conf.Key, "}", ": ", conf)
		queryString = insertConfig
	}
	query, args, _ := sqlx.Named(queryString, conf)
	_, err = db.db.Exec(query, args...)
	if err != nil {
		log.Error("Error saving config", err)
		return err
	}
	log.Debug("success saving config!")
	return nil
}

func (db *DB) GetConfigs() ([]*Config, error) {
	log.Debug("retrieving config... ")
	var configs []*Config
	rows, err := db.db.Queryx(getConfigs)
	if err == sql.ErrNoRows {
		return configs, nil
	}
	defer rows.Close()
	for rows.Next() {
		var conf Config

		if err := rows.StructScan(&conf); err != nil {
			return nil, err
		}

		configs = append(configs, &conf)
	}
	return configs, nil
}

func (db *DB) GetConfig(key string) (result Config, err error) {
	log.Debugf("retrieving config: %s ", key)
	rows, err := db.db.Queryx(getConfig, key)
	if err != nil {
		log.Fatalf("Error retrieving config: %v", err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			log.Fatalf("Error mapping config: %v", err)
			return
		}
	}
	return

}
