package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\devices\{device-id}\registeredOwners\{directoryObject-id}
type MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*MeDevicesItemRegisteredOwnersItemEndpointRequestBuilder) {
    return NewMeDevicesItemRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*MeDevicesItemRegisteredOwnersItemRefRequestBuilder) {
    return NewMeDevicesItemRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*MeDevicesItemRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewMeDevicesItemRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*MeDevicesItemRegisteredOwnersItemUserRequestBuilder) {
    return NewMeDevicesItemRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
