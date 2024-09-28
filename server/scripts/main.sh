#!/bin/bash

# Definindo variáveis de ambiente
export HTTP_PORT="4367"
export DB_NAME="ideas"
export DB_PASS="admin"
export DB_PORT="5432"
export DB_HOST="0.0.0.0"
export DB_USER="postgres"
export JWT_SECRET="segredo"

# Caso usem isso, só precisa rodar uma vez e depois comentar
# port=$HTTP_PORT
# ruleName="Allow HTTP_PORT $port"
# sudo iptables -A INPUT -p tcp --dport $port -j ACCEPT

SCRIPT_DIR="$(dirname "$(realpath "$0")")"

# Executando o programa Go
go run "$SCRIPT_DIR/../cmd/main.go"