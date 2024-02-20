package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder provides operations to call the count method.
type ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderInternal instantiates a new ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) {
    m := &ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/comments/{workbookComment%2Did}/task/comment/replies/{workbookCommentReply%2Did}/task/changes/count()", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder instantiates a new ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function count
// Deprecated: This method is obsolete. Use GetAsCountGetResponse instead.
// returns a ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountResponseable), nil
}
// GetAsCountGetResponse invoke function count
// returns a ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) GetAsCountGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountGetResponseable), nil
}
// ToGetRequestInformation invoke function count
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder when successful
func (m *ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder) {
    return NewItemItemsItemWorkbookCommentsItemTaskCommentRepliesItemTaskChangesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
