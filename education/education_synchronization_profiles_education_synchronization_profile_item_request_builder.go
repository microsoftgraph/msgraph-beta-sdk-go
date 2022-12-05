package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder provides operations to manage the synchronizationProfiles property of the microsoft.graph.educationRoot entity.
type EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters get synchronizationProfiles from education
type EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetQueryParameters
}
// EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewEducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) {
    m := &EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewEducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get synchronizationProfiles from education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property synchronizationProfiles in education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Errors provides operations to manage the errors property of the microsoft.graph.educationSynchronizationProfile entity.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Errors()(*EducationSynchronizationProfilesItemErrorsRequestBuilder) {
    return NewEducationSynchronizationProfilesItemErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ErrorsById provides operations to manage the errors property of the microsoft.graph.educationSynchronizationProfile entity.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ErrorsById(id string)(*EducationSynchronizationProfilesItemErrorsEducationSynchronizationErrorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationError%2Did"] = id
    }
    return NewEducationSynchronizationProfilesItemErrorsEducationSynchronizationErrorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get synchronizationProfiles from education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSynchronizationProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable), nil
}
// Patch update the navigation property synchronizationProfiles in education
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSynchronizationProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable), nil
}
// Pause provides operations to call the pause method.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Pause()(*EducationSynchronizationProfilesItemPauseRequestBuilder) {
    return NewEducationSynchronizationProfilesItemPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProfileStatus provides operations to manage the profileStatus property of the microsoft.graph.educationSynchronizationProfile entity.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) ProfileStatus()(*EducationSynchronizationProfilesItemProfileStatusRequestBuilder) {
    return NewEducationSynchronizationProfilesItemProfileStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reset provides operations to call the reset method.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Reset()(*EducationSynchronizationProfilesItemResetRequestBuilder) {
    return NewEducationSynchronizationProfilesItemResetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resume provides operations to call the resume method.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Resume()(*EducationSynchronizationProfilesItemResumeRequestBuilder) {
    return NewEducationSynchronizationProfilesItemResumeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start provides operations to call the start method.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) Start()(*EducationSynchronizationProfilesItemStartRequestBuilder) {
    return NewEducationSynchronizationProfilesItemStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadUrl provides operations to call the uploadUrl method.
func (m *EducationSynchronizationProfilesEducationSynchronizationProfileItemRequestBuilder) UploadUrl()(*EducationSynchronizationProfilesItemUploadUrlRequestBuilder) {
    return NewEducationSynchronizationProfilesItemUploadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
