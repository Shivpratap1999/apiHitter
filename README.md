# apiHitter
Simple go 1.18 package to consume an API through http methods like GET,POST,PUT,DELETE,PACH etc using Curl. (created to consume APIs, and could be used for other things) Methods build API URL params, connection url, and parse expected JSON response.

<h1>Installation</h1>
<p>run command <br>
<a href="github.com/Shivpratap1999/apiHitter">go get "github.com/Shivpratap1999/apiHitter"</a><br>
After running above comond import package :<br><br>
import "github.com/Shivpratap1999/apiHitter"<br>

<h2>How to use apiHItter</h2>
payload := apiHitter.Payload{<br>
		URL:         "http.google.com", <br>
		Method:      "POST", //this may any http method<br>
		Body:        Data, // Data should by of type []byte<br>
		ContentType: apiHitter.ApplicationJSON,<br>
	}<br>
	resp, err := payload.Hit() <br>
	if err != nil || resp.Code != http.StatusOK{<br>
		fmt.Printf(error:%s while hitting url",err )<br>
	}<br>
  fmt.Printf("Response Code: %d and Response Data :%s",resp.Code , resp.Data")<br>
  //resp is of type HitResponse which is a custom response having fields - Data []byte ,resp *http.Response , Code int.<br>
