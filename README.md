# Railway volume dump

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/EBwdAh?referralCode=C0tigH)

Easily download your Railway volume data as a ZIP file.

## Usage

The following `curl` will download the ZIP file on your machine, replace `<GENERATED_ENDPOINT>` and `<GENERATED_PASSWORD>` with the generated when deploying the template.

```bash
curl -OJ <GENERATED_ENDPOINT> -H "password: <GENERATED_PASSWORD>"
```

## Notes

- The template will automatically pick up your volume path and name.
- Downloading will occur egress fees.
- It might takes some time to compress the volume data, you can check the service logs for progress.
