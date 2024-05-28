package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder provides operations to manage the historyDefinitions property of the microsoft.graph.accessReviewSet entity.
type AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetQueryParameters retrieve an accessReviewHistoryDefinition object by its identifier. All of the properties of the access review history definition object are returned. If the definition is 30 days or older, a 404 Not Found error is returned.
type AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetQueryParameters
}
// AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderInternal instantiates a new AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) {
    m := &AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/historyDefinitions/{accessReviewHistoryDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder instantiates a new AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property historyDefinitions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve an accessReviewHistoryDefinition object by its identifier. All of the properties of the access review history definition object are returned. If the definition is 30 days or older, a 404 Not Found error is returned.
// returns a AccessReviewHistoryDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewhistorydefinition-get?view=graph-rest-beta
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewHistoryDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable), nil
}
// Instances provides operations to manage the instances property of the microsoft.graph.accessReviewHistoryDefinition entity.
// returns a *AccessreviewsHistorydefinitionsItemInstancesRequestBuilder when successful
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) Instances()(*AccessreviewsHistorydefinitionsItemInstancesRequestBuilder) {
    return NewAccessreviewsHistorydefinitionsItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property historyDefinitions in identityGovernance
// returns a AccessReviewHistoryDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewHistoryDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property historyDefinitions for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve an accessReviewHistoryDefinition object by its identifier. All of the properties of the access review history definition object are returned. If the definition is 30 days or older, a 404 Not Found error is returned.
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property historyDefinitions in identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewHistoryDefinitionable, requestConfiguration *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder when successful
func (m *AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) {
    return NewAccessreviewsHistorydefinitionsAccessReviewHistoryDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
