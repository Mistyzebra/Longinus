OWASP TOP 10(2021) + CWE
## 目录结构
```
Longinus/
├── cmd/
│   └── start/
│       └── main.go  # 程序入口点
├── internal/
│   ├── rules/
│   │   └── cwe_22_java.go  # CWE-22针对Java代码的审计逻辑
│   └── web/
│       └── server.go  # 使用Gin设置Web服务器
├── pkg/
│   └── cwe/
│       └── common.go  # 定义CWE相关的通用结构和接口
├── go.mod
└── go.sum
```


## 参考

https://owasp.org/Top10/
https://github.com/golang-standards/project-layout
