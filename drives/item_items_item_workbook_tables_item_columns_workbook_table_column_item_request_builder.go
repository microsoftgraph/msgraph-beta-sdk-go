package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
type ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters retrieve the properties and relationships of tablecolumn object.
type ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal instantiates a new WorkbookTableColumnItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder instantiates a new WorkbookTableColumnItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) DataBodyRange()(*ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes the column from the table.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumn-delete?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Filter provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Filter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of tablecolumn object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumn-get?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) HeaderRowRange()(*ItemItemsItemWorkbookTablesItemColumnsItemHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of tablecolumn object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumn-update?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable), nil
}
// RangeEscaped provides operations to call the range method.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookTablesItemColumnsItemRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes the column from the table.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of tablecolumn object.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of tablecolumn object.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookTableColumnable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// TotalRowRange provides operations to call the totalRowRange method.
func (m *ItemItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) TotalRowRange()(*ItemItemsItemWorkbookTablesItemColumnsItemTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
