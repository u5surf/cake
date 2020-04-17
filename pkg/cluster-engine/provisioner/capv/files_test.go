package capv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestExtractLocalArchive(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}
	fileLoc := filepath.Join(path, "testdata/test_data.tar.gz")
	name := strings.Replace(tt.name, " ", "_", -1) + "_test_"
	dir, err := ioutil.TempDir("/tmp", name)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	targetDir, err := extractLocalArchive(tt.archiveLoc, dir)
	if err != nil {
		t.Fatal(err.Error())
	}
	// TODO add some actual tests
	fmt.Println(targetDir)

}

func TestExtractRemoteArchive(t *testing.T) {
	url := "https://github.com/kubernetes/kubernetes/releases/download/v1.18.1/kubernetes.tar.gz"
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}
	fileLoc := filepath.Join(path, "testdata/test_data.tar.gz")
	name := strings.Replace(tt.name, " ", "_", -1) + "_test_"
	dir, err := ioutil.TempDir("/tmp", name)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	targetDir, err := extractRemoteArchive(url, dir)
	if err != nil {
		t.Fatal(err.Error())
	}
	// TODO add some actual tests
	fmt.Println(targetDir)
}
