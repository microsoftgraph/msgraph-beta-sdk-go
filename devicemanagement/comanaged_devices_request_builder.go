package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type ComanagedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesRequestBuilderGetQueryParameters the list of co-managed devices report
type ComanagedDevicesRequestBuilderGetQueryParameters struct {
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
// ComanagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesRequestBuilderGetQueryParameters
}
// ComanagedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppDiagnosticsWithUpn provides operations to call the appDiagnostics method.
// returns a *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) AppDiagnosticsWithUpn(upn *string)(*ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    return NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, upn)
}
// BulkReprovisionCloudPc provides operations to call the bulkReprovisionCloudPc method.
// returns a *ComanagedDevicesBulkReprovisionCloudPcRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*ComanagedDevicesBulkReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesBulkReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkRestoreCloudPc provides operations to call the bulkRestoreCloudPc method.
// returns a *ComanagedDevicesBulkRestoreCloudPcRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) BulkRestoreCloudPc()(*ComanagedDevicesBulkRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesBulkRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkSetCloudPcReviewStatus provides operations to call the bulkSetCloudPcReviewStatus method.
// returns a *ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*ComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByManagedDeviceId provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ComanagedDevicesManagedDeviceItemRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) ByManagedDeviceId(managedDeviceId string)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceId != "" {
        urlTplParams["managedDevice%2Did"] = managedDeviceId
    }
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewComanagedDevicesRequestBuilderInternal instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    m := &ComanagedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewComanagedDevicesRequestBuilder instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ComanagedDevicesCountRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) Count()(*ComanagedDevicesCountRequestBuilder) {
    return NewComanagedDevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DownloadAppDiagnostics provides operations to call the downloadAppDiagnostics method.
// returns a *ComanagedDevicesDownloadAppDiagnosticsRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) DownloadAppDiagnostics()(*ComanagedDevicesDownloadAppDiagnosticsRequestBuilder) {
    return NewComanagedDevicesDownloadAppDiagnosticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DownloadPowerliftAppDiagnostic provides operations to call the downloadPowerliftAppDiagnostic method.
// returns a *ComanagedDevicesDownloadPowerliftAppDiagnosticRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) DownloadPowerliftAppDiagnostic()(*ComanagedDevicesDownloadPowerliftAppDiagnosticRequestBuilder) {
    return NewComanagedDevicesDownloadPowerliftAppDiagnosticRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExecuteAction provides operations to call the executeAction method.
// returns a *ComanagedDevicesExecuteActionRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) ExecuteAction()(*ComanagedDevicesExecuteActionRequestBuilder) {
    return NewComanagedDevicesExecuteActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of co-managed devices report
// returns a ManagedDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
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
// returns a *ComanagedDevicesMoveDevicesToOURequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) MoveDevicesToOU()(*ComanagedDevicesMoveDevicesToOURequestBuilder) {
    return NewComanagedDevicesMoveDevicesToOURequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to comanagedDevices for deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// RetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalName provides operations to call the retrievePowerliftAppDiagnosticsDetails method.
// returns a *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) RetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalName(userPrincipalName *string)(*ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    return NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, userPrincipalName)
}
// ToGetRequestInformation the list of co-managed devices report
// returns a *RequestInformation when successful
func (m *ComanagedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to comanagedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ComanagedDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesRequestBuilder when successful
func (m *ComanagedDevicesRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesRequestBuilder) {
    return NewComanagedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
