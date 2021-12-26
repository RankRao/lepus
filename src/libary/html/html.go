/*
Copyright 2014-2021 The Lepus Team Group, website: https://www.lepus.cc
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

package html

import "reflect"

func CreateTable(title string, tableTitle string, header []string, dataList [][]string) string {
	var table, tableHeader, tableFooter, tableTh, tableTd, tableTr string
	tableHeader = "<style type='text/css'>table{border-collapse: collapse; margin: 0 auto; text-align: center;} table td, table th{border: 1px solid #cad9ea; color: #666;height: 30px;} table thead th{background-color: #CCE8EB;width: 100px;}table tr:nth-child(odd) {background: #fff;} table tr:nth-child(even) { background: #F5FAFA;} </style><table width='80%' class='table' >" +
		"<caption><center><h2>" + tableTitle + "</h2></center></caption>"
	for _, td := range header {
		tableTh = tableTh + "<th>&nbsp;" + string(td) + "</th>"
	}
	tableTh = "<thead><tr>" + tableTh + "</tr></thead>"
	for _, item := range dataList {
		tableTd = ""
		for _, col := range item {
			tableTd = tableTd + "<td>&nbsp;" + string(col) + "</td>"
		}
		tableTr = tableTr + "<tr>" + tableTd + "</tr>"

	}
	tableFooter = "</table>"
	table = tableHeader + tableTh + tableTr + tableFooter
	return table
}

func CreateTableFromSliceMap(tableTitle string, eventDetail []interface{}) string {
	var table, tableHeader, tableFooter, tableTh, tableTd, tableTr string
	tableHeader = "<style type='text/css'>table{border-collapse: collapse; margin: 1 auto; text-align: center;} table td, table th{border: 1px solid #cad9ea; color: #666;height:30px;} table thead th{background-color: #CCE8EB;width: 100px;}table tr:nth-child(odd) {background: #fff;} table tr:nth-child(even) { background: #F5FAFA;} </style><table width='80%' class='table' >" +
		"<caption><center><h2>" + tableTitle + "</h2></center></caption>"
	keys := reflect.ValueOf(eventDetail[0]).MapKeys()
	for _, key := range keys {
		tableTh = tableTh + "<th>&nbsp;" + key.String() + "</th>"
	}
	tableTh = "<thead><tr>" + tableTh + "</tr></thead>"

	for _, item := range eventDetail {
		tableTd = ""
		dataMap := item.(map[string]interface{})
		for _, key := range keys {
			var tdData = "---"
			if dataMap[key.String()] != nil {
				tdData = dataMap[key.String()].(string)
			}
			tableTd = tableTd + "<td><span style=\"white-space: pre-line\">&nbsp;" + tdData + "</span></td>"
		}
		tableTr = tableTr + "<tr>" + tableTd + "</tr>"

	}
	tableFooter = "</table>"
	table = tableHeader + tableTh + tableTr + tableFooter
	return table
}
