package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*EducationUsersItemAssignmentsItemCategoriesRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.categories.item collection
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) CategoriesById(id string)(*EducationUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return NewEducationUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for education
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of assignments for the user. Nullable.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignments in education
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property assignments for education
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of assignments for the user. Nullable.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Patch update the navigation property assignments in education
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Publish provides operations to call the publish method.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Publish()(*EducationUsersItemAssignmentsItemPublishRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*EducationUsersItemAssignmentsItemResourcesRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) ResourcesById(id string)(*EducationUsersItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return NewEducationUsersItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*EducationUsersItemAssignmentsItemRubricRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) SetUpFeedbackResourcesFolder()(*EducationUsersItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*EducationUsersItemAssignmentsItemSetUpResourcesFolderRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*EducationUsersItemAssignmentsItemSubmissionsRequestBuilder) {
    return NewEducationUsersItemAssignmentsItemSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *EducationUsersItemAssignmentsEducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*EducationUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return NewEducationUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
