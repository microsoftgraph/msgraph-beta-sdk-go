package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder provides operations to call the getSubscriptions method.
type VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetQueryParameters invoke function getSubscriptions
type VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderInternal instantiates a new GetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder) {
    m := &VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/microsoft.graph.getSubscriptions(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder instantiates a new GetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getSubscriptions
func (m *VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsResponseable), nil
}
// ToGetRequestInformation invoke function getSubscriptions
func (m *VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsMicrosoftGraphGetSubscriptionsGetSubscriptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
