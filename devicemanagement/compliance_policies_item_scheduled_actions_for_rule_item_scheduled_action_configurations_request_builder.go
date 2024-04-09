package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
type CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
type CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters struct {
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
// CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters
}
// CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementComplianceActionItemId provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
// returns a *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder when successful
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) ByDeviceManagementComplianceActionItemId(deviceManagementComplianceActionItemId string)(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementComplianceActionItemId != "" {
        urlTplParams["deviceManagementComplianceActionItem%2Did"] = deviceManagementComplianceActionItemId
    }
    return NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal instantiates a new CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    m := &CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder instantiates a new CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsCountRequestBuilder when successful
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Count()(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsCountRequestBuilder) {
    return NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// returns a DeviceManagementComplianceActionItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemCollectionResponseable), nil
}
// Post create new navigation property to scheduledActionConfigurations for deviceManagement
// returns a DeviceManagementComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable), nil
}
// ToGetRequestInformation the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// returns a *RequestInformation when successful
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scheduledActionConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder when successful
func (m *CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) WithUrl(rawUrl string)(*CompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    return NewCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
