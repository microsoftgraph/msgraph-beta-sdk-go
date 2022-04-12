package item

import (
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
// OnPremisesPublishingProfileItemRequestBuilderDeleteOptions options for Delete
type OnPremisesPublishingProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnPremisesPublishingProfileItemRequestBuilderGetOptions options for Get
type OnPremisesPublishingProfileItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters get entity from onPremisesPublishingProfiles by key
type OnPremisesPublishingProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnPremisesPublishingProfileItemRequestBuilderPatchOptions options for Patch
type OnPremisesPublishingProfileItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AgentGroups the agentGroups property
func (m *OnPremisesPublishingProfileItemRequestBuilder) AgentGroups()(*i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.AgentGroupsRequestBuilder) {
    return i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.NewAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item collection
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
// Agents the agents property
func (m *OnPremisesPublishingProfileItemRequestBuilder) Agents()(*i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.AgentsRequestBuilder) {
    return i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agents.item collection
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
// ConnectorGroups the connectorGroups property
func (m *OnPremisesPublishingProfileItemRequestBuilder) ConnectorGroups()(*i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.ConnectorGroupsRequestBuilder) {
    return i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.NewConnectorGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectorGroups.item collection
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
// Connectors the connectors property
func (m *OnPremisesPublishingProfileItemRequestBuilder) Connectors()(*i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.ConnectorsRequestBuilder) {
    return i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectors.item collection
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
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreateDeleteRequestInformation(options *OnPremisesPublishingProfileItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreateGetRequestInformation(options *OnPremisesPublishingProfileItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) CreatePatchRequestInformation(options *OnPremisesPublishingProfileItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) Delete(options *OnPremisesPublishingProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileItemRequestBuilder) Get(options *OnPremisesPublishingProfileItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesPublishingProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesPublishingProfileable), nil
}
// Patch update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileItemRequestBuilder) Patch(options *OnPremisesPublishingProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PublishedResources the publishedResources property
func (m *OnPremisesPublishingProfileItemRequestBuilder) PublishedResources()(*if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.PublishedResourcesRequestBuilder) {
    return if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item collection
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
