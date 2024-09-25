$env:HTTP_PORT = "4367"
$env:DB_NAME = "ideas"
$env:DB_PASS = "123456"
$env:DB_PORT = "5432"
$env:DB_HOST = "0.0.0.0"
$env:DB_USER = "admin"
$env:JWT_SECRET = "segredo"


# Caso usem isso, só precisa rodar uma vez e depois recomentar
# $port = $env:HTTP_PORT
# $ruleName = "Allow HTTP_PORT $port"
# netsh advfirewall firewall add rule name=$ruleName dir=in action=allow protocol=TCP localport=$port

go run $PSScriptRoot\..\cmd\main.go
