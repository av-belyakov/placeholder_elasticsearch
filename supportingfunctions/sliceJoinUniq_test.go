package supportingfunctions_test

import (
	"testing"

	"placeholder_elasticsearch/supportingfunctions"

	"github.com/stretchr/testify/assert"
)

func TestSliceJoinUniq(t *testing.T) {
	listOne := []string{"red", "green", "yellow", "white"}
	listTwo := []string{"black", "white", "blue", "grey", "orange", "green"}

	assert.True(t, supportingfunctions.SliceElemIsExist("white", listTwo))

	list := supportingfunctions.SliceJoinUniq[string](listOne, listTwo)

	assert.Equal(t, len(list), 8)
}
