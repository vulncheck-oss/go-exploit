package wordpress

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/vulncheck-oss/go-exploit/config"
	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/protocol"
	"github.com/vulncheck-oss/go-exploit/random"
)

var LoginPath = `wp-login.php`

// Attempts to log into the WordPress instance and if successful return the cookies set by
// WordPress.
func Login(conf *config.Config, username, password string) ([]*http.Cookie, bool) {
	form := url.Values{}
	form.Add("log", username)
	form.Add("pwd", password)
	form.Add("wp-submit", "Login")
	url := protocol.GenerateURL(conf.Rhost, conf.Rport, conf.SSL, "/"+LoginPath)
	form.Add("redirect_to", url+"#"+random.RandLettersRange(10, 20))
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	resp, body, ok := protocol.HTTPSendAndRecvWithHeadersNoRedirect("POST", url, form.Encode(), headers)
	if !ok {
		output.PrintFrameworkError("WordPress login failed")
		output.PrintfFrameworkDebug("resp=%#v body=%q", resp, body)

		return []*http.Cookie{}, false
	}
	location, err := resp.Location()
	if err != nil {
		output.PrintFrameworkError("WordPress did not return a redirect")

		return []*http.Cookie{}, false
	}
	if location.String() != form["redirect_to"][0] {
		output.PrintFrameworkError("WordPress did not redirect to the expected location")

		return []*http.Cookie{}, false
	}
	if len(resp.Cookies()) < 1 {
		output.PrintFrameworkError("WordPress did respond with cookies")

		return []*http.Cookie{}, false
	}
	for _, cookie := range resp.Cookies() {
		if strings.Contains(strings.ToLower(cookie.Name), "wordpress") {
			return resp.Cookies(), true
		}
	}

	output.PrintFrameworkError("WordPress cookie not found")

	return []*http.Cookie{}, false
}
