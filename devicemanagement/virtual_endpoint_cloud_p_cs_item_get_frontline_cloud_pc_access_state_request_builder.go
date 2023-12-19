package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3e7bfbc115fa935e142d212a729dbd28cef8afacc4f26b50744a689586890d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/getfrontlinecloudpcaccessstate"
)

// VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder provides operations to call the getFrontlineCloudPcAccessState method.
type VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal instantiates a new GetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/getFrontlineCloudPcAccessState()", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder instantiates a new GetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFrontlineCloudPcAccessState
func (m *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i3e7bfbc115fa935e142d212a729dbd28cef8afacc4f26b50744a689586890d96.GetFrontlineCloudPcAccessStateGetResponse, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, i3e7bfbc115fa935e142d212a729dbd28cef8afacc4f26b50744a689586890d96.ParseGetFrontlineCloudPcAccessStateGetResponse, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*i3e7bfbc115fa935e142d212a729dbd28cef8afacc4f26b50744a689586890d96.GetFrontlineCloudPcAccessStateGetResponse), nil
}
// ToGetRequestInformation invoke function getFrontlineCloudPcAccessState
func (m *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
