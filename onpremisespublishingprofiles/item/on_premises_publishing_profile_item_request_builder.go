package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups"
    i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents"
    i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups"
    i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors"
    if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources"
    i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item"
    i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item"
    i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item"
    i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item"
    ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors/item"
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
    Headers map[string]string
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters
}
// OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AgentGroups provides operations to manage the agentGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentGroups()(*i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.AgentGroupsRequestBuilder) {
    return i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.NewAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById provides operations to manage the agentGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentGroupsById(id string)(*i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.OnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup%2Did"] = id
    }
    return i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.NewOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Agents provides operations to manage the agents property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) Agents()(*i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.AgentsRequestBuilder) {
    return i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById provides operations to manage the agents property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentsById(id string)(*i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.OnPremisesAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent%2Did"] = id
    }
    return i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.NewOnPremisesAgentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConnectorGroups provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorGroups()(*i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.ConnectorGroupsRequestBuilder) {
    return i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.NewConnectorGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorGroupsById provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorGroupsById(id string)(*i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.ConnectorGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectorGroup%2Did"] = id
    }
    return i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.NewConnectorGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) Connectors()(*i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.ConnectorsRequestBuilder) {
    return i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorsById(id string)(*ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.ConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector%2Did"] = id
    }
    return ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.NewConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
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
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
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
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
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
func (m *OnPremisesPublishingProfileItemRequestBuilder) PublishedResources()(*if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.PublishedResourcesRequestBuilder) {
    return if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById provides operations to manage the publishedResources property of the microsoft.graph.onPremisesPublishingProfile entity.
func (m *OnPremisesPublishingProfileItemRequestBuilder) PublishedResourcesById(id string)(*i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.PublishedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource%2Did"] = id
    }
    return i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.NewPublishedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
