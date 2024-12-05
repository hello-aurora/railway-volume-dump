# Railway volume dump

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/EBwdAh?referralCode=C0tigH)

Easily download your Railway volume data as a ZIP file.

## Usage

1. Mount volume

You will first need to "disconnect the volume" from the original service you want to dump the volume from, and "mount" it to this service.

2. Run `curl` command

The following `curl` command will download a ZIP file on your machine, replace `<GENERATED_ENDPOINT>` and `<GENERATED_PASSWORD>` with the generated values from this service.

```bash
curl -OJ https://<GENERATED_ENDPOINT> -H "password: <GENERATED_PASSWORD>"
```

3. Re-mount volume

Once you have downloaded the ZIP file, you can re-mount the volume to the original service.

## Notes

- The template will automatically pick up your volume path and name.
- Downloading the ZIP file will occur egress fees.
- Check the service deploy logs for progress and additional information.
