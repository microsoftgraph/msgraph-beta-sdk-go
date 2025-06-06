// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
type DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetQueryParameters the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
type DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetQueryParameters struct {
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
// DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetQueryParameters
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceComplianceScheduledActionForRuleId provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder when successful
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) ByDeviceComplianceScheduledActionForRuleId(deviceComplianceScheduledActionForRuleId string)(*DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceComplianceScheduledActionForRuleId != "" {
        urlTplParams["deviceComplianceScheduledActionForRule%2Did"] = deviceComplianceScheduledActionForRuleId
    }
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal instantiates a new DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    m := &DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder instantiates a new DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceCompliancePoliciesItemScheduledActionsForRuleCountRequestBuilder when successful
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) Count()(*DeviceCompliancePoliciesItemScheduledActionsForRuleCountRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
// returns a DeviceComplianceScheduledActionForRuleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScheduledActionForRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleCollectionResponseable), nil
}
// Post create new navigation property to scheduledActionsForRule for deviceManagement
// returns a DeviceComplianceScheduledActionForRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleable), nil
}
// ToGetRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
// returns a *RequestInformation when successful
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScheduledActionForRuleable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder when successful
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) WithUrl(rawUrl string)(*DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
