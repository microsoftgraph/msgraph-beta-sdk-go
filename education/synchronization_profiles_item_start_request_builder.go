package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesItemStartRequestBuilder provides operations to call the start method.
type SynchronizationProfilesItemStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationProfilesItemStartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesItemStartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesItemStartRequestBuilderInternal instantiates a new SynchronizationProfilesItemStartRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemStartRequestBuilder) {
    m := &SynchronizationProfilesItemStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/start", pathParameters),
    }
    return m
}
// NewSynchronizationProfilesItemStartRequestBuilder instantiates a new SynchronizationProfilesItemStartRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesItemStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a SynchronizationProfilesItemStartResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-start?view=graph-rest-1.0
func (m *SynchronizationProfilesItemStartRequestBuilder) Post(ctx context.Context, requestConfiguration *SynchronizationProfilesItemStartRequestBuilderPostRequestConfiguration)(SynchronizationProfilesItemStartResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationProfilesItemStartResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationProfilesItemStartResponseable), nil
}
// PostAsStartPostResponse verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// returns a SynchronizationProfilesItemStartPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-start?view=graph-rest-1.0
func (m *SynchronizationProfilesItemStartRequestBuilder) PostAsStartPostResponse(ctx context.Context, requestConfiguration *SynchronizationProfilesItemStartRequestBuilderPostRequestConfiguration)(SynchronizationProfilesItemStartPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationProfilesItemStartPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationProfilesItemStartPostResponseable), nil
}
// ToPostRequestInformation verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// returns a *RequestInformation when successful
func (m *SynchronizationProfilesItemStartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesItemStartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SynchronizationProfilesItemStartRequestBuilder when successful
func (m *SynchronizationProfilesItemStartRequestBuilder) WithUrl(rawUrl string)(*SynchronizationProfilesItemStartRequestBuilder) {
    return NewSynchronizationProfilesItemStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
