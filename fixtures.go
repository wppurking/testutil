package testutil

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/go-testfixtures/testfixtures/v3"
)

// 提供 db 的 fixtures 的辅助方法
var fixturePath = "fixtures"

// FixturesHelper 某一个 Fixutres 的 helper
type FixturesHelper struct {
	fixtures *testfixtures.Loader
}

// NewFixtureHelper 初始化一个 FixturesHelper
func NewFixtureHelper(db *sql.DB, dialect string) *FixturesHelper {
	var folderName string
	guessFolders := []string{fixturePath, "../" + fixturePath, "../../" + fixturePath}
	for _, fd := range guessFolders {
		folderName, _ = filepath.Abs(fd)
		_, err := ioutil.ReadDir(folderName)
		if err == nil {
			break
		}
	}
	if folderName == "" {
		panic("无法确定 fixutres 文件目录")
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect(dialect),
		testfixtures.Directory(folderName),
	)
	if err != nil {
		panic(err)
	}
	return &FixturesHelper{fixtures}
}

// ResetDB 用于重置 DB
func (fh FixturesHelper) ResetDB() {
	if err := fh.fixtures.Load(); err != nil {
		panic(errors.Wrap(err, "数据库 Reset 失败:"))
	}
}
