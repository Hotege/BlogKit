package render

func RenderConfirmFailed() string {
    result :=
`<!DOCTYPE html>
<html>
<head>
    <title>BlogKit - Error</title>
    <script type='text/javascript'>
        function count(sec, url) {
            var jump = document.getElementById('jump');
            jump.innerHTML = sec;
            if (--sec > 0) {
                setTimeout("count(" + sec + ",'" + url + "')", 1000);
            } else {
                location.href = url;
            }
        }
    </script>
</head>
<body>
    <p>Cannot confirm password.</p>
    <span>The page will redirect after <span id='jump'>5</span> seconds</span>
    <script type='text/javascript'>
        count(5, '/');
    </script>
</body>
</html>`
    return result
}

func RenderNotFound() string {
    result :=
`<!DOCTYPE html>
<html>
<head>
    <title>BlogKit - Error</title>
    <script type='text/javascript'>
        function count(sec, url) {
            var jump = document.getElementById('jump');
            jump.innerHTML = sec;
            if (--sec > 0) {
                setTimeout("count(" + sec + ",'" + url + "')", 1000);
            } else {
                location.href = url;
            }
        }
    </script>
</head>
<body>
    <p>Page not found.</p>
    <span>The page will redirect after <span id='jump'>5</span> seconds</span>
    <script type='text/javascript'>
        count(5, '/');
    </script>
</body>
</html>`
    return result
}
