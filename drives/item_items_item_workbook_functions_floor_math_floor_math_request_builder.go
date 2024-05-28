package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder provides operations to call the floor_Math method.
type ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/floor_Math", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action floor_Math
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action floor_Math
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsFloor_mathFloor_MathRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
