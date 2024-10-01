package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemPermanentDeleteRequestBuilder provides operations to call the permanentDelete method.
type FileStorageDeletedContainersItemPermanentDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemPermanentDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemPermanentDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemPermanentDeleteRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) {
    m := &FileStorageDeletedContainersItemPermanentDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/permanentDelete", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilder instantiates a new FileStorageDeletedContainersItemPermanentDeleteRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete a fileStorageContainer object. Deleting a file storage container with this method permanently removes it. It isn't sent to the deleted container collection. Therefore, it can't be restored later.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/filestoragecontainer-permanentdelete?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemPermanentDeleteRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation permanently delete a fileStorageContainer object. Deleting a file storage container with this method permanently removes it. It isn't sent to the deleted container collection. Therefore, it can't be restored later.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemPermanentDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemPermanentDeleteRequestBuilder when successful
func (m *FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemPermanentDeleteRequestBuilder) {
    return NewFileStorageDeletedContainersItemPermanentDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
