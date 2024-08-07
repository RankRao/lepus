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

package query

import (
	"dbmcloud/src/database"
	"fmt"
	"net/http"
	_ "reflect"

	"github.com/gin-gonic/gin"
)

func DataSourceList(c *gin.Context) {
	method := c.Request.Method
	if method == "GET" {
		datasourceType := c.Query("type")
		if datasourceType == "" {
			c.JSON(http.StatusOK, gin.H{"success": false, "msg": "Params Error:"})
			return
		}

		sql := fmt.Sprintf("select id,host,port,name,status from datasource where enable=1 and type='%s' order by name asc", datasourceType)
		dataList, _ := database.QueryAll(sql)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "OK",
			"data":    dataList,
			"total":   len(dataList),
		})
		return
	}
}
