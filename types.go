package apiHitter
import (
	"net/http"
	"time"
)

//ContentType specify Http request Content-Type
type ContentType string //Alias of string Type

type Payload struct {
	URL         string
	Method      string // defoult GET
	Header      map[string]string
	Params      map[string]string
	Body        []byte
	ContentType ContentType // for specify Content-Type
	Timeout     time.Duration

	// unexported var
	client *http.Client
	req    *http.Request
}

//HitResponse contains *http.Response with two extra field Data and Code
//where Data = ioutild.ReadAll(*http.Response.Body)  and Code = *http.Response.StatusCode)
type HitResponse struct {
	resp *http.Response
	Data []byte
	Code int
}
