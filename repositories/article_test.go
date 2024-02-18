package repositories_test

import (
	"testing"

	"github.com/yutak03/my-api-server/models"
	"github.com/yutak03/my-api-server/repositories"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID: 1,
				Title: "First post",
				Contents: "This is my first blog",
				UserName: "Yusuke",
				NiceNum: 2,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID: 2,
				Title: "2nd",
				Contents: "Second blog post",
				UserName: "Yusuke",
				NiceNum: 4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数を実行
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}
