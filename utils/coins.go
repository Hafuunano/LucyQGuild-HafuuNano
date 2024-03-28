package utils

import (
	sql "github.com/FloatTech/sqlite"
	"strconv"
	"sync"
	"time"
)

type CoinFunction struct {
	UserID uint64 `db:"userid"` // due to qq guild use longer id uint so change it.
	Coin   int64  `db:"coins"`  // to return user real token.
	Status bool   `db:"status"` // check if user use protect mode or not.
}

// AwakenFunction Use for Sign Function.
type AwakenFunction struct {
	UserID     uint64 `db:"userid"` // UserID For here.
	AwakenTime int64  `db:"awaken"` // AwakenTime is for int64 UNIX.
	SignDay    string `db:"sign"`   // SignDay Is for user to check today is signed.
}

type TotalCountFunction struct {
	TimeDay string `db:"time"`  // TimeDay Use "YYYYMMDD"
	Count   int64  `db:"Count"` // Count Count This Day How Many Users Logined.
}

var (
	userCoinsService = sql.Sqlite{}
	userCoinsLocker  = sync.Mutex{}
)

func init() {
	userCoinsService.DBPath = ResLoader("coin") + "coins.db"
	userCoinsService.Open(time.Hour * 24)
	InitUserWalletDataBase()
}

func ModifyUserCoinFunction(fixedUsage CoinFunction) {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	userCoinsService.Insert("coins", &fixedUsage)
}

func QueryUserCoinsFunction(queryID uint64) CoinFunction {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	var queryResult CoinFunction
	userCoinsService.Find("coins", &queryResult, "WHERE userid is "+strconv.FormatUint(queryID, 10))
	return queryResult
}

func ModifyUserAwakenFunction(fixedUsage AwakenFunction) {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	userCoinsService.Insert("signs", &fixedUsage)
}

func QueryUserAwakenFunction(queryID uint64) AwakenFunction {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	var queryResult AwakenFunction
	userCoinsService.Find("signs", &queryResult, "WHERE userid is "+strconv.FormatUint(queryID, 10))
	return queryResult
}

func ModifyTotalAwakenFunction(fixedUsage TotalCountFunction) {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	userCoinsService.Insert("awaken", &fixedUsage)
}

func QueryTotalAwakenFunction() TotalCountFunction {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	var queryResult TotalCountFunction
	getTime := time.Now().Format("20060102")
	userCoinsService.Find("awaken", &queryResult, "WHERE time is '"+getTime+"'")
	return queryResult
}

func InitUserWalletDataBase() {
	userCoinsLocker.Lock()
	defer userCoinsLocker.Unlock()
	userCoinsService.Create("coins", &CoinFunction{})
	userCoinsService.Create("signs", &AwakenFunction{})
	userCoinsService.Create("awaken", &TotalCountFunction{})
}
