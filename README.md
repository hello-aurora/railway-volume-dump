# Railway volume dump

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/igq6yW?referralCode=C0tigH)

Easily download your Railway volume data as a ZIP file.

## Prerequisites

Make sure to mount the volume you want to download at `/data`. After deploying this template, copy the generated password and endpoint.

## Usage

The following `curl` will download the ZIP file on your machine.

```bash
curl <GENERATED_ENDPOINT>/download --header "password: <GENERATED_PASSWORD>" --output ./data.zip
```
