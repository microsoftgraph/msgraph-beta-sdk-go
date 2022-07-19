package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i3c9974dff4a549b5aa5b112949aa9254b8a07cd6ea35bad7b8fa734bbd1678a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/ref"
    i4310d9213d9c21a0c2b7685297393949ff281addd5c1758b12f24051c4cc04bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/endpoint"
    ibe2fab5c5f873feac42f8c946cb359ddd7ad4069e1b8a277be080be1960be9d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/serviceprincipal"
    ifbe6fdf38ec35168c9fa3e401e5e87906eadca693bddc1a428bf64f9ba2a7b55 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\microsoftAuthenticatorMethods\{microsoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint the endpoint property
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i4310d9213d9c21a0c2b7685297393949ff281addd5c1758b12f24051c4cc04bd.EndpointRequestBuilder) {
    return i4310d9213d9c21a0c2b7685297393949ff281addd5c1758b12f24051c4cc04bd.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i3c9974dff4a549b5aa5b112949aa9254b8a07cd6ea35bad7b8fa734bbd1678a8.RefRequestBuilder) {
    return i3c9974dff4a549b5aa5b112949aa9254b8a07cd6ea35bad7b8fa734bbd1678a8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*ibe2fab5c5f873feac42f8c946cb359ddd7ad4069e1b8a277be080be1960be9d2.ServicePrincipalRequestBuilder) {
    return ibe2fab5c5f873feac42f8c946cb359ddd7ad4069e1b8a277be080be1960be9d2.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*ifbe6fdf38ec35168c9fa3e401e5e87906eadca693bddc1a428bf64f9ba2a7b55.UserRequestBuilder) {
    return ifbe6fdf38ec35168c9fa3e401e5e87906eadca693bddc1a428bf64f9ba2a7b55.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
