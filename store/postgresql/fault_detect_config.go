/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package postgresql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

var _ store.FaultDetectRuleStore = (*faultDetectRuleStore)(nil)

type faultDetectRuleStore struct {
	master *BaseDB
	slave  *BaseDB
}

const (
	labelCreateFaultDetectRule = "createFaultDetectRule"
	labelUpdateFaultDetectRule = "updateFaultDetectRule"
	labelDeleteFaultDetectRule = "deleteFaultDetectRule"
)

const (
	insertFaultDetectSql = "insert into fault_detect_rule(id, name, namespace, revision, description, " +
		"dst_service, dst_namespace, dst_method, config, ctime, mtime) " +
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,current_timestamp,current_timestamp)"

	updateFaultDetectSql = "update fault_detect_rule set name = $1, namespace = $2, revision = $3, " +
		"description = $4, dst_service = $5, dst_namespace = $6, dst_method = $7, " +
		"config = $8, mtime = current_timestamp where id = $9"

	deleteFaultDetectSql = "update fault_detect_rule set flag = 1, mtime = current_timestamp where id = $1"

	countFaultDetectSql = "select count(*) from fault_detect_rule where flag = 0"

	queryFaultDetectFullSql = "select id, name, namespace, revision, description, dst_service, " +
		"dst_namespace, dst_method, config, ctime, mtime from fault_detect_rule where flag = 0"

	queryFaultDetectBriefSql = "select id, name, namespace, revision, description, dst_service, " +
		"dst_namespace, dst_method, ctime, mtime from fault_detect_rule where flag = 0"

	queryFaultDetectCacheSql = "select id, name, namespace, revision, description, dst_service, " +
		"dst_namespace, dst_method, config, flag, ctime, mtime from fault_detect_rule where mtime > $1"
)

// CreateFaultDetectRule create fault detect rule
func (f *faultDetectRuleStore) CreateFaultDetectRule(fdRule *model.FaultDetectRule) error {
	err := RetryTransaction(labelCreateFaultDetectRule, func() error {
		return f.createFaultDetectRule(fdRule)
	})
	return store.Error(err)
}

func (f *faultDetectRuleStore) createFaultDetectRule(fdRule *model.FaultDetectRule) error {
	return f.master.processWithTransaction(labelCreateFaultDetectRule, func(tx *BaseTx) error {
		stmt, err := tx.Prepare(insertFaultDetectSql)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(fdRule.ID, fdRule.Name, fdRule.Namespace, fdRule.Revision,
			fdRule.Description, fdRule.DstService, fdRule.DstNamespace, fdRule.DstMethod,
			fdRule.Rule); err != nil {
			log.Errorf("[Store][database] fail to %s exec sql, rule(%+v), err: %s",
				labelCreateFaultDetectRule, fdRule, err.Error())
			return err
		}

		if err := tx.Commit(); err != nil {
			log.Errorf("[Store][database] fail to %s commit tx, rule(%+v), err: %s",
				labelCreateFaultDetectRule, fdRule, err.Error())
			return err
		}
		return nil
	})
}

// UpdateFaultDetectRule update fault detect rule
func (f *faultDetectRuleStore) UpdateFaultDetectRule(fdRule *model.FaultDetectRule) error {
	err := RetryTransaction(labelUpdateFaultDetectRule, func() error {
		return f.updateFaultDetectRule(fdRule)
	})
	return store.Error(err)
}

func (f *faultDetectRuleStore) updateFaultDetectRule(fdRule *model.FaultDetectRule) error {
	return f.master.processWithTransaction(labelUpdateFaultDetectRule, func(tx *BaseTx) error {
		stmt, err := tx.Prepare(updateFaultDetectSql)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(fdRule.Name, fdRule.Namespace, fdRule.Revision, fdRule.Description,
			fdRule.DstService, fdRule.DstNamespace, fdRule.DstMethod, fdRule.Rule,
			fdRule.ID); err != nil {
			log.Errorf("[Store][database] fail to %s exec sql, rule(%+v), err: %s",
				labelUpdateFaultDetectRule, fdRule, err.Error())
			return err
		}

		if err := tx.Commit(); err != nil {
			log.Errorf("[Store][database] fail to %s commit tx, rule(%+v), err: %s",
				labelUpdateFaultDetectRule, fdRule, err.Error())
			return err
		}
		return nil
	})
}

// DeleteFaultDetectRule delete fault detect rule
func (f *faultDetectRuleStore) DeleteFaultDetectRule(id string) error {
	err := RetryTransaction(labelDeleteFaultDetectRule, func() error {
		return f.deleteFaultDetectRule(id)
	})
	return store.Error(err)
}

