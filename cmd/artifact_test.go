package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestListArtifactsCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	e := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetOut(e)
	rootCmd.SetArgs([]string{"artifact", "ls"})
	err := rootCmd.Execute()
	if !assert.Nil(t, err) {
		return
	}

	stdOut, err := io.ReadAll(b)
	if !assert.Nil(t, err) {
		return
	}

	stdErr, err := io.ReadAll(e)
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, "", string(stdErr))
	assert.Equal(t, "listing artifacts\n", string(stdOut))
}
