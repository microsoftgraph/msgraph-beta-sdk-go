package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ClearFilters()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Columns()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ConvertToRange()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) DataBodyRange()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function itemAt
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) HeaderRowRange()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) RangeEscaped()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ReapplyFilters()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Rows()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Sort()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function itemAt
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) TotalRowRange()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Worksheet()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
