package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfileItemRequestBuilder provides operations to manage the collection of onPremisesPublishingProfile entities.
type OnPremisesPublishingProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters get entity from onPremisesPublishingProfiles by key
type OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnPremisesPublishingProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters
}
// OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AgentGroups provides operations to manage the agentGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentGroups()(*ItemAgentGroupsRequestBuilder) {
    return NewItemAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById provides operations to manage the agentGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentGroupsById(id string)(*ItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup%2Did"] = id
    }
    return NewItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Agents provides operations to manage the agents property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) Agents()(*ItemAgentsRequestBuilder) {
    return NewItemAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById provides operations to manage the agents property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentsById(id string)(*ItemAgentsOnPremisesAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent%2Did"] = id
    }
    return NewItemAgentsOnPremisesAgentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConnectorGroups provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorGroups()(*ItemConnectorGroupsRequestBuilder) {
    return NewItemConnectorGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorGroupsById provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorGroupsById(id string)(*ItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectorGroup%2Did"] = id
    }
    return NewItemConnectorGroupsConnectorGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) Connectors()(*ItemConnectorsRequestBuilder) {
    return NewItemConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorsById(id string)(*ItemConnectorsConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector%2Did"] = id
    }
    return NewItemConnectorsConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesPublishingProfileItemRequestBuilderInternal instantiates a new OnPremisesPublishingProfileItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfileItemRequestBuilder) {
    m := &OnPremisesPublishingProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfileItemRequestBuilder instantiates a new OnPremisesPublishingProfileItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesPublishingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable), nil
}
// Patch update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable, requestConfiguration *OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesPublishingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable), nil
}
// PublishedResources provides operations to manage the publishedResources property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) PublishedResources()(*ItemPublishedResourcesRequestBuilder) {
    return NewItemPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById provides operations to manage the publishedResources property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) PublishedResourcesById(id string)(*ItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource%2Did"] = id
    }
    return NewItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
