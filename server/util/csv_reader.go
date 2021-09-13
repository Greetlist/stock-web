package util

import (
    "os"
    "path"
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

func ReadMarketRawData(stockCode, period string, dataLenGap int) model.MarketRawData {
    if dataLenGap < dataLenGapMin {
        dataLenGap = dataLenGapMin
    }
    var res model.MarketRawData
    marketRawDataPath := path.Join(conf.StockRawBaseDir, period, selectExchange(stockCode), stockCode+".csv")
    f, _ := os.OpenFile(marketRawDataPath, os.O_RDWR, os.ModePerm)
    df := dataframe.ReadCSV(f)
    startIndex := df.Nrow() - dataLenGap
    if startIndex < 0 {
        startIndex = 0
    }
    res.Date = df.Col("trade_date").Records()[startIndex:]
    res.Open = df.Col("open").Float()[startIndex:]
    res.Close = df.Col("close").Float()[startIndex:]
    res.High = df.Col("high").Float()[startIndex:]
    res.Low = df.Col("low").Float()[startIndex:]
    res.Volume = df.Col("vol").Float()[startIndex:]
    res.Money = df.Col("amount").Float()[startIndex:]
    return res
}

func ReadFactorRawData(stockCode, period string, dataLenGap int) model.FactorRawData {
    if dataLenGap < dataLenGapMin {
        dataLenGap = dataLenGapMin
    }
    var res model.FactorRawData
    maDataPath := path.Join(conf.StockMaBaseDir, period, selectExchange(stockCode), stockCode+".ma.csv")
    maFile, _ := os.OpenFile(maDataPath, os.O_RDWR, os.ModePerm)
    madf := dataframe.ReadCSV(maFile)
    startIndex := madf.Nrow() - dataLenGap
    if startIndex < 0 {
        startIndex = 0
    }
    //emaDataPath := path.Join(conf.StockEmaBaseDir, period, selectExchange(stockCode), stockCode+".ema.csv")
    //emaFile, _ := os.OpenFile(emaDataPath, os.O_RDWR, os.ModePerm)
    //emadf := dataframe.ReadCSV(emaFile)

    res.MA13 = madf.Col("MA13").Float()[startIndex:]
    res.MA34 = madf.Col("MA34").Float()[startIndex:]
    res.MA55 = madf.Col("MA55").Float()[startIndex:]

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
    for idx := 0; idx < dataLen; idx++ {
        var curStockItem model.StockBasicInfo
        curStockItem.StockCode = stockCodeList.Elem(idx).String()
        curStockItem.StockName = stockNameList.Elem(idx).String()
        res = append(res, curStockItem)
    }
    return res
}

func ReadPredictionData(filePath string) model.SinglePredictRecord {
    var res model.SinglePredictRecord
    predFile, _ := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
    preddf := dataframe.ReadCSV(predFile)

    res.Date = preddf.Col("day_to_prediction").Records()
    res.Open = preddf.Col("open").Float()
    res.Close = preddf.Col("close").Float()
    res.High = preddf.Col("high").Float()
    res.Low = preddf.Col("low").Float()
    return res
}
