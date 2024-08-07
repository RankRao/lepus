/*
Copyright 2014-2024 The Lepus Team Group, website: https://www.lepus.cc
Licensed under the GNU General Public License, Version 3.0 (the "GPLv3 License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.gnu.org/licenses/gpl-3.0.html
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
Special note:
Please do not use this source code for any commercial purpose,
or use it for commercial purposes after secondary development, otherwise you may bear legal risks.
*/

package alarm

import (
	"dbmcloud/src/libary/mail"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DoSendEmailTest(c *gin.Context) {
	method := c.Request.Method

	if method == "POST" {
		params := make(map[string]interface{})
		c.BindJSON(&params)
		fmt.Print(params)
		emailList := params["email_list"].(string)
		mailTo := strings.Split(emailList, ";")
		mailTitle := "Lepus告警通知测试邮件"
		mailContent := "当您收到这份邮件，表明您的邮箱网关配置是正确的，可以正常发送告警邮件."
		if err := mail.Send(mailTo, mailTitle, mailContent); err != nil {
			c.JSON(200, gin.H{"success": false, "msg": "发送邮件失败:" + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "发送邮件成功",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"msg":     "method not allow.",
	})
	return
}
