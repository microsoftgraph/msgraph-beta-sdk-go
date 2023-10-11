package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemPowerOnRequestBuilder provides operations to call the powerOn method.
type VirtualEndpointCloudPCsItemPowerOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemPowerOnRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemPowerOnRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemPowerOnRequestBuilderInternal instantiates a new PowerOnRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemPowerOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemPowerOnRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemPowerOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/powerOn", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemPowerOnRequestBuilder instantiates a new PowerOnRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemPowerOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemPowerOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemPowerOnRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power on a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered on, it is allocated to a user, and licenses are assigned immediately. Only IT admin users can perform this action.  This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-poweron?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsItemPowerOnRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemPowerOnRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation power on a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered on, it is allocated to a user, and licenses are assigned immediately. Only IT admin users can perform this action.  This API is available in the following national cloud deployments.
func (m *VirtualEndpointCloudPCsItemPowerOnRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemPowerOnRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointCloudPCsItemPowerOnRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemPowerOnRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemPowerOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
