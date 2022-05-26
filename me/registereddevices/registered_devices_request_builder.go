package registereddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5213ea311e0df4f90d2c4dd380d4b7e5f4a37c4c373c882446a596a80320480e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/registereddevices/device"
    iae1984fb3fc141b0f414161e859e280d2f4ae41744529837c3922e8d5ba99818 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/registereddevices/count"
    ic0f72fe7856899dc82b6839158f8473ef255e8bafe9d49bd69089e026e3559f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/registereddevices/endpoint"
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
    m.urlTemplate = "{+baseurl}/me/registeredDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the count property
func (m *RegisteredDevicesRequestBuilder) Count()(*iae1984fb3fc141b0f414161e859e280d2f4ae41744529837c3922e8d5ba99818.CountRequestBuilder) {
    return iae1984fb3fc141b0f414161e859e280d2f4ae41744529837c3922e8d5ba99818.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *RegisteredDevicesRequestBuilder) Device()(*i5213ea311e0df4f90d2c4dd380d4b7e5f4a37c4c373c882446a596a80320480e.DeviceRequestBuilder) {
    return i5213ea311e0df4f90d2c4dd380d4b7e5f4a37c4c373c882446a596a80320480e.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint the endpoint property
func (m *RegisteredDevicesRequestBuilder) Endpoint()(*ic0f72fe7856899dc82b6839158f8473ef255e8bafe9d49bd69089e026e3559f4.EndpointRequestBuilder) {
    return ic0f72fe7856899dc82b6839158f8473ef255e8bafe9d49bd69089e026e3559f4.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
