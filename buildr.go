package main
import (
	"fmt"
	"io"
	"os"
	"template"
)


const signTemplateHTML = `
<html>
  <body>
		{{html .}}
	</body>
</html>
`
var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

func sign(w io.Writer, d string) {
	err := signTemplate.Execute(w,d)
	if err != nil {
		fmt.Println("ERROR:",err)
	}
}

func main() {
	fmt.Println("Hello, world");
	sign(os.Stdout,"barf")
}
