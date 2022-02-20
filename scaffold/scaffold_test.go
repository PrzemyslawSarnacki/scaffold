package scaffold

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaffold(t *testing.T) {
	testCases := []struct {
		template string
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

func Test_scaffold_unmarshalJson(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "correct unmarshalling example",
			args:    args{
				jsonString: `{"name": "John", "surname": "Doe"}`,
			},
			want:    map[string]interface{}{"name": "John", "surname": "Doe"},
			wantErr: false,
		},
		{
			name:    "incorrect example",
			args:    args{
				jsonString: `{"surname": Doe}`,
			},
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scaffold{
				debug: false,
			}
			got, err := s.unmarshalJson(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("scaffold.unmarshalJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want){
				t.Errorf("scaffold.unmarshalJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

