package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder provides operations to manage the synchronizationProfiles property of the microsoft.graph.educationRoot entity.
type SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters retrieve a school data synchronization profile in the tenant based on the identifier.
type SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters
}
// SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) {
    m := &SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a school data synchronization profile in the tenant based on the identifier.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-delete?view=graph-rest-1.0
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Errors provides operations to manage the errors property of the microsoft.graph.educationSynchronizationProfile entity.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Errors()(*SynchronizationProfilesItemErrorsRequestBuilder) {
    return NewSynchronizationProfilesItemErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a school data synchronization profile in the tenant based on the identifier.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsynchronizationprofile-get?view=graph-rest-1.0
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Pause()(*SynchronizationProfilesItemPauseRequestBuilder) {
    return NewSynchronizationProfilesItemPauseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProfileStatus provides operations to manage the profileStatus property of the microsoft.graph.educationSynchronizationProfile entity.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ProfileStatus()(*SynchronizationProfilesItemProfileStatusRequestBuilder) {
    return NewSynchronizationProfilesItemProfileStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reset provides operations to call the reset method.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Reset()(*SynchronizationProfilesItemResetRequestBuilder) {
    return NewSynchronizationProfilesItemResetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resume provides operations to call the resume method.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Resume()(*SynchronizationProfilesItemResumeRequestBuilder) {
    return NewSynchronizationProfilesItemResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Start()(*SynchronizationProfilesItemStartRequestBuilder) {
    return NewSynchronizationProfilesItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a school data synchronization profile in the tenant based on the identifier.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a school data synchronization profile in the tenant based on the identifier.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property synchronizationProfiles in education
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UploadUrl provides operations to call the uploadUrl method.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) UploadUrl()(*SynchronizationProfilesItemUploadUrlRequestBuilder) {
    return NewSynchronizationProfilesItemUploadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) WithUrl(rawUrl string)(*SynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) {
    return NewSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
