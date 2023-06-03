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
	"testing"

	_ "github.com/lib/pq"
)

func TestDemo(t *testing.T) {
	// 创建连接
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "192.168.31.19", 5432, "postgres", "aaaaaa", "postgres"))
	//db, err := sql.Open("postgres", fmt.Sprintf("%s:%s@tcp(%s)/%s", "postgres", "aaaaaa", "192.168.31.19:5432", "postgres"))
	if err != nil {
		fmt.Printf("sql.Open, err: %+v\n", err)
		return
	}
	defer db.Close()

	if pingErr := db.Ping(); pingErr != nil {
		fmt.Printf("ping.err: %+v\n", pingErr)
		return
	}

	rows := db.QueryRow("SELECT id, user_name FROM demo;")
	var id *sql.NullInt16
	var userName *sql.NullString
	err = rows.Scan(&id, &userName)
	if err != nil {
		fmt.Printf("Scan, err: %+v\n", err)
		return
	}
	if id.Int16 == 0 || userName.String == "" {
		fmt.Printf("Scan.data.null\n")
		return
	}

	fmt.Println("id:", id.Int16)
	fmt.Println("user_name:", userName.String)
}
