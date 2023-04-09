package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
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
	f     *os.File
	n     *maelstrom.Node
	mutex sync.RWMutex
}

func NewLogger(n *maelstrom.Node) *Logger {
	return &Logger{
		f: nil,
		n: n,
	}
}

func (l *Logger) Debug(s string) {
	if l.f == nil {
		f, err := os.Create(fmt.Sprintf("/tmp/dat2-%s", l.n.ID()))
		check(err)
		l.f = f
	}

	l.mutex.Lock()
	var buf bytes.Buffer
	logger := log.New(&buf, "INFO: ", log.Lshortfile)
	logger.Output(2, s)

	w := bufio.NewWriter(l.f)
	_, err := w.WriteString(buf.String())
	check(err)
	w.Flush()
	l.mutex.Unlock()
}

func (l *Logger) Close() {
	l.f.Close()
}
