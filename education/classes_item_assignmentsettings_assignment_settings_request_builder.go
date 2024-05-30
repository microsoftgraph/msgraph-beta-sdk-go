package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder provides operations to manage the assignmentSettings property of the microsoft.graph.educationClass entity.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters read the properties and relationships of an educationAssignmentSettings object. Only teachers can perform this operation.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal instantiates a new ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    m := &ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignmentSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder instantiates a new ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultGradingScheme provides operations to manage the defaultGradingScheme property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) DefaultGradingScheme()(*ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) {
    return NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property assignmentSettings for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an educationAssignmentSettings object. Only teachers can perform this operation.
// returns a EducationAssignmentSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignmentsettings-get?view=graph-rest-beta
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable), nil
}
// GradingCategories provides operations to manage the gradingCategories property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) GradingCategories()(*ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GradingSchemes provides operations to manage the gradingSchemes property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) GradingSchemes()(*ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an educationAssignmentSettings object. Only teachers can update these settings.
// returns a EducationAssignmentSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignmentsettings-update?view=graph-rest-beta
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property assignmentSettings for education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an educationAssignmentSettings object. Only teachers can perform this operation.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an educationAssignmentSettings object. Only teachers can update these settings.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentSettingsable, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    return NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
