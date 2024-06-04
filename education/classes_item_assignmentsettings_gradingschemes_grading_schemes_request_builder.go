package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder provides operations to manage the gradingSchemes property of the microsoft.graph.educationAssignmentSettings entity.
type ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetQueryParameters read the properties and relationships of an educationGradingScheme object.
type ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEducationGradingSchemeId provides operations to manage the gradingSchemes property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsGradingschemesEducationGradingSchemeItemRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) ByEducationGradingSchemeId(educationGradingSchemeId string)(*ClassesItemAssignmentsettingsGradingschemesEducationGradingSchemeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if educationGradingSchemeId != "" {
        urlTplParams["educationGradingScheme%2Did"] = educationGradingSchemeId
    }
    return NewClassesItemAssignmentsettingsGradingschemesEducationGradingSchemeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderInternal instantiates a new ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) {
    m := &ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignmentSettings/gradingSchemes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder instantiates a new ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ClassesItemAssignmentsettingsGradingschemesCountRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) Count()(*ClassesItemAssignmentsettingsGradingschemesCountRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingschemesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an educationGradingScheme object.
// returns a EducationGradingSchemeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationGradingSchemeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeCollectionResponseable), nil
}
// Post create a new educationGradingScheme object.
// returns a EducationGradingSchemeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationgradingscheme-post?view=graph-rest-beta
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeable, requestConfiguration *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation read the properties and relationships of an educationGradingScheme object.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new educationGradingScheme object.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationGradingSchemeable, requestConfiguration *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingschemesGradingSchemesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
