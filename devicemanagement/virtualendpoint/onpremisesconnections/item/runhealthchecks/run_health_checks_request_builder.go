package runhealthchecks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RunHealthChecksRequestBuilder provides operations to call the runHealthChecks method.
type RunHealthChecksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RunHealthChecksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunHealthChecksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunHealthChecksRequestBuilderInternal instantiates a new RunHealthChecksRequestBuilder and sets the default values.
func NewRunHealthChecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunHealthChecksRequestBuilder) {
    m := &RunHealthChecksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/microsoft.graph.runHealthChecks";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRunHealthChecksRequestBuilder instantiates a new RunHealthChecksRequestBuilder and sets the default values.
func NewRunHealthChecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunHealthChecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunHealthChecksRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
func (m *RunHealthChecksRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
func (m *RunHealthChecksRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RunHealthChecksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post run health checks on the cloudPcOnPremisesConnection object. This will trigger a new health check for this cloudPcOnPremisesConnection object and change the healthCheckStatus and healthCheckStatusDetails properties when check finished.
func (m *RunHealthChecksRequestBuilder) Post(ctx context.Context, requestConfiguration *RunHealthChecksRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
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
