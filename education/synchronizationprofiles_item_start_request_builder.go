package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationprofilesItemStartRequestBuilder provides operations to call the start method.
type SynchronizationprofilesItemStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationprofilesItemStartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationprofilesItemStartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationprofilesItemStartRequestBuilderInternal instantiates a new SynchronizationprofilesItemStartRequestBuilder and sets the default values.
func NewSynchronizationprofilesItemStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesItemStartRequestBuilder) {
    m := &SynchronizationprofilesItemStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/start", pathParameters),
    }
    return m
}
// NewSynchronizationprofilesItemStartRequestBuilder instantiates a new SynchronizationprofilesItemStartRequestBuilder and sets the default values.
func NewSynchronizationprofilesItemStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesItemStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationprofilesItemStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// Deprecated: This method is obsolete. Use PostAsStartPostResponse instead.
// returns a SynchronizationprofilesItemStartResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-start?view=graph-rest-beta
func (m *SynchronizationprofilesItemStartRequestBuilder) Post(ctx context.Context, requestConfiguration *SynchronizationprofilesItemStartRequestBuilderPostRequestConfiguration)(SynchronizationprofilesItemStartResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationprofilesItemStartResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationprofilesItemStartResponseable), nil
}
// PostAsStartPostResponse verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// returns a SynchronizationprofilesItemStartPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-start?view=graph-rest-beta
func (m *SynchronizationprofilesItemStartRequestBuilder) PostAsStartPostResponse(ctx context.Context, requestConfiguration *SynchronizationprofilesItemStartRequestBuilderPostRequestConfiguration)(SynchronizationprofilesItemStartPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationprofilesItemStartPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationprofilesItemStartPostResponseable), nil
}
// ToPostRequestInformation verify the files uploaded to a specific school data synchronization profile in the tenant. If the verification is successful, synchronization starts on the profile. Otherwise, the response contains errors and warnings. If the response contains errors, the synchronization won't start. If the response contains only warnings, synchronization starts.
// returns a *RequestInformation when successful
func (m *SynchronizationprofilesItemStartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SynchronizationprofilesItemStartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SynchronizationprofilesItemStartRequestBuilder when successful
func (m *SynchronizationprofilesItemStartRequestBuilder) WithUrl(rawUrl string)(*SynchronizationprofilesItemStartRequestBuilder) {
    return NewSynchronizationprofilesItemStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
