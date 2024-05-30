package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
type DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters summary of all certificates for all devices.
type DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters struct {
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
// DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters
}
// DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedAllDeviceCertificateStateId provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationsallmanageddevicecertificatestatesManagedAllDeviceCertificateStateItemRequestBuilder when successful
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) ByManagedAllDeviceCertificateStateId(managedAllDeviceCertificateStateId string)(*DeviceconfigurationsallmanageddevicecertificatestatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedAllDeviceCertificateStateId != "" {
        urlTplParams["managedAllDeviceCertificateState%2Did"] = managedAllDeviceCertificateStateId
    }
    return NewDeviceconfigurationsallmanageddevicecertificatestatesManagedAllDeviceCertificateStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal instantiates a new DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder and sets the default values.
func NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    m := &DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder instantiates a new DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder and sets the default values.
func NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceconfigurationsallmanageddevicecertificatestatesCountRequestBuilder when successful
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Count()(*DeviceconfigurationsallmanageddevicecertificatestatesCountRequestBuilder) {
    return NewDeviceconfigurationsallmanageddevicecertificatestatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get summary of all certificates for all devices.
// returns a ManagedAllDeviceCertificateStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAllDeviceCertificateStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateCollectionResponseable), nil
}
// Post create new navigation property to deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
// returns a ManagedAllDeviceCertificateStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, requestConfiguration *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAllDeviceCertificateStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable), nil
}
// ToGetRequestInformation summary of all certificates for all devices.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, requestConfiguration *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder when successful
func (m *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
