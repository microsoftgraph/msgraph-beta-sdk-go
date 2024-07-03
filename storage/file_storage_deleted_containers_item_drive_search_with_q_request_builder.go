package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder provides operations to call the search method.
type FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetQueryParameters search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
type FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetQueryParameters struct {
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
// FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetQueryParameters
}
// NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, q *string)(*FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/search(q='{q}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if q != nil {
        m.BaseRequestBuilder.PathParameters["q"] = *q
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// Deprecated: This method is obsolete. Use GetAsSearchWithQGetResponse instead.
// returns a FileStorageDeletedContainersItemDriveSearchWithQResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveSearchWithQResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveSearchWithQResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveSearchWithQResponseable), nil
}
// GetAsSearchWithQGetResponse search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a FileStorageDeletedContainersItemDriveSearchWithQGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) GetAsSearchWithQGetResponse(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveSearchWithQGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveSearchWithQGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveSearchWithQGetResponseable), nil
}
// ToGetRequestInformation search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
