package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder provides operations to call the isPublished method.
type FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function isPublished
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse invoke function isPublished
// returns a FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable), nil
}
// ToGetRequestInformation invoke function isPublished
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
