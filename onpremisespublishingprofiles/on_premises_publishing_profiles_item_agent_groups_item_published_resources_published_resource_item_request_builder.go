package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder provides operations to manage the publishedResources property of the microsoft.graph.onPremisesAgentGroup entity.
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters
}
// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AgentGroups provides operations to manage the agentGroups property of the microsoft.graph.publishedResource entity.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) AgentGroups()(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.publishedResources.item.agentGroups.item collection
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) AgentGroupsById(id string)(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup%2Did1"] = id
    }
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal instantiates a new PublishedResourceItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources/{publishedResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder instantiates a new PublishedResourceItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property publishedResources for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property publishedResources in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property publishedResources for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublishedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable), nil
}
// Patch update the navigation property publishedResources in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublishedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable), nil
}
