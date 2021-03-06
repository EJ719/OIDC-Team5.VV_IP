package main

import (
	"io"
	"log"
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/utils"
)

var Logger *log.Logger

type User struct {
	Name		string	`json:"name"`
	ID		string	`json:"id"`
	Password	string	`json:"passwd"`
	TenantID	string	`json:"tenant_id"`
}

func CheckNexit(err error) {
	if err != nil {
		Logger.Println(err)
		os.Exit(1)
	}
}

func GetData(c *gin.Context) {
	var user User

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"info": user,
		"status": "you are logged in",
	})
}

func main() {
	gin.DisableConsoleColor()

	t := time.Now()
	startTime := t.Format("2021-08-09 15:04:05")
	logFile := "log/ngleLog- " + startTime

	f, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err)
	}

	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.Any("/controller", GetData)

	opts := gophercloud.AuthOptions {
		IdentityEndpoint: "https://openstack.example.com:5000/v2.0",
		TokenID:	"{token_id}",
		Username:	"{username}",
		Password:	"{password}",
		TenantID:	"{tenant_id}",
	}

	provider, err := openstack.AuthenticatedClient(opts)

	r.Run(":8080")
}
	
