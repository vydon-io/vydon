package login_cmd

import (
	"fmt"
	"html/template"
	"io"
)

const (
	header = `
<!DOCTYPE html>
<head>
    <title>{{ .Title }}</title>
    <style>
        body {
            background-color: #101010;
            font-family: system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
            color: #EBECFA;
        }

        h1 {
            font-size: 32px;
            font-weight: 500;
            text-align: center;
            padding-top: 40px;
        }

        p {
            font-size: 16px;
            text-align: center;
            line-height: 1.5;
        }

        #content {
            max-width: 600px;
            margin: 0 auto;
            padding: 24px;
        }

        #footer {
            font-size: 0.8em;
            text-align: center;
            padding: 16px;
            color: #888;
        }

        .error-text {
            font-weight: 600;
        }
    </style>
</head>

<body>
    <div id="content">
	`

	footer = `
	</div>
  <div id="footer"></div>
</body>
</html>
	`

	loginPageSuccess = `
  <h1>Login Success</h1>
  <p>You've successfully logged in to the Vydon CLI.</p>
  <p>You may now close this window and return to your terminal.</p>
	`

	loginPageError = `
    <h1>There was a problem logging you in</h1>
    <p class="error-text">Error Code: {{ .ErrorCode }}</p>
    <p class="error-text">Error Description: {{ .ErrorDescription }}</p>
    <p>You may close this window and try again from your terminal.</p>
	`
)

// wraps page with header and footer
func wrapPage(contents string) string {
	return fmt.Sprintf(
		`
{{ template "header" . }}
%s
{{ template "footer" . }}
`, contents,
	)
}

type loginPageData struct {
	Title string
}

func renderLoginSuccessPage(wr io.Writer, data loginPageData) error {
	pageTmpl, err := getHtmlPage()
	if err != nil {
		return err
	}
	pageTmpl, err = pageTmpl.New("login").Parse(wrapPage(loginPageSuccess))
	if err != nil {
		return err
	}
	return pageTmpl.ExecuteTemplate(wr, "login", data)
}

type loginPageErrorData struct {
	Title string

	ErrorCode        string
	ErrorDescription string
}

func renderLoginErrorPage(wr io.Writer, data loginPageErrorData) error {
	pageTmpl, err := getHtmlPage()
	if err != nil {
		return err
	}
	pageTmpl, err = pageTmpl.New("login").Parse(wrapPage(loginPageError))
	if err != nil {
		return err
	}
	return pageTmpl.ExecuteTemplate(wr, "login", data)
}

// returns a template with the header and footer templates added in
func getHtmlPage() (*template.Template, error) {
	templ, err := template.New("header").Parse(header)
	if err != nil {
		return nil, err
	}
	templ, err = templ.New("footer").Parse(footer)
	if err != nil {
		return nil, err
	}
	return templ, nil
}
