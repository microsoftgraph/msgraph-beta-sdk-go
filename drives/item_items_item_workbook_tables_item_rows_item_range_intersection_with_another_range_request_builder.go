package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder provides operations to call the intersection method.
type ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, anotherRange *string)(*ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/rows/{workbookTableRow%2Did}/range()/intersection(anotherRange='{anotherRange}')", pathParameters),
    }
    if anotherRange != nil {
        m.BaseRequestBuilder.PathParameters["anotherRange"] = *anotherRange
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function intersection
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookRangeable, error) {
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
// ToGetRequestInformation invoke function intersection
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemRowsItemRangeIntersectionWithAnotherRangeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
