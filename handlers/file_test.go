package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/phanikumarps/sample-go/db"
	"github.com/phanikumarps/sample-go/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFileHandler(t *testing.T) {

	type input struct {
		id string
	}
	type tests struct {
		name      string
		inputData input
		expected  db.File
	}

	testCases := []tests{
		{
			name:      "text file with no data",
			inputData: input{id: "1cac3dc4-4759-4340-8cb8-222c174e8d6d"},
			expected: db.File{
				ID:        uuid.MustParse("1cac3dc4-4759-4340-8cb8-222c174e8d6d"),
				Name:      "test.txt",
				Data:      nil,
				CreatedBy: uuid.MustParse("53f134f5-ae99-4a16-b48d-885b95aa134b"),
				CreatedAt: time.Now(),
				MimeType:  "text/plain; charset=utf-8",
			}},
		{
			name:      "pdf file with no data",
			inputData: input{id: "1cac3dc4-4759-4340-8cb8-222c174e8d6d"},
			expected: db.File{
				ID:        uuid.MustParse("1cac3dc4-4759-4340-8cb8-222c174e8d6d"),
				Name:      "test.pdf",
				Data:      nil,
				CreatedBy: uuid.MustParse("53f134f5-ae99-4a16-b48d-885b95aa134b"),
				CreatedAt: time.Now(),
				MimeType:  "pdf",
			}},
		{
			name:      "xls file with no data",
			inputData: input{id: "1cac3dc4-4759-4340-8cb8-222c174e8d6d"},
			expected: db.File{
				ID:        uuid.MustParse("1cac3dc4-4759-4340-8cb8-222c174e8d6d"),
				Name:      "test.txt",
				Data:      nil,
				CreatedBy: uuid.MustParse("53f134f5-ae99-4a16-b48d-885b95aa134b"),
				CreatedAt: time.Now(),
				MimeType:  "xls",
			}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/file/%s", tc.inputData.id)
			req, _ := http.NewRequest("GET", url, nil)
			http.Handler(FileHandler(context.TODO(), mockStore)).ServeHTTP(rec, req)

			var actualResults []db.File
			err := json.Unmarshal(rec.Body.Bytes(), &actualResults)
			if err != nil {
				t.Errorf("unmarshal failed for actual = %s", actualResults)
			}

			for _, actual := range actualResults {
				if !utils.IsEqualIgnoringTags(tc.expected, actual, "ignore") {
					t.Errorf("\nexpected = %v, \nactual = %v", tc.expected, actual)
				}
			}
		})
	}
}

var mockStore mockFile

type mockFile struct{}

func (m mockFile) ReadFile(ctx context.Context, id string) (*[]db.File, error) {
	var f []db.File

	c := "53f134f5-ae99-4a16-b48d-885b95aa134b"
	f = append(f, db.File{
		ID:        uuid.MustParse(id),
		Name:      "test.txt",
		Data:      nil,
		CreatedBy: uuid.MustParse(c),
		CreatedAt: time.Now(),
		MimeType:  "text/plain; charset=utf-8",
	})
	return &f, nil
}

func testFiles() []db.File {

	return []db.File{
		{
			ID:        uuid.MustParse("1cac3dc4-4759-4340-8cb8-222c174e8d6d"),
			Name:      "test.txt",
			Data:      nil,
			CreatedBy: uuid.MustParse("53f134f5-ae99-4a16-b48d-885b95aa134b"),
			CreatedAt: time.Now(),
			MimeType:  "text/plain; charset=utf-8",
		},
		{
			ID:        uuid.MustParse("1cac3dc4-4759-4340-8cb8-222c174e8d6d"),
			Name:      "test.pdf",
			Data:      nil,
			CreatedBy: uuid.MustParse("53f134f5-ae99-4a16-b48d-885b95aa134b"),
			CreatedAt: time.Now(),
			MimeType:  "pdf",
		},
	}
}
