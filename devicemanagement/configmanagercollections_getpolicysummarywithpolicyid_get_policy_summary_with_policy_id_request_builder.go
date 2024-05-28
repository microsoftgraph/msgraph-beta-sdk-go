package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder provides operations to call the getPolicySummary method.
type ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderInternal instantiates a new ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder and sets the default values.
func NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, policyId *string)(*ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) {
    m := &ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configManagerCollections/getPolicySummary(policyId='{policyId}')", pathParameters),
    }
    if policyId != nil {
        m.BaseRequestBuilder.PathParameters["policyId"] = *policyId
    }
    return m
}
// NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder instantiates a new ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder and sets the default values.
func NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getPolicySummary
// returns a ConfigManagerPolicySummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigManagerPolicySummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable), nil
}
// ToGetRequestInformation invoke function getPolicySummary
// returns a *RequestInformation when successful
func (m *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder when successful
func (m *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) WithUrl(rawUrl string)(*ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) {
    return NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
