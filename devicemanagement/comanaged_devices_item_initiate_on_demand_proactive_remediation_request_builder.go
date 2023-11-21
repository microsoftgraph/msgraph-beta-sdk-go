package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder provides operations to call the initiateOnDemandProactiveRemediation method.
type ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal instantiates a new InitiateOnDemandProactiveRemediationRequestBuilder and sets the default values.
func NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    m := &ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/initiateOnDemandProactiveRemediation", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder instantiates a new InitiateOnDemandProactiveRemediationRequestBuilder and sets the default values.
func NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform On Demand Proactive Remediation
func (m *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) Post(ctx context.Context, body ComanagedDevicesItemInitiateOnDemandProactiveRemediationPostRequestBodyable, requestConfiguration *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation perform On Demand Proactive Remediation
func (m *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesItemInitiateOnDemandProactiveRemediationPostRequestBodyable, requestConfiguration *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
