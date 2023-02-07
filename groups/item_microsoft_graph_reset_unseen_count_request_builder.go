package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphResetUnseenCountRequestBuilder provides operations to call the resetUnseenCount method.
type ItemMicrosoftGraphResetUnseenCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphResetUnseenCountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphResetUnseenCountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphResetUnseenCountRequestBuilderInternal instantiates a new MicrosoftGraphResetUnseenCountRequestBuilder and sets the default values.
func NewItemMicrosoftGraphResetUnseenCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphResetUnseenCountRequestBuilder) {
    m := &ItemMicrosoftGraphResetUnseenCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/microsoft.graph.resetUnseenCount";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphResetUnseenCountRequestBuilder instantiates a new MicrosoftGraphResetUnseenCountRequestBuilder and sets the default values.
func NewItemMicrosoftGraphResetUnseenCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphResetUnseenCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphResetUnseenCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset the unseenCount of all the posts that the current user has not seen since their last visit. Supported for Microsoft 365 groups only.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/group-resetunseencount?view=graph-rest-1.0
func (m *ItemMicrosoftGraphResetUnseenCountRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphResetUnseenCountRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reset the unseenCount of all the posts that the current user has not seen since their last visit. Supported for Microsoft 365 groups only.
func (m *ItemMicrosoftGraphResetUnseenCountRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphResetUnseenCountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
