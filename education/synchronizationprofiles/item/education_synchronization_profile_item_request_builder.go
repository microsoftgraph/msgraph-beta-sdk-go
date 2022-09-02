package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/profilestatus"
    i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors"
    ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/resume"
    ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/pause"
    icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/uploadurl"
    id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/reset"
    if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/start"
    i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors/item"
)

// EducationSynchronizationProfileItemRequestBuilder provides operations to manage the synchronizationProfiles property of the microsoft.graph.educationRoot entity.
type EducationSynchronizationProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationSynchronizationProfileItemRequestBuilderGetQueryParameters get synchronizationProfiles from education
type EducationSynchronizationProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationSynchronizationProfileItemRequestBuilderGetQueryParameters
}
// EducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEducationSynchronizationProfileItemRequestBuilderInternal instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewEducationSynchronizationProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSynchronizationProfileItemRequestBuilder) {
    m := &EducationSynchronizationProfileItemRequestBuilder{
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
// NewEducationSynchronizationProfileItemRequestBuilder instantiates a new EducationSynchronizationProfileItemRequestBuilder and sets the default values.
func NewEducationSynchronizationProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSynchronizationProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSynchronizationProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfileItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfileItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EducationSynchronizationProfileItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get synchronizationProfiles from education
func (m *EducationSynchronizationProfileItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EducationSynchronizationProfileItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property synchronizationProfiles in education
func (m *EducationSynchronizationProfileItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *EducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationSynchronizationProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Errors the errors property
func (m *EducationSynchronizationProfileItemRequestBuilder) Errors()(*i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.ErrorsRequestBuilder) {
    return i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.NewErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ErrorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.synchronizationProfiles.item.errors.item collection
func (m *EducationSynchronizationProfileItemRequestBuilder) ErrorsById(id string)(*i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.EducationSynchronizationErrorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationError%2Did"] = id
    }
    return i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.NewEducationSynchronizationErrorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get synchronizationProfiles from education
func (m *EducationSynchronizationProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationSynchronizationProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *EducationSynchronizationProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSynchronizationProfileable, requestConfiguration *EducationSynchronizationProfileItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Pause the pause property
func (m *EducationSynchronizationProfileItemRequestBuilder) Pause()(*ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.PauseRequestBuilder) {
    return ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProfileStatus the profileStatus property
func (m *EducationSynchronizationProfileItemRequestBuilder) ProfileStatus()(*i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.ProfileStatusRequestBuilder) {
    return i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.NewProfileStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reset the reset property
func (m *EducationSynchronizationProfileItemRequestBuilder) Reset()(*id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.ResetRequestBuilder) {
    return id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.NewResetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resume the resume property
func (m *EducationSynchronizationProfileItemRequestBuilder) Resume()(*ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.ResumeRequestBuilder) {
    return ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.NewResumeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start the start property
func (m *EducationSynchronizationProfileItemRequestBuilder) Start()(*if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.StartRequestBuilder) {
    return if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadUrl provides operations to call the uploadUrl method.
func (m *EducationSynchronizationProfileItemRequestBuilder) UploadUrl()(*icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.UploadUrlRequestBuilder) {
    return icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.NewUploadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
