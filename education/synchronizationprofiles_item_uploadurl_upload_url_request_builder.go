package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder provides operations to call the uploadUrl method.
type SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationprofilesItemUploadurlUploadUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationprofilesItemUploadurlUploadUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilderInternal instantiates a new SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) {
    m := &SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/uploadUrl()", pathParameters),
    }
    return m
}
// NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilder instantiates a new SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve a shared access signature (SAS) for uploading source files to Azure blob storage for a specific school data synchronization profile in the tenant. The SAS token has a validity of one hour. The upload URL is provided only for the CSV data provider.
// Deprecated: This method is obsolete. Use GetAsUploadUrlGetResponse instead.
// returns a SynchronizationprofilesItemUploadurlUploadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-uploadurl?view=graph-rest-beta
func (m *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilderGetRequestConfiguration)(SynchronizationprofilesItemUploadurlUploadUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationprofilesItemUploadurlUploadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationprofilesItemUploadurlUploadUrlResponseable), nil
}
// GetAsUploadUrlGetResponse retrieve a shared access signature (SAS) for uploading source files to Azure blob storage for a specific school data synchronization profile in the tenant. The SAS token has a validity of one hour. The upload URL is provided only for the CSV data provider.
// returns a SynchronizationprofilesItemUploadurlUploadUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-uploadurl?view=graph-rest-beta
func (m *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) GetAsUploadUrlGetResponse(ctx context.Context, requestConfiguration *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilderGetRequestConfiguration)(SynchronizationprofilesItemUploadurlUploadUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSynchronizationprofilesItemUploadurlUploadUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationprofilesItemUploadurlUploadUrlGetResponseable), nil
}
// ToGetRequestInformation retrieve a shared access signature (SAS) for uploading source files to Azure blob storage for a specific school data synchronization profile in the tenant. The SAS token has a validity of one hour. The upload URL is provided only for the CSV data provider.
// returns a *RequestInformation when successful
func (m *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder when successful
func (m *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) WithUrl(rawUrl string)(*SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) {
    return NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
