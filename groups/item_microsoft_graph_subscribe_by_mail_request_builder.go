package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphSubscribeByMailRequestBuilder provides operations to call the subscribeByMail method.
type ItemMicrosoftGraphSubscribeByMailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphSubscribeByMailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphSubscribeByMailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphSubscribeByMailRequestBuilderInternal instantiates a new MicrosoftGraphSubscribeByMailRequestBuilder and sets the default values.
func NewItemMicrosoftGraphSubscribeByMailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphSubscribeByMailRequestBuilder) {
    m := &ItemMicrosoftGraphSubscribeByMailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/microsoft.graph.subscribeByMail";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphSubscribeByMailRequestBuilder instantiates a new MicrosoftGraphSubscribeByMailRequestBuilder and sets the default values.
func NewItemMicrosoftGraphSubscribeByMailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphSubscribeByMailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphSubscribeByMailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post calling this method will enable the current user to receive email notifications for this group, about new posts, events, and files in that group. Supported for Microsoft 365 groups only.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/group-subscribebymail?view=graph-rest-1.0
func (m *ItemMicrosoftGraphSubscribeByMailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphSubscribeByMailRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation calling this method will enable the current user to receive email notifications for this group, about new posts, events, and files in that group. Supported for Microsoft 365 groups only.
func (m *ItemMicrosoftGraphSubscribeByMailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphSubscribeByMailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
