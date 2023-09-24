package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func Import() time.Duration {
	start := time.Now()

	buf := bytes.NewBuffer([]byte(Inda)) //Inda from json-file
	dec := json.NewDecoder(buf)

	if err := dec.Decode(&Data); err != nil {
		fmt.Println(err)
	}

	return time.Since(start)
}
