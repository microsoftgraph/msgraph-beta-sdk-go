package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemInviteRequestBuilder provides operations to call the invite method.
type FileStorageContainersItemDriveItemsItemInviteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemInviteRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemInviteRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemInviteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemInviteRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemInviteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/invite", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemInviteRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemInviteRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemInviteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemInviteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemInviteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// Deprecated: This method is obsolete. Use PostAsInvitePostResponse instead.
// returns a FileStorageContainersItemDriveItemsItemInviteResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-invite?view=graph-rest-beta
func (m *FileStorageContainersItemDriveItemsItemInviteRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(FileStorageContainersItemDriveItemsItemInviteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemInviteResponseable), nil
}
// PostAsInvitePostResponse sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// returns a FileStorageContainersItemDriveItemsItemInvitePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-invite?view=graph-rest-beta
func (m *FileStorageContainersItemDriveItemsItemInviteRequestBuilder) PostAsInvitePostResponse(ctx context.Context, body FileStorageContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(FileStorageContainersItemDriveItemsItemInvitePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemInvitePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemInvitePostResponseable), nil
}
// ToPostRequestInformation sends a sharing invitation for a driveItem.A sharing invitation provides permissions to the recipients and optionally sends an email to the recipients to notify them the item was shared.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemInviteRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemInvitePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemInviteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemInviteRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemInviteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
