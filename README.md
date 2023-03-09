# Microsoft Graph SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/microsoftgraph/msgraph-beta-sdk-go/)](https://pkg.go.dev/github.com/microsoftgraph/msgraph-beta-sdk-go/)

Get started with the Microsoft Graph SDK for Go by integrating the [Microsoft Graph API](https://docs.microsoft.com/graph/overview) into your Go application!

> **Note:** this SDK allows you to build applications using the [beta](https://docs.microsoft.com/en-us/graph/use-the-api#version) of Microsoft Graph. If you want to use the production supported Microsoft Graph APIs under v1.0, use our [v1.0 SDK](https://github.com/microsoftgraph/msgraph-sdk-go) instead.
>
> **Note:** The Microsoft Graph Go SDK is currently in Release Candidate (RC) version starting from version 0.55.0. The SDK is still undergoing testing but minimum breaking changes should be expected. Checkout the [known limitations](https://github.com/microsoftgraph/msgraph-sdk-go-core/issues/1).

## 1. Installation

```Shell
go get github.com/microsoftgraph/msgraph-beta-sdk-go
go get github.com/microsoft/kiota-authentication-azure-go
```

## 2. Getting started

### 2.1 Register your application

Register your application by following the steps at [Register your app with the Microsoft Identity Platform](https://docs.microsoft.com/graph/auth-register-app-v2).

### 2.2 Create an AuthenticationProvider object

An instance of the **GraphRequestAdapter** class handles building client. To create a new instance of this class, you need to provide an instance of **AuthenticationProvider**, which can authenticate requests to Microsoft Graph.

For an example of how to get an authentication provider, see [choose a Microsoft Graph authentication provider](https://docs.microsoft.com/graph/sdks/choose-authentication-providers?tabs=Go).

> Note: we are working to add the getting started information for Go to our public documentation, in the meantime the following sample should help you getting started.

```Golang
import (
    azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "context"
)

cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
    TenantID: "<the tenant id from your app registration>",
    ClientID: "<the client id from your app registration>",
    UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
        fmt.Println(message.Message)
        return nil
    },
})

if err != nil {
    fmt.Printf("Error creating credentials: %v\n", err)
}

```

### 2.3 Get a Graph Service Client Adapter object

You must get a **GraphRequestAdapter** object to make requests against the service.

```Golang
import msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

client , err  := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})
if err != nil {
    fmt.Printf("Error creating client: %v\n", err)
    return
}
```

## 3. Make requests against the service

After you have a GraphServiceClient that is authenticated, you can begin making calls against the service. The requests against the service look like our [REST API](https://docs.microsoft.com/graph/api/overview?view=graph-rest-beta).

### 3.1 Get the user's drive

To retrieve the user's drive:

```Golang
import (
    "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

result, err := client.Me().Drive().Get()
if err != nil {
    fmt.Printf("Error getting the drive: ")
    printOdataError(err)
}


fmt.Printf("Found Drive : %v\n", result.GetId())

// omitted for brevity
func printOdataError(err error) {
	switch err.(type) {
	case *odataerrors.ODataError:
		typed := err.(*odataerrors.ODataError)
		fmt.Printf("error:", typed.Error())
		if terr := typed.GetError(); terr != nil {
			fmt.Printf("code: %s", *terr.GetCode())
			fmt.Printf("msg: %s", *terr.GetMessage())
		}
	default:
		fmt.Printf("%T > error: %#v", err, err)
	}
}
```

## 4. Getting results that span across multiple pages

Items in a collection response can span across multiple pages. To get the complete set of items in the collection, your application must make additional calls to get the subsequent pages until no more next link is provided in the response.

### 4.1 Get all the users in an environment

To retrieve the users:

```Golang
import (
    msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
    "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
    "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

result, err := client.Users().Get(context.Background(), nil)
if err != nil {
    fmt.Printf("Error getting users: ")
    printOdataError(err error)
    return err
}

// Use PageIterator to iterate through all users
pageIterator, err := msgraphcore.NewPageIterator(result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)

err = pageIterator.Iterate(func(pageItem interface{}) bool {
    user := pageItem.(models.Userable)
    fmt.Printf("%s\n", *user.GetDisplayName())
    // Return true to continue the iteration
    return true
})

// omitted for brevity
func printOdataError(err error) {
	switch err.(type) {
	case *odataerrors.ODataError:
		typed := err.(*odataerrors.ODataError)
		fmt.Printf("error:", typed.Error())
		if terr := typed.GetError(); terr != nil {
			fmt.Printf("code: %s", *terr.GetCode())
			fmt.Printf("msg: %s", *terr.GetMessage())
		}
	default:
		fmt.Printf("%T > error: %#v", err, err)
	}
}
```

## 5. Documentation

For more detailed documentation, see:

* [Overview](https://docs.microsoft.com/graph/overview)
* [Collections](https://docs.microsoft.com/graph/sdks/paging)
* [Making requests](https://docs.microsoft.com/graph/sdks/create-requests)
* [Known issues](https://github.com/MicrosoftGraph/msgraph-beta-sdk-go/issues)
* [Contributions](https://github.com/microsoftgraph/msgraph-beta-sdk-go/blob/main/CONTRIBUTING.md)

## 6. Issues

For known issues, see [issues](https://github.com/MicrosoftGraph/msgraph-beta-sdk-go/issues).

## 7. Contributions

The Microsoft Graph SDK is open for contribution. To contribute to this project, see [Contributing](https://github.com/microsoftgraph/msgraph-beta-sdk-go/blob/main/CONTRIBUTING.md).

## 8. License

Copyright (c) Microsoft Corporation. All Rights Reserved. Licensed under the [MIT license](LICENSE).

## 9. Third-party notices

[Third-party notices](THIRD%20PARTY%20NOTICES)
