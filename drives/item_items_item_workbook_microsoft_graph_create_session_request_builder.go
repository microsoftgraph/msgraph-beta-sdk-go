package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder provides operations to call the createSession method.
type ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderInternal instantiates a new MicrosoftGraphCreateSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder) {
    m := &ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/microsoft.graph.createSession";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder instantiates a new MicrosoftGraphCreateSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new workbook session.  Excel APIs can be called in one of two modes:  To represent the session in the API, use the `workbook-session-id: {session-id}` header.  In some cases, creating a new session requires an indeterminate time to complete. Microsoft Graph also provides a long running operations pattern. This pattern provides a way to poll for creation status updates, without waiting for the creation to complete. The following are the steps:
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/workbook-createsession?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookMicrosoftGraphCreateSessionCreateSessionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookSessionInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookSessionInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookSessionInfoable), nil
}
// ToPostRequestInformation create a new workbook session.  Excel APIs can be called in one of two modes:  To represent the session in the API, use the `workbook-session-id: {session-id}` header.  In some cases, creating a new session requires an indeterminate time to complete. Microsoft Graph also provides a long running operations pattern. This pattern provides a way to poll for creation status updates, without waiting for the creation to complete. The following are the steps:
func (m *ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookMicrosoftGraphCreateSessionCreateSessionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
