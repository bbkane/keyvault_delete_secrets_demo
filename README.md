# How can I exclude deleted certificates from GetSecretsComplete ?

When I call GetSecretsComplete, it returns secrets who's certificates have been deleted.
How do I exclude them? GetDeletedSecretsComplete seems to always return an empty list

It's possible I don't have the latest update of the library?

- GetSecretsComplete - https://godoc.org/github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault#BaseClient.GetSecretsComplete

- GetDeletedSecretsComplete - https://godoc.org/github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault#BaseClient.GetDeletedSecretsComplete

# demo run

```
$ go run . with-soft-delete-553
# Starting demo: https://with-soft-delete-553.vault.azure.net
# Creating certificate: soft-delete-demo-664
createdID: https://with-soft-delete-553.vault.azure.net/certificates/soft-delete-demo-664/pending, status: inProgress, statusDetails: Pending certificate created. Certificate request is in progress. This may take some time based on the issuer provider. Please check again later.
# Create Certificate Done
# Listing secrets
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-254
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-515
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-525
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-664
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-80
# Listing secrets done
# Listing Deleted Secrets
# Listing Deletes Secrets done
# Deleting certificate
deletion status: 200 OK
# Deleting certificate done
# This should not list the secret for the deleted certificates
# Listing secrets
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-254
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-515
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-525
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-664
https://with-soft-delete-553.vault.azure.net/secrets/soft-delete-demo-80
# Listing secrets done
# Listing Deleted Secrets
# Listing Deletes Secrets done
# Demo done
```
