// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder provides operations to call the retrieveCloudPcCountByStatus method.
type VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters retrieve the Cloud PC count grouped by status.
type VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    m := &VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/retrieveCloudPcCountByStatus(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder instantiates a new VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the Cloud PC count grouped by status.
// Deprecated: This method is obsolete. Use GetAsRetrieveCloudPcCountByStatusGetResponse instead.
// returns a VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpccountbystatus?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusResponseable), nil
}
// GetAsRetrieveCloudPcCountByStatusGetResponse retrieve the Cloud PC count grouped by status.
// returns a VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpccountbystatus?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) GetAsRetrieveCloudPcCountByStatusGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusGetResponseable), nil
}
// ToGetRequestInformation retrieve the Cloud PC count grouped by status.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder when successful
func (m *VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    return NewVirtualEndpointCloudPCsRetrieveCloudPcCountByStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
