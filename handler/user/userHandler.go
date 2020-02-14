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

func Login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	result := model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}

	claims := &model.JWTClaims{
		StandardClaims: jwt.StandardClaims{},
		UserID:         1,
		Password:       password,
		Username:       username,
		FullName:       username,
		Permissions:    []string{},
	}

	signedToken,err := getToken(claims)
	if err != nil {
		result.Message = "登录失败"
		c.JSON(http.StatusOK, gin.H{"result":result})
	}
	if initDB.TokenExist(signedToken) == true {
		result.Data = "Bearer" + signedToken
		c.JSON(200, gin.H{"result": result})
	}else {
		result.Message = "登陆失败，用户名与密码不匹配"
		c.JSON(200, gin.H{"result": result})
	}

}


func getToken(claims *model.JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "",errors.New(config.ErrorReasonServerBusy)
	}
	return signedToken,nil
}


func Register(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	result := model.Result{
		Code:    200,
		Message: "注册成功",
		Data:    nil,
	}

	claims := &model.JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}

	signedToken,err:=getToken(claims)
	if err != nil || initDB.TokenExist(signedToken) {
		result.Message = "注册失败，账号已存在"
		c.JSON(200, gin.H{"result":result})
	}  else {
		user := initDB.User{
			Token:    signedToken,
			Username: username,
			Sex:      "nan",
		}
		initDB.Db.Create(&user)

		c.JSON(200, gin.H{"result":result})
	}

}