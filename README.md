# Railway volume dump

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/EBwdAh?referralCode=C0tigH)

Easily download your Railway volume data as a ZIP file.

## Prerequisites

After deploying this template, copy the generated password and endpoint.

> [!NOTE]
> The template will automatically pick up your volume path and name.

## Usage

The following `curl` will download the ZIP file on your machine.

```bash
curl -sLD - -JO <GENERATED_ENDPOINT> -H "password: <GENERATED_PASSWORD>"
```
