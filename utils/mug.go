// Package utils For MAIMAI | CHUN SQL DATABASE, SAVED AS USERNAME => USER UINT ID.
package utils

import (
	sql "github.com/FloatTech/sqlite"
	"strconv"
	"sync"
	"time"
)

type MaiStructSQL struct {
	UserID     uint64 `db:"userid"` // userID
	Plate      string `db:"plate"`  // plate
	Background string `db:"bg"`     // bg
}

// GeneralUserdatabase For DivingFish Service
type GeneralUserdatabase struct {
	UserID   uint64 `db:"userid"`   // userID
	UserName string `db:"username"` // UserName
}

var (
	generalDatabase = &sql.Sqlite{}
	generalLocker   = sync.Mutex{}
)

func init() {
	generalDatabase.DBPath = ResLoader("mug") + "ry.db"
	err := generalDatabase.Open(time.Hour * 24)
	if err != nil {
		return
	}
	_ = InitDataBase()
}

func BindUserGenaralInfo(userid uint64, username string) error {
	generalLocker.Lock()
	defer generalLocker.Unlock()
	return generalDatabase.Insert("userinfo", &GeneralUserdatabase{UserID: userid, UserName: username})
}

func QueryUserGeneralInfo(userid uint64) GeneralUserdatabase {
	generalLocker.Lock()
	defer generalLocker.Unlock()
	var UserinfoStruct GeneralUserdatabase
	generalDatabase.Find("userinfo", &UserinfoStruct, "WHERE userid is "+strconv.FormatUint(userid, 10))
	return UserinfoStruct
}

func QueryUserbaseMaiData(userid uint64) MaiStructSQL {
	generalLocker.Lock()
	defer generalLocker.Unlock()
	var UserinfoStruct MaiStructSQL
	generalDatabase.Find("maiuserinfo", &UserinfoStruct, "WHERE userid is "+strconv.FormatUint(userid, 10))
	return UserinfoStruct
}

func InsertUserDataModifier(userData MaiStructSQL) error {
	generalLocker.Lock()
	defer generalLocker.Unlock()
	return generalDatabase.Insert("maiuserinfo", userData)
}

func InitDataBase() error {
	generalLocker.Lock()
	defer generalLocker.Unlock()
	generalDatabase.Create("userinfo", &GeneralUserdatabase{})
	generalDatabase.Create("maiuserinfo", &MaiStructSQL{})
	return nil
}
