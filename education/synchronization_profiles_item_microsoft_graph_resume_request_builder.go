package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder provides operations to call the resume method.
type SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderInternal instantiates a new MicrosoftGraphResumeRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder) {
    m := &SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/microsoft.graph.resume";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder instantiates a new MicrosoftGraphResumeRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resume the sync of a specific school data synchronization profile in the tenant.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsynchronizationprofile-resume?view=graph-rest-1.0
func (m *SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder) Post(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation resume the sync of a specific school data synchronization profile in the tenant.
func (m *SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
