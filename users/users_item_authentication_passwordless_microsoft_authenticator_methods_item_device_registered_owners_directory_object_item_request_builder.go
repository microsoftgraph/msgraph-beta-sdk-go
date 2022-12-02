package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\passwordlessMicrosoftAuthenticatorMethods\{passwordlessMicrosoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
