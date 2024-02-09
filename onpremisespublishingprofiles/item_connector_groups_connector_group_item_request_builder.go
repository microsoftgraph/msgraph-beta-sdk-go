package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorGroupsConnectorGroupItemRequestBuilder provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemConnectorGroupsConnectorGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters retrieve the properties of a connectorGroup.
type ItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters
}
// ItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Applications provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorGroupsItemApplicationsRequestBuilder when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) Applications()(*ItemConnectorGroupsItemApplicationsRequestBuilder) {
    return NewItemConnectorGroupsItemApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplicationsWithAppId provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorGroupsItemApplicationsWithAppIdRequestBuilder when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) ApplicationsWithAppId(appId *string)(*ItemConnectorGroupsItemApplicationsWithAppIdRequestBuilder) {
    return NewItemConnectorGroupsItemApplicationsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// ApplicationsWithUniqueName provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorGroupsItemApplicationsWithUniqueNameRequestBuilder when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) ApplicationsWithUniqueName(uniqueName *string)(*ItemConnectorGroupsItemApplicationsWithUniqueNameRequestBuilder) {
    return NewItemConnectorGroupsItemApplicationsWithUniqueNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, uniqueName)
}
// NewItemConnectorGroupsConnectorGroupItemRequestBuilderInternal instantiates a new ItemConnectorGroupsConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    m := &ItemConnectorGroupsConnectorGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemConnectorGroupsConnectorGroupItemRequestBuilder instantiates a new ItemConnectorGroupsConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorGroupsConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a connectorGroup. All connectors and applications must be removed from the connector group before a connector group can be deleted.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectorgroup-delete?view=graph-rest-1.0
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties of a connectorGroup.
// returns a ConnectorGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectorgroup-get?view=graph-rest-1.0
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorGroupsItemMembersRequestBuilder when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) Members()(*ItemConnectorGroupsItemMembersRequestBuilder) {
    return NewItemConnectorGroupsItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a connectorGroup object.
// returns a ConnectorGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectorgroup-update?view=graph-rest-1.0
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable), nil
}
// ToDeleteRequestInformation delete a connectorGroup. All connectors and applications must be removed from the connector group before a connector group can be deleted.
// returns a *RequestInformation when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties of a connectorGroup.
// returns a *RequestInformation when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a connectorGroup object.
// returns a *RequestInformation when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *ItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemConnectorGroupsConnectorGroupItemRequestBuilder when successful
func (m *ItemConnectorGroupsConnectorGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    return NewItemConnectorGroupsConnectorGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
