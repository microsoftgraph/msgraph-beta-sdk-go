package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters collection of tables that are part of the worksheet. Read-only.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Columns()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ColumnsById(id string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal instantiates a new WorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder instantiates a new WorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tables for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of tables that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable), nil
}
// MicrosoftGraphClearFilters provides operations to call the clearFilters method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphClearFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphConvertToRange provides operations to call the convertToRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphConvertToRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphConvertToRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDataBodyRange provides operations to call the dataBodyRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphDataBodyRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHeaderRowRange provides operations to call the headerRowRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphHeaderRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRange provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReapplyFilters provides operations to call the reapplyFilters method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphReapplyFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphReapplyFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTotalRowRange provides operations to call the totalRowRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) MicrosoftGraphTotalRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemMicrosoftGraphTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property tables in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable), nil
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Rows()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RowsById provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) RowsById(id string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsWorkbookTableRowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsWorkbookTableRowItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Sort()(*ItemItemsItemWorkbookWorksheetsItemTablesItemSortRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemSortRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property tables for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property tables in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemTablesItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
