package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RemoteAssistancePartnersItemBeginOnboardingRequestBuilder provides operations to call the beginOnboarding method.
type RemoteAssistancePartnersItemBeginOnboardingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) {
    m := &RemoteAssistancePartnersItemBeginOnboardingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}/microsoft.graph.beginOnboarding";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilder instantiates a new BeginOnboardingRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteAssistancePartnersItemBeginOnboardingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post a request to start onboarding.  Must be coupled with the appropriate TeamViewer account information
func (m *RemoteAssistancePartnersItemBeginOnboardingRequestBuilder) Post(ctx context.Context, requestConfiguration *RemoteAssistancePartnersItemBeginOnboardingRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
