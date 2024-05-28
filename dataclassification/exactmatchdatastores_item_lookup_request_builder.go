package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactmatchdatastoresItemLookupRequestBuilder provides operations to call the lookup method.
type ExactmatchdatastoresItemLookupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactmatchdatastoresItemLookupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresItemLookupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactmatchdatastoresItemLookupRequestBuilderInternal instantiates a new ExactmatchdatastoresItemLookupRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemLookupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemLookupRequestBuilder) {
    m := &ExactmatchdatastoresItemLookupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/lookup", pathParameters),
    }
    return m
}
// NewExactmatchdatastoresItemLookupRequestBuilder instantiates a new ExactmatchdatastoresItemLookupRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemLookupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemLookupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactmatchdatastoresItemLookupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action lookup
// Deprecated: This method is obsolete. Use PostAsLookupPostResponse instead.
// returns a ExactmatchdatastoresItemLookupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresItemLookupRequestBuilder) Post(ctx context.Context, body ExactmatchdatastoresItemLookupPostRequestBodyable, requestConfiguration *ExactmatchdatastoresItemLookupRequestBuilderPostRequestConfiguration)(ExactmatchdatastoresItemLookupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExactmatchdatastoresItemLookupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExactmatchdatastoresItemLookupResponseable), nil
}
// PostAsLookupPostResponse invoke action lookup
// returns a ExactmatchdatastoresItemLookupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresItemLookupRequestBuilder) PostAsLookupPostResponse(ctx context.Context, body ExactmatchdatastoresItemLookupPostRequestBodyable, requestConfiguration *ExactmatchdatastoresItemLookupRequestBuilderPostRequestConfiguration)(ExactmatchdatastoresItemLookupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExactmatchdatastoresItemLookupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExactmatchdatastoresItemLookupPostResponseable), nil
}
// ToPostRequestInformation invoke action lookup
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresItemLookupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExactmatchdatastoresItemLookupPostRequestBodyable, requestConfiguration *ExactmatchdatastoresItemLookupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExactmatchdatastoresItemLookupRequestBuilder when successful
func (m *ExactmatchdatastoresItemLookupRequestBuilder) WithUrl(rawUrl string)(*ExactmatchdatastoresItemLookupRequestBuilder) {
    return NewExactmatchdatastoresItemLookupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
