package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder provides operations to call the hypGeom_Dist method.
type ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderInternal instantiates a new MicrosoftGraphHypGeom_DistRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/microsoft.graph.hypGeom_Dist";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder instantiates a new MicrosoftGraphHypGeom_DistRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hypGeom_Dist
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action hypGeom_Dist
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
