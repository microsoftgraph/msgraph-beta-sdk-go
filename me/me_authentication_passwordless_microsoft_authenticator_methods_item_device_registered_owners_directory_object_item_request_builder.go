package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\passwordlessMicrosoftAuthenticatorMethods\{passwordlessMicrosoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
