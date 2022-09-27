package manageddevices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04cfb1258ecd2605a74af4f6db4fb79ce07d7a1dd318ff32459c8640ebba651c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/appdiagnosticswithupn"
    i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/movedevicestoou"
    i110b95d5c92c30c95fc9cb7f533d46c23d6b53abfed1ab0c57efbcd3adc67edb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/bulksetcloudpcreviewstatus"
    i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/bulkrestorecloudpc"
    i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/bulkreprovisioncloudpc"
    iae6e25a8eb88b77410f0b3fa5c5df3b4c37bcc35afc1330f4f6200d047c6f8e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/downloadappdiagnostics"
    iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/count"
    ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/executeaction"
)

// ManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesRequestBuilderGetQueryParameters the list of managed devices.
type ManagedDevicesRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesRequestBuilderGetQueryParameters
}
// ManagedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppDiagnosticsWithUpn provides operations to call the appDiagnostics method.
func (m *ManagedDevicesRequestBuilder) AppDiagnosticsWithUpn(upn *string)(*i04cfb1258ecd2605a74af4f6db4fb79ce07d7a1dd318ff32459c8640ebba651c.AppDiagnosticsWithUpnRequestBuilder) {
    return i04cfb1258ecd2605a74af4f6db4fb79ce07d7a1dd318ff32459c8640ebba651c.NewAppDiagnosticsWithUpnRequestBuilderInternal(m.pathParameters, m.requestAdapter, upn);
}
// BulkReprovisionCloudPc the bulkReprovisionCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a.BulkReprovisionCloudPcRequestBuilder) {
    return i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkRestoreCloudPc the bulkRestoreCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e.BulkRestoreCloudPcRequestBuilder) {
    return i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkSetCloudPcReviewStatus the bulkSetCloudPcReviewStatus property
func (m *ManagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*i110b95d5c92c30c95fc9cb7f533d46c23d6b53abfed1ab0c57efbcd3adc67edb.BulkSetCloudPcReviewStatusRequestBuilder) {
    return i110b95d5c92c30c95fc9cb7f533d46c23d6b53abfed1ab0c57efbcd3adc67edb.NewBulkSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDevicesRequestBuilder instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ManagedDevicesRequestBuilder) Count()(*iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62.CountRequestBuilder) {
    return iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of managed devices.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of managed devices.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to managedDevices for deviceManagement
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to managedDevices for deviceManagement
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DownloadAppDiagnostics the downloadAppDiagnostics property
func (m *ManagedDevicesRequestBuilder) DownloadAppDiagnostics()(*iae6e25a8eb88b77410f0b3fa5c5df3b4c37bcc35afc1330f4f6200d047c6f8e1.DownloadAppDiagnosticsRequestBuilder) {
    return iae6e25a8eb88b77410f0b3fa5c5df3b4c37bcc35afc1330f4f6200d047c6f8e1.NewDownloadAppDiagnosticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecuteAction the executeAction property
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa.ExecuteActionRequestBuilder) {
    return ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of managed devices.
func (m *ManagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// MoveDevicesToOU the moveDevicesToOU property
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599.MoveDevicesToOURequestBuilder) {
    return i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to managedDevices for deviceManagement
func (m *ManagedDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
