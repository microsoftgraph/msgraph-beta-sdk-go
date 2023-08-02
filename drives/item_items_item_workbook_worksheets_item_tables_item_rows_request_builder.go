package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetQueryParameters retrieve a list of tablerow objects.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Add provides operations to call the add method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) Add()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsAddRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByWorkbookTableRowId provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) ByWorkbookTableRowId(workbookTableRowId string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsWorkbookTableRowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookTableRowId != "" {
        urlTplParams["workbookTableRow%2Did"] = workbookTableRowId
    }
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsWorkbookTableRowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderInternal instantiates a new RowsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/rows{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder instantiates a new RowsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to call the count method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) Count()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of tablerow objects.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablerow-list?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableRowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowCollectionResponseable), nil
}
// ItemAtWithIndex provides operations to call the itemAt method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) ItemAtWithIndex(index *int32)(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsItemAtWithIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, index)
}
// Post add rows to the end of a table.  Note that the API can accept multiple rows data using this API. Adding one row at a time can lead to performance degradation. The recommended approach is to batch the rows together in a single call rather than doing single row insertion. For best results, collect the rows to be inserted on the application side and perform a single row add operation. Experiment with the number of rows to determine the ideal number of rows to use in a single API call.  This request might occasionally result in a 504 HTTP error. The appropriate response to this error is to repeat the request.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/table-post-rows?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation retrieve a list of tablerow objects.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation add rows to the end of a table.  Note that the API can accept multiple rows data using this API. Adding one row at a time can lead to performance degradation. The recommended approach is to batch the rows together in a single call rather than doing single row insertion. For best results, collect the rows to be inserted on the application side and perform a single row add operation. Experiment with the number of rows to determine the ideal number of rows to use in a single API call.  This request might occasionally result in a 504 HTTP error. The appropriate response to this error is to repeat the request.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableRowable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
