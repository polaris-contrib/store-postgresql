package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
	"strconv"
	"time"
)

type grayStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateGrayResourceTx 创建灰度资源
func (g *grayStore) CreateGrayResourceTx(tx store.Tx, data *model.GrayResource) error {
	if tx == nil {
		return ErrTxIsNil
	}
	dbTx := tx.GetDelegateTx().(*BaseTx)

	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	// 使用 ON CONFLICT 替代 MySQL 的 ON DUPLICATE KEY UPDATE
	s := "INSERT INTO gray_resource(name, match_rule, create_time, create_by, modify_time, modify_by) " +
		"VALUES ($1, $2, CURRENT_TIMESTAMP, $3, CURRENT_TIMESTAMP, $4) " +
		"ON CONFLICT (name) DO UPDATE SET " +
		"match_rule = EXCLUDED.match_rule, create_time = CURRENT_TIMESTAMP, create_by = EXCLUDED.create_by, " +
		"modify_time = CURRENT_TIMESTAMP, modify_by = EXCLUDED.modify_by"

	args := []interface{}{
		data.Name, data.MatchRule,
		data.CreateBy, data.ModifyBy,
	}
	if _, err := dbTx.Exec(s, args...); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

func (g *grayStore) CleanGrayResource(tx store.Tx, data *model.GrayResource) error {
	if tx == nil {
		return ErrTxIsNil
	}
	dbTx := tx.GetDelegateTx().(*BaseTx)
	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()
	s := "UPDATE gray_resource SET flag = 1, modify_time = CURRENT_TIMESTAMP WHERE name = $1"
	args := []interface{}{data.Name}
	if _, err := dbTx.Exec(s, args...); err != nil {
		return store.Error(err)
	}
	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}
	return nil
}

// GetMoreGrayResouces  获取最近更新的灰度资源, 此方法用于 cache 增量更新，需要注意 modifyTime 应为数据库时间戳
func (g *grayStore) GetMoreGrayResouces(firstUpdate bool,
	modifyTime time.Time) ([]*model.GrayResource, error) {

	if firstUpdate {
		modifyTime = time.Time{}
	}

	// 构造 PostgreSQL 查询
	s := "SELECT name, match_rule, create_time, COALESCE(create_by, ''), modify_time, COALESCE(modify_by, ''), flag " +
		"FROM gray_resource WHERE modify_time > TO_TIMESTAMP($1)"
	if firstUpdate {
		s += " AND flag = 0"
	}
	rows, err := g.slave.Query(s, timeToTimestamp(modifyTime))
	if err != nil {
		return nil, err
	}
	grayResources, err := g.fetchGrayResourceRows(rows)
	if err != nil {
		return nil, err
	}
	return grayResources, nil
}

func (g *grayStore) fetchGrayResourceRows(rows *sql.Rows) ([]*model.GrayResource, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	grayResources := make([]*model.GrayResource, 0, 32)
	for rows.Next() {
		var (
			ctimeStr, mtimeStr string
			valid              int64
		)
		grayResource := &model.GrayResource{}
		if err := rows.Scan(&grayResource.Name, &grayResource.MatchRule, &ctimeStr,
			&grayResource.CreateBy, &mtimeStr, &grayResource.ModifyBy, &valid); err != nil {
			return nil, err
		}
		grayResource.Valid = valid == 0
		// 将字符串转换为int64
		ctimeFloat, err := strconv.ParseFloat(ctimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse create_time: %v", err)
		}
		mtimeFloat, err := strconv.ParseFloat(mtimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse modify_time: %v", err)
		}
		grayResource.CreateTime = time.Unix(int64(ctimeFloat), 0)
		grayResource.ModifyTime = time.Unix(int64(mtimeFloat), 0)
		grayResources = append(grayResources, grayResource)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return grayResources, nil
}

// DeleteGrayResource 删除灰度资源
func (g *grayStore) DeleteGrayResource(tx store.Tx, data *model.GrayResource) error {
	s := "DELETE FROM  gray_resource  WHERE name= $1"
	_, err := g.master.Exec(s, data.Name)
	if err != nil {
		return store.Error(err)
	}
	return nil
}
