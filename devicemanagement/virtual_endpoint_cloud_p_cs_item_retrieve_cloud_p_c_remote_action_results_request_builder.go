package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder provides operations to call the retrieveCloudPCRemoteActionResults method.
type VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
type VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/retrieveCloudPCRemoteActionResults(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder instantiates a new VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// Deprecated: This method is obsolete. Use GetAsRetrieveCloudPCRemoteActionResultsGetResponse instead.
// returns a VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpcremoteactionresults?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable), nil
}
// GetAsRetrieveCloudPCRemoteActionResultsGetResponse retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// returns a VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpcremoteactionresults?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) GetAsRetrieveCloudPCRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
