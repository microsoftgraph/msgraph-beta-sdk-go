package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder provides operations to manage the exactMatchDataStores property of the microsoft.graph.dataClassificationService entity.
type ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetQueryParameters get exactMatchDataStores from dataClassification
type ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetQueryParameters
}
// ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilderInternal instantiates a new ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder and sets the default values.
func NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) {
    m := &ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilder instantiates a new ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder and sets the default values.
func NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exactMatchDataStores for dataClassification
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get exactMatchDataStores from dataClassification
// returns a ExactMatchDataStoreable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchDataStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable), nil
}
// Lookup provides operations to call the lookup method.
// returns a *ExactmatchdatastoresItemLookupRequestBuilder when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) Lookup()(*ExactmatchdatastoresItemLookupRequestBuilder) {
    return NewExactmatchdatastoresItemLookupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property exactMatchDataStores in dataClassification
// returns a ExactMatchDataStoreable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchDataStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable), nil
}
// Sessions provides operations to manage the sessions property of the microsoft.graph.exactMatchDataStore entity.
// returns a *ExactmatchdatastoresItemSessionsRequestBuilder when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) Sessions()(*ExactmatchdatastoresItemSessionsRequestBuilder) {
    return NewExactmatchdatastoresItemSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property exactMatchDataStores for dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get exactMatchDataStores from dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property exactMatchDataStores in dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchDataStoreable, requestConfiguration *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder when successful
func (m *ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) WithUrl(rawUrl string)(*ExactmatchdatastoresExactMatchDataStoreItemRequestBuilder) {
    return NewExactmatchdatastoresExactMatchDataStoreItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
