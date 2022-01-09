package dictionary

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const defaultlist = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"

type Dictionary struct {
	Words map[string]interface{}
}

func New(ctx context.Context) (*Dictionary, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		defaultlist,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching word list: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error building dictionary from list: %w", err)
	}

	bodyStr := string(body)

	dict := &Dictionary{
		Words: make(map[string]interface{}),
	}

	for _, str := range strings.Split(bodyStr, "\n") {
		dict.Words[strings.TrimSpace(str)] = struct{}{}
	}

	return dict, nil
}
