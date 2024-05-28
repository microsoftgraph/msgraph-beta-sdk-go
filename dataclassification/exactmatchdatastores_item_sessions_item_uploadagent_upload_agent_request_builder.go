package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder provides operations to manage the uploadAgent property of the microsoft.graph.exactMatchSession entity.
type ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetQueryParameters get uploadAgent from dataClassification
type ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetQueryParameters
}
// ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderInternal instantiates a new ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) {
    m := &ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}/uploadAgent{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder instantiates a new ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder and sets the default values.
func NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property uploadAgent for dataClassification
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) Get(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, error) {
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
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, error) {
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
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get uploadAgent from dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder when successful
func (m *ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) WithUrl(rawUrl string)(*ExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder) {
    return NewExactmatchdatastoresItemSessionsItemUploadagentUploadAgentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
