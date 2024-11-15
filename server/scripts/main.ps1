$env:HTTP_PORT = "4367"
$env:DB_NAME = "ideas"
$env:DB_PASS = "123456"
$env:DB_PORT = "5432"
$env:DB_HOST = "0.0.0.0"
$env:DB_USER = "admin"
$env:JWT_SECRET = "segredo"
$env:RESEND_API_KEY = "re_itnpGX5k_Jmt7jM11Zh834afryDH56iBP"

go run $PSScriptRoot\..\cmd\main.go
