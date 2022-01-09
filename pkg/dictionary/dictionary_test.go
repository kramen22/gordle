package dictionary_test

import (
	"context"
	"testing"

	"github.com/kramen22/gordle/pkg/dictionary"
)

func TestEverything(t *testing.T) {
	_, _ = dictionary.New(context.TODO())
}
