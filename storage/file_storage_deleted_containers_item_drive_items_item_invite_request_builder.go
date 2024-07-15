package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder provides operations to call the invite method.
type FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/invite", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// Deprecated: This method is obsolete. Use PostAsInvitePostResponse instead.
// returns a FileStorageDeletedContainersItemDriveItemsItemInviteResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-invite?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) Post(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemInviteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemInviteResponseable), nil
}
// PostAsInvitePostResponse sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// returns a FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-invite?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) PostAsInvitePostResponse(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemInvitePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable), nil
}
// ToPostRequestInformation sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
