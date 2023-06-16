package Handlers

import (
	"main/DB"
	"main/helpers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var query, field bson.M

func IndexHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	var role, user string
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	} else {
		role, user, _ = helpers.FindRole(c)
	}
	if role == "user" {
		c.HTML(http.StatusOK, "dayout.html", user)
	} else {
		c.Redirect(http.StatusFound, "/admin")
	}
}

func SignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPost(c *gin.Context) {
	userName := c.Request.FormValue("UserName")
	userEmail := c.Request.FormValue("Emailid")
	password := c.Request.FormValue("Password")
	user := models.User{Role: "user", Name: userName, Email: userEmail, Password: password}
	_, err := DB.InsertOne("sample", "user", user)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusFound, "/login")

}

func LoginHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func LoginPost(c *gin.Context) {
	var e models.Invalid
	newmail := c.Request.FormValue("Email")
	newpassword := c.Request.FormValue("Password")
	query = bson.M{"email": newmail}
	field = bson.M{"role": 1, "email": 1, "password": 1}
	result, err := DB.FindUser(query, field, "sample", "user")
	if err != nil {
		e.Errmail = "invalid email id"
		c.HTML(http.StatusOK, "login.html", e)
	} else if result.Password != newpassword {
		e.Errpass = "invalid password"
		c.HTML(http.StatusOK, "login.html", e)
	} else if result.Role == "user" {
		helpers.CreateToken(*result, c)
		c.Redirect(http.StatusFound, "/")
	} else {
		helpers.CreateToken(*result, c)
		c.Redirect(http.StatusFound, "/admin")
	}
}

func HomeHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := helpers.ValidateCookie(c)
	if ok == false {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func LogoutHandler(c *gin.Context) {
	helpers.DeleteCookie(c)
	c.Redirect(http.StatusFound, "/login")
}
