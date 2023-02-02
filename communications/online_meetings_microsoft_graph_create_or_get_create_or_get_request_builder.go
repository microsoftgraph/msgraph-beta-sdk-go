package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder provides operations to call the createOrGet method.
type OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderInternal instantiates a new CreateOrGetRequestBuilder and sets the default values.
func NewOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder) {
    m := &OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/onlineMeetings/microsoft.graph.createOrGet";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder instantiates a new CreateOrGetRequestBuilder and sets the default values.
func NewOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an onlineMeeting object with a custom specified external ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/onlinemeeting-createorget?view=graph-rest-1.0
func (m *OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder) Post(ctx context.Context, body OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBodyable, requestConfiguration *OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// ToPostRequestInformation create an onlineMeeting object with a custom specified external ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID. 
func (m *OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilder) ToPostRequestInformation(ctx context.Context, body OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBodyable, requestConfiguration *OnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
