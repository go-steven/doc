<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Markdown Document: Index</title>
</head>
<body>
<div class="container">
    <ul>
    {{ range .docs }}
        <li><a href="/doc/md/{{ .Path }}">{{ .Filename }}</a></li>
    {{ end }}
    </ul>
</div>
</body>
</html>