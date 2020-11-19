# Az Acme

az-acme is an open soure tool to manage TLS certificates in Azure from ACME compatible certificate authorities such as Let's Encrypt. While there are many other tools that support ACME, az-acme is based on the following key design goals to provide an experience highly aligned with Azure:

- Standalone binary for execution with a small compiled size
- Externalise configuration to support flexible execution scenarios and GitOps flows
- Utilise Azure Key Vault to store certificates and sensitive state / configuration.
- Integrate with Azure DNS for ACME challenges

