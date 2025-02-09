project-name/
├── api/                    # API 接口定义
│   └── v1/                # API 版本
│       ├── user.go        # 用户相关接口
│       └── product.go     # 产品相关接口
│
├── cmd/                    # 主要应用程序入口
│   └── server/            
│       └── main.go        # 主程序入口点
│
├── configs/                # 配置文件目录
│   ├── config.yaml        # 主配置文件
│   └── config.go          # 配置结构定义
│
├── internal/              # 私有应用程序代码
│   ├── dao/              # 数据访问层
│   │   ├── mysql/        # MySQL 相关
│   │   │   ├── user.go
│   │   │   └── product.go
│   │   └── redis/        # Redis 相关
│   │       └── cache.go
│   │
│   ├── middleware/       # HTTP 中间件
│   │   ├── auth.go      # 认证中间件
│   │   ├── cors.go      # CORS 中间件
│   │   └── logger.go    # 日志中间件
│   │
│   ├── model/           # 数据模型
│   │   ├── user.go      # 用户模型
│   │   └── product.go   # 产品模型
│   │
│   ├── service/         # 业务逻辑层
│   │   ├── user.go      # 用户服务
│   │   └── product.go   # 产品服务
│   │
│   └── router/          # 路由相关
│       └── router.go    # 路由配置
│
├── pkg/                  # 公共代码包
│   ├── auth/            # 认证相关
│   │   └── jwt.go
│   │
│   ├── conf/            # 配置加载
│   │   └── viper.go
│   │
│   ├── database/        # 数据库初始化
│   │   ├── mysql.go
│   │   └── redis.go
│   │
│   ├── errcode/         # 错误码定义
│   │   └── errcode.go
│   │
│   ├── logger/          # 日志相关
│   │   └── logger.go
│   │
│   └── util/            # 通用工具函数
│       ├── hash.go
│       └── upload.go
│
├── scripts/             # 构建、安装、分析等操作的脚本
│   ├── build.sh
│   └── deploy.sh
│
├── test/               # 测试文件目录
│   ├── api_test.go
│   └── mock/
│
├── web/               # Web 静态资源
│   ├── static/
│   └── template/
│
├── .gitignore        # Git 忽略文件
├── Dockerfile        # Docker 构建文件
├── go.mod           # Go 模块定义
├── go.sum           # Go 模块校验和
├── Makefile         # 项目管理工具
└── README.md        # 项目说明文档
