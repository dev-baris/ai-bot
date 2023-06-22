# AI-Bot

## Build
Install poppler
`sudo apt-get install -y poppler-utils on Ubuntu, or brew install poppler on Mac`

```bash
yarn install
```

```bash
yarn start

Webserver: http://localhost:8100
```

```bash
yarn dev

Run: webpack
```

```bash
sudo yarn clean
```

```bash
sudo npm list -g --depth 0
sudo npm outdated -g --depth=0
sudo npm update -g
sudo yarn cache clean -f --unsafe-perm=true --allow-root
```

## Update outdated Go
```
go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null
```

## Generate private key (.key)

```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048
    
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

##### Generation of self-signed(x509) public key (PEM-encodings `.pem`|`.crt`) based on the private (`.key`)

```sh
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

```sh
curl -sL https://localhost:443 | xxd
```

```go
http.HandleFunc("/hello", HelloServer)
err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
if err != nil {
    log.Fatal("ListenAndServe: ", err)
}
```

## Supported Filetypes
```bash
.pdf, .txt, .rtf, .docx, .epub, and plaintext.
```

## Uploading larger files
```shell
fileupload.go

const MAX_FILE_SIZE int64 = 25 << 20         // 3 MB
const MAX_TOTAL_UPLOAD_SIZE int64 = 50 << 20 // 3 MB
```

## Setup
```bash
https://platform.openai.com/docs/api-reference/authentication
https://app.pinecone.io/organizations/
```