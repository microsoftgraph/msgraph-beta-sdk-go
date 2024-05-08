package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LearningProvidersLearningProviderItemRequestBuilder provides operations to manage the learningProviders property of the microsoft.graph.employeeExperience entity.
type LearningProvidersLearningProviderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningProvidersLearningProviderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningProvidersLearningProviderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LearningProvidersLearningProviderItemRequestBuilderGetQueryParameters read the properties and relationships of a learningProvider object.
type LearningProvidersLearningProviderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LearningProvidersLearningProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningProvidersLearningProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningProvidersLearningProviderItemRequestBuilderGetQueryParameters
}
// LearningProvidersLearningProviderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningProvidersLearningProviderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLearningProvidersLearningProviderItemRequestBuilderInternal instantiates a new LearningProvidersLearningProviderItemRequestBuilder and sets the default values.
func NewLearningProvidersLearningProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningProvidersLearningProviderItemRequestBuilder) {
    m := &LearningProvidersLearningProviderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLearningProvidersLearningProviderItemRequestBuilder instantiates a new LearningProvidersLearningProviderItemRequestBuilder and sets the default values.
func NewLearningProvidersLearningProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningProvidersLearningProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningProvidersLearningProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a learningProvider resource and remove its registration in Viva Learning for the tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/employeeexperience-delete-learningproviders?view=graph-rest-beta
func (m *LearningProvidersLearningProviderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a learningProvider object.
// returns a LearningProviderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-get?view=graph-rest-beta
func (m *LearningProvidersLearningProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLearningProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable), nil
}
// LearningContents provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
// returns a *LearningProvidersItemLearningContentsRequestBuilder when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) LearningContents()(*LearningProvidersItemLearningContentsRequestBuilder) {
    return NewLearningProvidersItemLearningContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningContentsWithExternalId provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
// returns a *LearningProvidersItemLearningContentsWithExternalIdRequestBuilder when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) LearningContentsWithExternalId(externalId *string)(*LearningProvidersItemLearningContentsWithExternalIdRequestBuilder) {
    return NewLearningProvidersItemLearningContentsWithExternalIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalId)
}
// LearningCourseActivities provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
// returns a *LearningProvidersItemLearningCourseActivitiesRequestBuilder when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) LearningCourseActivities()(*LearningProvidersItemLearningCourseActivitiesRequestBuilder) {
    return NewLearningProvidersItemLearningCourseActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningCourseActivitiesWithExternalcourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
// returns a *LearningProvidersItemLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) LearningCourseActivitiesWithExternalcourseActivityId(externalcourseActivityId *string)(*LearningProvidersItemLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewLearningProvidersItemLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalcourseActivityId)
}
// Patch update the properties of a learningProvider object.
// returns a LearningProviderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-update?view=graph-rest-beta
func (m *LearningProvidersLearningProviderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLearningProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable), nil
}
// ToDeleteRequestInformation delete a learningProvider resource and remove its registration in Viva Learning for the tenant.
// returns a *RequestInformation when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a learningProvider object.
// returns a *RequestInformation when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a learningProvider object.
// returns a *RequestInformation when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningProviderable, requestConfiguration *LearningProvidersLearningProviderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningProvidersLearningProviderItemRequestBuilder when successful
func (m *LearningProvidersLearningProviderItemRequestBuilder) WithUrl(rawUrl string)(*LearningProvidersLearningProviderItemRequestBuilder) {
    return NewLearningProvidersLearningProviderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
