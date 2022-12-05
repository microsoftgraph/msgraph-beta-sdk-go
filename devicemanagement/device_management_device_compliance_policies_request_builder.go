package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceCompliancePoliciesRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceCompliancePoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceCompliancePoliciesRequestBuilderGetQueryParameters the device compliance policies.
type DeviceManagementDeviceCompliancePoliciesRequestBuilderGetQueryParameters struct {
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
// DeviceManagementDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceCompliancePoliciesRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceCompliancePoliciesRequestBuilderInternal instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesRequestBuilder) {
    m := &DeviceManagementDeviceCompliancePoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceCompliancePoliciesRequestBuilder instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceCompliancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) Count()(*DeviceManagementDeviceCompliancePoliciesCountRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deviceCompliancePolicies for deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the device compliance policies.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable), nil
}
// GetDevicesScheduledToRetire provides operations to call the getDevicesScheduledToRetire method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) GetDevicesScheduledToRetire()(*DeviceManagementDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNoncompliantDevicesToRetire provides operations to call the getNoncompliantDevicesToRetire method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) GetNoncompliantDevicesToRetire()(*DeviceManagementDeviceCompliancePoliciesGetNoncompliantDevicesToRetireRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesGetNoncompliantDevicesToRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) HasPayloadLinks()(*DeviceManagementDeviceCompliancePoliciesHasPayloadLinksRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deviceCompliancePolicies for deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// RefreshDeviceComplianceReportSummarization provides operations to call the refreshDeviceComplianceReportSummarization method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) RefreshDeviceComplianceReportSummarization()(*DeviceManagementDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetScheduledRetireState provides operations to call the setScheduledRetireState method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) SetScheduledRetireState()(*DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStateRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateComplianceScript provides operations to call the validateComplianceScript method.
func (m *DeviceManagementDeviceCompliancePoliciesRequestBuilder) ValidateComplianceScript()(*DeviceManagementDeviceCompliancePoliciesValidateComplianceScriptRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesValidateComplianceScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
