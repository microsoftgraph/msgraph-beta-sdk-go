package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredOwners\{directoryObject-id}
type UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*UsersItemDevicesItemRegisteredOwnersItemEndpointRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*UsersItemDevicesItemRegisteredOwnersItemRefRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*UsersItemDevicesItemRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *UsersItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*UsersItemDevicesItemRegisteredOwnersItemUserRequestBuilder) {
    return NewUsersItemDevicesItemRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
