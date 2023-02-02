package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder provides operations to call the resume method.
type SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderInternal instantiates a new ResumeRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder) {
    m := &SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder{
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
// NewSynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder instantiates a new ResumeRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resume the sync of a specific school data synchronization profile in the tenant.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsynchronizationprofile-resume?view=graph-rest-1.0
func (m *SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder) Post(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderPostRequestConfiguration)(error) {
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
func (m *SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphResumeResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
