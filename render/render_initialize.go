package render

func RenderInitialize() string {
    result := 
`
<!DOCTYPE html>
<html>
<head>
    <title>BlogKit - Initialize</title>
</head>
<body>
    <form action='initialize' method='POST'>
        <span>Mail: <input name='init_mail' /></span><br>
        <span>Username: <input name='init_name' /></span><br>
        <span>Password: <input name='init_password' type='password' /></span><br>
        <span>Confirm: <input name='init_confirm' type='password' /></span><br>
        <input type='submit' value='submit' />
    </form>
</body>
</html>`
    return result
}
