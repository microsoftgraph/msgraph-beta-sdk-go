package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemPowerOffRequestBuilder provides operations to call the powerOff method.
type VirtualEndpointCloudPCsItemPowerOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemPowerOffRequestBuilderInternal instantiates a new PowerOffRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemPowerOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemPowerOffRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemPowerOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/powerOff", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemPowerOffRequestBuilder instantiates a new PowerOffRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemPowerOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemPowerOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemPowerOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it is de-allocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-poweroff?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsItemPowerOffRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it is de-allocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
func (m *VirtualEndpointCloudPCsItemPowerOffRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
