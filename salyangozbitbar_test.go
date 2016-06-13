package main

import "testing"

func Test(t *testing.T) {
	var posts = new(Posts)
	posts.URL = POSTURL
	// Sunucudan json datasını alalım
	if posts.Get() != nil {
		t.Error("Sunucudan json datası alınamadı", POSTURL)
	}
	if posts.Parse() != nil {
		t.Error("Json parse edilemedi")
	}
	if posts.CheckLength() != nil {
		t.Error("json parse sonucu post sayısı 0")
	}
	if posts.Posts[0].Title == "" {
		t.Error("Post formatında Title bulunamadı")
	}
	if posts.Posts[0].URL == "" {
		t.Error("Post formatında URL bulunamadı")
	}
}
