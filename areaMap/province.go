// Package areaMap 该文件是由go generate自动生成的，请勿直接修改代码！！！
// 如需更新请更新/data文件的数据源，然后在/generate下执行 make all
package areaMap

type ProvinceId struct {
Name string `json:"name"`
Pid int `json:"pid"`
}

type ProvincePid struct {
Name string `json:"name"`
Id int `json:"id"`
}

type ProvinceName struct {
Name string `json:"name"`
Id int `json:"id"`
Pid int `json:"pid"`
}

var ProvinceById = map[int]ProvinceId{
110000:{Name:"北京",Pid:0},
120000:{Name:"天津",Pid:0},
130000:{Name:"河北省",Pid:0},
140000:{Name:"山西省",Pid:0},
150000:{Name:"内蒙古自治区",Pid:0},
210000:{Name:"辽宁省",Pid:0},
220000:{Name:"吉林省",Pid:0},
230000:{Name:"黑龙江省",Pid:0},
310000:{Name:"上海",Pid:0},
320000:{Name:"江苏省",Pid:0},
330000:{Name:"浙江省",Pid:0},
340000:{Name:"安徽省",Pid:0},
350000:{Name:"福建省",Pid:0},
360000:{Name:"江西省",Pid:0},
370000:{Name:"山东省",Pid:0},
410000:{Name:"河南省",Pid:0},
420000:{Name:"湖北省",Pid:0},
430000:{Name:"湖南省",Pid:0},
440000:{Name:"广东省",Pid:0},
450000:{Name:"广西壮族自治区",Pid:0},
460000:{Name:"海南省",Pid:0},
500000:{Name:"重庆",Pid:0},
510000:{Name:"四川省",Pid:0},
520000:{Name:"贵州省",Pid:0},
530000:{Name:"云南省",Pid:0},
540000:{Name:"西藏自治区",Pid:0},
610000:{Name:"陕西省",Pid:0},
620000:{Name:"甘肃省",Pid:0},
630000:{Name:"青海省",Pid:0},
640000:{Name:"宁夏回族自治区",Pid:0},
650000:{Name:"新疆维吾尔自治区",Pid:0},
710000:{Name:"台湾",Pid:0},
810000:{Name:"香港特别行政区",Pid:0},
820000:{Name:"澳门特别行政区",Pid:0},
990000:{Name:"海外",Pid:0},
}

var ProvinceByPid = map[int][]ProvincePid{
0:{{Name:"北京",Id:110000},{Name:"天津",Id:120000},{Name:"河北省",Id:130000},{Name:"山西省",Id:140000},{Name:"内蒙古自治区",Id:150000},{Name:"辽宁省",Id:210000},{Name:"吉林省",Id:220000},{Name:"黑龙江省",Id:230000},{Name:"上海",Id:310000},{Name:"江苏省",Id:320000},{Name:"浙江省",Id:330000},{Name:"安徽省",Id:340000},{Name:"福建省",Id:350000},{Name:"江西省",Id:360000},{Name:"山东省",Id:370000},{Name:"河南省",Id:410000},{Name:"湖北省",Id:420000},{Name:"湖南省",Id:430000},{Name:"广东省",Id:440000},{Name:"广西壮族自治区",Id:450000},{Name:"海南省",Id:460000},{Name:"重庆",Id:500000},{Name:"四川省",Id:510000},{Name:"贵州省",Id:520000},{Name:"云南省",Id:530000},{Name:"西藏自治区",Id:540000},{Name:"陕西省",Id:610000},{Name:"甘肃省",Id:620000},{Name:"青海省",Id:630000},{Name:"宁夏回族自治区",Id:640000},{Name:"新疆维吾尔自治区",Id:650000},{Name:"台湾",Id:710000},{Name:"香港特别行政区",Id:810000},{Name:"澳门特别行政区",Id:820000},{Name:"海外",Id:990000}},
}

var ProvinceByName = map[string][]ProvinceName{
"澳门特别行政区":{{Name:"澳门特别行政区",Id:820000,Pid:0}},
"海外":{{Name:"海外",Id:990000,Pid:0}},
"山西省":{{Name:"山西省",Id:140000,Pid:0}},
"内蒙古自治区":{{Name:"内蒙古自治区",Id:150000,Pid:0}},
"吉林省":{{Name:"吉林省",Id:220000,Pid:0}},
"江苏省":{{Name:"江苏省",Id:320000,Pid:0}},
"江西省":{{Name:"江西省",Id:360000,Pid:0}},
"贵州省":{{Name:"贵州省",Id:520000,Pid:0}},
"甘肃省":{{Name:"甘肃省",Id:620000,Pid:0}},
"青海省":{{Name:"青海省",Id:630000,Pid:0}},
"北京":{{Name:"北京",Id:110000,Pid:0}},
"黑龙江省":{{Name:"黑龙江省",Id:230000,Pid:0}},
"山东省":{{Name:"山东省",Id:370000,Pid:0}},
"河南省":{{Name:"河南省",Id:410000,Pid:0}},
"西藏自治区":{{Name:"西藏自治区",Id:540000,Pid:0}},
"陕西省":{{Name:"陕西省",Id:610000,Pid:0}},
"宁夏回族自治区":{{Name:"宁夏回族自治区",Id:640000,Pid:0}},
"辽宁省":{{Name:"辽宁省",Id:210000,Pid:0}},
"湖南省":{{Name:"湖南省",Id:430000,Pid:0}},
"四川省":{{Name:"四川省",Id:510000,Pid:0}},
"河北省":{{Name:"河北省",Id:130000,Pid:0}},
"香港特别行政区":{{Name:"香港特别行政区",Id:810000,Pid:0}},
"安徽省":{{Name:"安徽省",Id:340000,Pid:0}},
"海南省":{{Name:"海南省",Id:460000,Pid:0}},
"新疆维吾尔自治区":{{Name:"新疆维吾尔自治区",Id:650000,Pid:0}},
"湖北省":{{Name:"湖北省",Id:420000,Pid:0}},
"广东省":{{Name:"广东省",Id:440000,Pid:0}},
"天津":{{Name:"天津",Id:120000,Pid:0}},
"福建省":{{Name:"福建省",Id:350000,Pid:0}},
"重庆":{{Name:"重庆",Id:500000,Pid:0}},
"上海":{{Name:"上海",Id:310000,Pid:0}},
"浙江省":{{Name:"浙江省",Id:330000,Pid:0}},
"广西壮族自治区":{{Name:"广西壮族自治区",Id:450000,Pid:0}},
"云南省":{{Name:"云南省",Id:530000,Pid:0}},
"台湾":{{Name:"台湾",Id:710000,Pid:0}},
}
