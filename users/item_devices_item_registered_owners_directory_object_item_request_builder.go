package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredOwners\{directoryObject-id}
type ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MicrosoftGraphEndpoint casts the previous resource to endpoint.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphEndpoint()(*ItemDevicesItemRegisteredOwnersItemMicrosoftGraphEndpointEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemMicrosoftGraphEndpointEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemDevicesItemRegisteredOwnersItemMicrosoftGraphServicePrincipalServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemMicrosoftGraphServicePrincipalServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*ItemDevicesItemRegisteredOwnersItemMicrosoftGraphUserUserRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemMicrosoftGraphUserUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of user entities.
func (m *ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemDevicesItemRegisteredOwnersItemRefRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
