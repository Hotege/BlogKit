package render

func RenderSignUp() string {
    result := 
`
<!DOCTYPE html>
<html>
<head>
    <title>BlogKit - Sign Up</title>
</head>
<body>
    <form action='signup' method='POST'>
        <span>Mail: <input name='signup_mail' /></span><br>
        <span>Username: <input name='signup_name' /></span><br>
        <span>Password: <input name='signup_password' type='password' /></span><br>
        <span>Confirm: <input name='signup_confirm' type='password' /></span><br>
        <input type='submit' value='submit' />
    </form>
</body>
</html>`
    return result
}
