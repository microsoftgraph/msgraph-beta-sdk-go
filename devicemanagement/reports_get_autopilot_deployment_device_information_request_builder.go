package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder provides operations to call the getAutopilotDeploymentDeviceInformation method.
type ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetAutopilotDeploymentDeviceInformationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetAutopilotDeploymentDeviceInformationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetAutopilotDeploymentDeviceInformationRequestBuilderInternal instantiates a new GetAutopilotDeploymentDeviceInformationRequestBuilder and sets the default values.
func NewReportsGetAutopilotDeploymentDeviceInformationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder) {
    m := &ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getAutopilotDeploymentDeviceInformation", pathParameters),
    }
    return m
}
// NewReportsGetAutopilotDeploymentDeviceInformationRequestBuilder instantiates a new GetAutopilotDeploymentDeviceInformationRequestBuilder and sets the default values.
func NewReportsGetAutopilotDeploymentDeviceInformationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetAutopilotDeploymentDeviceInformationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getAutopilotDeploymentDeviceInformation
func (m *ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder) Post(ctx context.Context, body ReportsGetAutopilotDeploymentDeviceInformationPostRequestBodyable, requestConfiguration *ReportsGetAutopilotDeploymentDeviceInformationRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getAutopilotDeploymentDeviceInformation
func (m *ReportsGetAutopilotDeploymentDeviceInformationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetAutopilotDeploymentDeviceInformationPostRequestBodyable, requestConfiguration *ReportsGetAutopilotDeploymentDeviceInformationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
