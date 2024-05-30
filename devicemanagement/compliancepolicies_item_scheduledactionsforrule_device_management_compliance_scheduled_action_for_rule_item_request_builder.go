package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
type CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters the list of scheduled action for this rule
type CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters
}
// CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal instantiates a new CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    m := &CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder instantiates a new CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property scheduledActionsForRule for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of scheduled action for this rule
// returns a DeviceManagementComplianceScheduledActionForRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable), nil
}
// Patch update the navigation property scheduledActionsForRule in deviceManagement
// returns a DeviceManagementComplianceScheduledActionForRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable), nil
}
// ScheduledActionConfigurations provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
// returns a *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurations()(*CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) {
    return NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property scheduledActionsForRule for deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of scheduled action for this rule
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property scheduledActionsForRule in deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) WithUrl(rawUrl string)(*CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    return NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
