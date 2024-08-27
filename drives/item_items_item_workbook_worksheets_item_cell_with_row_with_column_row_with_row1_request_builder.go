package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder provides operations to call the row method.
type ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, row1 *int32)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/cell(row={row},column={column})/row(row={row1})", pathParameters),
    }
    if row1 != nil {
        m.BaseRequestBuilder.PathParameters["row1"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*row1), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function row
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookRangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable), nil
}
// ToGetRequestInformation invoke function row
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRowWithRow1RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
