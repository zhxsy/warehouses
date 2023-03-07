package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

func GetLastName(path string) string {
	splicedList := strings.Split(path, "/")
	// if len(splicedList) == 0 {
	//	return ""
	// }
	return splicedList[len(splicedList)-1]
}

func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ChunkInt64Slice(items []int64, chunkSize int) (chunks [][]int64) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}

func ChunkStringSlice(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if fmt.Sprintf("%d", RandInt(65, 90)) != temp {
			temp = fmt.Sprintf("%d", RandInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func RandomStringLetter(l int, letters string) bytes.Buffer {
	var res bytes.Buffer
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		r := rand.Intn(len(letters) - 1)
		res.WriteByte(letters[r])
		//res.WriteString(string(letters[r]))

	}
	//fmt.Println(res.Len())
	//fmt.Println(res.String())
	return res
}
func Int32ToString(a int32) string {
	return strconv.FormatInt(int64(a), 10)
}

func Int64ToString(a int64) string {
	return strconv.FormatInt(a, 10)
}

func Int64ToStringIgnore0(a int64) string {
	if a == 0 {
		return ""
	}
	return strconv.FormatInt(a, 10)
}

func Int64ToStringPoint(a int64) *string {
	s := strconv.FormatInt(a, 10)
	return &s
}

func PointerStringListToStringList(a []*string) []string {
	respList := make([]string, 0, len(a))
	for _, val := range a {
		if len(*val) == 0 {
			continue
		}
		respList = append(respList, *val)
	}
	return respList
}

func PointerInt32ListToInt32List(a []*int) []int32 {
	respList := make([]int32, 0, len(a))
	for _, val := range a {
		respList = append(respList, int32(*val))
	}
	return respList
}

func Int32ListToPointerInt32List(a []int32) []*int {
	respList := make([]*int, 0, len(a))
	for _, val := range a {
		intVal := int(val)
		respList = append(respList, &intVal)
	}
	return respList
}

func StringToInt64(a string) int64 {
	int64Val, int64ValErr := strconv.ParseInt(a, 10, 64)
	if int64ValErr != nil {
		return 0
	}
	return int64Val
}

func StringToInt32(a string) int32 {
	int64Val, _ := strconv.ParseInt(a, 10, 64)
	return int32(int64Val)
}

func StringToInt(a string) int {
	intVal, _ := strconv.Atoi(a)
	return intVal
}

func StringToFloat32(a string) float32 {
	val, _ := strconv.ParseFloat(a, 32)
	return float32(val)
}

func StringToFloat64(a string) float32 {
	val, _ := strconv.ParseFloat(a, 64)
	return float32(val)
}

func StringToFloat64Fix(a string) float64 {
	val, _ := strconv.ParseFloat(a, 64)
	return val
}

func IntToString(a int) string {
	return strconv.FormatInt(int64(a), 10)
}

func RandomStringFromString(l int, str string) string {
	var result bytes.Buffer
	var temp string
	length := len(str)
	for i := 0; i < l; {
		index := RandInt(0, length)

		temp = string(str[index])
		result.WriteString(temp)
		i++
	}
	return result.String()
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + int64(rand.Intn(int(max-min)))
}

func RandFloat(min, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Float64()*(max-min)
}

func Now() int {
	return int(time.Now().Unix())
}

func NowInt64() int64 {
	return time.Now().Unix()
}

func NowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func FormatTime(timestamp int64, formatStr string) string {
	if formatStr == "" {
		formatStr = "2006-01-02 15:04:05"
	}
	t := time.Unix(timestamp, 0)
	return t.Format(formatStr)
}

// Time2Unix 字符串转13位时间戳
func Time2Unix(t string, f string) int64 {
	stamp, err := time.ParseInLocation(f, t, time.Local)
	if err != nil {
		return 0
	}
	return stamp.UnixNano() / 1e6
}

func Md5(input []byte) string {
	hash := md5.New()

	// Get the 16 bytes hash
	hash.Write(input)
	hashInBytes := hash.Sum(nil)[:16]

	// Convert the bytes to a string

	return hex.EncodeToString(hashInBytes)
}

func Md5UpperCase(input []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(input)
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

func Sha1(input []byte) string {
	hash := sha1.New()

	hash.Write(input)

	return hex.EncodeToString(hash.Sum(nil))
}

func InArray(val, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

func IsArray(val interface{}) bool {
	rt := reflect.TypeOf(val)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2JsonMap(obj interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	var structByte, err = json.Marshal(obj)
	if err != nil {
		return nil
	}

	_ = json.Unmarshal(structByte, &data)

	return data
}

// 从一个指针创建该类型的slice
func MakeSliceFromPtr(ptr interface{}) interface{} {
	elemType := reflect.TypeOf(ptr)
	slice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 10)
	data := reflect.New(slice.Type())
	data.Elem().Set(slice)
	return data.Interface()
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(ints ...int) int {
	L := len(ints)
	if L == 0 {
		return 0
	}
	m := ints[0]
	for i := 1; i < L; i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
}

func MinInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

// Indirect returns last value that v points to
func Indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	return v
}

func GetFieldValue(v interface{}, field string) interface{} {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(field)
			fmt.Println(v)
		}
	}()

	if v == nil {
		return nil
	}
	r := reflect.ValueOf(v)
	if r.IsNil() {
		return nil
	}
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}

func NewStruct(model interface{}) interface{} {
	if model == nil {
		return nil
	}
	return reflect.New(Indirect(reflect.ValueOf(model)).Type()).Interface()
}

func NewSlice(model interface{}) interface{} {
	if model == nil {
		return nil
	}
	sliceType := reflect.SliceOf(reflect.TypeOf(model))
	slice := reflect.MakeSlice(sliceType, 0, 0)
	slicePtr := reflect.New(sliceType)
	slicePtr.Elem().Set(slice)
	return slicePtr.Interface()
}

func Ufirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func ArrayIsEmpty(arr interface{}) bool {
	switch arr.(type) {
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []string, []interface{}:
		if reflect.ValueOf(arr).Len() > 0 {
			return false
		} else {
			return true
		}
	default:
		return true
	}
	return true
}

func addSegment(inrune, segment []rune) []rune {
	if len(segment) == 0 {
		return inrune
	}
	if len(inrune) != 0 {
		inrune = append(inrune, '_')
	}
	inrune = append(inrune, segment...)
	return inrune
}

func CamelCaseToUnderscore(str string) string {
	var output []rune
	var segment []rune
	for _, r := range str {
		// not treat number as separate segment
		if !unicode.IsLower(r) && string(r) != "_" && !unicode.IsNumber(r) {
			output = addSegment(output, segment)
			segment = nil
		}
		segment = append(segment, unicode.ToLower(r))
	}
	output = addSegment(output, segment)
	return string(output)
}

func UnderscoreToCamelCase(inputUnderScoreStr string) (camelCase string) {
	isToUpper := false
	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return
}

func SubStr(str string, start, end int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong")
	}

	if end < start {
		return "", errors.New("end is wrong")
	}
	if end > length {
		return string(rs[start:]), nil
	}

	return string(rs[start:end]), nil
}

func Base64InBytes(max int) (string, error) {
	b, err := Bytes(maximumBytes(max))
	return base64.StdEncoding.EncodeToString(b), err
}

func maximumBytes(size int) int {
	return int((float64(size) / 4) * 3)
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func GetErrorCode(err error) string {
	type MicroError struct {
		Id     string `json:"id"`
		Detail string `json:"detail"`
		Status string `json:"status"`
	}
	if err != nil {
		microError := new(MicroError)
		_ = json.Unmarshal([]byte(err.Error()), microError)
		return microError.Detail
	}
	return ""
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func GetValueByField(field string, s interface{}) interface{} {
	elem := reflect.ValueOf(s).Elem()
	value := elem.FieldByName(field)
	switch fmt.Sprint(value.Type()) {
	case "bool":
		return value.Bool()
	case "int32":
		return int32(value.Int())
	case "int64":
		return value.Int()
	case "string":
		return value.String()
	default:
		return value.Int()
	}
}

func UniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0, len(intSlice))
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueInt64Slice(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := make([]int64, 0, len(intSlice))
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueInt32Slice(intSlice []int32) []int32 {
	keys := make(map[int32]bool)
	list := make([]int32, 0, len(intSlice))
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func PhoneFormat(phone string) string {
	return phone[:3] + "****" + phone[len(phone)-4:]
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func ToInt64(value interface{}) int64 {
	switch value := value.(type) {
	case []byte:
		n, _ := strconv.ParseInt(string(value), 10, 64)
		return n
	case string:
		n, _ := strconv.ParseInt(value, 10, 64)
		return n
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	default:
		return value.(int64)
	}
}

func ToFloat64(value interface{}) float64 {
	switch value := value.(type) {
	case string:
		val, _ := strconv.ParseFloat(value, 64)
		return val
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case float64:
		return value
	case float32:
		return float64(value)
	default:
		return float64(0)
	}
}

func ToInt64Array(input interface{}) []int64 {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	int64Array := make([]int64, len(items))
	for i, i2 := range items {
		int64Array[i] = ToInt64(i2)
	}
	return int64Array
}

// input like []*Shop or []Shop
func GetStringIds(input interface{}) []string {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}

	// Populate the rest of the items into <ids>
	var ids []string
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		id := val.FieldByName("Id").String()
		ids = append(ids, id)
	}
	return ids
}

func GetStringFields(input interface{}, name string) []string {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	var fields []string
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		field := val.FieldByName(name).String()
		fields = append(fields, field)
	}
	return fields
}

func ArrayToStringKeyMap(input interface{}, keyName, valueName string) map[string]interface{} {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	m := make(map[string]interface{})
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			elem := val.Elem()
			key := elem.FieldByName(keyName).String()
			if valueName != "" {
				m[key] = elem.FieldByName(valueName).Interface()
			} else {
				m[key] = val.Interface()
			}
		} else {
			key := val.FieldByName(keyName).String()
			if valueName != "" {
				m[key] = val.FieldByName(valueName).Interface()
			} else {
				m[key] = val.Interface()
			}
		}
	}
	return m
}

