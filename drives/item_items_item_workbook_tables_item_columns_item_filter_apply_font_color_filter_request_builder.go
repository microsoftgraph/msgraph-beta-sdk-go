package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder provides operations to call the applyFontColorFilter method.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyFontColorFilter", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyFontColorFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyFontColorFilter
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
