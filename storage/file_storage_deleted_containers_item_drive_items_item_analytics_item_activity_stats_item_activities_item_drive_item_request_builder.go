package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivity entity.
type FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters exposes the driveItem that was the target of this activity.
type FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters
}
// NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}/driveItem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) Content()(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ContentStream provides operations to manage the media for the storage entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentStreamRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) ContentStream()(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentStreamRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the driveItem that was the target of this activity.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ToGetRequestInformation exposes the driveItem that was the target of this activity.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
