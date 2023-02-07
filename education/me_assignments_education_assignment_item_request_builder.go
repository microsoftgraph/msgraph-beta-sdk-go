package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type MeAssignmentsEducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*MeAssignmentsItemCategoriesRequestBuilder) {
    return NewMeAssignmentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.categories.item collection
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) CategoriesById(id string)(*MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &MeAssignmentsEducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of assignments for the user. Nullable.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// MicrosoftGraphPublish provides operations to call the publish method.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphPublish()(*MeAssignmentsItemMicrosoftGraphPublishRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpFeedbackResourcesFolder()(*MeAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*MeAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property assignments in education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*MeAssignmentsItemResourcesRequestBuilder) {
    return NewMeAssignmentsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ResourcesById(id string)(*MeAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return NewMeAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*MeAssignmentsItemRubricRequestBuilder) {
    return NewMeAssignmentsItemRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*MeAssignmentsItemSubmissionsRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation list of assignments for the user. Nullable.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
