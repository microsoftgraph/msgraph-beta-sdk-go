package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder provides operations to call the runHealthChecks method.
type VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderInternal instantiates a new MicrosoftGraphRunHealthChecksRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder) {
    m := &VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/microsoft.graph.runHealthChecks";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder instantiates a new MicrosoftGraphRunHealthChecksRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpconpremisesconnection-runhealthcheck?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
func (m *VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemMicrosoftGraphRunHealthChecksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
