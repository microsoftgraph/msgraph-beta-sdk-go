package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder provides operations to manage the originatingKnowledgeBaseArticle property of the microsoft.graph.windowsUpdates.knownIssue entity.
type WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetQueryParameters knowledge base article associated with the release when the known issue was first reported.
type WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetQueryParameters
}
// WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderInternal instantiates a new WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) {
    m := &WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/{product%2Did}/knownIssues/{knownIssue%2Did}/originatingKnowledgeBaseArticle{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder instantiates a new WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property originatingKnowledgeBaseArticle for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get knowledge base article associated with the release when the known issue was first reported.
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
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
// Patch update the navigation property originatingKnowledgeBaseArticle in admin
// returns a KnowledgeBaseArticleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, error) {
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
// ToDeleteRequestInformation delete navigation property originatingKnowledgeBaseArticle for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/admin/windows/updates/products/{product%2Did}/knownIssues/{knownIssue%2Did}/originatingKnowledgeBaseArticle", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation knowledge base article associated with the release when the known issue was first reported.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property originatingKnowledgeBaseArticle in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnowledgeBaseArticleable, requestConfiguration *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/admin/windows/updates/products/{product%2Did}/knownIssues/{knownIssue%2Did}/originatingKnowledgeBaseArticle", m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder when successful
func (m *WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder) {
    return NewWindowsUpdatesProductsItemKnownIssuesItemOriginatingKnowledgeBaseArticleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