func (f *faultDetectRuleStore) deleteFaultDetectRule(id string) error {
	return f.master.processWithTransaction(labelDeleteFaultDetectRule, func(tx *BaseTx) error {
		stmt, err := tx.Prepare(deleteFaultDetectSql)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(id); err != nil {
			log.Errorf("[Store][database] fail to %s exec sql, rule(%s), err: %s",
				labelDeleteFaultDetectRule, id, err.Error())
			return err
		}

		if err := tx.Commit(); err != nil {
			log.Errorf("[Store][database] fail to %s commit tx, rule(%s), err: %s",
				labelDeleteFaultDetectRule, id, err.Error())
			return err
		}
		return nil
	})
}

// HasFaultDetectRule check fault detect rule exists
func (f *faultDetectRuleStore) HasFaultDetectRule(id string) (bool, error) {
	queryParams := map[string]string{"id": id}
	count, err := f.getFaultDetectRulesCount(queryParams)
	if nil != err {
		return false, err
	}
	return count > 0, nil
}

// HasFaultDetectRuleByName check fault detect rule exists by name
func (f *faultDetectRuleStore) HasFaultDetectRuleByName(name string, namespace string) (bool, error) {
	queryParams := map[string]string{exactName: name, "namespace": namespace}
	count, err := f.getFaultDetectRulesCount(queryParams)
	if nil != err {
		return false, err
	}
	return count > 0, nil
}

// HasFaultDetectRuleByNameExcludeId check fault detect rule exists by name not this id
func (f *faultDetectRuleStore) HasFaultDetectRuleByNameExcludeId(
	name string, namespace string, id string) (bool, error) {
	queryParams := map[string]string{exactName: name, "namespace": namespace, excludeId: id}
	count, err := f.getFaultDetectRulesCount(queryParams)
	if nil != err {
		return false, err
	}
	return count > 0, nil
}

// GetFaultDetectRules get all fault detect rules by query and limit
func (f *faultDetectRuleStore) GetFaultDetectRules(
	filter map[string]string, offset uint32, limit uint32) (uint32, []*model.FaultDetectRule, error) {
	var out []*model.FaultDetectRule
	var err error

	bValue, ok := filter[briefSearch]
	var isBrief = ok && strings.ToLower(bValue) == "true"
	delete(filter, briefSearch)

	if isBrief {
		out, err = f.getBriefFaultDetectRules(filter, offset, limit)
	} else {
		out, err = f.getFullFaultDetectRules(filter, offset, limit)
	}
	if err != nil {
		return 0, nil, err
	}
	num, err := f.getFaultDetectRulesCount(filter)
	if err != nil {
		return 0, nil, err
	}
	return num, out, nil
}

// GetFaultDetectRulesForCache get increment circuitbreaker rules
func (f *faultDetectRuleStore) GetFaultDetectRulesForCache(
	mtime time.Time, firstUpdate bool) ([]*model.FaultDetectRule, error) {
	str := queryFaultDetectCacheSql
	if firstUpdate {
		str += " and flag != 1"
	}
	rows, err := f.slave.Query(str, mtime)
	if err != nil {
		log.Errorf("[Store][database] query fault detect rules with mtime err: %s", err.Error())
		return nil, err
	}
	fdRules, err := fetchFaultDetectRulesRows(rows)
	if err != nil {
		return nil, err
	}
	return fdRules, nil
}

func fetchFaultDetectRulesRows(rows *sql.Rows) ([]*model.FaultDetectRule, error) {
	defer rows.Close()
	var out []*model.FaultDetectRule
	for rows.Next() {
		var fdRule model.FaultDetectRule
		var flag int
		err := rows.Scan(&fdRule.ID, &fdRule.Name, &fdRule.Namespace, &fdRule.Revision,
			&fdRule.Description, &fdRule.DstService, &fdRule.DstNamespace,
			&fdRule.DstMethod, &fdRule.Rule, &flag, &fdRule.CreateTime, &fdRule.ModifyTime)
		if err != nil {
			log.Errorf("[Store][database] fetch brief fault detect rule scan err: %s", err.Error())
			return nil, err
		}
		fdRule.Valid = true
		if flag == 1 {
			fdRule.Valid = false
		}
		out = append(out, &fdRule)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[Store][database] fetch brief fault detect rule next err: %s", err.Error())
		return nil, err
	}
	return out, nil
}

