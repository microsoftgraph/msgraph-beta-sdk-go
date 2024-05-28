package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder provides operations to manage the exactMatchUploadAgents property of the microsoft.graph.dataClassificationService entity.
type ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetQueryParameters get exactMatchUploadAgents from dataClassification
type ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetQueryParameters
}
// ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByExactMatchUploadAgentId provides operations to manage the exactMatchUploadAgents property of the microsoft.graph.dataClassificationService entity.
// returns a *ExactmatchuploadagentsExactMatchUploadAgentItemRequestBuilder when successful
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) ByExactMatchUploadAgentId(exactMatchUploadAgentId string)(*ExactmatchuploadagentsExactMatchUploadAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if exactMatchUploadAgentId != "" {
        urlTplParams["exactMatchUploadAgent%2Did"] = exactMatchUploadAgentId
    }
    return NewExactmatchuploadagentsExactMatchUploadAgentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderInternal instantiates a new ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder and sets the default values.
func NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) {
    m := &ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchUploadAgents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder instantiates a new ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder and sets the default values.
func NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ExactmatchuploadagentsCountRequestBuilder when successful
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) Count()(*ExactmatchuploadagentsCountRequestBuilder) {
    return NewExactmatchuploadagentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exactMatchUploadAgents from dataClassification
// returns a ExactMatchUploadAgentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchUploadAgentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentCollectionResponseable), nil
}
// Post create new navigation property to exactMatchUploadAgents for dataClassification
// returns a ExactMatchUploadAgentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get exactMatchUploadAgents from dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to exactMatchUploadAgents for dataClassification
// returns a *RequestInformation when successful
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchUploadAgentable, requestConfiguration *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder when successful
func (m *ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) WithUrl(rawUrl string)(*ExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder) {
    return NewExactmatchuploadagentsExactMatchUploadAgentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
