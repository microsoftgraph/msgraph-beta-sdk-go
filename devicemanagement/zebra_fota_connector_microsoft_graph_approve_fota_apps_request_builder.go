package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder provides operations to call the approveFotaApps method.
type ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderInternal instantiates a new MicrosoftGraphApproveFotaAppsRequestBuilder and sets the default values.
func NewZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder) {
    m := &ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/zebraFotaConnector/microsoft.graph.approveFotaApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder instantiates a new MicrosoftGraphApproveFotaAppsRequestBuilder and sets the default values.
func NewZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action approveFotaApps
func (m *ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorMicrosoftGraphApproveFotaAppsApproveFotaAppsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorMicrosoftGraphApproveFotaAppsApproveFotaAppsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorMicrosoftGraphApproveFotaAppsApproveFotaAppsResponseable), nil
}
// ToPostRequestInformation invoke action approveFotaApps
func (m *ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorMicrosoftGraphApproveFotaAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
