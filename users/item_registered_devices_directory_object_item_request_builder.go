package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRegisteredDevicesDirectoryObjectItemRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal instantiates a new ItemRegisteredDevicesDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredDevicesDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/registeredDevices/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRegisteredDevicesDirectoryObjectItemRequestBuilder instantiates a new ItemRegisteredDevicesDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredDevicesDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// GraphDevice casts the previous resource to device.
// returns a *ItemRegisteredDevicesItemGraphDeviceRequestBuilder when successful
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) GraphDevice()(*ItemRegisteredDevicesItemGraphDeviceRequestBuilder) {
    return NewItemRegisteredDevicesItemGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegisteredDevicesItemGraphEndpointRequestBuilder when successful
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemRegisteredDevicesItemGraphEndpointRequestBuilder) {
    return NewItemRegisteredDevicesItemGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder when successful
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    return NewItemRegisteredDevicesDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
