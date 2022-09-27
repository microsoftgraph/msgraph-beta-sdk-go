package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assignments"
    i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assign"
    i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/settings"
    i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/setscheduledactions"
    iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule"
    i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/assignments/item"
    i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule/item"
    ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/settings/item"
)

// DeviceManagementCompliancePolicyItemRequestBuilder provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceManagementCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters list of all compliance policies
type DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Assign()(*i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.AssignRequestBuilder) {
    return i0742bff39d39edf9e7fdaa5802f2ceda48231b0df3ad173faa313f5c87241b5a.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Assignments()(*i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.AssignmentsRequestBuilder) {
    return i02451624685ebfdb81d6b1930c79509f0df4294168ed18436aaa09cc42c5c946.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.assignments.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.DeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return i1d9a96a6a2f04ecb11762741e92d99faf72d8bfcea65eb2f15cca45f47882da5.NewDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePolicyItemRequestBuilder) {
    m := &DeviceManagementCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementCompliancePolicyItemRequestBuilder instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property compliancePolicies for deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property compliancePolicies for deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of all compliance policies
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration list of all compliance policies
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property compliancePolicies in deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property compliancePolicies in deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property compliancePolicies for deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get list of all compliance policies
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable), nil
}
// Patch update the navigation property compliancePolicies in deviceManagement
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, requestConfiguration *DeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable), nil
}
// ScheduledActionsForRule the scheduledActionsForRule property
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.ScheduledActionsForRuleRequestBuilder) {
    return iffa40ac9a6da23297b60087932a288a73ea0e6ab171a35dec05137c5d4ed2f62.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.scheduledActionsForRule.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceScheduledActionForRule%2Did"] = id
    }
    return i6cd919cf4ce1e5e402de893f996a02ad42aff8ec4ade9971ee45c71e8fbcb47a.NewDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetScheduledActions the setScheduledActions property
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) SetScheduledActions()(*i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.SetScheduledActionsRequestBuilder) {
    return i1fb378ecdc09aaabf5c55da82fe6def98379ca93c0d2715669e8fa5761d1ed11.NewSetScheduledActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) Settings()(*i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.SettingsRequestBuilder) {
    return i15585ed0e13327df9171fd8499ee3db205ffd397a580c4ed84e065a5d220e628.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.compliancePolicies.item.settings.item collection
func (m *DeviceManagementCompliancePolicyItemRequestBuilder) SettingsById(id string)(*ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.DeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return ie4d81e74e3624280f0e519971ef2c7635b27540efa521cec95913734fcecc59f.NewDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
