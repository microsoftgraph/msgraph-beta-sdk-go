package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\devices\{device-id}\registeredOwners\{directoryObject-id}
type DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MicrosoftGraphEndpoint casts the previous resource to endpoint.
func (m *DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphEndpoint()(*DevicesItemRegisteredOwnersItemMicrosoftGraphEndpointRequestBuilder) {
    return NewDevicesItemRegisteredOwnersItemMicrosoftGraphEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*DevicesItemRegisteredOwnersItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewDevicesItemRegisteredOwnersItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*DevicesItemRegisteredOwnersItemMicrosoftGraphUserRequestBuilder) {
    return NewDevicesItemRegisteredOwnersItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of user entities.
func (m *DevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*DevicesItemRegisteredOwnersItemRefRequestBuilder) {
    return NewDevicesItemRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
