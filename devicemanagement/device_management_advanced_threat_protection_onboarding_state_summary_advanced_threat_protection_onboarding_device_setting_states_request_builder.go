package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters not yet documented
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters struct {
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
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    m := &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Count()(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCountRequestBuilder) {
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation not yet documented
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get not yet documented
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateCollectionResponseable), nil
}
// Post create new navigation property to advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable), nil
}
