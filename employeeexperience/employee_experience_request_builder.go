package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmployeeExperienceRequestBuilder provides operations to manage the employeeExperience singleton.
type EmployeeExperienceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmployeeExperienceRequestBuilderGetQueryParameters get employeeExperience
type EmployeeExperienceRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmployeeExperienceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmployeeExperienceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmployeeExperienceRequestBuilderGetQueryParameters
}
// EmployeeExperienceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmployeeExperienceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Communities provides operations to manage the communities property of the microsoft.graph.employeeExperience entity.
// returns a *CommunitiesRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) Communities()(*CommunitiesRequestBuilder) {
    return NewCommunitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmployeeExperienceRequestBuilderInternal instantiates a new EmployeeExperienceRequestBuilder and sets the default values.
func NewEmployeeExperienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmployeeExperienceRequestBuilder) {
    m := &EmployeeExperienceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience{?%24select}", pathParameters),
    }
    return m
}
// NewEmployeeExperienceRequestBuilder instantiates a new EmployeeExperienceRequestBuilder and sets the default values.
func NewEmployeeExperienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmployeeExperienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmployeeExperienceRequestBuilderInternal(urlParams, requestAdapter)
}
// EngagementAsyncOperations provides operations to manage the engagementAsyncOperations property of the microsoft.graph.employeeExperience entity.
// returns a *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) EngagementAsyncOperations()(*EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) {
    return NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get employeeExperience
// returns a EmployeeExperienceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmployeeExperienceRequestBuilder) Get(ctx context.Context, requestConfiguration *EmployeeExperienceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmployeeExperienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable), nil
}
// Goals provides operations to manage the goals property of the microsoft.graph.employeeExperience entity.
// returns a *GoalsRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) Goals()(*GoalsRequestBuilder) {
    return NewGoalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningCourseActivities provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperience entity.
// returns a *LearningcourseactivitiesLearningCourseActivitiesRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) LearningCourseActivities()(*LearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    return NewLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningCourseActivitiesWithExternalcourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperience entity.
// returns a *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) LearningCourseActivitiesWithExternalcourseActivityId(externalcourseActivityId *string)(*LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalcourseActivityId)
}
// LearningProviders provides operations to manage the learningProviders property of the microsoft.graph.employeeExperience entity.
// returns a *LearningprovidersLearningProvidersRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) LearningProviders()(*LearningprovidersLearningProvidersRequestBuilder) {
    return NewLearningprovidersLearningProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update employeeExperience
// returns a EmployeeExperienceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmployeeExperienceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable, requestConfiguration *EmployeeExperienceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmployeeExperienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable), nil
}
// ToGetRequestInformation get employeeExperience
// returns a *RequestInformation when successful
func (m *EmployeeExperienceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmployeeExperienceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update employeeExperience
// returns a *RequestInformation when successful
func (m *EmployeeExperienceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmployeeExperienceable, requestConfiguration *EmployeeExperienceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmployeeExperienceRequestBuilder when successful
func (m *EmployeeExperienceRequestBuilder) WithUrl(rawUrl string)(*EmployeeExperienceRequestBuilder) {
    return NewEmployeeExperienceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
