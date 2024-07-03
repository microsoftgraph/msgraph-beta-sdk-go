package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder provides operations to manage the publishedResources property of the microsoft.graph.onPremisesAgentGroup entity.
type ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
type ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetQueryParameters
}
// ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AgentGroups provides operations to manage the agentGroups property of the microsoft.graph.publishedResource entity.
// returns a *ItemAgentGroupsItemPublishedResourcesItemAgentGroupsRequestBuilder when successful
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) AgentGroups()(*ItemAgentGroupsItemPublishedResourcesItemAgentGroupsRequestBuilder) {
    return NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal instantiates a new ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    m := &ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources/{publishedResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder instantiates a new ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property publishedResources for onPremisesPublishingProfiles
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
// returns a PublishedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublishedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable), nil
}
// Patch update the navigation property publishedResources in onPremisesPublishingProfiles
// returns a PublishedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublishedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable), nil
}
// ToDeleteRequestInformation delete navigation property publishedResources for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property publishedResources in onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder when successful
func (m *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) WithUrl(rawUrl string)(*ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    return NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
