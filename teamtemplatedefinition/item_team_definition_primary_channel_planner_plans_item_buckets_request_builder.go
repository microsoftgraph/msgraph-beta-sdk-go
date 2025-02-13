package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetQueryParameters collection of buckets in the plan. Read-only. Nullable.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetQueryParameters struct {
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
// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerBucketId provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
// returns a *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsPlannerBucketItemRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) ByPlannerBucketId(plannerBucketId string)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsPlannerBucketItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerBucketId != "" {
        urlTplParams["plannerBucket%2Did"] = plannerBucketId
    }
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsPlannerBucketItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderInternal instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) {
    m := &ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/planner/plans/{plannerPlan%2Did}/buckets{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsCountRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) Count()(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsCountRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsDeltaRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) Delta()(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsDeltaRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of buckets in the plan. Read-only. Nullable.
// returns a PlannerBucketCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerBucketCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketCollectionResponseable), nil
}
// Post create new navigation property to buckets for teamTemplateDefinition
// returns a PlannerBucketable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketable, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerBucketFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketable), nil
}
// ToGetRequestInformation collection of buckets in the plan. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to buckets for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerBucketable, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
