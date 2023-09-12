package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
type CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters the list of scheduled action for this rule
type CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters
}
// CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal instantiates a new DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    m := &CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder instantiates a new DeviceManagementComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property scheduledActionsForRule for deviceManagement
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of scheduled action for this rule
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurations()(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    return NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property scheduledActionsForRule for deviceManagement
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of scheduled action for this rule
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property scheduledActionsForRule in deviceManagement
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) WithUrl(rawUrl string)(*CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    return NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
