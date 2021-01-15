package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/go-autorest/autorest/to"
	// "github.com/Azure/go-autorest/autorest"
)

//createCertificate creates a test certificate we can show then delete
func createCertificate(kvClient keyvault.BaseClient, vaultURL string, certName string) error {
	commonName := "bbkane.com"
	san := []string{"bbkane.com", "www.bbkane.com"}
	result, err := kvClient.CreateCertificate(
		context.Background(),
		vaultURL,
		certName,
		keyvault.CertificateCreateParameters{
			CertificateAttributes: nil, // godocs say it can be nil and the REST API example omits it
			CertificatePolicy: &keyvault.CertificatePolicy{
				// Not adding Response field, we'll use the default value
				ID: nil, // this is only useful in a response
				KeyProperties: &keyvault.KeyProperties{
					Exportable: to.BoolPtr(true),
					KeyType:    to.StringPtr("RSA"),
					KeySize:    to.Int32Ptr(2048),
					ReuseKey:   to.BoolPtr(false),
				},
				SecretProperties: &keyvault.SecretProperties{
					ContentType: to.StringPtr("application/x-pkcs12"),
				},
				X509CertificateProperties: &keyvault.X509CertificateProperties{
					Subject: to.StringPtr("CN=" + commonName),
					Ekus:    nil,
					SubjectAlternativeNames: &keyvault.SubjectAlternativeNames{
						DNSNames: &san,
					},
					KeyUsage:         nil,
					ValidityInMonths: to.Int32Ptr(6),
				},
				LifetimeActions: &[]keyvault.LifetimeAction{
					{
						Trigger: &keyvault.Trigger{
							LifetimePercentage: nil,
							DaysBeforeExpiry:   to.Int32Ptr(30),
						},
						Action: &keyvault.Action{
							ActionType: keyvault.AutoRenew,
						},
					},
				},
				IssuerParameters: &keyvault.IssuerParameters{
					Name: to.StringPtr("Self"),
					// NOTE: az keyvault show shows a "certificateTransparency"
					// field that's not in the Go API
					CertificateType: nil,
				},
				// Not in the REST API and it looks like a repeat of CertificateAttributes
				Attributes: nil,
			},
			Tags: map[string]*string{"key": to.StringPtr("value")},
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("createdID: %v, status: %v, statusDetails: %v\n", result.ID, result.Status, result.StatusDetails)
	return nil
}

func demo(kvClient keyvault.BaseClient, vaultURL string, certName string) error {
	fmt.Printf("Starting demo: %v\n", vaultURL)

	// create a certificate
	err := createCertificate(kvClient, vaultURL, certName)
	if err != nil {
		return err
	}

	// list secrets
	{
		secrets, err := kvClient.GetSecretsComplete(context.Background(), vaultURL, nil)
		for secrets.NotDone() {
			secret := secrets.Value()
			fmt.Println(secret.ID)
			err = secrets.NextWithContext(context.Background())
			if err != nil {
				return err
			}
		}
	}

	// delete certificate
	result, err := kvClient.DeleteCertificate(context.Background(), vaultURL, certName)
	if err != nil {
		return err
	}
	fmt.Printf("deletion status: %v\n", result.Status)

	// list secrets
	{
		secrets, err := kvClient.GetSecretsComplete(context.Background(), vaultURL, nil)
		for secrets.NotDone() {
			secret := secrets.Value()
			fmt.Println(secret.ID)
			err = secrets.NextWithContext(context.Background())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func run() error {

	kvClient := keyvault.New()
	var err error
	kvClient.Authorizer, err = kvauth.NewAuthorizerFromCLI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "This is probably an auth error. Use `az login` to fix")
		return err
	}

	if len(os.Args) != 3 {
		err = errors.New("Usage: keyvault_delete_secrets_demo keyvault-without-soft-delete keyvault-with-soft-delete")
		return err
	}

	vaultNoSoftDeleteURL := "https://" + os.Args[1] + ".vault.azure.net"
	vaultSoftDeleteURL := "https://" + os.Args[2] + ".vault.azure.net"
	certName := "soft-delete-demo"

	err = demo(kvClient, vaultNoSoftDeleteURL, certName)
	if err != nil {
		return err
	}

	err = demo(kvClient, vaultSoftDeleteURL, certName)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
