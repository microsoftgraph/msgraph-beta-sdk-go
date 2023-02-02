package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder provides operations to call the grant method.
type ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderInternal instantiates a new GrantRequestBuilder and sets the default values.
func NewItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder) {
    m := &ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/permissions/{permission%2Did}/microsoft.graph.grant";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder instantiates a new GrantRequestBuilder and sets the default values.
func NewItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post grant users access to a link represented by a [permission][].
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permission-grant?view=graph-rest-1.0
func (m *ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder) Post(ctx context.Context, body ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantPostRequestBodyable, requestConfiguration *ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderPostRequestConfiguration)(ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemsItemPermissionsItemMicrosoftGraphGrantGrantResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantResponseable), nil
}
// ToPostRequestInformation grant users access to a link represented by a [permission][].
func (m *ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantPostRequestBodyable, requestConfiguration *ItemItemsItemPermissionsItemMicrosoftGraphGrantGrantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
