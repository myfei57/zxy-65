基于 Go 实现的港口岸电接驳管理平台项目，一款后端服务，管理船舶插接确认、相序核验、频率同步、断路器合闸顺序与用电计量。

## 运行

项目使用文件持久化，无外部数据库。启动前先初始化 Go 依赖（离线 vendor）：

```
go build -mod=vendor -o portpower.exe ./cmd/portpower
./portpower.exe
```

服务默认监听 8080 端口，健康检查地址为 /health。

## 依赖

- github.com/go-chi/chi/v5
- github.com/google/uuid
