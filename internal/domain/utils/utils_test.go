package utils_test

import (
	"testing"

	"github.com/Lefthander/otus-go-calendar/internal/domain/utils"
	"github.com/stretchr/testify/assert"
)

func Test_NewID(t *testing.T) {
	one := utils.NewID()
	two := utils.NewID()
	assert.NotEqual(t, one, two, "One=%v should not be equals to Two=%v", one, two)
}
