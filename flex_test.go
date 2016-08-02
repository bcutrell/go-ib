// http://keighl.com/post/mocking-http-responses-in-golang/
// http://stackoverflow.com/questions/36786068/golang-copy-remote-file-to-local-folder-using-sftp-golang-library

package go_ib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func testFlex() {
	assert(true == true)
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
