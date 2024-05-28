package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder provides operations to manage the historyItems property of the microsoft.graph.userActivity entity.
type ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetQueryParameters optional. NavigationProperty/Containment; navigation property to the activity's activityHistoryItems.
type ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetQueryParameters
}
// ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activity provides operations to manage the activity property of the microsoft.graph.activityHistoryItem entity.
// returns a *ItemActivitiesItemHistoryitemsItemActivityRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) Activity()(*ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) {
    return NewItemActivitiesItemHistoryitemsItemActivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderInternal instantiates a new ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) {
    m := &ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/activities/{userActivity%2Did}/historyItems/{activityHistoryItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder instantiates a new ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property historyItems for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get optional. NavigationProperty/Containment; navigation property to the activity's activityHistoryItems.
// returns a ActivityHistoryItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityHistoryItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable), nil
}
// Patch update the navigation property historyItems in users
// returns a ActivityHistoryItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityHistoryItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable), nil
}
// ToDeleteRequestInformation delete navigation property historyItems for users
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation optional. NavigationProperty/Containment; navigation property to the activity's activityHistoryItems.
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property historyItems in users
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityHistoryItemable, requestConfiguration *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) {
    return NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
