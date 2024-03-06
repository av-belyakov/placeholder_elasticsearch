package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

func NewTtpsMessageEs() *TtpsMessageEs {
	return &TtpsMessageEs{
		Ttp: make(map[string][]TtpMessage),
	}
}

func (ttp *TtpsMessageEs) GetTtp() map[string][]TtpMessage {
	return ttp.Ttp
}

func (ttp *TtpsMessageEs) SetTtp(list map[string][]TtpMessage) {
	ttp.Ttp = list
}

func (ttp *TtpsMessageEs) GetKeyTtp(k string) ([]TtpMessage, bool) {
	if value, ok := ttp.Ttp[k]; ok {
		return value, true
	}

	return nil, false
}

func (ttp *TtpsMessageEs) SetKeyTtp(k string, v []TtpMessage) {
	ttp.Ttp[k] = v
}

// SetTtp устанавливает значение для поля Ttp
func (ttp *TtpsMessageEs) SetValueTtp(v map[string][]TtpMessage) {
	ttp.Ttp = v
}

// AddValueTtp устанавливает значение для поля Ttp
func (ttp *TtpsMessageEs) AddValueTtp(k string, v TtpMessage) {
	if _, ok := ttp.Ttp[k]; !ok {
		ttp.Ttp[k] = []TtpMessage(nil)
	}

	ttp.Ttp[k] = append(ttp.Ttp[k], v)
}

func (ttp TtpsMessageEs) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	for key, value := range ttp.Ttp {
		str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
		for k, v := range value {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+2), k))
			str.WriteString(v.ToStringBeautiful(num + 3))
		}
	}

	return str.String()
}
