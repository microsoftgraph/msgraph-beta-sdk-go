package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters the summary state of ATP onboarding state for this account.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters
}
// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedThreatProtectionOnboardingDeviceSettingStates provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates()(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdvancedThreatProtectionOnboardingDeviceSettingStatesById provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStatesById(id string)(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["advancedThreatProtectionOnboardingDeviceSettingState%2Did"] = id
    }
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    m := &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the summary state of ATP onboarding state for this account.
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the summary state of ATP onboarding state for this account.
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, requestConfiguration *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
