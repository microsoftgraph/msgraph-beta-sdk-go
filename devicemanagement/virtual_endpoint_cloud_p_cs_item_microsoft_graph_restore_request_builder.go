package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder provides operations to call the restore method.
type VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderInternal instantiates a new MicrosoftGraphRestoreRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/microsoft.graph.restore";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder instantiates a new MicrosoftGraphRestoreRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a specific Cloud PC. Use this API to trigger a remote action that restores a Cloud PC device to a previous state.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpc-restore?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore a specific Cloud PC. Use this API to trigger a remote action that restores a Cloud PC device to a previous state.
func (m *VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
