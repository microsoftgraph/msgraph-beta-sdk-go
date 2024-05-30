package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder provides operations to call the getSupportedCloudPcRemoteActions method.
type VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
type VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters
}
// NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    m := &VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/getSupportedCloudPcRemoteActions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder instantiates a new VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// Deprecated: This method is obsolete. Use GetAsGetSupportedCloudPcRemoteActionsGetResponse instead.
// returns a VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getsupportedcloudpcremoteactions?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable), nil
}
// GetAsGetSupportedCloudPcRemoteActionsGetResponse get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// returns a VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getsupportedcloudpcremoteactions?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) GetAsGetSupportedCloudPcRemoteActionsGetResponse(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable), nil
}
// ToGetRequestInformation get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