func GetInt64Ids(input interface{}) []int64 {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	// Populate the rest of the items into <ids>
	var ids []int64
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		id := val.FieldByName("Id").Interface()
		ids = append(ids, id.(int64))
	}
	return ids
}

// StringSliceIntersect get unique intersect of two []string
func StringSliceIntersect(strs1, strs2 []string) []string {
	EMPTY := struct{}{}
	map1 := map[string]interface{}{}
	for _, s := range strs1 {
		map1[s] = EMPTY
	}
	result := make([]string, 0, MinInt(len(strs1), len(strs2)))
	for _, s := range strs2 {
		if _, ok := map1[s]; ok {
			result = append(result, s)
		}
	}

	return result
}

// StringSliceUnion merge all duplicated items
func StringSliceUnion(strs1, strs2 []string) []string {
	result := make([]string, 0, len(strs1)+len(strs2))

	EMPTY := struct{}{}
	map1 := map[string]interface{}{}
	for _, s := range strs1 {
		if _, ok := map1[s]; !ok {
			result = append(result, s)
			map1[s] = EMPTY
		}
	}

	for _, s := range strs2 {
		if _, ok := map1[s]; !ok {
			result = append(result, s)
			map1[s] = EMPTY
		}
	}

	return result
}

func ContainsAnyInt32(s1, s2 []int32) bool {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

// FirstValid return first not empty str
func FirstValid(str ...string) string {
	for _, s := range str {
		if s != "" {
			return s
		}
	}
	return ""
}

func ConvertCent2Yuan(cent int64) string {
	positive := ""
	if 0 > cent {
		positive = "-"
		cent = 0 - cent
	}
	if 10 > cent {
		return positive + "0.0" + strconv.Itoa(int(cent))
	}

	return positive + strconv.Itoa(int(cent/100)) + "." + strconv.Itoa(int(cent%100))
}

// RoundBankFloat format float, keep at most p digits
func RoundBankFloat(y float64, p int32) string {
	return decimal.NewFromFloat(y).RoundBank(p).String()
}

func Int32SliceToStringSlice(a []int32) []string {
	result := make([]string, 0, len(a))
	for _, v := range a {
		result = append(result, strconv.Itoa(int(v)))
	}
	return result
}

func StringSliceToInt32Slice(a []string) []int32 {
	result := make([]int32, 0, len(a))
	for _, v := range a {
		val, _ := strconv.Atoi(v)
		result = append(result, int32(val))
	}
	return result
}

func Time0ForTimeNow(timeNow int64) int64 {
	timeTemplate3 := "2006-01-02"
	time0Str := time.Unix(timeNow, 0).Format(timeTemplate3)
	cst := time.FixedZone("CST", 8*3600)
	timeOT, _ := time.ParseInLocation(timeTemplate3, time0Str, cst)
	return timeOT.Unix()
}

// 利用反射机制，把string转成指针的type。
// 注意valPtr不要传入空指针。
func AutoUnboxValue(str string, valPtr interface{}) (err error) {
	if str == "" {
		return
	}
	str = strings.TrimSpace(str)
	directVal := reflect.Indirect(reflect.ValueOf(valPtr))
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			if e0, ok := e.(error); ok {
				err = e0
			} else {
				err = fmt.Errorf("%v", e)
			}
		}
	}()
	var newValPtr interface{}
	switch directVal.Type().Kind() {
	case reflect.String:
		newValPtr = &str
	case reflect.Int32, reflect.Int64, reflect.Int, reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		newValPtr = toTargetPtr(f, directVal.Type().Kind())
	case reflect.Bool:
		var b bool
		if str == "" || str == "0" || strings.EqualFold(str, "false") || strings.EqualFold(str, "off") {
			b = false
		} else if str == "1" || strings.EqualFold(str, "true") || strings.EqualFold(str, "on") {
			b = true
		} else {
			return errors.New("invalid bool type str:" + str)
		}
		newValPtr = &b
	default:
		// reflect accessible struct or map ptr
		newValPtr = reflect.New(directVal.Type()).Interface()
		err := json.Unmarshal([]byte(str), &newValPtr)
		if err != nil {
			return err
		}
	}
	// type convert for type alias
	newValPtr = reflect.ValueOf(newValPtr).Convert(reflect.TypeOf(valPtr)).Interface()
	if ii, ok := newValPtr.(interface{ Init() error }); ok {
		if err := ii.Init(); err != nil {
			return err
		}
	}
	directVal.Set(reflect.ValueOf(newValPtr).Elem())
	return
}

