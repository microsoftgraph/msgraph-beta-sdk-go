package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalaccessEvaluateRequestBuilder provides operations to call the evaluate method.
type ConditionalaccessEvaluateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessEvaluateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessEvaluateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessEvaluateRequestBuilderInternal instantiates a new ConditionalaccessEvaluateRequestBuilder and sets the default values.
func NewConditionalaccessEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessEvaluateRequestBuilder) {
    m := &ConditionalaccessEvaluateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/evaluate", pathParameters),
    }
    return m
}
// NewConditionalaccessEvaluateRequestBuilder instantiates a new ConditionalaccessEvaluateRequestBuilder and sets the default values.
func NewConditionalaccessEvaluateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessEvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluate
// Deprecated: This method is obsolete. Use PostAsEvaluatePostResponse instead.
// returns a ConditionalaccessEvaluateResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessEvaluateRequestBuilder) Post(ctx context.Context, body ConditionalaccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalaccessEvaluateRequestBuilderPostRequestConfiguration)(ConditionalaccessEvaluateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalaccessEvaluateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalaccessEvaluateResponseable), nil
}
// PostAsEvaluatePostResponse invoke action evaluate
// returns a ConditionalaccessEvaluatePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessEvaluateRequestBuilder) PostAsEvaluatePostResponse(ctx context.Context, body ConditionalaccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalaccessEvaluateRequestBuilderPostRequestConfiguration)(ConditionalaccessEvaluatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalaccessEvaluatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalaccessEvaluatePostResponseable), nil
}
// ToPostRequestInformation invoke action evaluate
// returns a *RequestInformation when successful
func (m *ConditionalaccessEvaluateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConditionalaccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalaccessEvaluateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessEvaluateRequestBuilder when successful
func (m *ConditionalaccessEvaluateRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessEvaluateRequestBuilder) {
    return NewConditionalaccessEvaluateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
