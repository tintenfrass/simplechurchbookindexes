package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func Import() time.Duration {
	start := time.Now()

	buf2 := bytes.NewBuffer([]byte(Inda))
	dec := json.NewDecoder(buf2)

	if err := dec.Decode(&Data); err != nil {
		fmt.Println(err)
	}

	return time.Since(start)
}
