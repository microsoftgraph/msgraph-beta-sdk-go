package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
type DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters the collection property of MobileAppTroubleshootingEvent.
type DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters
}
// DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppLogCollectionRequests provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequests()(*DeviceManagementMobileAppTroubleshootingEventsItemAppLogCollectionRequestsRequestBuilder) {
    return NewDeviceManagementMobileAppTroubleshootingEventsItemAppLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppLogCollectionRequestsById provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequestsById(id string)(*DeviceManagementMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appLogCollectionRequest%2Did"] = id
    }
    return NewDeviceManagementMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal instantiates a new MobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    m := &DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder instantiates a new MobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileAppTroubleshootingEvents for deviceManagement
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileAppTroubleshootingEvents in deviceManagement
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property mobileAppTroubleshootingEvents for deviceManagement
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable), nil
}
// Patch update the navigation property mobileAppTroubleshootingEvents in deviceManagement
func (m *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, requestConfiguration *DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable), nil
}
