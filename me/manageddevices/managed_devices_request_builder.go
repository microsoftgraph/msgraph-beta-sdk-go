package manageddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/bulkreprovisioncloudpc"
    i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/executeaction"
    i09dab353f35c6e8c23d8c077e3ae5ed0ae54c83a7bd6e2663817b587f499b3af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/bulksetcloudpcreviewstatus"
    i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/bulkrestorecloudpc"
    i683808f21d3e066affe118ad2cfd89e4a74edfa2ad53ba3de252833220b09f88 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/appdiagnosticswithupn"
    i7737024fb5d89c929b3b03de5dd1234e52923060aab0a2ee3f345636d382081c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/downloadappdiagnostics"
    i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/movedevicestoou"
    i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/count"
)

// ManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesRequestBuilderGetQueryParameters the managed devices associated with the user.
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
func (m *ManagedDevicesRequestBuilder) AppDiagnosticsWithUpn(upn *string)(*i683808f21d3e066affe118ad2cfd89e4a74edfa2ad53ba3de252833220b09f88.AppDiagnosticsWithUpnRequestBuilder) {
    return i683808f21d3e066affe118ad2cfd89e4a74edfa2ad53ba3de252833220b09f88.NewAppDiagnosticsWithUpnRequestBuilderInternal(m.pathParameters, m.requestAdapter, upn);
}
// BulkReprovisionCloudPc the bulkReprovisionCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d.BulkReprovisionCloudPcRequestBuilder) {
    return i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkRestoreCloudPc the bulkRestoreCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633.BulkRestoreCloudPcRequestBuilder) {
    return i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkSetCloudPcReviewStatus the bulkSetCloudPcReviewStatus property
func (m *ManagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*i09dab353f35c6e8c23d8c077e3ae5ed0ae54c83a7bd6e2663817b587f499b3af.BulkSetCloudPcReviewStatusRequestBuilder) {
    return i09dab353f35c6e8c23d8c077e3ae5ed0ae54c83a7bd6e2663817b587f499b3af.NewBulkSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
func (m *ManagedDevicesRequestBuilder) Count()(*i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051.CountRequestBuilder) {
    return i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the managed devices associated with the user.
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
// CreatePostRequestInformation create new navigation property to managedDevices for me
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to managedDevices for me
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
func (m *ManagedDevicesRequestBuilder) DownloadAppDiagnostics()(*i7737024fb5d89c929b3b03de5dd1234e52923060aab0a2ee3f345636d382081c.DownloadAppDiagnosticsRequestBuilder) {
    return i7737024fb5d89c929b3b03de5dd1234e52923060aab0a2ee3f345636d382081c.NewDownloadAppDiagnosticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecuteAction the executeAction property
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18.ExecuteActionRequestBuilder) {
    return i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable), nil
}
// MoveDevicesToOU the moveDevicesToOU property
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784.MoveDevicesToOURequestBuilder) {
    return i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to managedDevices for me
func (m *ManagedDevicesRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to managedDevices for me
func (m *ManagedDevicesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
