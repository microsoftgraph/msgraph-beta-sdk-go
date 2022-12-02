package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder provides operations to manage the agents property of the microsoft.graph.onPremisesAgentGroup entity.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetQueryParameters list of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetQueryParameters
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AgentGroups provides operations to manage the agentGroups property of the microsoft.graph.onPremisesAgent entity.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) AgentGroups()(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.agents.item.agentGroups.item collection
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) AgentGroupsById(id string)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup%2Did1"] = id
    }
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderInternal instantiates a new OnPremisesAgentItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder instantiates a new OnPremisesAgentItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property agents for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property agents in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property agents for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable), nil
}
// Patch update the navigation property agents in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsOnPremisesAgentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentable), nil
}
