package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder provides operations to manage the resolvingKnowledgeBaseArticle property of the microsoft.graph.windowsUpdates.knownIssue entity.
type WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetQueryParameters knowledge base article associated with the release when the known issue was resolved or mitigated.
type WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetQueryParameters
}
// WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderInternal instantiates a new WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) {
    m := &WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/{product%2Did}/knownIssues/{knownIssue%2Did}/resolvingKnowledgeBaseArticle{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder instantiates a new WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resolvingKnowledgeBaseArticle for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get knowledge base article associated with the release when the known issue was resolved or mitigated.
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
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
// Patch update the navigation property resolvingKnowledgeBaseArticle in admin
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
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
// ToDeleteRequestInformation delete navigation property resolvingKnowledgeBaseArticle for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation knowledge base article associated with the release when the known issue was resolved or mitigated.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resolvingKnowledgeBaseArticle in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder) {
    return NewWindowsUpdatesProductsItemKnownIssuesItemResolvingKnowledgeBaseArticleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
