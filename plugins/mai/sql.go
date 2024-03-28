// Package mai COPY FROM REI LUCY ^^
package mai

import (
	sql "github.com/FloatTech/sqlite"
	"github.com/moyoez/HafuuNano/utils"
	"strconv"
	"sync"
	"time"
)

type UserSwitcherService struct {
	TGId   uint64 `db:"tgid"`
	IsUsed bool   `db:"isused"` // true == lxns service \ false == Diving Fish.
}

type UserIDToMaimaiFriendCode struct {
	TelegramId uint64 `db:"telegramid"`
	MaimaiID   int64  `db:"friendid"`
}

// UserIDToQQ TIPS: Onebot path, actually it refers to (*Telegram*) Nope This is QGUILD Userid.

type UserIDToQQ struct {
	QQ     uint64 `db:"user_qq"` // qq nums
	Userid string `db:"user_id"` // user_id
}

type UserIDToToken struct {
	UserID string `db:"user_id"`
	Token  string `db:"user_token"`
}

var (
	maiDatabase = &sql.Sqlite{}
	maiLocker   = sync.Mutex{}
)

func init() {
	maiDatabase.DBPath = utils.ResLoader("mai") + "maisql.db"
	err := maiDatabase.Open(time.Hour * 24)
	if err != nil {
		return
	}
	_ = InitDataBase()
}

func FormatUserToken(tgid string, token string) *UserIDToToken {
	return &UserIDToToken{Token: token, UserID: tgid}
}

func FormatUserIDDatabase(qq uint64, userid string) *UserIDToQQ {
	return &UserIDToQQ{QQ: qq, Userid: userid}
}

func FormatUserSwitcher(tgid uint64, isSwitcher bool) *UserSwitcherService {
	return &UserSwitcherService{TGId: tgid, IsUsed: isSwitcher}
}

func FormatMaimaiFriendCode(friendCode int64, tgid uint64) *UserIDToMaimaiFriendCode {
	return &UserIDToMaimaiFriendCode{TelegramId: tgid, MaimaiID: friendCode}
}

func InitDataBase() error {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	maiDatabase.Create("useridinfo", &UserIDToQQ{})
	maiDatabase.Create("userswitchinfo", &UserSwitcherService{})
	maiDatabase.Create("usermaifriendinfo", &UserIDToMaimaiFriendCode{})
	maiDatabase.Create("usertokenid", &UserIDToToken{})
	return nil
}

func GetUserMaiFriendID(userid uint64) UserIDToMaimaiFriendCode {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	var infosql UserIDToMaimaiFriendCode
	userIDStr := strconv.FormatUint(userid, 10)
	_ = maiDatabase.Find("usermaifriendinfo", &infosql, "where telegramid is "+userIDStr)
	return infosql
}

func GetUserSwitcherInfoFromDatabase(userid uint64) bool {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	var info UserSwitcherService
	userIDStr := strconv.FormatUint(userid, 10)
	err := maiDatabase.Find("userswitchinfo", &info, "where tgid is "+userIDStr)
	if err != nil {
		return false
	}
	return info.IsUsed
}

func (info *UserSwitcherService) ChangeUserSwitchInfoFromDataBase() error {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	return maiDatabase.Insert("userswitchinfo", info)
}

// GetUserIDFromDatabase Params: user qq id ==> user maimai id.
func GetUserIDFromDatabase(userID uint64) UserIDToQQ {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	var infosql UserIDToQQ
	userIDStr := strconv.FormatUint(userID, 10)
	_ = maiDatabase.Find("useridinfo", &infosql, "where user_qq is "+userIDStr)
	return infosql
}

func (info *UserIDToQQ) BindUserIDDataBase() error {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	return maiDatabase.Insert("useridinfo", info)
}

func (info *UserIDToMaimaiFriendCode) BindUserFriendCode() error {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	return maiDatabase.Insert("usermaifriendinfo", info)
}

func GetUserToken(userid string) string {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	var infosql UserIDToToken
	maiDatabase.Find("usertokenid", &infosql, "where user_id is "+userid)
	if infosql.Token == "" {
		return ""
	}
	return infosql.Token
}

func (info *UserIDToToken) BindUserToken() error {
	maiLocker.Lock()
	defer maiLocker.Unlock()
	return maiDatabase.Insert("usertokenid", info)
}
