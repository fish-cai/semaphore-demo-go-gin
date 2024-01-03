package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	//c.HTML(http.StatusOK, "index.html", gin.H{
	//	"title":   "Home Page",
	//	"payload": articles,
	//})
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")

}

func getArticle(ctx *gin.Context) {
	if articleId, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := getArticleById(articleId); err == nil {
			ctx.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

}

func render(ctx *gin.Context, h gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		ctx.JSON(http.StatusOK, h["payload"])
	case "application/xml":
		ctx.XML(http.StatusOK, h["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, h)
	}
}
