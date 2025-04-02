# 🌡️ Sistema de Temperatura por CEP com OpenTelemetry e Zipkin
Este projeto consiste em dois microsserviços (Serviço A e Serviço B ) que trabalham em conjunto para fornecer informações de temperatura com base em um CEP fornecido, com instrumentação de observabilidade usando OpenTelemetry 📊 e Zipkin 🔍.

## 🏗️ Visão Geral
O sistema recebe um CEP via HTTP POST, valida-o e retorna as informações de temperatura (em Celsius, Fahrenheit e Kelvin) junto com o nome da cidade associada ao CEP.

## 🏛️ Arquitetura
**Serviço 🅰️:** Recebe e valida o CEP, encaminha para o Serviço B

**Serviço 🅱️:** Consulta localização pelo CEP, obtém dados meteorológicos e retorna formatado

**OpenTelemetry 📊 + Zipkin 🔍:** Instrumentação para tracing distribuído

## 📋 Pré-requisitos
🐳 Docker

🐙 Docker Compose

## 🚀 Como Executar

### Clone o repositório:

```
git clone https://github.com/paulagates/cep-weather-tracing.git
cd cep-weather-tracing

```

### Execute os serviços com Docker Compose:

Criar seu .env

```
cd service-b
echo "WEATHER_API_KEY=83341080b0c04ea49a1140041250603" > .env
cd ..

```
### Execute os serviços com Docker Compose:

```
docker-compose up --d

```

### Os serviços estarão disponíveis em:

Serviço 🅰️: http://localhost:8080

Zipkin UI 🔍: http://localhost:9411

## 💻 Uso
Envie uma requisição POST para o Serviço A:

```
curl -X POST http://localhost:8080/cep \
  -H "Content-Type: application/json" \
  -d '{"cep":"01001000"}'
```

### ✅ Exemplo de Resposta de Sucesso
```
{
  "city": "São Paulo",
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```
### ❌ Possíveis Respostas de Erro

CEP inválido (422):
```
{"error":"invalid zipcode"}

```
CEP não encontrado (404):
```
{"error":"can not find zipcode"}
```

## 🔍 Observabilidade
O projeto está instrumentado com OpenTelemetry 📊 e os traces podem ser visualizados no Zipkin:

Acesse o Zipkin em http://localhost:9411 🔍

Selecione o serviço que deseja visualizar (service-a 🅰️ ou service-b 🅱️)

Clique em "Run Query" para ver os traces

## ⚙️ Variáveis de Ambiente

WEATHER_API_KEY= Chave da API WeatherAPI ⛅

OTEL_SERVICE_NAME= Nome do Serviço

OTEL_EXPORTER_ZIPKIN_ENDPOINT= Endpoint do Zipkin 🔍 (http://zipkin:9411/api/v2/spans)


