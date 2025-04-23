package utils

import (
	"bytes"
	"fmt"
	"iter"
	"strings"
)

type OrderedMap[K comparable, V any] struct {
	Keys []K
	Data map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		Keys: make([]K, 0),
		Data: make(map[K]V),
	}
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	if _, exists := m.Data[key]; !exists {
		m.Keys = append(m.Keys, key)
	}
	m.Data[key] = value
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, exists := m.Data[key]
	return value, exists
}

func (m *OrderedMap[K, V]) Items() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, key := range m.Keys {
			if !yield(key, m.Data[key]) {
				return
			}
		}
	}
}

// KVLM means Key-Value List Map which is a simplified version of the mail message
// format, as specified in RFC 2822.
func ParseKVLM(data []byte, parsed *OrderedMap[string, []string]) error {

	var remaining []byte
	keyValue, remaining, found := bytes.Cut(data, []byte("\n"))
	if !found {
		return fmt.Errorf("fatal: malformed kvlm data")
	}

	if string(keyValue) != "" {
		key, value, found := bytes.Cut(keyValue, []byte(" "))
		if !found {
			return fmt.Errorf("fatal: malformed kvlm data")
		}

		isMultiLineValue := bytes.HasPrefix(remaining, []byte(" "))
		for isMultiLineValue {
			var moreValue []byte
			moreValue, remaining, found = bytes.Cut(remaining, []byte("\n"))
			if !found {
				return fmt.Errorf("fatal: malformed kvlm data")
			}

			moreValue = bytes.TrimPrefix(moreValue, []byte(" "))
			moreValue = append([]byte("\n"), moreValue...)
			value = append(value, moreValue...)

			isMultiLineValue = bytes.HasPrefix(remaining, []byte(" "))
		}

		existingValue, exists := parsed.Get(string(key))
		if exists {
			parsed.Set(string(key), append(existingValue, string(value)))
		} else {
			parsed.Set(string(key), []string{string(value)})
		}
		return ParseKVLM(remaining, parsed)
	}

	// if we got here, it means that the last parsed line was just a line break
	// and the remaining data is the message
	parsed.Set("", []string{string(remaining)})

	return nil

}

func DumpKVLM(parsed *OrderedMap[string, []string]) string {
	var buffer bytes.Buffer

	for k, v := range parsed.Items() {

		if k != "" {
			for _, value := range v {
				value = strings.ReplaceAll(value, "\n", "\n ")
				buffer.WriteString(fmt.Sprintf("%s %s\n", k, value))
			}
			continue
		}

		buffer.WriteString(fmt.Sprintf("\n%s", v[0]))

	}
	return buffer.String()
}
