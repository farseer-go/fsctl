Database:
  default: "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
Redis:
  default: "Server=127.0.0.1:6379,DB=13,Password=,ConnectTimeout=600000,SyncTimeout=10000,ResponseTimeout=10000"
Rabbit:
  - Server: "Server=127.0.0.1:5672,UserName=farseer,Password=farseer"
    Exchange:
      - "ExchangeName=Ex1,RoutingKey=,ExchangeType=fanout,UseConfirmModel=false,AutoCreateExchange=true"
      - "ExchangeName=Ex1,RoutingKey=,ExchangeType=fanout,UseConfirmModel=false,AutoCreateExchange=true"
Etcd:
  default: "Server=127.0.0.1:2379|127.0.0.1:2379,ConnectTimeout=5000"
ElasticSearch:
  es: "Server=http://127.0.0.1:9200,Username=es,Password=123456,ReplicasCount=1,ShardsCount=1,RefreshInterval=5,IndexFormat=yyyy_MM"
  LinkTrack: "Server=http://127.0.0.1:9200,Username=es,Password=123456"
WebApi:
  Url: ":888"
  Session:
    Store: "Redis"
    StoreConfigName: "default"
    Age: 1800
Log:
  Default:
    LogLevel: "info"          # 只记录级别>=info的日志内容
    Format: "json"            # 默认使用json格式输出
  Console:
    LogLevel: "info"          # 只记录级别>=info的日志内容
    Format: "text"            # 控制台打印，使用text格式输出
  File:
    LogLevel: "info"          # 只记录级别>=info的日志内容
    Format: "text"            # 使用text格式写入日志文件
    Path: "./log"             # 日志文件存储在应用程序的./log目录中
    RollingInterval: "Hour"   # 滚动间隔（Hour|Day|Week|Month|Year）
    FileSizeLimitMb: 1        # 文件大小限制
    FileCountLimit: 20        # 文件数量限制
    RefreshInterval: 1        # 写入到文件的时间间隔，秒单位，最少为1
  Database:
  ElasticSearch:
  Component:
    task: true                # 打印task组件的日志
    cacheManage: true         # 打印cacheManage组件的日志
    webapi: true              # 打印webapi组件的日志
    event: true               # 打印event组件的日志
    httpRequest: true         # 打印httpRequest组件的日志
    queue: true               # 打印queue组件的日志
    fSchedule: true           # 打印fSchedule组件的日志