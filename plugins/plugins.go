package plugins

import "strconv"

type Page struct {
	MaxPage  int64
	MaxCount int64
}

func NumberPage() int64 {
	return 30
}

/**************分页数据******************/
func CountPage(value int64) Page {
	this := Page{}
	this.MaxCount = value
	switch {
	case value <= 0:
		this.MaxPage = 0
	case value < 30:
		this.MaxPage = 1
	default:
		this.MaxPage = value / 30
		if value%30 > 0 {
			this.MaxPage += 1
		}
	}
	return this
}

/**************字符串转int64******************/
func StringtoInt64(value string) int64 {
	if path, err := strconv.Atoi(value); err != nil {
		return 0
	} else {
		return int64(path)
	}
}

/**************int64转字符串******************/
func Int64toString(value int64) string {
	return strconv.FormatInt(value, 10)
}
