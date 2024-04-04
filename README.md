OWASP TOP 10(2021) + CWE
## 架构图
```markdown
                         +-------------------+
                         |    Web Interface  |
                         |    (Gin Framework)|
                         +--------+----------+
                                  |
                                  v
                    +--------------------------+
                    |      longinus API        |
                    |    (Go Application)      |
                    +------+-------------------+
                           |
                           v
           +-------------------------------+
           |      CWE Code Auditing Engine |
           |        (Go Package/Library)   |
           +------+------------------------+
                  |
                  v
    +-----------------------------------------+
    |       Language-specific Auditors       |
    |   (e.g., CWE22/Java Auditor Package)  |
    +-----------------------------------------+
```

## 目录结构
```markdown
longinus/
├── cmd
│   └── longinuscli
│       └── main.go
├── internal
│   ├── web
│   └── rules
        └── cwe22
│           └── java.go
├── pkg
├── api
└── configs
```


## 参考

https://owasp.org/Top10/
https://github.com/golang-standards/project-layout
