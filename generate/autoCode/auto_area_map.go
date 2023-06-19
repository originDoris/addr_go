package autoCode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type AreaInfoId struct {
	Name    string `json:"name"`
	Pid     int    `json:"pid"`
	Zipcode int    `json:"zipcode"`
}

type AreaInfoPid struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Zipcode int    `json:"zipcode"`
}

type AreaInfoName struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Zipcode int    `json:"zipcode"`
}

// 源数据文件地址
const dataFilePath = "../data/"

// 自动生成代码地址
const packageName = "areaMap"
const codePath = "../areaMap/"

// 定义需要生成map的行政等级
var ranks = []string{"city.json", "province.json", "region.json"}

func AutoAreaMap() {
	for _, rank := range ranks {
		// 读取数据文件地址
		data, err := os.ReadFile(dataFilePath + rank)
		if err != nil {
			fmt.Println("读取json文件失败,行政等级为："+rank+"，请检查文件！", err)
			return
		}

		// 将数据序列化为 AreaInfoId
		dataMap := make(map[int]AreaInfoId)
		err = json.Unmarshal(data, &dataMap)
		if err != nil {
			fmt.Println("json序列化数据源失败！请检出数据！path："+dataFilePath+rank+"---", err)
			return
		}

		// 将map key 放入keys切片中
		var keys []int
		for id := range dataMap {
			keys = append(keys, id)
		}

		// 排序从小到大
		sort.Ints(keys)
		rank = strings.Replace(rank, ".json", "", -1)
		generateFileContent(rank, keys, dataMap, err)
	}
}

// generateFileContent 生成索引文件
func generateFileContent(rank string, keys []int, dataMap map[int]AreaInfoId, err error) {

	content := ""
	// 构建package
	content += "// Package " + packageName + " 该文件是由go generate自动生成的，请勿直接修改代码！！！\n"
	content += "// 如需更新请更新/data文件的数据源，然后在/generate下执行 make all\n"
	content += "package areaMap\n\n"

	// 构建struct
	content += "type " + cases.Title(language.Dutch).String(rank) + "Id struct {\n"
	content += "Name string `json:\"name\"`\n"
	content += "Pid int `json:\"pid\"`\n"
	if rank != "province" {
		content += "Zipcode int `json:\"zipcode\"`\n"
	}
	content += "}\n\n"
	content += "type " + cases.Title(language.Dutch).String(rank) + "Pid struct {\n"
	content += "Name string `json:\"name\"`\n"
	content += "Id int `json:\"id\"`\n"
	if rank != "province" {
		content += "Zipcode int `json:\"zipcode\"`\n"
	}
	content += "}\n\n"

	content += "type " + cases.Title(language.Dutch).String(rank) + "Name struct {\n"
	content += "Name string `json:\"name\"`\n"
	content += "Id int `json:\"id\"`\n"
	content += "Pid int `json:\"pid\"`\n"
	if rank != "province" {
		content += "Zipcode int `json:\"zipcode\"`\n"
	}
	content += "}\n\n"

	idMapContent, mapIntInterface, mapStringInterface := buildIdMap(keys, dataMap, rank)
	pidMapContent := buildPidMap(rank, mapIntInterface)
	nameMapContent := buildNameMap(rank, mapStringInterface)

	content = content + idMapContent + pidMapContent + nameMapContent

	// 尝试创建此路径
	uploadDir := codePath
	mkdirErr := os.MkdirAll(uploadDir, os.ModePerm)
	if mkdirErr != nil {
		fmt.Println(err)
	}

	// 打开文件
	file, err := os.OpenFile(codePath+rank+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)

	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()

	//及时关闭file句柄
	func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭file句柄失败！，可能回导致内存泄漏，请care一下！---", err)
		}
	}(file)
}

