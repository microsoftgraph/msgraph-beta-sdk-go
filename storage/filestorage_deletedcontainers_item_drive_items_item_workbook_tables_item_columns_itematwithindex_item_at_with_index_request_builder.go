package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) DataBodyRange()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Filter provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) Filter()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a column based on its position in the collection.
// returns a WorkbookTableColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumncollection-itemat?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) HeaderRowRange()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) RangeEscaped()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a column based on its position in the collection.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) TotalRowRange()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
