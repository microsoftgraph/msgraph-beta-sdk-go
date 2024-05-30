package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder provides operations to call the sharedWithMe method.
type FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
type FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters
}
// NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/sharedWithMe(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// Deprecated: This method is obsolete. Use GetAsSharedWithMeGetResponse instead.
// returns a FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeResponseable), nil
}
// GetAsSharedWithMeGetResponse get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// returns a FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) GetAsSharedWithMeGetResponse(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeGetResponseable), nil
}
// ToGetRequestInformation get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they're items from a different drive.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveSharedwithmeSharedWithMeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
