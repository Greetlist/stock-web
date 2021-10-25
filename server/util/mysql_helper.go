package util

import (
    "greetlist/stock-web/server/model"
    "strings"
)

func ArgsContainIllegueWord(arg interface{}) bool {
    switch arg.(type) {
    case model.LoginRequest:
        curStruct, _ := arg.(model.LoginRequest)
        return containSqlKeyWord(curStruct.Account) || containSqlKeyWord(curStruct.Password)
    }
    return false
}

func containSqlKeyWord(str string) bool {
    lowerStr := strings.ToLower(str)
    checkKeyWordList := []string{
        "select", "delete", "drop",
        "update", "create", "alter",
        "index", "change", "commit"}
    for _, keyWord := range(checkKeyWordList) {
        if strings.Contains(lowerStr, keyWord) {
            return true
        }
    }
    return false
}
