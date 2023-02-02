package static

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"mime"
	"net/http"
	"strings"
)

// 其实不用这种方式放根路由, 而是就放个/web路由, 然后vite给个base: /web, 然后根路由重定向到/web也不是不行...
// 再有就是开两个路由, 虽然端口不一样, 但是本来放在根路由就没方便多少

//go:embed web/*
var web embed.FS

type EmbedFileSystem struct {
	fs http.FileSystem
}

func NewEmbedFileSystem(embedFS fs.FS) *EmbedFileSystem {
	return &EmbedFileSystem{http.FS(embedFS)}
}

func (e *EmbedFileSystem) Open(name string) (http.File, error) {
	_, err := e.fs.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	return e.fs.Open(name)
}

func (e *EmbedFileSystem) Exists(prefix string, filepath string) bool {
	p := strings.TrimPrefix(filepath, prefix)
	if len(p) >= len(filepath) {
		return false
	}
	_, err := e.fs.Open(p)
	if err != nil {
		return false
	}
	return true
}

func Web() gin.HandlerFunc {
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".mjs", "application/javascript")
	_ = mime.AddExtensionType(".css", "text/css")

	webSub, err := fs.Sub(web, "web")
	if err != nil {
		panic("fs.Sub(embed.FS)失败" + err.Error())
	}
	return static.Serve("/", NewEmbedFileSystem(webSub))
}

//r.Use(static.ServeRoot("/", "web")) // 不用embed
//r.StaticFS("/web", http.FS(webSub)) // gin自带的不能设置URL为/

func VueHistoryRouter(c *gin.Context) {
	if strings.Contains(c.Request.Header.Get("Accept"), "text/html") {
		content, err := web.ReadFile("web/index.html")
		if err != nil {
			panic("检查@/static/web下前端静态资源文件是否正常")
		}
		c.Data(http.StatusOK, "text/html", content)
	}
}
