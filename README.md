# polaris-store-postgresql

## 如何使用

#### 构建

执行构建脚本 `build.sh` 即可

```bash
# ${store_pg_plugin_version}: store-postgresql 插件版本，默认为 main 的最新 commit
# ${polaris_server_tag}: 北极星服务端版本信息，默认为 main 分支
make build STORE_PG_PLUGIN_VERSION=${store_pg_plugin_version} POLARIS_SERVER_VERSION=${polaris_server_tag}
```

#### 配置文件调整

修改 **conf/polaris-server.yaml** 文件，参考下列配置调整 store 的配置信息

```yaml
# Storage configuration
store:
  ## Database storage plugin
  name: defaultStore
  option:
    master:
      # 设置数据库类型为 postgresql
      dbType: postgresql
      dbName: polaris_server
      dbUser: ##DB_USER##
      dbPwd: ##DB_PWD##
      dbAddr: ##DB_ADDR##
      maxOpenConns: 300
      maxIdleConns: 50
      connMaxLifetime: 300 # Unit second
      txIsolationLevel: 2 #LevelReadCommitted
```


## 其他

- NACOS 中的 struct 数据结构定义大部份引用自 [nacos-sdk-go](https://github.com/nacos-group/nacos-sdk-go)