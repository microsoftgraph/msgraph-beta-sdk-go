package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder provides operations to call the image method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fittingMode *string, height *int32, width *int32)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')/image(width={width},height={height},fittingMode='{fittingMode}')", pathParameters),
    }
    if fittingMode != nil {
        m.BaseRequestBuilder.PathParameters["fittingMode"] = *fittingMode
    }
    if height != nil {
        m.BaseRequestBuilder.PathParameters["height"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*height), 10)
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthWithHeightWithFittingModeGetResponse instead.
// returns a FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseable), nil
}
// GetAsImageWithWidthWithHeightWithFittingModeGetResponse invoke function image
// returns a FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) GetAsImageWithWidthWithHeightWithFittingModeGetResponse(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
