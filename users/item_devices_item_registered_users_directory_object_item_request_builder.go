package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredUsers\{directoryObject-id}
type ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    m := &ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder instantiates a new ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemDevicesItemRegisteredUsersItemGraphEndpointRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemDevicesItemRegisteredUsersItemGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersItemGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemDevicesItemRegisteredUsersItemGraphServicePrincipalRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemDevicesItemRegisteredUsersItemGraphServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemDevicesItemRegisteredUsersItemGraphUserRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemDevicesItemRegisteredUsersItemGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of user entities.
// returns a *ItemDevicesItemRegisteredUsersItemRefRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) Ref()(*ItemDevicesItemRegisteredUsersItemRefRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
