package object

import (
	"fmt"
	"testing"

	"github.com/mrwbarg/go-git-clone/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Commit_Parent(t *testing.T) {
	commit := &Commit{}
	_ = commit.Deserialize(utils.KVLMFixture)

	parent := commit.Parent()
	assert.Equal(t, "206941306e8a8af65b66eaaaea388a7ae24d49a0", parent)
}

func Test_Commit_Message(t *testing.T) {
	commit := &Commit{}
	_ = commit.Deserialize(utils.KVLMFixture)

	message := commit.Message(false)
	assert.Equal(t, "With great power, \ncomes great responsibility.", message)

	firstLineMessage := commit.Message(true)
	assert.Equal(t, "With great power, ", firstLineMessage)
}

func Test_Commit_Author(t *testing.T) {
	commit := &Commit{}
	_ = commit.Deserialize(utils.KVLMFixture)

	author := commit.Author()
	assert.Equal(t, author, "Mauricio Barg <mbarg@email.com>")

}

func Test_Commit_Log(t *testing.T) {
	commit := &Commit{}
	_ = commit.Deserialize(utils.KVLMFixture)

	log := commit.Log()
	fmt.Println(log)
	assert.Equal(t, `commit d5343baae5a139b4f7dc39fc7a39b445d4ccad1b
Author: Mauricio Barg <mbarg@email.com>

	With great power, 
	comes great responsibility.`, log)
}
