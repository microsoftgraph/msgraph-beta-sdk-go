package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder provides operations to count the resources in the collection.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetQueryParameters
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderInternal instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) {
    m := &ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemBucketsItemTasksCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
