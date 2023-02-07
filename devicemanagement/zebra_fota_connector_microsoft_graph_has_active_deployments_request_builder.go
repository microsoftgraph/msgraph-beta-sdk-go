package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder provides operations to call the hasActiveDeployments method.
type ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderInternal instantiates a new MicrosoftGraphHasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder) {
    m := &ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/zebraFotaConnector/microsoft.graph.hasActiveDeployments";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder instantiates a new MicrosoftGraphHasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasActiveDeployments
func (m *ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsHasActiveDeploymentsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsHasActiveDeploymentsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsHasActiveDeploymentsResponseable), nil
}
// ToPostRequestInformation invoke action hasActiveDeployments
func (m *ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorMicrosoftGraphHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
