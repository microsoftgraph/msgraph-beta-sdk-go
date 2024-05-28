package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder provides operations to call the setScheduledActions method.
type CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderInternal instantiates a new CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder and sets the default values.
func NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) {
    m := &CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/setScheduledActions", pathParameters),
    }
    return m
}
// NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder instantiates a new CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder and sets the default values.
func NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setScheduledActions
// Deprecated: This method is obsolete. Use PostAsSetScheduledActionsPostResponse instead.
// returns a CompliancepoliciesItemSetscheduledactionsSetScheduledActionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) Post(ctx context.Context, body CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostRequestBodyable, requestConfiguration *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderPostRequestConfiguration)(CompliancepoliciesItemSetscheduledactionsSetScheduledActionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompliancepoliciesItemSetscheduledactionsSetScheduledActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompliancepoliciesItemSetscheduledactionsSetScheduledActionsResponseable), nil
}
// PostAsSetScheduledActionsPostResponse invoke action setScheduledActions
// returns a CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) PostAsSetScheduledActionsPostResponse(ctx context.Context, body CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostRequestBodyable, requestConfiguration *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderPostRequestConfiguration)(CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostResponseable), nil
}
// ToPostRequestInformation invoke action setScheduledActions
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CompliancepoliciesItemSetscheduledactionsSetScheduledActionsPostRequestBodyable, requestConfiguration *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder when successful
func (m *CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) WithUrl(rawUrl string)(*CompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder) {
    return NewCompliancepoliciesItemSetscheduledactionsSetScheduledActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
