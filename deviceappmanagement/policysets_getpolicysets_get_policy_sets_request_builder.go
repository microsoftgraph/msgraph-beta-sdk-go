package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicysetsGetpolicysetsGetPolicySetsRequestBuilder provides operations to call the getPolicySets method.
type PolicysetsGetpolicysetsGetPolicySetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PolicysetsGetpolicysetsGetPolicySetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicysetsGetpolicysetsGetPolicySetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilderInternal instantiates a new PolicysetsGetpolicysetsGetPolicySetsRequestBuilder and sets the default values.
func NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) {
    m := &PolicysetsGetpolicysetsGetPolicySetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/policySets/getPolicySets", pathParameters),
    }
    return m
}
// NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilder instantiates a new PolicysetsGetpolicysetsGetPolicySetsRequestBuilder and sets the default values.
func NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPolicySets
// Deprecated: This method is obsolete. Use PostAsGetPolicySetsPostResponse instead.
// returns a PolicysetsGetpolicysetsGetPolicySetsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) Post(ctx context.Context, body PolicysetsGetpolicysetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicysetsGetpolicysetsGetPolicySetsRequestBuilderPostRequestConfiguration)(PolicysetsGetpolicysetsGetPolicySetsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicysetsGetpolicysetsGetPolicySetsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicysetsGetpolicysetsGetPolicySetsResponseable), nil
}
// PostAsGetPolicySetsPostResponse invoke action getPolicySets
// returns a PolicysetsGetpolicysetsGetPolicySetsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) PostAsGetPolicySetsPostResponse(ctx context.Context, body PolicysetsGetpolicysetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicysetsGetpolicysetsGetPolicySetsRequestBuilderPostRequestConfiguration)(PolicysetsGetpolicysetsGetPolicySetsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicysetsGetpolicysetsGetPolicySetsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicysetsGetpolicysetsGetPolicySetsPostResponseable), nil
}
// ToPostRequestInformation invoke action getPolicySets
// returns a *RequestInformation when successful
func (m *PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicysetsGetpolicysetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicysetsGetpolicysetsGetPolicySetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PolicysetsGetpolicysetsGetPolicySetsRequestBuilder when successful
func (m *PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) WithUrl(rawUrl string)(*PolicysetsGetpolicysetsGetPolicySetsRequestBuilder) {
    return NewPolicysetsGetpolicysetsGetPolicySetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
