package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
// returns a *UsersItemAssignmentsItemActivateRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Activate()(*UsersItemAssignmentsItemActivateRequestBuilder) {
    return NewUsersItemAssignmentsItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemCategoriesRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*UsersItemAssignmentsItemCategoriesRequestBuilder) {
    return NewUsersItemAssignmentsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new UsersItemAssignmentsEducationAssignmentItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &UsersItemAssignmentsEducationAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUsersItemAssignmentsEducationAssignmentItemRequestBuilder instantiates a new UsersItemAssignmentsEducationAssignmentItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Deactivate provides operations to call the deactivate method.
// returns a *UsersItemAssignmentsItemDeactivateRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Deactivate()(*UsersItemAssignmentsItemDeactivateRequestBuilder) {
    return NewUsersItemAssignmentsItemDeactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property assignments for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of assignments for the user. Nullable.
// returns a EducationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// GradingCategory provides operations to manage the gradingCategory property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemGradingCategoryRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) GradingCategory()(*UsersItemAssignmentsItemGradingCategoryRequestBuilder) {
    return NewUsersItemAssignmentsItemGradingCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GradingScheme provides operations to manage the gradingScheme property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemGradingSchemeRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) GradingScheme()(*UsersItemAssignmentsItemGradingSchemeRequestBuilder) {
    return NewUsersItemAssignmentsItemGradingSchemeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignments in education
// returns a EducationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Publish provides operations to call the publish method.
// returns a *UsersItemAssignmentsItemPublishRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Publish()(*UsersItemAssignmentsItemPublishRequestBuilder) {
    return NewUsersItemAssignmentsItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemResourcesRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*UsersItemAssignmentsItemResourcesRequestBuilder) {
    return NewUsersItemAssignmentsItemResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemRubricRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*UsersItemAssignmentsItemRubricRequestBuilder) {
    return NewUsersItemAssignmentsItemRubricRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
// returns a *UsersItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) SetUpFeedbackResourcesFolder()(*UsersItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewUsersItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
// returns a *UsersItemAssignmentsItemSetUpResourcesFolderRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*UsersItemAssignmentsItemSetUpResourcesFolderRequestBuilder) {
    return NewUsersItemAssignmentsItemSetUpResourcesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
// returns a *UsersItemAssignmentsItemSubmissionsRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*UsersItemAssignmentsItemSubmissionsRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
// returns a *RequestInformation when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of assignments for the user. Nullable.
// returns a *RequestInformation when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in education
// returns a *RequestInformation when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UsersItemAssignmentsEducationAssignmentItemRequestBuilder when successful
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    return NewUsersItemAssignmentsEducationAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
