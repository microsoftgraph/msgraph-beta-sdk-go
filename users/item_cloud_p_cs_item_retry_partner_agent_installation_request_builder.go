package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder provides operations to call the retryPartnerAgentInstallation method.
type ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal instantiates a new RetryPartnerAgentInstallationRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    m := &ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/microsoft.graph.retryPartnerAgentInstallation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder instantiates a new RetryPartnerAgentInstallationRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retry installation for the partner agents which failed to install on the Cloud PC. Service side will check which agent installation failed firstly and retry.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpc-retrypartneragentinstallation?view=graph-rest-1.0
func (m *ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation retry installation for the partner agents which failed to install on the Cloud PC. Service side will check which agent installation failed firstly and retry.
func (m *ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
