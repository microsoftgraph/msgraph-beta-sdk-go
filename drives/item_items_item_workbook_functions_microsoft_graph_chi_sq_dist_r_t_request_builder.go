package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder provides operations to call the chiSq_Dist_RT method.
type ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderInternal instantiates a new MicrosoftGraphChiSq_Dist_RTRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/microsoft.graph.chiSq_Dist_RT";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder instantiates a new MicrosoftGraphChiSq_Dist_RTRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action chiSq_Dist_RT
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTChiSq_Dist_RTPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action chiSq_Dist_RT
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTChiSq_Dist_RTPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
