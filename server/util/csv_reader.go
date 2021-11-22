package util

import (
    "os"
    "path"
    "strings"
    "bufio"
    "fmt"
    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
    "greetlist/stock-web/server/model"
    "greetlist/stock-web/server/conf"
)

var totalDf dataframe.DataFrame
var dataLenGapMin int = 30

func init() {
    fileList := []string{"total_stock.csv", "total_index.csv"}
    for _, fileName := range(fileList) {
        f, _ := os.OpenFile(path.Join(conf.StockRawBaseDir, fileName), os.O_RDWR, os.ModePerm)
        df := dataframe.ReadCSV(f)
        totalDf = totalDf.Concat(df)
    }
}

func ReadMarketRawData(stockCode, period, endDate string, dataLenGap int) model.MarketRawData {
    if dataLenGap < dataLenGapMin {
        dataLenGap = dataLenGapMin
    }
    var res model.MarketRawData
    marketRawDataPath := path.Join(conf.StockRawBaseDir, period, selectExchange(stockCode), stockCode+".csv")
    f, _ := os.OpenFile(marketRawDataPath, os.O_RDWR, os.ModePerm)
    df := dataframe.ReadCSV(f)

    dateStr := strings.ReplaceAll(endDate, "-", "")
    indexList := getIndexRangeList(df, dateStr, dataLenGap)
    rangeDf := df.Subset(indexList)
    res.Date = rangeDf.Col("trade_date").Records()
    res.Open = rangeDf.Col("open").Float()
    res.Close = rangeDf.Col("close").Float()
    res.High = rangeDf.Col("high").Float()
    res.Low = rangeDf.Col("low").Float()
    res.Volume = rangeDf.Col("vol").Float()
    res.Money = rangeDf.Col("amount").Float()
    return res
}

func getIndexRangeList(df dataframe.DataFrame, endDate string, dataLenGap int) []int {
    endIndex := -1
    tradeDateList := df.Col("trade_date").Records()
    for index, value := range(tradeDateList) {
        if value == endDate {
            endIndex = index
            break
        }
    }
    startIndex := endIndex - dataLenGap
    var indexList []int
    if endIndex != -1 && startIndex > 0 && endIndex < df.Nrow() {
        for i := startIndex+1; i < endIndex+1; i++ {
            indexList = append(indexList, i)
        }
    } else {
        endIndex = df.Nrow()
        startIndex = endIndex - dataLenGap
        if startIndex < 0 {
            startIndex = 0
        }
        for i := startIndex; i < df.Nrow(); i++ {
            indexList = append(indexList, i)
        }
    }
    return indexList
}

func ReadFactorRawData(stockCode, period, endDate string, dataLenGap int) model.FactorRawData {
    if dataLenGap < dataLenGapMin {
        dataLenGap = dataLenGapMin
    }
    var res model.FactorRawData
    maDataPath := path.Join(conf.StockMaBaseDir, period, selectExchange(stockCode), stockCode+".ma.csv")
    maFile, _ := os.OpenFile(maDataPath, os.O_RDWR, os.ModePerm)
    madf := dataframe.ReadCSV(maFile)
    //emaDataPath := path.Join(conf.StockEmaBaseDir, period, selectExchange(stockCode), stockCode+".ema.csv")
    //emaFile, _ := os.OpenFile(emaDataPath, os.O_RDWR, os.ModePerm)
    //emadf := dataframe.ReadCSV(emaFile)

    dateStr := strings.ReplaceAll(endDate, "-", "")
    indexList := getIndexRangeList(madf, dateStr, dataLenGap)
    maRangeDf := madf.Subset(indexList)

    res.MA13 = maRangeDf.Col("MA13").Float()
    res.MA34 = maRangeDf.Col("MA34").Float()
    res.MA55 = maRangeDf.Col("MA55").Float()

    //res.EMA13 = emadf.Col("EMA13").Float()[startIndex:]
    //res.EMA34 = emadf.Col("EMA34").Float()[startIndex:]
    //res.EMA55 = emadf.Col("EMA55").Float()[startIndex:]
    return res
}

func ReadStockBasicInfo(stockCode string) []model.StockBasicInfo {
    var res []model.StockBasicInfo
    var filterDf dataframe.DataFrame
    if stockCode != "all" {
        filterDf = totalDf.Filter(
            dataframe.F{
                Colname: "ts_code",
                Comparator: series.Eq,
                Comparando: stockCode,
            },
        )
    } else {
        filterDf = totalDf
    }
    dataLen := filterDf.Nrow()
    stockCodeList := filterDf.Col("ts_code")
    stockNameList := filterDf.Col("name")
    stockIndustryList := filterDf.Col("industry")
    stockExchangeList := filterDf.Col("exchange")
    for idx := 0; idx < dataLen; idx++ {
        var curStockItem model.StockBasicInfo
        curStockItem.StockCode = stockCodeList.Elem(idx).String()
        curStockItem.StockName = stockNameList.Elem(idx).String()
        curStockItem.StockIndustry = stockIndustryList.Elem(idx).String()
        curStockItem.Exchange = stockExchangeList.Elem(idx).String()
        res = append(res, curStockItem)
    }
    return res
}

func ReadPredictionData(filePath string, getNext bool) model.SinglePredictRecord {
    var res model.SinglePredictRecord
    predFile, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
    if err != nil {
        fmt.Printf("Open File Error: %s.\n", err)
        return res
    }
    preddf := dataframe.ReadCSV(predFile)

    var compStr string
    if getNext {
        compStr = "True"
    } else {
        compStr = "False"
    }
    filterDf := preddf.Filter(
        dataframe.F{
            Colname: "is_next",
            Comparator: series.Eq,
            Comparando: compStr,
        },
    )
    res.Date = filterDf.Col("trade_date").Records()
    res.Open = filterDf.Col("open").Float()
    res.Close = filterDf.Col("close").Float()
    res.High = filterDf.Col("high").Float()
    res.Low = filterDf.Col("low").Float()
    return res
}

func ReadPredictionMsg(filePath string) string {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Open File Error: %s\n", err)
        return ""
    }
    defer func() {
        file.Close()
    }()

    var res string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        res += scanner.Text() + "\n"
    }
    return res
}
