package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder provides operations to call the resizedRange method.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deltaColumns *int32, deltaRows *int32)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/dataBodyRange()/resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})", pathParameters),
    }
    if deltaColumns != nil {
        m.BaseRequestBuilder.PathParameters["deltaColumns"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaColumns), 10)
    }
    if deltaRows != nil {
        m.BaseRequestBuilder.PathParameters["deltaRows"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaRows), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function resizedRange
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable, error) {
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
// ToGetRequestInformation invoke function resizedRange
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemDataBodyRangeResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
