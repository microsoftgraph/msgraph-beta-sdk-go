package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder provides operations to call the reapplyFilters method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/itemAt(index={index})/reapplyFilters", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reapplies all the filters currently on the table.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/table-reapplyfilters?view=graph-rest-beta
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reapplies all the filters currently on the table.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemAtWithIndexReapplyFiltersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
