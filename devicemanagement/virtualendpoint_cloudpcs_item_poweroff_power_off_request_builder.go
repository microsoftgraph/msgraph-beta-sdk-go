package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder provides operations to call the powerOff method.
type VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) {
    m := &VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/powerOff", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder instantiates a new VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it's deallocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-poweroff?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it's deallocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) {
    return NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
