package scaffold

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaffold(t *testing.T) {

	tempDir, err := ioutil.TempDir(filepath.Join(GoPath, "src", "scaffold"), "test")

	assert.Nil(t, err)
	if !filepath.IsAbs(tempDir) {
		tempDir, err = filepath.Abs(tempDir)
		assert.NoError(t, err)
	}

	fmt.Printf("tempDir:%s\n", tempDir)
	assert.NoError(t, New(true).Generate(tempDir, "proj"))

	// defer os.RemoveAll(tempDir) // clean up
}
 func TestCurrentDirectory(t *testing.T) {
	 currentDir, err := os.Getwd()
	// currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	assert.Nil(t, err)
	assert.Equal(t, "", currentDir)
 }