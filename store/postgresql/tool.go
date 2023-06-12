package postgresql

import (
	"github.com/polarismesh/polaris/common/log"
	"time"
)

// toolStore 实现了ToolStoreStore
type toolStore struct {
	db *BaseDB
}

const (
	nowSql           = `select CURRENT_TIMESTAMP`
	maxQueryInterval = time.Second
)

// GetUnixSecond 获取当前时间，单位秒
func (t *toolStore) GetUnixSecond(maxWait time.Duration) (int64, error) {
	startTime := time.Now()
	rows, err := t.db.Query(nowSql)
	if err != nil {
		log.Errorf("[Store][database] query now err: %s", err.Error())
		return 0, err
	}
	defer rows.Close()
	timePass := time.Since(startTime)
	if maxWait != 0 && timePass > maxWait {
		log.Infof("[Store][database] query now spend %s, exceed %s, skip", timePass, maxWait)
		return 0, nil
	}
	var value int64
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Errorf("[Store][database] get now rows scan err: %s", err.Error())
			return 0, err
		}
	}
	return value, nil
}
