package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
	"strings"
)

type configFileGroupStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateConfigFileGroup 创建配置文件组
func (fg *configFileGroupStore) CreateConfigFileGroup(
	fileGroup *model.ConfigFileGroup) (*model.ConfigFileGroup, error) {
	createSql := "insert into config_file_group(name, namespace,comment,create_time, create_by, " +
		" modify_time, modify_by, owner)" +
		"value ($1,$2,$3,$4,$5,$6,$7,$8)"
	stmt, err := fg.master.Prepare(createSql)
	if err != nil {
		return nil, store.Error(err)
	}
	_, err = stmt.Exec(fileGroup.Name, fileGroup.Namespace, fileGroup.Comment,
		fileGroup.CreateBy, GetCurrentTimeFormat(), fileGroup.ModifyBy, fileGroup.Owner)
	if err != nil {
		return nil, store.Error(err)
	}

	return fg.GetConfigFileGroup(fileGroup.Namespace, fileGroup.Name)
}

// GetConfigFileGroup 获取配置文件组
func (fg *configFileGroupStore) GetConfigFileGroup(namespace, name string) (*model.ConfigFileGroup, error) {
	querySql := fg.genConfigFileGroupSelectSql() + fmt.Sprintf(" where namespace=$1 and name=$2")
	rows, err := fg.master.Query(querySql, namespace, name)
	if err != nil {
		return nil, store.Error(err)
	}
	cfgs, err := fg.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(cfgs) > 0 {
		return cfgs[0], nil
	}
	return nil, nil
}

// QueryConfigFileGroups 翻页查询配置文件组, name 为模糊匹配关键字
func (fg *configFileGroupStore) QueryConfigFileGroups(namespace, name string,
	offset, limit uint32) (uint32, []*model.ConfigFileGroup, error) {
	name = "%" + name + "%"
	// 全部 namespace
	if namespace == "" {
		countSql := "select count(*) from config_file_group where name like $1"
		var count uint32
		err := fg.master.QueryRow(countSql, name).Scan(&count)
		if err != nil {
			return count, nil, err
		}

		s := fg.genConfigFileGroupSelectSql() + " where name like $1 order by id desc limit $2 offset $3"
		rows, err := fg.master.Query(s, name, limit, offset)
		if err != nil {
			return 0, nil, err
		}
		cfgs, err := fg.transferRows(rows)
		if err != nil {
			return 0, nil, err
		}

		return count, cfgs, nil
	}

	// 特定 namespace
	countSql := "select count(*) from config_file_group where namespace=$1 and name like $2"
	var count uint32
	err := fg.master.QueryRow(countSql, namespace, name).Scan(&count)
	if err != nil {
		return count, nil, err
	}

	s := fg.genConfigFileGroupSelectSql() + " where namespace=$1 and name like $2 order by id desc limit $3 offset $4 "
	rows, err := fg.master.Query(s, namespace, name, limit, offset)
	if err != nil {
		return 0, nil, err
	}
	cfgs, err := fg.transferRows(rows)
	if err != nil {
		return 0, nil, err
	}

	return count, cfgs, nil
}

// DeleteConfigFileGroup 删除配置文件组
func (fg *configFileGroupStore) DeleteConfigFileGroup(namespace, name string) error {
	deleteSql := "delete from config_file_group where namespace = $1 and name=$2"

	log.Infof("[Config][Storage] delete config file group(%s, %s)", namespace, name)
	if _, err := fg.master.Exec(deleteSql, namespace, name); err != nil {
		return err
	}

	return nil
}

// UpdateConfigFileGroup 更新配置文件组信息
func (fg *configFileGroupStore) UpdateConfigFileGroup(
	fileGroup *model.ConfigFileGroup) (*model.ConfigFileGroup, error) {
	updateSql := "update config_file_group set comment = $1, modify_time = $2, modify_by = $3 " +
		" where namespace = $4 and name = $5"
	stmt, err := fg.master.Prepare(updateSql)
	if err != nil {
		return nil, store.Error(err)
	}
	_, err = stmt.Exec(fileGroup.Comment, GetCurrentTimeFormat(), fileGroup.ModifyBy,
		fileGroup.Namespace, fileGroup.Name)
	if err != nil {
		return nil, store.Error(err)
	}
	return fg.GetConfigFileGroup(fileGroup.Namespace, fileGroup.Name)
}

// FindConfigFileGroups 获取一组配置文件组信息
func (fg *configFileGroupStore) FindConfigFileGroups(namespace string,
	names []string) ([]*model.ConfigFileGroup, error) {
	querySql := fg.genConfigFileGroupSelectSql()
	params := make([]interface{}, 0)
	idx := 1

	if namespace == "" {
		querySql += " where name in (%s)"
	} else {
		querySql += fmt.Sprintf(" where namespace = $%d", idx) + " and name in (%s)"
		idx++
		params = append(params, namespace)
	}

	inParamPlaceholders := make([]string, 0)
	for i := 0; i < len(names); i++ {
		inParamPlaceholders = append(inParamPlaceholders, fmt.Sprintf("$%d", idx))
		idx++
		params = append(params, names[i])
	}
	querySql = fmt.Sprintf(querySql, strings.Join(inParamPlaceholders, ","))

	rows, err := fg.master.Query(querySql, params...)
	if err != nil {
		return nil, err
	}
	cfgs, err := fg.transferRows(rows)
	if err != nil {
		return nil, err
	}
	return cfgs, nil
}

func (fg *configFileGroupStore) GetConfigFileGroupById(id uint64) (*model.ConfigFileGroup, error) {
	querySql := fg.genConfigFileGroupSelectSql()
	querySql += fmt.Sprintf(" where id = %d", id)

	rows, err := fg.master.Query(querySql)
	if err != nil {
		return nil, err
	}

	cfgs, err := fg.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(cfgs) == 0 {
		return nil, nil
	}

	return cfgs[0], nil
}

func (fg *configFileGroupStore) CountGroupEachNamespace() (map[string]int64, error) {
	metricsSql := "SELECT namespace, count(name) FROM config_file_group GROUP by namespace"
	rows, err := fg.slave.Query(metricsSql)
	if err != nil {
		return nil, store.Error(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	ret := map[string]int64{}
	for rows.Next() {
		var (
			namespce string
			cnt      int64
		)

		if err := rows.Scan(&namespce, &cnt); err != nil {
			return nil, err
		}
		ret[namespce] = cnt
	}

	return ret, nil
}

func (fg *configFileGroupStore) genConfigFileGroupSelectSql() string {
	return "select id,name,namespace,comment,create_time,create_by," +
		"modify_time,modify_by,owner from config_file_group"
}

func (fg *configFileGroupStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileGroup, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var fileGroups []*model.ConfigFileGroup

	for rows.Next() {
		fileGroup := &model.ConfigFileGroup{}
		err := rows.Scan(&fileGroup.Id, &fileGroup.Name, &fileGroup.Namespace, &fileGroup.Comment,
			&fileGroup.CreateTime, &fileGroup.CreateBy, &fileGroup.ModifyTime, &fileGroup.ModifyBy,
			&fileGroup.Owner)
		if err != nil {
			return nil, err
		}

		fileGroups = append(fileGroups, fileGroup)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fileGroups, nil
}