func genFaultDetectRuleSQL(query map[string]string) (string, []interface{}, int) {
	str := ""
	args := make([]interface{}, 0, len(query))
	var (
		svcNamespaceQueryValue string
		svcQueryValue          string
		idx                    = 1
	)
	for key, value := range query {
		if len(value) == 0 {
			continue
		}
		if key == svcSpecificQueryKeyService {
			svcQueryValue = value
			continue
		}
		if key == svcSpecificQueryKeyNamespace {
			svcNamespaceQueryValue = value
			continue
		}
		storeKey := toUnderscoreName(key)
		if _, ok := blurQueryKeys[key]; ok {
			str += fmt.Sprintf(" and %s like $%d", storeKey, idx)
			args = append(args, "%"+value+"%")
		} else if key == exactName {
			str += fmt.Sprintf(" and name = $%d", idx)
			args = append(args, value)
		} else if key == excludeId {
			str += fmt.Sprintf(" and id != $%d", idx)
			args = append(args, value)
		} else {
			str += fmt.Sprintf(" and %s = $%d", storeKey, idx)
			args = append(args, value)
		}
		idx++
	}
	if len(svcQueryValue) > 0 {
		str += fmt.Sprintf(" and (dst_service = $%d or dst_service = '*')", idx)
		idx++
		args = append(args, svcQueryValue)
	}
	if len(svcNamespaceQueryValue) > 0 {
		str += fmt.Sprintf(" and (dst_namespace = $%d or dst_namespace = '*')", idx)
		idx++
		args = append(args, svcNamespaceQueryValue)
	}
	return str, args, idx
}

func (f *faultDetectRuleStore) getFaultDetectRulesCount(filter map[string]string) (uint32, error) {
	queryStr, args, _ := genFaultDetectRuleSQL(filter)
	str := countFaultDetectSql + queryStr
	var total uint32
	err := f.master.QueryRow(str, args...).Scan(&total)
	switch {
	case err == sql.ErrNoRows:
		return 0, nil
	case err != nil:
		log.Errorf("[Store][database] get fault detect rule count err: %s", err.Error())
		return 0, err
	default:
	}
	return total, nil
}

func (f *faultDetectRuleStore) getBriefFaultDetectRules(
	filter map[string]string, offset uint32, limit uint32) ([]*model.FaultDetectRule, error) {
	queryStr, args, idx := genFaultDetectRuleSQL(filter)
	args = append(args, limit, offset)
	str := queryFaultDetectBriefSql + queryStr + fmt.Sprintf(` order by mtime desc limit $%d offset $%d`, idx, idx+1)

	rows, err := f.master.Query(str, args...)
	if err != nil {
		log.Errorf("[Store][database] query brief fault detect rule rules err: %s", err.Error())
		return nil, err
	}
	out, err := fetchBriefFaultDetectRules(rows)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func fetchBriefFaultDetectRules(rows *sql.Rows) ([]*model.FaultDetectRule, error) {
	defer rows.Close()
	var out []*model.FaultDetectRule
	for rows.Next() {
		var fdRule model.FaultDetectRule
		err := rows.Scan(&fdRule.ID, &fdRule.Name, &fdRule.Namespace, &fdRule.Revision,
			&fdRule.Description, &fdRule.DstService, &fdRule.DstNamespace,
			&fdRule.DstMethod, &fdRule.CreateTime, &fdRule.ModifyTime)
		if err != nil {
			log.Errorf("[Store][database] fetch brief fault detect rule scan err: %s", err.Error())
			return nil, err
		}
		out = append(out, &fdRule)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[Store][database] fetch brief fault detect rule next err: %s", err.Error())
		return nil, err
	}
	return out, nil
}

func (f *faultDetectRuleStore) getFullFaultDetectRules(
	filter map[string]string, offset uint32, limit uint32) ([]*model.FaultDetectRule, error) {
	queryStr, args, idx := genFaultDetectRuleSQL(filter)
	args = append(args, limit, offset)
	str := queryFaultDetectFullSql + queryStr + fmt.Sprintf(` order by mtime desc limit $%d offset $%d`, idx, idx+1)

	rows, err := f.master.Query(str, args...)
	if err != nil {
		log.Errorf("[Store][database] query brief fault detect rules err: %s", err.Error())
		return nil, err
	}
	out, err := fetchFullFaultDetectRules(rows)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func fetchFullFaultDetectRules(rows *sql.Rows) ([]*model.FaultDetectRule, error) {
	defer rows.Close()
	var out []*model.FaultDetectRule
	for rows.Next() {
		var fdRule model.FaultDetectRule
		err := rows.Scan(&fdRule.ID, &fdRule.Name, &fdRule.Namespace, &fdRule.Revision,
			&fdRule.Description, &fdRule.DstService, &fdRule.DstNamespace,
			&fdRule.DstMethod, &fdRule.Rule, &fdRule.CreateTime, &fdRule.ModifyTime)
		if err != nil {
			log.Errorf("[Store][database] fetch brief fault detect rule scan err: %s", err.Error())
			return nil, err
		}
		out = append(out, &fdRule)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[Store][database] fetch brief fault detect rule next err: %s", err.Error())
		return nil, err
	}
	return out, nil
}
