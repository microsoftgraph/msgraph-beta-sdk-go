package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder provides operations to call the count method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/comments/{workbookComment%2Did}/replies/{workbookCommentReply%2Did}/task/changes/count()", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function count
// Deprecated: This method is obsolete. Use GetAsCountGetResponse instead.
// returns a FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseable), nil
}
// GetAsCountGetResponse invoke function count
// returns a FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) GetAsCountGetResponse(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable), nil
}
// ToGetRequestInformation invoke function count
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
