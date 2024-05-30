package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder provides operations to manage the defaultGradingScheme property of the microsoft.graph.educationAssignmentSettings entity.
type ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetQueryParameters get defaultGradingScheme from education
type ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetQueryParameters
}
// NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderInternal instantiates a new ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) {
    m := &ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignmentSettings/defaultGradingScheme{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder instantiates a new ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get defaultGradingScheme from education
// returns a EducationGradingSchemeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationGradingSchemeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeable), nil
}
// ToGetRequestInformation get defaultGradingScheme from education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder) {
    return NewClassesItemAssignmentsettingsDefaultgradingschemeDefaultGradingSchemeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
