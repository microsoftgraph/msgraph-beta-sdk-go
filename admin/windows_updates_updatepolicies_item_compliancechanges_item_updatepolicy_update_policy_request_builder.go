package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder provides operations to manage the updatePolicy property of the microsoft.graph.windowsUpdates.complianceChange entity.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetQueryParameters the policy this compliance change is a member of.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetQueryParameters
}
// NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/complianceChanges/{complianceChange%2Did}/updatePolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the policy this compliance change is a member of.
// returns a UpdatePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatePolicyable), nil
}
// ToGetRequestInformation the policy this compliance change is a member of.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemCompliancechangesItemUpdatepolicyUpdatePolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
