package llm_test

import (
	"go-cleanarch/internal/llm"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLLMTargetJson(t *testing.T) {

	type LossItem struct {
		Brand string `json:"brand"`
		Color string `json:"color"`
		Item  string `json:"item"`
	}

	t.Run("Simple Request", func(t *testing.T) {
		plainText := "I lost a iphone 15 black"

		var lostItem LossItem
		err := llm.GetLLMTargetJson(plainText, &lostItem)
		require.NoError(t, err)

		t.Logf("%+v", lostItem)
	})

}
