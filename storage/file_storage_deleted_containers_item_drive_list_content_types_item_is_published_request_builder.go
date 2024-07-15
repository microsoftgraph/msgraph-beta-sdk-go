package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder provides operations to call the isPublished method.
type FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function isPublished
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse invoke function isPublished
// returns a FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedGetResponseable), nil
}
// ToGetRequestInformation invoke function isPublished
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
