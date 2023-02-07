package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder provides operations to call the uploadUrl method.
type SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderInternal instantiates a new MicrosoftGraphUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder) {
    m := &SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}/microsoft.graph.uploadUrl()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder instantiates a new MicrosoftGraphUploadUrlRequestBuilder and sets the default values.
func NewSynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function uploadUrl
func (m *SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderGetRequestConfiguration)(SynchronizationProfilesItemMicrosoftGraphUploadUrlUploadUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateSynchronizationProfilesItemMicrosoftGraphUploadUrlUploadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SynchronizationProfilesItemMicrosoftGraphUploadUrlUploadUrlResponseable), nil
}
// ToGetRequestInformation invoke function uploadUrl
func (m *SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesItemMicrosoftGraphUploadUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
