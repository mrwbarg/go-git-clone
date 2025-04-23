package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OrderedMap_Set_Order(t *testing.T) {
	om := NewOrderedMap[string, []string]()

	om.Set("key1", []string{"value1"})
	om.Set("key2", []string{"value2"})
	om.Set("key3", []string{"value3"})

	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, om.Keys)
}

func Test_OrderedMap_Set_ExistingValue(t *testing.T) {
	om := NewOrderedMap[string, []string]()

	om.Set("key1", []string{"value1"})
	om.Set("key2", []string{"value2"})
	om.Set("key1", []string{"value2"})

	assert.ElementsMatch(t, []string{"key1", "key2"}, om.Keys)
	assert.Equal(t, []string{"value2"}, om.Data["key1"])
}

func Test_OrderedMap_Get(t *testing.T) {
	om := NewOrderedMap[string, []string]()

	om.Set("key1", []string{"value1"})
	om.Set("key2", []string{"value2"})

	value, exists := om.Get("key1")
	assert.True(t, exists)
	assert.Equal(t, []string{"value1"}, value)

	value, exists = om.Get("key3")
	assert.False(t, exists)
	assert.Nil(t, value)
}

func Test_OrderedMap_Items(t *testing.T) {
	om := NewOrderedMap[string, []string]()

	om.Set("key1", []string{"value1"})
	om.Set("key2", []string{"value2a", "value2b"})
	om.Set("key3", []string{"value3"})

	seenKeys := make([]string, 0)
	seenValues := make([][]string, 0)

	for k, v := range om.Items() {
		seenKeys = append(seenKeys, k)
		seenValues = append(seenValues, v)
	}

	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, seenKeys)
	assert.ElementsMatch(t, [][]string{{"value1"}, {"value2a", "value2b"}, {"value3"}}, seenValues)

}

func Test_ParseKVLM(t *testing.T) {

	parsed := NewOrderedMap[string, []string]()
	err := ParseKVLM(kvlmFixture, parsed)

	assert.NoError(t, err)

	if assert.Equal(t, 6, len(parsed.Keys)) {
		assert.ElementsMatch(
			t,
			[]string{"tree", "parent", "author", "committer", "gpgsig", ""},
			parsed.Keys,
		)
	}

	if assert.Equal(t, len(parsed.Data["tree"]), 1) {
		assert.Equal(t, "29ff16c9c14e2652b22f8b78bb08a5a07930c147", parsed.Data["tree"][0])
	}

	if assert.Equal(t, len(parsed.Data["parent"]), 1) {
		assert.Equal(t, "206941306e8a8af65b66eaaaea388a7ae24d49a0", parsed.Data["parent"][0])
	}

	if assert.Equal(t, len(parsed.Data["author"]), 2) {
		assert.Equal(t, "Mauricio Barg <mbarg@email.com> 1527025023 +0200", parsed.Data["author"][0])
		assert.Equal(t, "Peter Parker <pparker@email.com> 1527025023 +0200", parsed.Data["author"][1])
	}

	if assert.Equal(t, len(parsed.Data["committer"]), 1) {
		assert.Equal(t, "Mauricio Barg <mbarg@email.com> 1527025044 +0200", parsed.Data["committer"][0])
	}

	if assert.Equal(t, len(parsed.Data["gpgsig"]), 1) {
		assert.Equal(t, 831, len(parsed.Data["gpgsig"][0]))
	}

	if assert.Equal(t, len(parsed.Data[""]), 1) {
		assert.Equal(t, "With great power, \ncomes great responsibility.", parsed.Data[""][0])
	}

}

func Test_DumpKVLM(t *testing.T) {
	data := NewOrderedMap[string, []string]()
	data.Set("tree", []string{"29ff16c9c14e2652b22f8b78bb08a5a07930c147"})
	data.Set("parent", []string{"206941306e8a8af65b66eaaaea388a7ae24d49a0"})
	data.Set("author", []string{
		"Mauricio Barg <mbarg@email.com> 1527025023 +0200",
		"Peter Parker <pparker@email.com> 1527025023 +0200",
	})
	data.Set("committer", []string{"Mauricio Barg <mbarg@email.com> 1527025044 +0200"})
	data.Set("gpgsig", []string{`-----BEGIN PGP SIGNATURE-----

iQIzBAABCAAdFiEExwXquOM8bWb4Q2zVGxM2FxoLkGQFAlsEjZQACgkQGxM2FxoL
kGQdcBAAqPP+ln4nGDd2gETXjvOpOxLzIMEw4A9gU6CzWzm+oB8mEIKyaH0UFIPh
rNUZ1j7/ZGFNeBDtT55LPdPIQw4KKlcf6kC8MPWP3qSu3xHqx12C5zyai2duFZUU
wqOt9iCFCscFQYqKs3xsHI+ncQb+PGjVZA8+jPw7nrPIkeSXQV2aZb1E68wa2YIL
3eYgTUKz34cB6tAq9YwHnZpyPx8UJCZGkshpJmgtZ3mCbtQaO17LoihnqPn4UOMr
V75R/7FjSuPLS8NaZF4wfi52btXMSxO/u7GuoJkzJscP3p4qtwe6Rl9dc1XC8P7k
NIbGZ5Yg5cEPcfmhgXFOhQZkD0yxcJqBUcoFpnp2vu5XJl2E5I/quIyVxUXi6O6c
/obspcvace4wy8uO0bdVhc4nJ+Rla4InVSJaUaBeiHTW8kReSFYyMmDCzLjGIu1q
doU61OM3Zv1ptsLu3gUE6GU27iWYj2RWN3e3HE4Sbd89IFwLXNdSuM0ifDLZk7AQ
WBhRhipCCgZhkj9g2NEk7jRVslti1NdN5zoQLaJNqSwO1MtxTmJ15Ksk3QP6kfLB
Q52UWybBzpaP9HEd4XnR+HuQ4k2K0ns2KgNImsNvIyFwbpMUyUWLMPimaV1DWUXo
5SBjDB/V/W2JBFR+XKHFJeFwYhj7DD/ocsGr4ZMx/lgc8rjIBkI=
=lgTX
-----END PGP SIGNATURE----`})
	data.Set("", []string{"With great power, \ncomes great responsibility."})

	dumped := DumpKVLM(data)

	assert.Equal(t, string(kvlmFixture), dumped)

}
