package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder provides operations to manage the uploadAgent property of the microsoft.graph.exactMatchSession entity.
type ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetQueryParameters get uploadAgent from dataClassification
type ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetQueryParameters
}
// ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderInternal instantiates a new ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) {
    m := &ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/uploadAgent{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder instantiates a new ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property uploadAgent for dataClassification
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get uploadAgent from dataClassification
// returns a ExactMatchUploadAgentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) Get(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchUploadAgentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable), nil
}
// Patch update the navigation property uploadAgent in dataClassification
// returns a ExactMatchUploadAgentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchUploadAgentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable), nil
}
// ToDeleteRequestInformation delete navigation property uploadAgent for dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/uploadAgent", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get uploadAgent from dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property uploadAgent in dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/uploadAgent", m.BaseRequestBuilder.PathParameters)
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
// returns a *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) WithUrl(rawUrl string)(*ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
