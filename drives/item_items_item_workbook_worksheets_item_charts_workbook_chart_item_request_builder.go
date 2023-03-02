package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters returns collection of charts that are part of the worksheet. Read-only.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Axes()(*ItemItemsItemWorkbookWorksheetsItemChartsItemAxesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) DataLabels()(*ItemItemsItemWorkbookWorksheetsItemChartsItemDataLabelsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property charts for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Format()(*ItemItemsItemWorkbookWorksheetsItemChartsItemFormatRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get returns collection of charts that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable), nil
}
// Image provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Image()(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ImageWithWidth provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ImageWithWidth(width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width)
}
// ImageWithWidthWithHeight provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ImageWithWidthWithHeight(height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, height, width)
}
// ImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, fittingMode, height, width)
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Legend()(*ItemItemsItemWorkbookWorksheetsItemChartsItemLegendRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property charts in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable), nil
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Series()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SeriesById provides operations to manage the series property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) SeriesById(id string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesWorkbookChartSeriesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesWorkbookChartSeriesItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SetData provides operations to call the setData method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) SetData()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSetDataRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SetPosition provides operations to call the setPosition method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) SetPosition()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Title()(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property charts for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property charts in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
