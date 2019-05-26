package timeline

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

//Tweet Tweetの構造体
type Tweet struct {
	TweetID   string    `json:"tweetID,omitempty"  db:"tweet_ID"`
	UserID    string    `json:"userID,omitempty"  db:"user_ID"`
	Tweet     string    `json:"tweet,omitempty"  db:"tweet"`
	CreatedAt time.Time `json:"createdAt,omitempty"  db:"created_at"`
	FavoNum   int       `json:"favoNum,omitempty"  db:"favo_num"`
}

var (
	db *sqlx.DB
)

//GetTimeLineHandler Get /timeline/:userName タイムライン
func GetTimeLineHandler(c echo.Context) error {
	userName := c.Param("userName")

	tweets := []Tweet{}
	var userID string
	db.Get(&userID, "SELECT ID FROM User WHERE name=?", userName)
	db.Select(&tweets, "SELECT * FROM Tweet WHERE user_ID=?", userID)
	return c.JSON(http.StatusOK, tweets)
}
