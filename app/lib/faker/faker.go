package faker

import (
	rands "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/syyongx/php2go"
	"io/ioutil"
	"math/big"
	"speed/app/lib/faker/dataStruct"
	"strconv"
	"strings"
)

type Faker struct {
	//数据路径
	dataPath      string
	provinceData  []dataStruct.Province           //省数据
	cityData      map[uint8][]dataStruct.City     //市数据
	countryData   map[uint16][]dataStruct.Country //县级数据
	townData      map[uint16][]dataStruct.Town    //镇级数据
	nameArray     []string                        //姓名
	iDArr         []string                        //身份证号段
	isInitAddr    bool                            //是否初始化了地址
	isInitNameArr bool                            //是否初始化了姓名
	isInitIDArr   bool                            //是否初始化身份证号段数组
}

//手机号段
var mobileSegment = []string{
	"133", "153", "180", "181", "189", "177", "173", "149",
	"130", "131", "132", "155", "156", "145", "185", "186", "176",
	"175", "135", "136", "137", "138", "139", "150", "151", "152",
	"157", "158", "159", "182", "183", "184", "187", "147", "178",
}
var mailLast = []string{"@126.com",
	"@163.com", "@sina.com",
	"@21cn.com", "@sohu.com",
	"@yahoo.com.cn", "@tom.com",
	"@qq.com", "@etang.com",
	"@eyou.com", "@56.com",
	"@hotmail.com", "@msn.com", "@yahoo.com", "@gmail.com", "@aim.com", "@aol.com", "@mail.com", "@walla.com", "@inbox.com", "@live.com"}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

//10个数字
var digit = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

//百家姓
var baijiaxing = []string{
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楮",
	"卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张",
	"孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏",
	"水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦",
	"昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史",
	"唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕",
	"郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍",
	"余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹", "姚", "邵",
	"湛", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成",
	"戴", "谈", "宋", "茅", "庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁",
	"杜", "阮", "蓝", "闽", "席", "季", "麻", "强", "贾", "路", "娄", "危", "江",
	"童", "颜", "郭", "梅", "盛", "林", "刁", "锺", "徐", "丘", "骆", "高", "夏",
	"蔡", "田", "樊", "胡", "凌", "霍", "虞", "万", "支", "柯", "昝", "管", "卢",
	"莫", "经", "房", "裘", "缪", "干", "解", "应", "宗", "丁", "宣", "贲", "邓",
	"郁", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "钮", "龚", "程",
	"嵇", "邢", "滑", "裴", "陆", "荣", "翁", "荀", "羊", "於", "惠", "甄", "麹",
	"家", "封", "芮", "羿", "储", "靳", "汲", "邴", "糜", "松", "井", "段", "富",
	"巫", "乌", "焦", "巴", "弓", "牧", "隗", "山", "谷", "车", "侯", "宓", "蓬",
	"全", "郗", "班", "仰", "秋", "仲", "伊", "宫", "宁", "仇", "栾", "暴", "甘",
	"斜", "厉", "戎", "祖", "武", "符", "刘", "景", "詹", "束", "龙", "叶", "幸",
	"司", "韶", "郜", "黎", "蓟", "薄", "印", "宿", "白", "怀", "蒲", "邰", "从",
	"鄂", "索", "咸", "籍", "赖", "卓", "蔺", "屠", "蒙", "池", "乔", "阴", "郁",
	"胥", "能", "苍", "双", "闻", "莘", "党", "翟", "谭", "贡", "劳", "逄", "姬",
	"申", "扶", "堵", "冉", "宰", "郦", "雍", "郤", "璩", "桑", "桂", "濮", "牛",
	"寿", "通", "边", "扈", "燕", "冀", "郏", "浦", "尚", "农", "温", "别", "庄",
	"晏", "柴", "瞿", "阎", "充", "慕", "连", "茹", "习", "宦", "艾", "鱼", "容",
	"向", "古", "易", "慎", "戈", "廖", "庾", "终", "暨", "居", "衡", "步", "都",
	"耿", "满", "弘", "匡", "国", "文", "寇", "广", "禄", "阙", "东", "欧", "殳",
	"沃", "利", "蔚", "越", "夔", "隆", "师", "巩", "厍", "聂", "晁", "勾", "敖",
	"融", "冷", "訾", "辛", "阚", "那", "简", "饶", "空", "曾", "毋", "沙", "乜",
	"养", "鞠", "须", "丰", "巢", "关", "蒯", "相", "查", "后", "荆", "红", "游",
	"竺", "权", "逑", "盖", "益", "桓", "公", "万俟", "司马", "上官", "欧阳", "夏侯",
	"诸葛", "闻人", "东方", "赫连", "皇甫", "尉迟", "公羊", "澹台", "公冶", "宗政", "濮阳",
	"淳于", "单于", "太叔", "申屠", "公孙", "仲孙", "轩辕", "令狐", "锺离", "宇文", "长孙",
	"慕容", "鲜于", "闾丘", "司徒", "司空", "丌官", "司寇", "仉", "督", "子车", "颛孙",
	"端木", "巫马", "公西", "漆雕", "乐正", "壤驷", "公良", "拓拔", "夹谷", "宰父", "谷梁",
	"晋", "楚", "阎", "法", "汝", "鄢", "涂", "钦", "段干", "百里", "东郭", "南门",
	"呼延", "归", "海", "羊舌", "微生", "岳", "帅", "缑", "亢", "况", "后", "有", "琴",
	"梁丘", "左丘", "东门", "西门", "商", "牟", "佘", "佴", "伯", "赏", "南宫", "墨",
	"哈", "谯", "笪", "年", "爱", "阳", "佟", "第五", "言", "福",
}

