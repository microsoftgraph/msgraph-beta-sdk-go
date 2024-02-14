package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactMatchDataStoresItemSessionsItemCommitRequestBuilder provides operations to call the commit method.
type ExactMatchDataStoresItemSessionsItemCommitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactMatchDataStoresItemSessionsItemCommitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsItemCommitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactMatchDataStoresItemSessionsItemCommitRequestBuilderInternal instantiates a new ExactMatchDataStoresItemSessionsItemCommitRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsItemCommitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) {
    m := &ExactMatchDataStoresItemSessionsItemCommitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/commit", pathParameters),
    }
    return m
}
// NewExactMatchDataStoresItemSessionsItemCommitRequestBuilder instantiates a new ExactMatchDataStoresItemSessionsItemCommitRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsItemCommitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchDataStoresItemSessionsItemCommitRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action commit
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) Post(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemCommitRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action commit
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemCommitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExactMatchDataStoresItemSessionsItemCommitRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) WithUrl(rawUrl string)(*ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemCommitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
