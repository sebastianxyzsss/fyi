package fyi

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	html := markdown.Render(doc, renderer)
	return string(html)
}

func GetHtmlContent() string {
	content := `<!DOCTYPE html>
<html>

<body class="darkmode">

<style>
{{style}}
</style>
<div class"cpcontainer">

{{content}}

</div>

<h4>** click links to copy/paste</h4>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
<script>
$(".copyable").click(function(event) {
	console.log("click Function copyable");
    var copyText = $(event.target).text();
	console.log(copyText);
	// Copy the text inside the text field
	navigator.clipboard.writeText(copyText);
});
</script>

</body>

</html>`
	return content
}
