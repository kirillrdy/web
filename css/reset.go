package css

import (
	"net/http"
	"path"
	"runtime"
)

const resetCSSFilename = "reset.css"

//ResetCSSPath is the path that you should use for reset css file
const ResetCSSPath = "/assets/" + resetCSSFilename

//ServeResetCSS as promised serves reset css file
func ServeResetCSS(response http.ResponseWriter, request *http.Request) {
	_, currentFile, _, _ := runtime.Caller(0)
	packageDir := path.Dir(currentFile)

	http.ServeFile(response, request, packageDir+"/"+resetCSSFilename)
}
