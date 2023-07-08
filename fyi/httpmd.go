package fyi

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var mdDefault = `## Default
No markdown file was given. This is the default.
sample "quoted text"
`

type HttpHandler struct {
	counter  int
	port     string
	filePath string
	style    string
}

func (ct *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct.counter++
	log.Println("count:", ct.counter)
	content, err := processingContent(ct.filePath, ct.style)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, *content)
}

func RunHttpServer(filePath string, style string, port string) {
	th := &HttpHandler{counter: 0, style: style, port: port, filePath: filePath}
	http.Handle("/", th)
	http.ListenAndServe(":"+port, nil)
}

func processingContent(filePath string, style string) (*string, error) {
	htmlContent := GetHtmlContent()

	var mdBytes []byte = nil

	if filePath == "" {
		mdBytes = []byte(mdDefault)
	} else {
		mdContent, err := ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		mdBytes = []byte(*mdContent)
	}

	html := MdToHTML(mdBytes)

	htmlContent = strings.Replace(htmlContent, "{{content}}", html, -1)

	htmlContent = strings.Replace(htmlContent, "{{style}}", style, -1)

	newHtml := ""

	var rgx = regexp.MustCompile(`&ldquo;(.*?)&rdquo;`)

	for _, line := range strings.Split(strings.TrimSuffix(htmlContent, "\n"), "\n") {

		// adding copyable to clipboard to anything quoted
		rs := rgx.FindStringSubmatch(line)
		if len(rs) > 0 {
			aMatch := rs[1]
			newTag := "<a href=\"#\" class=\"copyable\">" + aMatch + "</a>"
			line = strings.Replace(line, aMatch, newTag, -1)
		}

		// adding copyable to clipboard to code block
		if strings.Contains(line, "<code>") {
			line = strings.Replace(line, "<code>", "<code class=\"copyable\">", -1)
		}

		newHtml += line + "\n"
	}

	return &newHtml, nil
}