func toTargetPtr(f float64, kind reflect.Kind) interface{} {
	switch kind {
	case reflect.Int32:
		i := int32(f)
		return &i
	case reflect.Int64:
		i := int64(f)
		return &i
	case reflect.Int:
		i := int(f)
		return &i
	case reflect.Float32:
		i := float32(f)
		return &i
	default:
		return &f
	}
}

func In(slice []string, target string) bool {
	for _, n := range slice {
		if n == target {
			return true
		}
	}
	return false
}

func GetName(str string, limitLen int, Id int64) string {
	length := len(str)
	if length < limitLen {
		limitLen = length
	}
	index := length - limitLen
	horseName := str[index:] + strconv.Itoa(int(Id))
	return horseName
}

// GetNameForLength 截取字符串的长度
func GetNameForLength(str string, limitLen int) string {
	length := len(str)
	if length < limitLen {
		limitLen = length
	}
	index := length - limitLen
	horseName := str[index:]
	return horseName
}

func RandomStrBySlice(strArr []string) string {
	index := RandInt(0, len(strArr))
	return strArr[index]
}

// 获取页码偏移量
func GetOffsetByPage(page, pageSize int) int {
	return (page - 1) * pageSize
}

func GetDay() string {
	return time.Now().Format("2006-01-02")
}

// 切片中，包含某值的数量
func SliceContainsNum(s []int, val int) int {
	resp := 0
	for _, v := range s {
		if v == val {
			resp += 1
		}
	}
	return resp
}

func SliceSample(a []uint, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for index, v1 := range a {
		v2 := b[index]
		if v1 != v2 {
			return false
		}
	}
	return true
}

// 16进制字符串转int64
func Hex2Dec(val string) int64 {
	if val[0:2] == "0x" {
		val = val[2:]
	}
	n, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

// 10进制转16进制
func Dec2Hex(n int64, x bool) string {
	if x {
		return fmt.Sprintf("0x%x", n)
	}
	return fmt.Sprintf("%x", n)
}

// GenerateOrderNo 生成订单号
// 生成24位订单号
// 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func GenerateOrderNo(t time.Time) string {
	var orderNo int64
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&orderNo, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
