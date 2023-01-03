package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Assignments()(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) AssignmentsById(id string)(*AppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleEnrollmentProfileAssignment%2Did"] = id
    }
    return NewAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// CreatePatchRequestInformation update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// Patch update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// SetPriority provides operations to call the setPriority method.
func (m *AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) SetPriority()(*AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
