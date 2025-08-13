#!/bin/bash

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para log com cores
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

if [ ! -f "./.env" ]; then
    if [ -f "./.env.example" ]; then
        log_info "Creating .env file from .env.example..."
        cp ./.env.example ./.env
        log_success ".env file created successfully"
        log_warning "Please review and update the .env file with your configuration"
    else
        log_error ".env.example file not found!"
        exit 1
    fi
else
    log_info ".env file already exists"
fi

# Verificar se todas as dependências estão instaladas
log_info "Checking Go dependencies..."
go mod tidy
if [ $? -ne 0 ]; then
    log_error "Failed to resolve Go dependencies"
    exit 1
fi

# Iniciar o servidor com Air
log_info "Starting development server with Air..."
echo "----------------------------------------"

air -c .air.linux.conf
