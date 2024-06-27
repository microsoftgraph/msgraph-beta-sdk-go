package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder provides operations to call the applyCellColorFilter method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyCellColorFilter", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyCellColorFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) Post(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterPostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action applyCellColorFilter
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterPostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
