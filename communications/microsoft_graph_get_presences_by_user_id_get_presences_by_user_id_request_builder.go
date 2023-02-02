package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder provides operations to call the getPresencesByUserId method.
type MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderInternal instantiates a new GetPresencesByUserIdRequestBuilder and sets the default values.
func NewMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder) {
    m := &MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/microsoft.graph.getPresencesByUserId";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder instantiates a new GetPresencesByUserIdRequestBuilder and sets the default values.
func NewMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the presence information for multiple users.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudcommunications-getpresencesbyuserid?view=graph-rest-1.0
func (m *MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder) Post(ctx context.Context, body MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdPostRequestBodyable, requestConfiguration *MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseable), nil
}
// ToPostRequestInformation get the presence information for multiple users.
func (m *MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdPostRequestBodyable, requestConfiguration *MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
