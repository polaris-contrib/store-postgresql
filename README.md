# polaris-store-postgresql

## 如何使用

#### 构建

执行构建脚本 `Makefile` 即可

```bash
# ${store_pg_plugin_version}: store-postgresql 插件版本，默认为 main 的最新 commit
# ${polaris_server_tag}: 北极星服务端版本信息，默认为 main 分支
make build STORE_PG_PLUGIN_VERSION=latest POLARIS_SERVER_VERSION=${polaris_server_tag}
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
      dbType: "postgres"
      dbName: "polaris_server"
      dbUser: "改成自已有用户名" ##DB_USER##
      dbPwd: "改成自已的密码" ##DB_PWD##
      dbAddr: "改成自已的IP" ##DB_ADDR##
      dbPort: "改成自已的端口"
      maxOpenConns: -1
      maxIdleConns: -1
      connMaxLifetime: 300 # 单位秒
```
