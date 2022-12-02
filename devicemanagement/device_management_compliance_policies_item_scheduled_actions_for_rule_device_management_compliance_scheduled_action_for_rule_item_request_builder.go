package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters the list of scheduled action for this rule
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal instantiates a new DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    m := &DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder instantiates a new DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action for this rule
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action for this rule
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable), nil
}
// Patch update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable), nil
}
// ScheduledActionConfigurations provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurations()(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    return NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionConfigurationsById provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurationsById(id string)(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceActionItem%2Did"] = id
    }
    return NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
