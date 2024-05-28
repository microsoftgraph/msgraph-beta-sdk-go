package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder provides operations to manage the synchronizationProfiles property of the microsoft.graph.educationRoot entity.
type SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters retrieve a school data synchronization profile in the tenant based on the identifier.
type SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters
}
// SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderInternal instantiates a new SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) {
    m := &SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder instantiates a new SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a school data synchronization profile in the tenant based on the identifier.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-delete?view=graph-rest-beta
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Errors provides operations to manage the errors property of the microsoft.graph.educationSynchronizationProfile entity.
// returns a *SynchronizationprofilesItemErrorsRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Errors()(*SynchronizationprofilesItemErrorsRequestBuilder) {
    return NewSynchronizationprofilesItemErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a school data synchronization profile in the tenant based on the identifier.
// returns a EducationSynchronizationProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-get?view=graph-rest-beta
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSynchronizationProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable), nil
}
// Patch update the navigation property synchronizationProfiles in education
// returns a EducationSynchronizationProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSynchronizationProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable), nil
}
// Pause provides operations to call the pause method.
// returns a *SynchronizationprofilesItemPauseRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Pause()(*SynchronizationprofilesItemPauseRequestBuilder) {
    return NewSynchronizationprofilesItemPauseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProfileStatus provides operations to manage the profileStatus property of the microsoft.graph.educationSynchronizationProfile entity.
// returns a *SynchronizationprofilesItemProfilestatusProfileStatusRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) ProfileStatus()(*SynchronizationprofilesItemProfilestatusProfileStatusRequestBuilder) {
    return NewSynchronizationprofilesItemProfilestatusProfileStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reset provides operations to call the reset method.
// returns a *SynchronizationprofilesItemResetRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Reset()(*SynchronizationprofilesItemResetRequestBuilder) {
    return NewSynchronizationprofilesItemResetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resume provides operations to call the resume method.
// returns a *SynchronizationprofilesItemResumeRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Resume()(*SynchronizationprofilesItemResumeRequestBuilder) {
    return NewSynchronizationprofilesItemResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *SynchronizationprofilesItemStartRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) Start()(*SynchronizationprofilesItemStartRequestBuilder) {
    return NewSynchronizationprofilesItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a school data synchronization profile in the tenant based on the identifier.
// returns a *RequestInformation when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a school data synchronization profile in the tenant based on the identifier.
// returns a *RequestInformation when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property synchronizationProfiles in education
// returns a *RequestInformation when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UploadUrl provides operations to call the uploadUrl method.
// returns a *SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) UploadUrl()(*SynchronizationprofilesItemUploadurlUploadUrlRequestBuilder) {
    return NewSynchronizationprofilesItemUploadurlUploadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder when successful
func (m *SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) WithUrl(rawUrl string)(*SynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder) {
    return NewSynchronizationprofilesEducationSynchronizationProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