//生成手机号
func (f *Faker) MakeMobile() string {
	mobile := mobileSegment[f.rand(0, len(mobileSegment)-1)]

	for i := 0; i < 8; i++ {
		mobile += strconv.Itoa(f.rand(0, 9))
	}
	return mobile
}

func (f *Faker) initNameArr() error {
	if f.isInitNameArr {
		return nil
	}

	var filePath = ""
	if f.dataPath == "" {
		filePath = "/resources/data/faker/nameData"
	} else {
		filePath = f.dataPath + "nameData"
	}
	namebytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New(err.Error())
	}

	if err = json.Unmarshal(namebytes, &f.nameArray); err != nil {
		return err
	}

	return nil

}

//生成中文名字
func (f *Faker) MakeName() (string, error) {

	if !f.isInitNameArr {
		if err := f.initNameArr(); err != nil {
			return "", err
		}
	}

	xing := baijiaxing[f.rand(0, len(baijiaxing)-1)]
	name := f.nameArray[f.rand(0, len(f.nameArray)-1)]
	return xing + name, nil
}

//随机生成单个全国省市县乡地址
func (f *Faker) MakeAddress() string {

	if !f.isInitIDArr {
		f.initAddress()
	}

	var (
		pstr        = ""
		cityStr     = ""
		countryStr  = ""
		tStr        = ""
		cityData    dataStruct.City
		countryData dataStruct.Country
		tData       dataStruct.Town
	)

	pData := f.provinceData[f.rand(0, len(f.provinceData)-1)]
	pstr = pData.Pv

	if len(f.cityData[pData.Pk])-1 < 0 {
		cityStr = ""
	} else {
		cityData = f.cityData[pData.Pk][f.rand(0, len(f.cityData[pData.Pk])-1)]
		cityStr = cityData.Cv
	}

	if len(f.countryData[cityData.Ck])-1 < 0 {
		countryStr = ""
	} else {
		countryData = f.countryData[cityData.Ck][f.rand(0, len(f.countryData[cityData.Ck])-1)]
		countryStr = countryData.Cyv
	}

	if len(f.townData[countryData.Cyk])-1 < 0 {
		tStr = ""
	} else {
		tData = f.townData[countryData.Cyk][f.rand(0, len(f.townData[countryData.Cyk])-1)]
		tStr = tData.Tv

	}

	return pstr + cityStr + countryStr + tStr
}

//生成银行卡号
func (f *Faker) MakeBankCardId() string {
	bankArea := strconv.Itoa(f.rand(1, 800) + 622126)

	for i := 0; i < 12; i++ {
		bankArea += strconv.Itoa(f.rand(0, 9))
	}

	areaArr := strings.Split(bankArea, "")
	lastNum := 0
	for i, j := len(areaArr)-1, 0; i >= 0; j, i = j+1, i-1 {
		num, _ := strconv.Atoi(areaArr[i])

		if j%2 == 0 {
			num *= 2
			num = int(num/10) + int(num%10)
		}

		lastNum += num
	}

	if lastNum%10 == 0 {
		lastNum = 0
	} else {
		lastNum = 10 - lastNum%10
	}

	return bankArea + strconv.Itoa(lastNum)
}

