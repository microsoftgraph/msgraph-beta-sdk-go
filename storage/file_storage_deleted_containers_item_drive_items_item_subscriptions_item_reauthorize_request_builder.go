package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder provides operations to call the reauthorize method.
type FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/subscriptions/{subscription%2Did}/reauthorize", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reauthorize a subscription when you receive a reauthorizationRequired challenge.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subscription-reauthorize?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reauthorize a subscription when you receive a reauthorizationRequired challenge.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
