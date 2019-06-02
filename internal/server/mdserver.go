package server

import (
	"fmt"
	"net/http"
	"path"
	"io/ioutil"
	"log"
	"html/template"

	"internal/frontmatter"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// Post is holds simple blog post data
type Post struct {
	Slug string
	Title string
	Date string
	Filename string
	Content []byte
}

type serverResources struct {
	Posts []Post
}


// MarkdownServer starts a simple server that uses the templates given in templateDir
// to serve the Github Flavored Markdown files given in markdownDir
func MarkdownServer(port int, markdownDir string, templateDir string, staticDir string) {
	postsMap := make(map[string]Post)
	
	router := gin.Default()

	resources := loadAllResources(router, markdownDir, templateDir, staticDir)
	

	for _, post := range resources.Posts {
		postsMap[post.Slug] = post
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing.gohtml", nil)
	})

	router.GET("/posts/:postName", func(c *gin.Context) {
		postName := c.Param("postName")

		currentPost := postsMap[postName]
		postHTML := template.HTML(blackfriday.Run(currentPost.Content))
	  	  
		c.HTML(http.StatusOK, "post.gohtml", gin.H{
		  "Title":   currentPost.Title,
		  "Content": postHTML,
		})
	})

	fmt.Printf("Started server on port %d", port)
	router.Run(fmt.Sprintf(":%d", port))

}

func configureAllRoutes(router *gin.Engine) {

}

func loadAllResources(router *gin.Engine, markdownDir string, templateDir string, staticDir string) (serverResources) {
	var allResources serverResources

	loadStaticFiles(router, staticDir)
	loadTemplates(router, templateDir)
	posts, postErr := loadPosts(router, markdownDir)

	if postErr != nil {
		log.Printf("err while loading posts: %s", postErr)
	}

	allResources.Posts = posts
	return allResources
}

func loadPosts(router *gin.Engine, markdownDir string) ([]Post, error) {
	var posts []Post
	
	files, dirErr := ioutil.ReadDir(markdownDir)

	if dirErr != nil {
		log.Fatal(dirErr)
	}

	for _, file := range files {
		post := new(Post)

		filename := file.Name()
		
		mdfile, ioErr := ioutil.ReadFile(path.Join(markdownDir, filename))

		if ioErr != nil {
			return nil, ioErr
		}

		frontmatter, otherContent, frontmatterErr := frontmatter.ParseFrontmatter([]byte(mdfile))

		if frontmatterErr != nil {
			post.Title = "The Unknown Post"
			post.Date = "Mystery date"
			post.Slug = ""
		} else {
			post.Title = frontmatter["title"].(string)
			post.Date = frontmatter["date"].(string)
			post.Slug =  frontmatter["slug"].(string)
		}

		post.Filename = filename
		post.Content = otherContent

		log.Printf("%s %s %s", post.Title, post.Date, post.Filename)

		posts = append(posts, *post)
	}

	return posts, nil
}


func loadTemplates(router *gin.Engine, templateDir string) {
	router.LoadHTMLGlob(templateDir)
}

func loadStaticFiles(router *gin.Engine, staticDir string) {
	router.Static("/static", staticDir)
}