func (f *Faker) initIDArray() error {

	if f.isInitIDArr {
		return nil
	}
	var filePath = ""
	if f.dataPath == "" {
		filePath = "/resources/data/faker/cityidnumber"
	} else {
		filePath = f.dataPath + "cityidnumber"
	}
	cityIdNumBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(cityIdNumBytes, &f.iDArr); err != nil {
		return err
	}

	f.isInitIDArr = true
	return nil

}

//生成身份证号码
func (f *Faker) MakeIdentificationCard() (string, error) {

	if !f.isInitIDArr {
		if err := f.initIDArray(); err != nil {
			return "", err
		}
	}

	area := f.iDArr[f.rand(0, len(f.iDArr)-1)]

	//生成年
	year := 1900 + f.rand(50, 110)

	//生成月
	month := f.rand(1, 12)
	var monthStr string
	if month < 10 {
		monthStr = "0" + strconv.Itoa(month)
	} else {
		monthStr = strconv.Itoa(month)
	}

	var dayMax int

	if month == 2 {
		if year%4 <= 0 {
			dayMax = 29
		} else {
			dayMax = 28
		}
	}

	if php2go.InArray(monthStr, []string{"1", "3", "5", "7", "8", "10", "12"}) {
		dayMax = 31
	} else {
		dayMax = 30
	}
	day := f.rand(1, dayMax)

	//日期
	var dayStr = ""
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}

	//序号

	xuhao := f.rand(1, 999)
	//身份证号17位系数
	var xishu = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	tmpCode := area + strconv.Itoa(year) + monthStr + dayStr + fmt.Sprintf("%03d", xuhao)
	tmpCodeArr := strings.Split(tmpCode, "")

	var tmpSum int

	for i := 0; i < len(xishu); i++ {
		atoi, _ := strconv.Atoi(tmpCodeArr[i])
		tmpSum += xishu[i] * atoi
	}

	last := tmpSum % 11
	lastArr := []int{
		0:  1,
		1:  0,
		2:  999,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	last = lastArr[last]
	lastStr := ""
	if last >= 999 {
		lastStr = "X"
	} else {
		lastStr = strconv.Itoa(last)
	}

	return tmpCode + lastStr, nil

}

//生成电子邮箱
func (f *Faker) MakeEmail() string {
	last := mailLast[f.rand(0, len(mailLast)-1)]
	var stra, strd = "", ""
	alphabetMax := len(alphabet) - 1
	digitMax := len(digit) - 1
	for i := 0; i < 5; i++ {
		stra += alphabet[f.rand(0, alphabetMax)]
		strd += digit[f.rand(0, digitMax)]

	}
	return stra + strd + last
}

func (f *Faker) initAddress() error {

	if f.isInitAddr {
		return nil
	}
	file, err := ioutil.ReadFile(f.dataPath + "provinceData")
	if err != nil {
		return err
	}
	f.provinceData = []dataStruct.Province{}
	_ = json.Unmarshal(file, &f.provinceData)

	readFile, err := ioutil.ReadFile(f.dataPath + "cityData")
	if err != nil {
		return err
	}
	var city []dataStruct.City
	_ = json.Unmarshal(readFile, &city)

	f.cityData = map[uint8][]dataStruct.City{}

	for _, value := range city {
		f.cityData[value.Pk] = append(f.cityData[value.Pk], value)
	}

	var country []dataStruct.Country

	countryBytes, err := ioutil.ReadFile(f.dataPath + "countrydata")
	_ = json.Unmarshal(countryBytes, &country)

	f.countryData = map[uint16][]dataStruct.Country{}
	for _, value := range country {
		f.countryData[value.Ck] = append(f.countryData[value.Ck], value)
	}

	var town []dataStruct.Town

	townBytes, err := ioutil.ReadFile(f.dataPath + "townData")
	if err != nil {
		return err
	}
	_ = json.Unmarshal(townBytes, &town)
	f.townData = map[uint16][]dataStruct.Town{}
	for _, value := range town {
		f.townData[value.Cyk] = append(f.townData[value.Cyk], value)
	}
	f.isInitAddr = true
	return nil

}
func NewFaker(path string) *Faker {
	f := &Faker{dataPath: path, isInitAddr: false, isInitNameArr: false, isInitIDArr: false}
	f.initNameArr()
	f.initIDArray()
	f.initAddress()
	return f
}

func (f *Faker) rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}

	if min == max {
		return min
	}
	n, _ := rands.Int(rands.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
