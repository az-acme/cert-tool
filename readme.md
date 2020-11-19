# Az Acme

az-acme is an open soure tool to manage TLS certificates in Azure from ACME compatible certificate authorities. While there are other tools that support Acme, az-acme is aligned with the following design goals to achive an opinionatd experience that is highly aligned with Azure:

- Standalone binary for execution, with a small compiled size
- Externalise configuration to files to support flexible execution scenarios
- Utilise Azure KeyVault to persist certificates and sensitive state / configuration.
- Core integration for Azure DNS for ACME challenges

