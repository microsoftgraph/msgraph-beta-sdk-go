package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementComanagedDevicesRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type DeviceManagementComanagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementComanagedDevicesRequestBuilderGetQueryParameters the list of co-managed devices report
type DeviceManagementComanagedDevicesRequestBuilderGetQueryParameters struct {
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
// DeviceManagementComanagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementComanagedDevicesRequestBuilderGetQueryParameters
}
// DeviceManagementComanagedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppDiagnosticsWithUpn provides operations to call the appDiagnostics method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) AppDiagnosticsWithUpn(upn *string)(*DeviceManagementComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    return NewDeviceManagementComanagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(m.pathParameters, m.requestAdapter, upn);
}
// BulkReprovisionCloudPc provides operations to call the bulkReprovisionCloudPc method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*DeviceManagementComanagedDevicesBulkReprovisionCloudPcRequestBuilder) {
    return NewDeviceManagementComanagedDevicesBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkRestoreCloudPc provides operations to call the bulkRestoreCloudPc method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) BulkRestoreCloudPc()(*DeviceManagementComanagedDevicesBulkRestoreCloudPcRequestBuilder) {
    return NewDeviceManagementComanagedDevicesBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkSetCloudPcReviewStatus provides operations to call the bulkSetCloudPcReviewStatus method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*DeviceManagementComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilder) {
    return NewDeviceManagementComanagedDevicesBulkSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementComanagedDevicesRequestBuilderInternal instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesRequestBuilder) {
    m := &DeviceManagementComanagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementComanagedDevicesRequestBuilder instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementComanagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DeviceManagementComanagedDevicesRequestBuilder) Count()(*DeviceManagementComanagedDevicesCountRequestBuilder) {
    return NewDeviceManagementComanagedDevicesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of co-managed devices report
func (m *DeviceManagementComanagedDevicesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to comanagedDevices for deviceManagement
func (m *DeviceManagementComanagedDevicesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementComanagedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DownloadAppDiagnostics provides operations to call the downloadAppDiagnostics method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) DownloadAppDiagnostics()(*DeviceManagementComanagedDevicesDownloadAppDiagnosticsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesDownloadAppDiagnosticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecuteAction provides operations to call the executeAction method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) ExecuteAction()(*DeviceManagementComanagedDevicesExecuteActionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
func (m *DeviceManagementComanagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable), nil
}
// MoveDevicesToOU provides operations to call the moveDevicesToOU method.
func (m *DeviceManagementComanagedDevicesRequestBuilder) MoveDevicesToOU()(*DeviceManagementComanagedDevicesMoveDevicesToOURequestBuilder) {
    return NewDeviceManagementComanagedDevicesMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to comanagedDevices for deviceManagement
func (m *DeviceManagementComanagedDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementComanagedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
