package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder provides operations to call the setPosition method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')/setPosition", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post positions the chart relative to cells on the worksheet.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chart-setposition?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation positions the chart relative to cells on the worksheet.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
