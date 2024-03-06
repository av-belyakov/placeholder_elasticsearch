package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

func NewTtpsMessageTheHive() *TtpsMessageTheHive {
	return &TtpsMessageTheHive{
		Ttp: []TtpMessage(nil),
	}
}

func (ttps *TtpsMessageTheHive) SetTtps(list []TtpMessage) {
	ttps.Ttp = list
}

func (ttps *TtpsMessageTheHive) GetTtps() []TtpMessage {
	return ttps.Ttp
}

func (ttps *TtpsMessageTheHive) Set(v TtpMessage) {
	ttps.Ttp = append(ttps.Ttp, v)
}

func (tm TtpsMessageTheHive) ToStringBeautiful(num int) string {
	return fmt.Sprintf("%s'ttp': \n%s", supportingfunctions.GetWhitespace(num), func(l []TtpMessage) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}

		return str.String()
	}(tm.Ttp))
}
