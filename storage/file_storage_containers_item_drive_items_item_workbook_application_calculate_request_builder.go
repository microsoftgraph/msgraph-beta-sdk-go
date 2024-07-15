package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder provides operations to call the calculate method.
type FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/application/calculate", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recalculate all currently opened workbooks in Excel.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workbookapplication-calculate?view=graph-rest-beta
func (m *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculatePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation recalculate all currently opened workbooks in Excel.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculatePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookApplicationCalculateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
