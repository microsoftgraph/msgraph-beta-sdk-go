package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder provides operations to call the reprovision method.
type VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal instantiates a new ReprovisionRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/microsoft.graph.reprovision";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder instantiates a new ReprovisionRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprovision a specific Cloud PC.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpc-reprovision?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation reprovision a specific Cloud PC.
func (m *VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
