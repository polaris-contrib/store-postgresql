package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
	"strings"
)

type configFileTagStore struct {
	db *BaseDB
}

// CreateConfigFileTag 创建配置文件标签
func (t *configFileTagStore) CreateConfigFileTag(tx store.Tx, fileTag *model.ConfigFileTag) error {
	insertSql := "insert into config_file_tag(key,value,namespace,\"group\",file_name,create_time, " +
		" create_by,modify_time,modify_by)" +
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(insertSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(fileTag.Key, fileTag.Value, fileTag.Namespace,
			fileTag.Group, fileTag.FileName, GetCurrentTimeFormat(), fileTag.CreateBy,
			GetCurrentTimeFormat(), fileTag.ModifyBy)
	} else {
		stmt, err := t.db.Prepare(insertSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(fileTag.Key, fileTag.Value, fileTag.Namespace,
			fileTag.Group, fileTag.FileName, GetCurrentTimeFormat(), fileTag.CreateBy,
			GetCurrentTimeFormat(), fileTag.ModifyBy)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

// QueryConfigFileByTag 通过标签查询配置文件
func (t *configFileTagStore) QueryConfigFileByTag(namespace, group, fileName string,
	tags ...string) ([]*model.ConfigFileTag, error) {
	var idx = 1
	group = "%" + group + "%"
	fileName = "%" + fileName + "%"
	querySql := t.baseSelectSql() + fmt.Sprintf(" where namespace = $%d and \"group\" like $%d and file_name like $%d ", idx, idx+1, idx+2)
	idx += 3

	var tagWhereSql []string
	for i := 0; i < len(tags)/2; i++ {
		tagWhereSql = append(tagWhereSql, fmt.Sprintf("($%d,$%d)", idx, idx+1))
		idx += 2
	}
	tagIn := "and (key, value) in  (" + strings.Join(tagWhereSql, ",") + ")"
	querySql = querySql + tagIn

	params := []interface{}{namespace, group, fileName}
	for _, tag := range tags {
		params = append(params, tag)
	}
	rows, err := t.db.Query(querySql, params...)
	if err != nil {
		return nil, store.Error(err)
	}

	result, err := t.transferRows(rows)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTagByConfigFile 查询配置文件标签
func (t *configFileTagStore) QueryTagByConfigFile(namespace, group, fileName string) ([]*model.ConfigFileTag, error) {
	querySql := t.baseSelectSql() + fmt.Sprintf(" where namespace = $1 and \"group\" = $2 and file_name = $3")
	rows, err := t.db.Query(querySql, namespace, group, fileName)
	if err != nil {
		return nil, store.Error(err)
	}

	tags, err := t.transferRows(rows)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// DeleteConfigFileTag 删除配置文件标签
func (t *configFileTagStore) DeleteConfigFileTag(tx store.Tx, namespace, group, fileName, key, value string) error {
	deleteSql := "delete from config_file_tag where key = $1 and value = $2 and namespace = $3 " +
		" and \"group\" = $4 and file_name = $5"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(key, value, namespace, group, fileName)
	} else {
		stmt, err := t.db.Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(key, value, namespace, group, fileName)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

// DeleteTagByConfigFile 删除配置文件的标签
func (t *configFileTagStore) DeleteTagByConfigFile(tx store.Tx, namespace, group, fileName string) error {
	deleteSql := "delete from config_file_tag where namespace = $1 and \"group\" = $2 and file_name = $3"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(namespace, group, fileName)
	} else {
		stmt, err := t.db.Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(namespace, group, fileName)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

func (t *configFileTagStore) baseSelectSql() string {
	return "select id, key,value,namespace,\"group\",file_name,create_time, " +
		" create_by,modify_time,modify_by from config_file_tag"
}

func (t *configFileTagStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileTag, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var tags []*model.ConfigFileTag

	for rows.Next() {
		tag := &model.ConfigFileTag{}
		err := rows.Scan(&tag.Id, &tag.Key, &tag.Value, &tag.Namespace, &tag.Group, &tag.FileName,
			&tag.CreateTime, &tag.CreateBy, &tag.ModifyTime, &tag.ModifyBy)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
