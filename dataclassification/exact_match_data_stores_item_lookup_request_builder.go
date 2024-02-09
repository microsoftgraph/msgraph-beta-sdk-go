package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactMatchDataStoresItemLookupRequestBuilder provides operations to call the lookup method.
type ExactMatchDataStoresItemLookupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactMatchDataStoresItemLookupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemLookupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactMatchDataStoresItemLookupRequestBuilderInternal instantiates a new ExactMatchDataStoresItemLookupRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemLookupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemLookupRequestBuilder) {
    m := &ExactMatchDataStoresItemLookupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/lookup", pathParameters),
    }
    return m
}
// NewExactMatchDataStoresItemLookupRequestBuilder instantiates a new ExactMatchDataStoresItemLookupRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemLookupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemLookupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchDataStoresItemLookupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action lookup
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ExactMatchDataStoresItemLookupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemLookupRequestBuilder) Post(ctx context.Context, body ExactMatchDataStoresItemLookupPostRequestBodyable, requestConfiguration *ExactMatchDataStoresItemLookupRequestBuilderPostRequestConfiguration)(ExactMatchDataStoresItemLookupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExactMatchDataStoresItemLookupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExactMatchDataStoresItemLookupResponseable), nil
}
// PostAsLookupPostResponse invoke action lookup
// returns a ExactMatchDataStoresItemLookupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemLookupRequestBuilder) PostAsLookupPostResponse(ctx context.Context, body ExactMatchDataStoresItemLookupPostRequestBodyable, requestConfiguration *ExactMatchDataStoresItemLookupRequestBuilderPostRequestConfiguration)(ExactMatchDataStoresItemLookupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExactMatchDataStoresItemLookupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExactMatchDataStoresItemLookupPostResponseable), nil
}
// ToPostRequestInformation invoke action lookup
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemLookupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExactMatchDataStoresItemLookupPostRequestBodyable, requestConfiguration *ExactMatchDataStoresItemLookupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExactMatchDataStoresItemLookupRequestBuilder when successful
func (m *ExactMatchDataStoresItemLookupRequestBuilder) WithUrl(rawUrl string)(*ExactMatchDataStoresItemLookupRequestBuilder) {
    return NewExactMatchDataStoresItemLookupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
