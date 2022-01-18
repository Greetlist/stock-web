package conf

var StockResultBaseDir string = "/home/greetlist/macd/result/"
var StockStorageBaseDir string = "/home/greetlist/macd/data_storage/"
var DailyResultBaseDir string = "/data/tmp/daily_ma_result/"

var StockPredictionBaseDir string = "/home/greetlist/workspace/data_storage/"
var StockMaBaseDir string = "/home/greetlist/workspace/data_storage/ma"
var StockEmaBaseDir string = "/home/greetlist/workspace/data_storage/ema"
var StockRawBaseDir string = "/home/greetlist/workspace/data_storage/raw"

var DBConnectionStr string = "trader:trader@tcp(127.0.0.1:3306)/trade?charset=utf8mb4&parseTime=True&loc=Local"
var DBConnectionNum int = 10
var RedisConnectionNum int = 10
var SessionExpireTime int64 = 900
var RedisProtocol string = "tcp"
var RedisAddr string = ":6379"
