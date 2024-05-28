package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder provides operations to manage the knowledgeBaseArticle property of the microsoft.graph.windowsUpdates.productRevision entity.
type WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetQueryParameters the knowledge base article associated with the product revision.
type WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetQueryParameters
}
// WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderInternal instantiates a new WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) {
    m := &WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/{product%2Did}/revisions/{productRevision%2Did}/knowledgeBaseArticle{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder instantiates a new WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property knowledgeBaseArticle for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the knowledge base article associated with the product revision.
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateKnowledgeBaseArticleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable), nil
}
// Patch update the navigation property knowledgeBaseArticle in admin
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateKnowledgeBaseArticleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable), nil
}
// ToDeleteRequestInformation delete navigation property knowledgeBaseArticle for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the knowledge base article associated with the product revision.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property knowledgeBaseArticle in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder when successful
func (m *WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder) {
    return NewWindowsUpdatesProductsItemRevisionsItemKnowledgebasearticleKnowledgeBaseArticleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
