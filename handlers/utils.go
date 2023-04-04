package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func parseBody(j json.RawMessage) (map[string]any, error) {
	var body map[string]any
	if err := json.Unmarshal(j, &body); err != nil {
		return nil, err
	}
	return body, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Logger struct {
	f *os.File
}

func NewLogger(n string) *Logger {
	f, err := os.Create(fmt.Sprintf("/tmp/dat2-n%s", n))
	check(err)
	return &Logger{
		f,
	}
}

func (l *Logger) Debug(s string) {
	var buf bytes.Buffer
	logger := log.New(&buf, "INFO: ", log.Lshortfile)
	logger.Output(2, s)

	w := bufio.NewWriter(l.f)
	_, err := w.WriteString(buf.String())
	check(err)
	w.Flush()
}

func (l *Logger) Close() {
	l.f.Close()
}
