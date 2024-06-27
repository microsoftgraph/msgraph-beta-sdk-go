package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/series/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChartSeries entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexFormatRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) Format()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexFormatRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieves a series based on its position in the collection
// returns a WorkbookChartSeriesable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartseriescollection-itemat?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartSeriesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartSeriesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartSeriesable), nil
}
// Points provides operations to manage the points property of the microsoft.graph.workbookChartSeries entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexPointsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) Points()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexPointsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexPointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieves a series based on its position in the collection
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
