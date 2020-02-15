package user

import (
	"app/config"
	"app/initDB"
	"app/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Token string
	Role string
)

func Login(c *gin.Context) {
	userID := c.Param("userID")
	password := c.Param("password")
	role := c.Param("role")
	result := model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}

	claims := &model.JWTClaims{
		StandardClaims: jwt.StandardClaims{},
		Password:       password,
		Username:       userID,
		FullName:       userID,
		Permissions:    []string{},
		Role:           role,
	}

	signedToken, err := getToken(claims)
	if err != nil {
		result.Message = config.ErrorReasonReLogin
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
	if initDB.TokenExist(signedToken) == true {
		result.Data = "Bearer" + signedToken
		c.JSON(200, gin.H{"result": result})
		Token = signedToken
	} else {
		result.Message = "登陆失败，用户名与密码不匹配"
		c.JSON(200, gin.H{"result": result})
	}

}

func getToken(claims *model.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", errors.New(config.ErrorReasonServerBusy)
	}
	return signedToken, nil
}

func Register(c *gin.Context) {
	userID := c.Param("userID")
	password := c.Param("password")
	role := c.Param("role")

	result := model.Result{
		Code:    200,
		Message: "注册成功",
		Data:    nil,
	}

	claims := &model.JWTClaims{
		StandardClaims: jwt.StandardClaims{},
		Password:       password,
		Username:       userID,
		FullName:       userID,
		Permissions:    []string{},
		Role:           role,
	}

	signedToken, err := getToken(claims)
	if err != nil || initDB.TokenExist(signedToken) {
		result.Message = "注册失败，账号已存在"
		c.JSON(200, gin.H{"result": result})
	} else {
		Role = role
		
		user := initDB.User{
			Token:  signedToken,
			UserID: userID,
			Roles:  role,
		}
		err := initDB.Db.Create(&user).Error
		if err != nil {
			c.JSON(200, gin.H{"result": "角色重复"})

		} else {
			c.JSON(301, gin.H{"result": result})
		}
	}
} 


func InsertMsg(c *gin.Context) {
	switch Role {
	case "student":
		student := model.Student{
			StudentID: c.PostForm("StudentID"),
			Name:      c.PostForm("Name"),
			ClassID:   c.PostForm("ClassID"),
			QQ:        c.PostForm("QQ"),
			Address:   c.PostForm("Address"),
			Phone:     c.PostForm("Phone"),
			//Families:  nil,
			Major:     c.PostForm("Major"),
			College:   c.PostForm("College"),
		}
		initDB.Db.Create(&student)
		c.String(200, config.InsertMsgSusses)
	case "teacher":
		teacher := model.Teacher{
			TeacherID: c.PostForm("TeacherID"),
			Name:      c.PostForm("Name"),
			ClassID:   c.PostForm("ClassID"),
			QQ:        c.PostForm("QQ"),
			Phone:     c.PostForm("Phone"),
			Major:     c.PostForm("Major"),
		}
		initDB.Db.Create(&teacher)
		c.String(200, config.InsertMsgSusses)
	case "admin":
		admin := model.Admin{
			AdminID: c.PostForm("AdminID"),
			Name:    c.PostForm("Name"),
		}
		initDB.Db.Create(&admin)
		c.String(200, config.InsertMsgSusses)
	default:
		c.String(200, config.ErrorInsert)
	}
	Role = "empty"
}


func Logout(c *gin.Context) {
	Token = ""
	c.String(200,config.Logout)
}
