package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorgroupsConnectorGroupItemRequestBuilder provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemConnectorgroupsConnectorGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorgroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemConnectorgroupsConnectorGroupItemRequestBuilderGetQueryParameters list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
type ItemConnectorgroupsConnectorGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorgroupsConnectorGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupsConnectorGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorgroupsConnectorGroupItemRequestBuilderGetQueryParameters
}
// ItemConnectorgroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Applications provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorgroupsItemApplicationsRequestBuilder when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) Applications()(*ItemConnectorgroupsItemApplicationsRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplicationsWithAppId provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorgroupsItemApplicationswithappidApplicationsWithAppIdRequestBuilder when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) ApplicationsWithAppId(appId *string)(*ItemConnectorgroupsItemApplicationswithappidApplicationsWithAppIdRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationswithappidApplicationsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// ApplicationsWithUniqueName provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
// returns a *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) ApplicationsWithUniqueName(uniqueName *string)(*ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, uniqueName)
}
// NewItemConnectorgroupsConnectorGroupItemRequestBuilderInternal instantiates a new ItemConnectorgroupsConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsConnectorGroupItemRequestBuilder) {
    m := &ItemConnectorgroupsConnectorGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemConnectorgroupsConnectorGroupItemRequestBuilder instantiates a new ItemConnectorgroupsConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorgroupsConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property connectorGroups for onPremisesPublishingProfiles
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
// returns a ConnectorGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
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
// returns a *ItemConnectorgroupsItemMembersRequestBuilder when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) Members()(*ItemConnectorgroupsItemMembersRequestBuilder) {
    return NewItemConnectorgroupsItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property connectorGroups in onPremisesPublishingProfiles
// returns a ConnectorGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
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
// ToDeleteRequestInformation delete navigation property connectorGroups for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property connectorGroups in onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *ItemConnectorgroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConnectorgroupsConnectorGroupItemRequestBuilder when successful
func (m *ItemConnectorgroupsConnectorGroupItemRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorgroupsConnectorGroupItemRequestBuilder) {
    return NewItemConnectorgroupsConnectorGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
