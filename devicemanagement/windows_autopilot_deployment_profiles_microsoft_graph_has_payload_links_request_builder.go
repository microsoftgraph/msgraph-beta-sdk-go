package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderInternal instantiates a new MicrosoftGraphHasPayloadLinksRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder) {
    m := &WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/microsoft.graph.hasPayloadLinks";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder instantiates a new MicrosoftGraphHasPayloadLinksRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
func (m *WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder) Post(ctx context.Context, body WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyable, requestConfiguration *WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderPostRequestConfiguration)(WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateWindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
func (m *WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyable, requestConfiguration *WindowsAutopilotDeploymentProfilesMicrosoftGraphHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
