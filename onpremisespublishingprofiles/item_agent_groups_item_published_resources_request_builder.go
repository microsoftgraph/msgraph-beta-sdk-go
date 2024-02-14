package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAgentGroupsItemPublishedResourcesRequestBuilder provides operations to manage the publishedResources property of the microsoft.graph.onPremisesAgentGroup entity.
type ItemAgentGroupsItemPublishedResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAgentGroupsItemPublishedResourcesRequestBuilderGetQueryParameters list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
type ItemAgentGroupsItemPublishedResourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemAgentGroupsItemPublishedResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemPublishedResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAgentGroupsItemPublishedResourcesRequestBuilderGetQueryParameters
}
// ItemAgentGroupsItemPublishedResourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemPublishedResourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPublishedResourceId provides operations to manage the publishedResources property of the microsoft.graph.onPremisesAgentGroup entity.
// returns a *ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder when successful
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) ByPublishedResourceId(publishedResourceId string)(*ItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if publishedResourceId != "" {
        urlTplParams["publishedResource%2Did"] = publishedResourceId
    }
    return NewItemAgentGroupsItemPublishedResourcesPublishedResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAgentGroupsItemPublishedResourcesRequestBuilderInternal instantiates a new ItemAgentGroupsItemPublishedResourcesRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesRequestBuilder) {
    m := &ItemAgentGroupsItemPublishedResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAgentGroupsItemPublishedResourcesRequestBuilder instantiates a new ItemAgentGroupsItemPublishedResourcesRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentGroupsItemPublishedResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAgentGroupsItemPublishedResourcesCountRequestBuilder when successful
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) Count()(*ItemAgentGroupsItemPublishedResourcesCountRequestBuilder) {
    return NewItemAgentGroupsItemPublishedResourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
// returns a PublishedResourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublishedResourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceCollectionResponseable), nil
}
// Post create new navigation property to publishedResources for onPremisesPublishingProfiles
// returns a PublishedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *ItemAgentGroupsItemPublishedResourcesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgentGroupsItemPublishedResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to publishedResources for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublishedResourceable, requestConfiguration *ItemAgentGroupsItemPublishedResourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemAgentGroupsItemPublishedResourcesRequestBuilder when successful
func (m *ItemAgentGroupsItemPublishedResourcesRequestBuilder) WithUrl(rawUrl string)(*ItemAgentGroupsItemPublishedResourcesRequestBuilder) {
    return NewItemAgentGroupsItemPublishedResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
