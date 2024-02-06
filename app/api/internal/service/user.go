package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"winter_test/app/api/global"
	"winter_test/app/api/internal/consts"
	"winter_test/app/api/internal/dao/mysql"
	"winter_test/app/api/internal/dao/redis"
	"winter_test/app/api/internal/model"
	"winter_test/utils"
)

// Register 用户注册
func Register(c *gin.Context) {
	var u *model.User
	if err := c.ShouldBind(&u); err != nil {
		global.Logger.Error("user Register bind parameter failed," + err.Error())
		c.JSON(400, gin.H{
			"code": consts.ShouldBindFailed,
			"msg":  "user Register bind parameter failed," + err.Error(),
		})
		return
	}

	flag, err := redis.CheckUser(u.Username)
	if err != nil {
		global.Logger.Error("user Register bind parameter failed," + err.Error())
		c.JSON(400, gin.H{
			"code": consts.RedisCheckUserFailed,
			"msg":  "user Register bind parameter failed," + err.Error(),
		})
		return
	}

	if !flag {
		flag, err = mysql.CheckUser(u.Username)
		if err != nil {
			global.Logger.Error("user Register bind parameter failed," + err.Error())
			c.JSON(400, gin.H{
				"code": consts.MysqlCheckUserFailed,
				"msg":  "user Register bind parameter failed," + err.Error(),
			})
			return
		}

		if flag {
			c.JSON(400, gin.H{
				"code": consts.UserAlreadyExists,
				"msg":  "user already exists",
			})
			return
		}
		password := utils.Crypto(u.Password)
		err := mysql.AddUser(u.Username, password)
		if err != nil {
			return
		}

		err = redis.AddUser(u.Username, password)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "user register success",
		})
		return

	}

	c.JSON(400, gin.H{
		"code": consts.UserAlreadyExists,
		"msg":  "user already exists",
	})
	return
}

//用户个人主页和用户信息修改：
// 假设已登录，可以在用户登录后获取用户信息并展示个人主页

func userProfile(username string) {
	users := groups
	user, exists := users[username]
	if exists {
		fmt.Printf("Username: %s\nPassword: %s\n", user.Username, user.Password)
	} else {
		fmt.Println("User not found.")
	}
}

func updatePassword(username, newPassword string) {
	users := groups
	user, exists := users[username]
	if exists {
		user.Password = newPassword
		users[username] = user
		fmt.Println("Password updated successfully!")
	} else {
		fmt.Println("User not found.")
	}
}

// User 好友功能：
type User struct {
	Name     string
	Age      int
	Location string
}

func (u *User) AddFriend(friend *User) {
	fmt.Printf("%s added %s as a friend\n", u.Name, friend.Name)
}

// Group 群聊功能：
type Group struct {
	Name     string
	Members  []string
	Username string
	Password string
}

var groups = make(map[string]Group)

// Message 消息功能：
type Message struct {
	Sender    string
	Recipient string
	Content   string
}

var messages []Message

func sendMessage(sender, recipient, content string) {
	message := Message{Sender: sender, Recipient: recipient, Content: content}
	messages = append(messages, message)
	fmt.Printf("Message sent from %s to %s: %s\n", sender, recipient, content)
}
