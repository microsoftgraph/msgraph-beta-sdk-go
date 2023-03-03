package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters retrieve the filter applied to the column. Read-only.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apply provides operations to call the apply method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Apply()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyBottomItemsFilter provides operations to call the applyBottomItemsFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyBottomItemsFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyBottomItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyBottomPercentFilter provides operations to call the applyBottomPercentFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyBottomPercentFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyCellColorFilter provides operations to call the applyCellColorFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyCellColorFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyCustomFilter provides operations to call the applyCustomFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyCustomFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyCustomFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyDynamicFilter provides operations to call the applyDynamicFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyDynamicFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyDynamicFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyFontColorFilter provides operations to call the applyFontColorFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyFontColorFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyIconFilter provides operations to call the applyIconFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyIconFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyIconFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyTopItemsFilter provides operations to call the applyTopItemsFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyTopItemsFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyTopPercentFilter provides operations to call the applyTopPercentFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyTopPercentFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyTopPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyValuesFilter provides operations to call the applyValuesFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ApplyValuesFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyValuesFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Clear provides operations to call the clear method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Clear()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterClearRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterClearRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filter for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the filter applied to the column. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable), nil
}
// Patch update the navigation property filter in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable), nil
}
// ToDeleteRequestInformation delete navigation property filter for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the filter applied to the column. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property filter in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
