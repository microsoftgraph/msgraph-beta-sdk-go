package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactmatchdatastoresItemSessionsItemRenewRequestBuilder provides operations to call the renew method.
type ExactmatchdatastoresItemSessionsItemRenewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactmatchdatastoresItemSessionsItemRenewRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresItemSessionsItemRenewRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactmatchdatastoresItemSessionsItemRenewRequestBuilderInternal instantiates a new ExactmatchdatastoresItemSessionsItemRenewRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemSessionsItemRenewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) {
    m := &ExactmatchdatastoresItemSessionsItemRenewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/renew", pathParameters),
    }
    return m
}
// NewExactmatchdatastoresItemSessionsItemRenewRequestBuilder instantiates a new ExactmatchdatastoresItemSessionsItemRenewRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemSessionsItemRenewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactmatchdatastoresItemSessionsItemRenewRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action renew
// returns a ExactMatchSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) Post(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemRenewRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable), nil
}
// ToPostRequestInformation invoke action renew
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemRenewRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExactmatchdatastoresItemSessionsItemRenewRequestBuilder when successful
func (m *ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) WithUrl(rawUrl string)(*ExactmatchdatastoresItemSessionsItemRenewRequestBuilder) {
    return NewExactmatchdatastoresItemSessionsItemRenewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
