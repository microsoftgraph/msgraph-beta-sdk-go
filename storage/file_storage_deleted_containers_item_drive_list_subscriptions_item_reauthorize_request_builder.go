package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder provides operations to call the reauthorize method.
type FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list/subscriptions/{subscription%2Did}/reauthorize", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reauthorize a subscription when you receive a reauthorizationRequired challenge.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subscription-reauthorize?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(error) {
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
func (m *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListSubscriptionsItemReauthorizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
