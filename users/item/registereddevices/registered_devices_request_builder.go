package registereddevices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4797ee7d9bee8ef70e9c258e15fc18b01a77425ef5d538712f652324c35b292c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/registereddevices/device"
    i8ec2cdfea557d8381fec87953bd2e19cdb86356ee71b9a468f71e6e802d67d3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/registereddevices/endpoint"
    icc2c9ed085aed4af5333f00888f47851a6efc61442b13587eed6798a7ffffe79 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/registereddevices/count"
)

// RegisteredDevicesRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type RegisteredDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RegisteredDevicesRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand.
type RegisteredDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RegisteredDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RegisteredDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RegisteredDevicesRequestBuilderGetQueryParameters
}
// NewRegisteredDevicesRequestBuilderInternal instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    m := &RegisteredDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/registeredDevices{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredDevicesRequestBuilder instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RegisteredDevicesRequestBuilder) Count()(*icc2c9ed085aed4af5333f00888f47851a6efc61442b13587eed6798a7ffffe79.CountRequestBuilder) {
    return icc2c9ed085aed4af5333f00888f47851a6efc61442b13587eed6798a7ffffe79.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *RegisteredDevicesRequestBuilder) Device()(*i4797ee7d9bee8ef70e9c258e15fc18b01a77425ef5d538712f652324c35b292c.DeviceRequestBuilder) {
    return i4797ee7d9bee8ef70e9c258e15fc18b01a77425ef5d538712f652324c35b292c.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint the endpoint property
func (m *RegisteredDevicesRequestBuilder) Endpoint()(*i8ec2cdfea557d8381fec87953bd2e19cdb86356ee71b9a468f71e6e802d67d3b.EndpointRequestBuilder) {
    return i8ec2cdfea557d8381fec87953bd2e19cdb86356ee71b9a468f71e6e802d67d3b.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
