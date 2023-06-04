#!/bin/bash
# Tencent is pleased to support the open source community by making Polaris available.
#
# Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
#
# Licensed under the BSD 3-Clause License (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# https://opensource.org/licenses/BSD-3-Clause
#
# Unless required by applicable law or agreed to in writing, software distributed
# under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
# CONDITIONS OF ANY KIND, either express or implied. See the License for the
# specific language governing permissions and limitations under the License.

store_pg_plugin_version=$1
polaris_server_tag=$2

if [ "${store_pg_plugin_version}" == "" ]; then
    echo -e "you must set store_pg_plugin_version info, eg bash build.sh v1.0.0"
    exit 1
fi

workdir=$(pwd)

## 清理之前的临时资源
rm -rf build_resource

## 准备临时构建资源文件夹
mkdir build_resource
cp plugin.go.temp ./build_resource
cd build_resource

git clone https://github.com/polarismesh/polaris.git
cd polaris
if [ "${polaris_server_tag}" != "" ]; then
    git checkout ${polaris_server_tag}
fi

cat ../plugin_store_pg.go.temp >plugin_store_pg.go

go clean --modcache
go get -u github.com/polaris-contrib/polaris-store-postgresql@${store_pg_plugin_version}
go mod tidy

make build VERSION=${polaris_server_tag}

release_file=$(ls -lstrh | grep ".zip" | awk '{print $10}' | grep -v "md5")
cp ${release_file} ${workdir}

