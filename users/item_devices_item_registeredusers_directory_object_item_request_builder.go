package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredUsers\{directoryObject-id}
type ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) {
    m := &ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder instantiates a new ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemDevicesItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemDevicesItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemDevicesItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemDevicesItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemDevicesItemRegisteredusersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemDevicesItemRegisteredusersItemGraphuserGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredusersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of user entities.
// returns a *ItemDevicesItemRegisteredusersItemRefRequestBuilder when successful
func (m *ItemDevicesItemRegisteredusersDirectoryObjectItemRequestBuilder) Ref()(*ItemDevicesItemRegisteredusersItemRefRequestBuilder) {
    return NewItemDevicesItemRegisteredusersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
