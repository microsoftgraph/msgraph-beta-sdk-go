package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder provides operations to call the tableRowOperationResult method.
type ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, key *string)(*ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) {
    m := &ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tableRowOperationResult(key='{key}')", pathParameters),
    }
    if key != nil {
        m.BaseRequestBuilder.PathParameters["key"] = *key
    }
    return m
}
// NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder instantiates a new ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function tableRowOperationResult
// returns a WorkbookTableRowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableRowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowable), nil
}
// ToGetRequestInformation invoke function tableRowOperationResult
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder when successful
func (m *ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) {
    return NewItemItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
