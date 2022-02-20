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
	testCases := []struct {
		template	string
		
	}{
		{
			template: "c9",
		},
		{
			template: "fiber",
		},
		{
			template: "modern",
		},
		{
			template: "template",
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("generate code for template named: %s", tC.template), func(t *testing.T) {
			
			tempDir, err := ioutil.TempDir(filepath.Join(GoPath, "src", "scaffold"), "test")
		
			assert.Nil(t, err)
			if !filepath.IsAbs(tempDir) {
				tempDir, err = filepath.Abs(tempDir)
				assert.NoError(t, err)
			}
		
			fmt.Printf("tempDir:%s\n", tempDir)
			assert.NoError(t, New(true).Generate(tempDir, "project", tC.template, "module"))
			
			defer os.RemoveAll(tempDir) // clean up
		})
	}
}
 func TestCurrentDirectory(t *testing.T) {
	 currentDir, err := os.Getwd()
	// currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	assert.Nil(t, err)
	assert.Equal(t, "", currentDir)
 }