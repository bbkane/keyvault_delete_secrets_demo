# Bug

There is a change of behaviour for GetSecretsComplete when soft-delete is not enabled vs when it is enabled.

## No Soft Delete

- Create a certificate (call it `certName`)
- GetSecretsComplete - shows secret for `certName`
- Delete certificate
- GetSecretsComplete - does *not* show secret for `certName`

## With Soft Delete

- Create a certificate (call it `certName`)
- GetSecretsComplete - shows secret for `certName`
- Delete certificate
- GetSecretsComplete - shows secret for `certName`

This in spite of the fact that there is a GetDeletedSecretsComplete method which

# Repro

Unfortunately, I can't create a keyvault with soft delete disabled now, so I can't demo this

```
+ az keyvault create --location westus2 --name no-soft-delete-117 --resource-group soft-delete-demo --enable-soft-delete false --sku standard
Argument 'enable_soft_delete' has been deprecated and will be removed in a future release.
"--enable-soft-delete false" has been deprecated, you cannot disable Soft Delete via CLI. The value will be changed to true.
```

# Behavior of GetSecretsComplete vs GetDeletedSecretsComplete

What is the difference between GetSecretsComplete and GetDeletedSecretsComplete ?

It looks like GetDeletedSecretsComplete always returns an empty list?
