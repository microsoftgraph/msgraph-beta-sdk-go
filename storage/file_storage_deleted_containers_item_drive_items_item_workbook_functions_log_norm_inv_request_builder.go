package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder provides operations to call the logNorm_Inv method.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/logNorm_Inv", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action logNorm_Inv
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) Post(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvPostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action logNorm_Inv
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvPostRequestBodyable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
