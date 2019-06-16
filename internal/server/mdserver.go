package server

import (
	"fmt"
	"net/http"
	"path"
	"io/ioutil"
	"log"
	"html/template"
	"hash/crc32"

	"internal/frontmatter"

	"github.com/gin-gonic/gin"
	"github.com/shurcooL/github_flavored_markdown"
)

// Post is holds simple blog post data
type Post struct {
	Slug string
	Title string
	Date string
	Filename string
	Content []byte
	// CRC32 hash for quick identification
	Hash uint32
}

type serverResources struct {
	Posts []Post
}


// MarkdownServer starts a simple server that uses the templates given in templateDir
// to serve the Github Flavored Markdown files given in markdownDir
func MarkdownServer(port int, markdownDir string, templateDir string, staticDir string) {
	gin.SetMode(gin.ReleaseMode)

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
		freshCurrentPost, err := cachedLoadPost(&currentPost, markdownDir)

		if err != nil {
			fmt.Printf("Error while loading post: %s", err)
		}


		postHTML := template.HTML(github_flavored_markdown.Markdown(freshCurrentPost.Content))
	  	  
		c.HTML(http.StatusOK, "post.gohtml", gin.H{
		  "Title":   freshCurrentPost.Title,
		  "Content": postHTML,
		})
	})

	fmt.Printf("Started server on port %d", port)
	router.Run(fmt.Sprintf(":%d", port))

}

// Loads all the resources necessary for a static site
func loadAllResources(router *gin.Engine, markdownDir string, templateDir string, staticDir string) (serverResources) {
	var allResources serverResources

	loadStaticFiles(router, staticDir)
	loadTemplates(router, templateDir)
	posts, postErr := loadPosts(markdownDir)

	if postErr != nil {
		log.Printf("err while loading posts: %s", postErr)
	}

	allResources.Posts = posts
	return allResources
}

// Only reloads the post if the hash has changed
func cachedLoadPost(currentPost *Post, markdownDir string) (*Post, error) {

	mdfile, ioErr := ioutil.ReadFile(path.Join(markdownDir, currentPost.Filename))

	if ioErr != nil {
		return nil, ioErr
	}

	existingHash := currentPost.Hash 
	currentHash := crc32.ChecksumIEEE(mdfile)

	if currentHash != existingHash {
		fmt.Printf("Found an update to %s (%d != %d)\n", currentPost.Filename, currentHash, existingHash)
		currentPost = loadPost(mdfile, currentPost.Filename)
	}

	return currentPost, nil
}

// Loads an individual post given the markdown file bytes
func loadPost(mdfile []byte, filename string) (*Post) {
	post := new(Post)

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
	post.Hash = crc32.ChecksumIEEE(mdfile)

	log.Printf("Loaded: %s %s %s %d", post.Title, post.Date, post.Filename, post.Hash)

	return post
}

// Loads all the posts at once
func loadPosts(markdownDir string) ([]Post, error) {
	var posts []Post
	
	files, dirErr := ioutil.ReadDir(markdownDir)

	if dirErr != nil {
		log.Fatal(dirErr)
	}

	for _, file := range files {
		filename := file.Name()

		mdfile, ioErr := ioutil.ReadFile(path.Join(markdownDir, filename))

		if ioErr != nil {
			return nil, ioErr
		}

		post := loadPost(mdfile, filename)
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
