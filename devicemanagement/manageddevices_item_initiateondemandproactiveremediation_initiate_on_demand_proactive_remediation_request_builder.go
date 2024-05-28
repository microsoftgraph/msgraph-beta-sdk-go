package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder provides operations to call the initiateOnDemandProactiveRemediation method.
type ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal instantiates a new ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder and sets the default values.
func NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    m := &ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/initiateOnDemandProactiveRemediation", pathParameters),
    }
    return m
}
// NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder instantiates a new ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder and sets the default values.
func NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform On Demand Proactive Remediation
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) Post(ctx context.Context, body ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationPostRequestBodyable, requestConfiguration *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation perform On Demand Proactive Remediation
// returns a *RequestInformation when successful
func (m *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationPostRequestBodyable, requestConfiguration *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
