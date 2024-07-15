package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetQueryParameters summary of all certificates for all devices.
type DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetQueryParameters
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderInternal instantiates a new DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder and sets the default values.
func NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    m := &DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/{managedAllDeviceCertificateState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder instantiates a new DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder and sets the default values.
func NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get summary of all certificates for all devices.
// returns a ManagedAllDeviceCertificateStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property deviceConfigurationsAllManagedDeviceCertificateStates in deviceManagement
// returns a ManagedAllDeviceCertificateStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summary of all certificates for all devices.
// returns a *RequestInformation when successful
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceConfigurationsAllManagedDeviceCertificateStates in deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, requestConfiguration *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder when successful
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) WithUrl(rawUrl string)(*DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
