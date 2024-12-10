package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder provides operations to call the retrieveSnapshots method.
type VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters list all cloudPcSnapshot resources for a Cloud PC.
type VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/retrieveSnapshots(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder instantiates a new VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all cloudPcSnapshot resources for a Cloud PC.
// Deprecated: This method is obsolete. Use GetAsRetrieveSnapshotsGetResponse instead.
// returns a VirtualEndpointCloudPCsItemRetrieveSnapshotsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievesnapshots?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemRetrieveSnapshotsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemRetrieveSnapshotsResponseable), nil
}
// GetAsRetrieveSnapshotsGetResponse list all cloudPcSnapshot resources for a Cloud PC.
// returns a VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievesnapshots?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) GetAsRetrieveSnapshotsGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemRetrieveSnapshotsGetResponseable), nil
}
// ToGetRequestInformation list all cloudPcSnapshot resources for a Cloud PC.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRetrieveSnapshotsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
