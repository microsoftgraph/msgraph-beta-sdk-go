package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder provides operations to count the resources in the collection.
type DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetQueryParameters get the number of the resource
type DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetQueryParameters
}
// NewDevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewDevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder) {
    m := &DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/transitiveMemberOf/microsoft.graph.application/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewDevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicesItemTransitiveMemberOfMicrosoftGraphApplicationCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
