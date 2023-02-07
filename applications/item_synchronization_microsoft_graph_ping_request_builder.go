package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationMicrosoftGraphPingRequestBuilder provides operations to call the Ping method.
type ItemSynchronizationMicrosoftGraphPingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSynchronizationMicrosoftGraphPingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationMicrosoftGraphPingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationMicrosoftGraphPingRequestBuilderInternal instantiates a new MicrosoftGraphPingRequestBuilder and sets the default values.
func NewItemSynchronizationMicrosoftGraphPingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationMicrosoftGraphPingRequestBuilder) {
    m := &ItemSynchronizationMicrosoftGraphPingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/synchronization/microsoft.graph.Ping()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSynchronizationMicrosoftGraphPingRequestBuilder instantiates a new MicrosoftGraphPingRequestBuilder and sets the default values.
func NewItemSynchronizationMicrosoftGraphPingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationMicrosoftGraphPingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationMicrosoftGraphPingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function Ping
func (m *ItemSynchronizationMicrosoftGraphPingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynchronizationMicrosoftGraphPingRequestBuilderGetRequestConfiguration)(ItemSynchronizationMicrosoftGraphPingPingResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSynchronizationMicrosoftGraphPingPingResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSynchronizationMicrosoftGraphPingPingResponseable), nil
}
// ToGetRequestInformation invoke function Ping
func (m *ItemSynchronizationMicrosoftGraphPingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationMicrosoftGraphPingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
