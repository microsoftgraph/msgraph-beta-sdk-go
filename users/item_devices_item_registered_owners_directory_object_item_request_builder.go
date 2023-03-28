package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredOwners\{directoryObject-id}
type ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemDevicesItemRegisteredOwnersItemGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemDevicesItemRegisteredOwnersItemGraphServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemDevicesItemRegisteredOwnersItemGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of user entities.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemDevicesItemRegisteredOwnersItemRefRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
