package admin

import (
	"github.com/gin-gonic/gin"
	"time"
	"ws/app/auth"
	"ws/app/databases"
	"ws/app/file"
	"ws/app/http/requests"
	"ws/app/models"
	"ws/app/util"
)

func Me(c *gin.Context) {
	admin := auth.GetAdmin(c)
	util.RespSuccess(c, gin.H{
		"username": admin.GetUsername(),
		"id":       admin.GetPrimaryKey(),
		"avatar":   admin.GetAvatarUrl(),
	})
}
func GetChatSetting(c *gin.Context) {
	admin := auth.GetAdmin(c)
	setting := &models.AdminChatSetting{}
	databases.Db.Model(admin).Association("Setting").Find(setting)
	if setting.Id == 0 {
		setting = &models.AdminChatSetting{
			AdminId:        admin.GetPrimaryKey(),
			Background:     "",
			IsAutoAccept:   false,
			WelcomeContent: "",
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		}
		databases.Db.Save(setting)
	}
	util.RespSuccess(c, setting)
}
func UpdateChatSetting(c *gin.Context)  {
	admin := auth.GetAdmin(c)
	form  := requests.AdminChatSettingForm{}
	err := c.ShouldBind(&form)
	if err != nil {
		util.RespValidateFail(c, err.Error())
		return
	}
	setting := &models.AdminChatSetting{}
	databases.Db.Model(admin).Association("Setting").Find(setting)
	setting.Background = form.Background
	setting.IsAutoAccept = form.IsAutoAccept
	setting.WelcomeContent = form.WelcomeContent
	databases.Db.Save(setting)
	util.RespSuccess(c, gin.H{})
}
func ChatSettingImage(c *gin.Context) {
	f, _ := c.FormFile("file")
	ff, err := file.Save(f, "chat-settings")
	if err != nil {
		util.RespFail(c, err.Error(), 500)
	} else {
		util.RespSuccess(c, gin.H{
			"url": ff.FullUrl,
		})
	}
}
// 更新头像
func Avatar(c *gin.Context) {
	f, _ := c.FormFile("file")
	storage := file.Disk("local")
	fileInfo, err := storage.Save(f, "avatar")
	if err != nil {
		util.RespError(c, err.Error())
	} else {
		admin := auth.GetAdmin(c)
		admin.Avatar = fileInfo.Path
		databases.Db.Save(admin)
		util.RespSuccess(c, gin.H{})
	}
}
