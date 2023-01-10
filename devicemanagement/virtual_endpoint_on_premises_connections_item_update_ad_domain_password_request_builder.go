package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder provides operations to call the updateAdDomainPassword method.
type VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderInternal instantiates a new UpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) {
    m := &VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/microsoft.graph.updateAdDomainPassword";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder instantiates a new UpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Active Directory domain password for a cloudPcOnPremisesConnection object. This API is supported when the type of the **cloudPcOnPremisesConnection** object is `hybridAzureADJoin`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpconpremisesconnection-updateaddomainpassword?view=graph-rest-1.0
func (m *VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) Post(ctx context.Context, body VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordPostRequestBodyable, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation update the Active Directory domain password for a cloudPcOnPremisesConnection object. This API is supported when the type of the **cloudPcOnPremisesConnection** object is `hybridAzureADJoin`.
func (m *VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordPostRequestBodyable, requestConfiguration *VirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
