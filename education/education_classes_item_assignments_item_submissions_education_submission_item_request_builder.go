package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
type EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters
}
// EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    m := &EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property submissions for education
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property submissions in education
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property submissions for education
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable), nil
}
// Outcomes provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Outcomes()(*EducationClassesItemAssignmentsItemSubmissionsItemOutcomesRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) OutcomesById(id string)(*EducationClassesItemAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome%2Did"] = id
    }
    return NewEducationClassesItemAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property submissions in education
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable), nil
}
// Reassign provides operations to call the reassign method.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Reassign()(*EducationClassesItemAssignmentsItemSubmissionsItemReassignRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Resources()(*EducationClassesItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ResourcesById(id string)(*EducationClassesItemAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewEducationClassesItemAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Return_escaped provides operations to call the return method.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Return_escaped()(*EducationClassesItemAssignmentsItemSubmissionsItemReturnRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*EducationClassesItemAssignmentsItemSubmissionsItemSetUpResourcesFolderRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submit provides operations to call the submit method.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Submit()(*EducationClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResources provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResources()(*EducationClassesItemAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*EducationClassesItemAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewEducationClassesItemAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubmit provides operations to call the unsubmit method.
func (m *EducationClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Unsubmit()(*EducationClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) {
    return NewEducationClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
