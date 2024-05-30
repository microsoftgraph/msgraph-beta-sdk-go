package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder provides operations to call the retryPartnerAgentInstallation method.
type VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) {
    m := &VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/retryPartnerAgentInstallation", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder instantiates a new VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retry installation for the partner agents that failed to install on the Cloud PC. Service side checks which agent installation failed firstly and retry.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrypartneragentinstallation?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation retry installation for the partner agents that failed to install on the Cloud PC. Service side checks which agent installation failed firstly and retry.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
