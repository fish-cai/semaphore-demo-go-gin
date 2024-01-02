package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	articles := getAllArticles()
	if len(articles) != len(articleList) {
		t.Fail()
	}

	for i, v := range articles {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
