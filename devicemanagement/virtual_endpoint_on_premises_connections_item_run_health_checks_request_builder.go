package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder provides operations to call the runHealthChecks method.
type VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderInternal instantiates a new RunHealthChecksRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) {
    m := &VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/runHealthChecks", pathParameters),
    }
    return m
}
// NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder instantiates a new RunHealthChecksRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-runhealthcheck?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished. This API is available in the following national cloud deployments.
func (m *VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
