package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/syyongx/php2go"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(MakeEmail())
	name, _ := MakeName()
	fmt.Println(name)
	card, _ := MakeIdentificationCard()

	fmt.Println(card,len(card))
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
func MakeMobile() string {
	mobile := mobileSegment[php2go.Rand(0, len(mobileSegment))-1]

	for i := 0; i < 8; i++ {
		mobile += strconv.Itoa(php2go.Rand(0, 9))
	}
	return mobile
}

//生成中文名字
func MakeName() (string, error) {

	namebytes, err := ioutil.ReadFile("resources/data/faker/nameData")
	if err != nil {
		return "", errors.New(err.Error())
	}

	var nameArray []string

	if err = json.Unmarshal(namebytes, &nameArray); err != nil {
		return "", err
	}

	xing := baijiaxing[php2go.Rand(0, len(baijiaxing)-1)]
	name := nameArray[php2go.Rand(0, len(nameArray)-1)]
	return xing + name, nil

}

//随机生成单个全国省市县乡地址
func MakeAddress() {

}

//生成银行卡号
func MakeBankCardId() {

}

//生成身份证号码
func MakeIdentificationCard() (string, error) {
	cityIdNumBytes, err := ioutil.ReadFile("resources/data/faker/cityidnumber")
	if err != nil {
		return "", err
	}

	var cityIdNums []string

	if err := json.Unmarshal(cityIdNumBytes, &cityIdNums); err != nil {
		return "", err
	}

	area := cityIdNums[php2go.Rand(0, len(cityIdNums)-1)]

	//生成年
	year := 1900 + php2go.Rand(50, 110)

	//生成月
	month := php2go.Rand(1, 12)
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
	day := php2go.Rand(1, dayMax)

	//日期
	var dayStr = ""
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}

	//序号

	xuhao := php2go.Rand(1, 999)
	//身份证号17位系数
	var xishu = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	tmpCode := area + strconv.Itoa(year) + monthStr + dayStr + fmt.Sprintf("%03d",xuhao)
	tmpCodeArr := strings.Split(tmpCode,"")

	var tmpSum int

	for i := 0; i < len(xishu); i ++ {
		atoi, _ := strconv.Atoi(tmpCodeArr[i])
		tmpSum += xishu[i] * atoi
	}

	last := tmpSum % 11
	switch last {
	case 0:
		last = 1
		break
	case 1:
		last = 0
		break
	case 2:
		last = 999
		break;
	case 3:
		last = 9
		break
	case 4:
		last = 8
		break
	case 5:
		last = 7
		break
	case 6:
		last = 6
		break;
	case 7:
		last = 5
		break;
	case 8:
		last = 4
		break;
	case 9:
		last = 3
		break;
	case 10:
		last = 2
		break
	}
	lastStr := ""
	if last >= 999 {
		lastStr = "X"
	}else  {
		lastStr = strconv.Itoa(last)
	}


	return tmpCode + lastStr ,nil


}

//生成电子邮箱
func MakeEmail() string {
	last := mailLast[(php2go.Rand(0, len(mailLast)-1))]
	var stra, strd = "", ""
	alphabetMax := len(alphabet) - 1
	digitMax := len(digit) - 1
	for i := 0; i < 5; i++ {
		stra += alphabet[php2go.Rand(0, alphabetMax)]
		strd += digit[php2go.Rand(0, digitMax)]

	}
	return stra + strd + last
}