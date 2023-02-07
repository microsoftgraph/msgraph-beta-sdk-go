package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal instantiates a new MicrosoftGraphCheckMemberObjectsRequestBuilder and sets the default values.
func NewItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    m := &ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/microsoft.graph.checkMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder instantiates a new MicrosoftGraphCheckMemberObjectsRequestBuilder and sets the default values.
func NewItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
func (m *ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemDevicesItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemDevicesItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemDevicesItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDevicesItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
func (m *ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDevicesItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemDevicesItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
