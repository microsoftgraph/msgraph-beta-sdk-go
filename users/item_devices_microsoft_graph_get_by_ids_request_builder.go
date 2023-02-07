package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesMicrosoftGraphGetByIdsRequestBuilder provides operations to call the getByIds method.
type ItemDevicesMicrosoftGraphGetByIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDevicesMicrosoftGraphGetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesMicrosoftGraphGetByIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDevicesMicrosoftGraphGetByIdsRequestBuilderInternal instantiates a new MicrosoftGraphGetByIdsRequestBuilder and sets the default values.
func NewItemDevicesMicrosoftGraphGetByIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesMicrosoftGraphGetByIdsRequestBuilder) {
    m := &ItemDevicesMicrosoftGraphGetByIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/microsoft.graph.getByIds";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDevicesMicrosoftGraphGetByIdsRequestBuilder instantiates a new MicrosoftGraphGetByIdsRequestBuilder and sets the default values.
func NewItemDevicesMicrosoftGraphGetByIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesMicrosoftGraphGetByIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesMicrosoftGraphGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-1.0
func (m *ItemDevicesMicrosoftGraphGetByIdsRequestBuilder) Post(ctx context.Context, body ItemDevicesMicrosoftGraphGetByIdsGetByIdsPostRequestBodyable, requestConfiguration *ItemDevicesMicrosoftGraphGetByIdsRequestBuilderPostRequestConfiguration)(ItemDevicesMicrosoftGraphGetByIdsGetByIdsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemDevicesMicrosoftGraphGetByIdsGetByIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDevicesMicrosoftGraphGetByIdsGetByIdsResponseable), nil
}
// ToPostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *ItemDevicesMicrosoftGraphGetByIdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDevicesMicrosoftGraphGetByIdsGetByIdsPostRequestBodyable, requestConfiguration *ItemDevicesMicrosoftGraphGetByIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
