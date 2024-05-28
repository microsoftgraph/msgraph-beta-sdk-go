package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
type CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetQueryParameters the list of scheduled action for this rule
type CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetQueryParameters struct {
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
// CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetQueryParameters
}
// CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementComplianceScheduledActionForRuleId provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
// returns a *CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) ByDeviceManagementComplianceScheduledActionForRuleId(deviceManagementComplianceScheduledActionForRuleId string)(*CompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementComplianceScheduledActionForRuleId != "" {
        urlTplParams["deviceManagementComplianceScheduledActionForRule%2Did"] = deviceManagementComplianceScheduledActionForRuleId
    }
    return NewCompliancepoliciesItemScheduledactionsforruleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderInternal instantiates a new CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) {
    m := &CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder instantiates a new CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompliancepoliciesItemScheduledactionsforruleCountRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) Count()(*CompliancepoliciesItemScheduledactionsforruleCountRequestBuilder) {
    return NewCompliancepoliciesItemScheduledactionsforruleCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of scheduled action for this rule
// returns a DeviceManagementComplianceScheduledActionForRuleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceScheduledActionForRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleCollectionResponseable), nil
}
// Post create new navigation property to scheduledActionsForRule for deviceManagement
// returns a DeviceManagementComplianceScheduledActionForRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the list of scheduled action for this rule
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scheduledActionsForRule for deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceScheduledActionForRuleable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) WithUrl(rawUrl string)(*CompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder) {
    return NewCompliancepoliciesItemScheduledactionsforruleScheduledActionsForRuleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