// buildIdMap 构建区域信息map
func buildIdMap(keys []int, dataMap map[int]AreaInfoId, rank string) (string, map[int][]interface{}, map[string][]interface{}) {
	content := "var " + cases.Title(language.Dutch).String(rank) + "ById = map[int]" + cases.Title(language.Dutch).String(rank) + "Id{\n"
	// 为构建pid索引树创造条件
	str1Arr := make(map[int][]interface{})
	// 为构建name索引树创造条件
	str2Arr := make(map[string][]interface{})
	for _, key := range keys {
		var infoPid AreaInfoPid
		infoPid.Id = key
		infoPid.Name = dataMap[key].Name
		infoPid.Zipcode = dataMap[key].Zipcode
		str1Arr[dataMap[key].Pid] = append(str1Arr[dataMap[key].Pid], infoPid)

		var infoName AreaInfoName
		infoName.Id = key
		infoName.Pid = dataMap[key].Pid
		infoName.Name = dataMap[key].Name
		infoName.Zipcode = dataMap[key].Zipcode
		str2Arr[dataMap[key].Name] = append(str2Arr[dataMap[key].Name], infoName)

		keyType := reflect.TypeOf(dataMap[key])
		keyValue := reflect.ValueOf(dataMap[key])
		name := keyType.Field(0).Name
		name1 := keyValue.Field(0).String()
		pid := keyType.Field(1).Name
		pid1 := keyValue.Field(1).Int()
		if rank == "province" {
			content += strconv.Itoa(key) + ":{" + name + ":\"" + name1 + "\"," + pid + ":" + strconv.Itoa(int(pid1)) + "},\n"
			continue
		}
		zipCode := keyType.Field(2).Name
		zipCode1 := keyValue.Field(2).Int()
		content += strconv.Itoa(key) + ":{" + name + ":\"" + name1 + "\"," + pid + ":" + strconv.Itoa(int(pid1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},\n"
	}
	content += "}\n\n"
	return content, str1Arr, str2Arr
}

// buildPidMap 构建根据行政父id(pid)来生成的索引树
func buildPidMap(rank string, mapIntInterface map[int][]interface{}) string {
	str1 := "var " + cases.Title(language.Dutch).String(rank) + "ByPid = map[int][]" + cases.Title(language.Dutch).String(rank) + "Pid{\n"
	for key, value := range mapIntInterface {
		str1 += strconv.Itoa(key) + ":{"
		for _, value2 := range value {
			t := reflect.TypeOf(value2)
			v := reflect.ValueOf(value2)
			name := t.Field(0).Name
			name1 := v.Field(0).String()
			id := t.Field(1).Name
			id1 := v.Field(1).Int()
			if rank == "province" {
				str1 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "},"
				continue
			}
			zipCode := t.Field(2).Name
			zipCode1 := v.Field(2).Int()
			str1 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},"
		}
		str1 = strings.TrimRight(str1, ",")
		str1 += "},\n"
	}
	str1 += "}\n\n"
	return str1
}

// buildNameMap 构建根据地名(name)来生成的索引树
func buildNameMap(rank string, mapStringInterface map[string][]interface{}) string {
	str2 := "var " + cases.Title(language.Dutch).String(rank) + "ByName = map[string][]" + cases.Title(language.Dutch).String(rank) + "Name{\n"
	for key, value := range mapStringInterface {
		str2 += "\"" + key + "\":{"
		for _, value2 := range value {
			t := reflect.TypeOf(value2)
			v := reflect.ValueOf(value2)
			name := t.Field(0).Name
			name1 := v.Field(0).String()
			id := t.Field(1).Name
			id1 := v.Field(1).Int()
			pid := t.Field(2).Name
			pid1 := v.Field(2).Int()
			if rank == "province" {
				str2 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + pid + ":" + strconv.Itoa(int(pid1)) + "},"
				continue
			}
			zipCode := t.Field(3).Name
			zipCode1 := v.Field(3).Int()
			str2 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + pid + ":" + strconv.Itoa(int(pid1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},"
		}
		str2 = strings.TrimRight(str2, ",")
		str2 += "},\n"
	}
	str2 += "}\n"
	return str2
}
