package main

import "html/template"

var mainHeader *template.Template = template.Must(template.New("head").Parse(`
<html>
<head>
<link rel="stylesheet" href="/static/css/main.css">
`))

var redirectTemplate *template.Template = template.Must(template.New("err_redirect").Parse(`
<meta http-equiv="Refresh" content="5; url={{ .Url }}" />
</head>
<body>
<p>{{ .Msg }}: Please follow <a href="{{ .Url }}">this link</a> if the page doesn't refresh.</p>
</body>
</html>
`))

var userTemplate *template.Template = template.Must(template.New("user").Parse(`
</head>
<body>
<p>Recruiter: {{.}}</p>
</body>
</html>
`))

var userSuccessTemplate *template.Template = template.Must(template.New("user_success").Parse(`
</head>
<body>
<p>Successfully added user!</p>
</body>
</html>
`))

var addUserTemplate *template.Template = template.Must(template.New("add_user").Parse(`
</head>
<body>
<p>
<form action="/adduser" method="get">
User name: <input type="text" name="user">
<input type="submit" value="Submit">
</p>
<p>
Hint: If you're getting 'Invalid User' errors, log into the offical forums <a href="https://boards.faeria.com">here</a>.
</p>
</form>
</body>
</html>
`))

var homeTemplate *template.Template = template.Must(template.New("root").Parse(`
</head>
<body>
<p><a href="/adduser">Click here to add yourself as a recruiter.</a></p>
<p><a href="/getuser">Or here to get a random recruiter.</a></p>
</body>
</html>
`))
