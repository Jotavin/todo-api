#!/bin/sh

# Verifica se as variáveis obrigatórias estão definidas
if [ -z "$GRAFANA_CLOUD_URL" ]; then
    echo "ERRO: GRAFANA_CLOUD_URL não está definida"
    exit 1
fi

if [ -z "$GRAFANA_CLOUD_USERNAME" ]; then
    echo "ERRO: GRAFANA_CLOUD_USERNAME não está definida"
    exit 1
fi

if [ -z "$GRAFANA_CLOUD_PASSWORD" ]; then
    echo "ERRO: GRAFANA_CLOUD_PASSWORD não está definida"
    exit 1
fi

# Define valores padrão para variáveis opcionais
API_TARGET_URL=${API_TARGET_URL:-todo-api:8080}
SCRAPE_INTERVAL_SECS=${SCRAPE_INTERVAL_SECS:-30}
EVAL_INTERVAL_SECS=${EVAL_INTERVAL_SECS:-30}

echo "Configurando Prometheus..."
echo "- Remote Write URL: $GRAFANA_CLOUD_URL"
echo "- User ID: $GRAFANA_CLOUD_USERNAME"
echo "- API Target: $API_TARGET_URL"
echo "- Scrape Interval: ${SCRAPE_INTERVAL_SECS}s"

# Cria diretório temporário para o arquivo de configuração
CONFIG_DIR="/tmp/prometheus"
mkdir -p "$CONFIG_DIR"

echo "Template original:"
cat /etc/prometheus/prometheus.yml.template

# Gera o arquivo de configuração final usando substituição manual mais robusta
cat > "$CONFIG_DIR/prometheus.yml" << EOF
global:
  scrape_interval: ${SCRAPE_INTERVAL_SECS}s
  evaluation_interval: ${EVAL_INTERVAL_SECS}s

# Configuração para enviar métricas ao Grafana Cloud
remote_write:
  - url: ${GRAFANA_CLOUD_URL}
    basic_auth:
      username: ${GRAFANA_CLOUD_USERNAME}
      password: ${GRAFANA_CLOUD_PASSWORD}

# Configuração dos alvos para monitoramento
scrape_configs:
  # Monitoramento do próprio Prometheus
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  # Monitoramento da sua API Go
  - job_name: 'todo-api'
    static_configs:
      - targets: ['${API_TARGET_URL}']
    metrics_path: '/metrics'
    scrape_interval: ${SCRAPE_INTERVAL_SECS}s
EOF

echo ""
echo "Configuração final gerada:"
cat "$CONFIG_DIR/prometheus.yml"

echo ""
echo "Configuração gerada com sucesso em $CONFIG_DIR/prometheus.yml"
echo "Iniciando Prometheus..."

# Inicia o Prometheus com o arquivo de configuração do diretório temporário
exec /bin/prometheus \
    --config.file="$CONFIG_DIR/prometheus.yml" \
    --storage.tsdb.path=/prometheus \
    --web.console.libraries=/etc/prometheus/console_libraries \
    --web.console.templates=/etc/prometheus/consoles \
    --web.enable-lifecycle