package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredOwners\{directoryObject-id}
type ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) {
    m := &ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder instantiates a new ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemDevicesItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemDevicesItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemDevicesItemRegisteredownersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemDevicesItemRegisteredownersItemGraphuserGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredownersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of user entities.
// returns a *ItemDevicesItemRegisteredownersItemRefRequestBuilder when successful
func (m *ItemDevicesItemRegisteredownersDirectoryObjectItemRequestBuilder) Ref()(*ItemDevicesItemRegisteredownersItemRefRequestBuilder) {
    return NewItemDevicesItemRegisteredownersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
