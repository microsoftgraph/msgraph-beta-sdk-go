package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManageddevicesManagedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesManagedDevicesRequestBuilderGetQueryParameters the list of managed devices.
type ManageddevicesManagedDevicesRequestBuilderGetQueryParameters struct {
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
// ManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesManagedDevicesRequestBuilderGetQueryParameters
}
// ManageddevicesManagedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppDiagnosticsWithUpn provides operations to call the appDiagnostics method.
// returns a *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) AppDiagnosticsWithUpn(upn *string)(*ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    return NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, upn)
}
// BulkReprovisionCloudPc provides operations to call the bulkReprovisionCloudPc method.
// returns a *ManageddevicesBulkreprovisioncloudpcBulkReprovisionCloudPcRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*ManageddevicesBulkreprovisioncloudpcBulkReprovisionCloudPcRequestBuilder) {
    return NewManageddevicesBulkreprovisioncloudpcBulkReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkRestoreCloudPc provides operations to call the bulkRestoreCloudPc method.
// returns a *ManageddevicesBulkrestorecloudpcBulkRestoreCloudPcRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*ManageddevicesBulkrestorecloudpcBulkRestoreCloudPcRequestBuilder) {
    return NewManageddevicesBulkrestorecloudpcBulkRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkSetCloudPcReviewStatus provides operations to call the bulkSetCloudPcReviewStatus method.
// returns a *ManageddevicesBulksetcloudpcreviewstatusBulkSetCloudPcReviewStatusRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*ManageddevicesBulksetcloudpcreviewstatusBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewManageddevicesBulksetcloudpcreviewstatusBulkSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByManagedDeviceId provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) ByManagedDeviceId(managedDeviceId string)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceId != "" {
        urlTplParams["managedDevice%2Did"] = managedDeviceId
    }
    return NewManageddevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManageddevicesManagedDevicesRequestBuilderInternal instantiates a new ManageddevicesManagedDevicesRequestBuilder and sets the default values.
func NewManageddevicesManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDevicesRequestBuilder) {
    m := &ManageddevicesManagedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManageddevicesManagedDevicesRequestBuilder instantiates a new ManageddevicesManagedDevicesRequestBuilder and sets the default values.
func NewManageddevicesManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManageddevicesCountRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) Count()(*ManageddevicesCountRequestBuilder) {
    return NewManageddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DownloadAppDiagnostics provides operations to call the downloadAppDiagnostics method.
// returns a *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) DownloadAppDiagnostics()(*ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) {
    return NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExecuteAction provides operations to call the executeAction method.
// returns a *ManageddevicesExecuteactionExecuteActionRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) ExecuteAction()(*ManageddevicesExecuteactionExecuteActionRequestBuilder) {
    return NewManageddevicesExecuteactionExecuteActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of managed devices.
// returns a ManagedDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable), nil
}
// MoveDevicesToOU provides operations to call the moveDevicesToOU method.
// returns a *ManageddevicesMovedevicestoouMoveDevicesToOURequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) MoveDevicesToOU()(*ManageddevicesMovedevicestoouMoveDevicesToOURequestBuilder) {
    return NewManageddevicesMovedevicestoouMoveDevicesToOURequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to managedDevices for deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// ToGetRequestInformation the list of managed devices.
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesManagedDevicesRequestBuilder when successful
func (m *ManageddevicesManagedDevicesRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesManagedDevicesRequestBuilder) {
    return NewManageddevicesManagedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
