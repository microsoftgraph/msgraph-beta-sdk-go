package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder provides operations to call the delta method.
type DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetQueryParameters get newly created, updated, or deleted buckets in a Planner plan without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
type DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetQueryParameters struct {
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
// DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetQueryParameters
}
// NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderInternal instantiates a new DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}/buckets/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder instantiates a new DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get newly created, updated, or deleted buckets in a Planner plan without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerbucket-delta?view=graph-rest-beta
func (m *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetRequestConfiguration)(DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaResponseable), nil
}
// GetAsDeltaGetResponse get newly created, updated, or deleted buckets in a Planner plan without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerbucket-delta?view=graph-rest-beta
func (m *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetRequestConfiguration)(DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaGetResponseable), nil
}
// ToGetRequestInformation get newly created, updated, or deleted buckets in a Planner plan without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemBucketsDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
