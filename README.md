## Nats and JetStream

### Install

1. kubectl apply -f https://raw.githubusercontent.com/nats-io/nack/v0.6.0/deploy/crds.yml
2. helm repo add nats https://nats-io.github.io/k8s/helm/charts/
   helm install nats nats/nats --set=nats.jetstream.enabled=true
   helm install nack nats/nack --set jetstream.nats.url=nats://nats:4222
3. Port forward to 4222


# Working with .env files

### Add godotnev
go get github.com/joho/godotenv
### Read .env file
godotenv.Load()
### Read multiple .env files
godotenv.Load(".env.local", ".env")

* .env.local will override .env

# Build image
* nerdctl build --namespace k8s.io  -t [image]:[version]

# Deploy (without Helm)
* kubectl apply -f [Deployment yaml file]