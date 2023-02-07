package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder provides operations to call the complete method.
type ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderInternal instantiates a new MicrosoftGraphCompleteRequestBuilder and sets the default values.
func NewItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder) {
    m := &ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook/tasks/{outlookTask%2Did}/microsoft.graph.complete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder instantiates a new MicrosoftGraphCompleteRequestBuilder and sets the default values.
func NewItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action complete
func (m *ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderPostRequestConfiguration)(ItemOutlookTasksItemMicrosoftGraphCompleteCompleteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemOutlookTasksItemMicrosoftGraphCompleteCompleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookTasksItemMicrosoftGraphCompleteCompleteResponseable), nil
}
// ToPostRequestInformation invoke action complete
func (m *ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookTasksItemMicrosoftGraphCompleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
