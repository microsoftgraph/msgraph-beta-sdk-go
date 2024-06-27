package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessEvaluateRequestBuilder provides operations to call the evaluate method.
type ConditionalAccessEvaluateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessEvaluateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessEvaluateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalAccessEvaluateRequestBuilderInternal instantiates a new ConditionalAccessEvaluateRequestBuilder and sets the default values.
func NewConditionalAccessEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessEvaluateRequestBuilder) {
    m := &ConditionalAccessEvaluateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/evaluate", pathParameters),
    }
    return m
}
// NewConditionalAccessEvaluateRequestBuilder instantiates a new ConditionalAccessEvaluateRequestBuilder and sets the default values.
func NewConditionalAccessEvaluateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessEvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluate
// Deprecated: This method is obsolete. Use PostAsEvaluatePostResponse instead.
// returns a ConditionalAccessEvaluateResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalAccessEvaluateRequestBuilder) Post(ctx context.Context, body ConditionalAccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalAccessEvaluateRequestBuilderPostRequestConfiguration)(ConditionalAccessEvaluateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalAccessEvaluateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalAccessEvaluateResponseable), nil
}
// PostAsEvaluatePostResponse invoke action evaluate
// returns a ConditionalAccessEvaluatePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalAccessEvaluateRequestBuilder) PostAsEvaluatePostResponse(ctx context.Context, body ConditionalAccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalAccessEvaluateRequestBuilderPostRequestConfiguration)(ConditionalAccessEvaluatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConditionalAccessEvaluatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConditionalAccessEvaluatePostResponseable), nil
}
// ToPostRequestInformation invoke action evaluate
// returns a *RequestInformation when successful
func (m *ConditionalAccessEvaluateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConditionalAccessEvaluatePostRequestBodyable, requestConfiguration *ConditionalAccessEvaluateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalAccessEvaluateRequestBuilder when successful
func (m *ConditionalAccessEvaluateRequestBuilder) WithUrl(rawUrl string)(*ConditionalAccessEvaluateRequestBuilder) {
    return NewConditionalAccessEvaluateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
