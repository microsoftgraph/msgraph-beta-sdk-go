package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesItemUploadUrlRequestBuilder provides operations to call the uploadUrl method.
type SynchronizationProfilesItemUploadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationProfilesItemUploadUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesItemUploadUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesItemUploadUrlRequestBuilderInternal instantiates a new SynchronizationProfilesItemUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemUploadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemUploadUrlRequestBuilder) {
    m := &SynchronizationProfilesItemUploadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/uploadUrl()", pathParameters),
    }
    return m
}
// NewSynchronizationProfilesItemUploadUrlRequestBuilder instantiates a new SynchronizationProfilesItemUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemUploadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemUploadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesItemUploadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function uploadUrl
// Deprecated: This method is obsolete. Use GetAsUploadUrlGetResponse instead.
// returns a SynchronizationProfilesItemUploadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SynchronizationProfilesItemUploadUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationProfilesItemUploadUrlRequestBuilderGetRequestConfiguration)(SynchronizationProfilesItemUploadUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationProfilesItemUploadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationProfilesItemUploadUrlResponseable), nil
}
// GetAsUploadUrlGetResponse invoke function uploadUrl
// Deprecated: The Education Sync Profile API is deprecated and will stop returning data on December 31, 2024. Please transition to the new IndustryData API. as of 2024-06/Deprecated:SynchronizationProfiles
// returns a SynchronizationProfilesItemUploadUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SynchronizationProfilesItemUploadUrlRequestBuilder) GetAsUploadUrlGetResponse(ctx context.Context, requestConfiguration *SynchronizationProfilesItemUploadUrlRequestBuilderGetRequestConfiguration)(SynchronizationProfilesItemUploadUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationProfilesItemUploadUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationProfilesItemUploadUrlGetResponseable), nil
}
// ToGetRequestInformation invoke function uploadUrl
// Deprecated: The Education Sync Profile API is deprecated and will stop returning data on December 31, 2024. Please transition to the new IndustryData API. as of 2024-06/Deprecated:SynchronizationProfiles
// returns a *RequestInformation when successful
func (m *SynchronizationProfilesItemUploadUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesItemUploadUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The Education Sync Profile API is deprecated and will stop returning data on December 31, 2024. Please transition to the new IndustryData API. as of 2024-06/Deprecated:SynchronizationProfiles
// returns a *SynchronizationProfilesItemUploadUrlRequestBuilder when successful
func (m *SynchronizationProfilesItemUploadUrlRequestBuilder) WithUrl(rawUrl string)(*SynchronizationProfilesItemUploadUrlRequestBuilder) {
    return NewSynchronizationProfilesItemUploadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
