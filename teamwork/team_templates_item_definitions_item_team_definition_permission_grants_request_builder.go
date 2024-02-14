package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetQueryParameters list all resource-specific permission grants on the team. This list specifies the Microsoft Entra apps that have access to the team, along with the corresponding kind of resource-specific access that each app has.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetQueryParameters struct {
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
// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByResourceSpecificPermissionGrantId provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) ByResourceSpecificPermissionGrantId(resourceSpecificPermissionGrantId string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if resourceSpecificPermissionGrantId != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = resourceSpecificPermissionGrantId
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/permissionGrants{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsCountRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) Count()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsCountRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list all resource-specific permission grants on the team. This list specifies the Microsoft Entra apps that have access to the team, along with the corresponding kind of resource-specific access that each app has.
// returns a ResourceSpecificPermissionGrantCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-list-permissiongrants?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantCollectionResponseable), nil
}
// Post create new navigation property to permissionGrants for teamwork
// returns a ResourceSpecificPermissionGrantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// ToGetRequestInformation list all resource-specific permission grants on the team. This list specifies the Microsoft Entra apps that have access to the team, along with the corresponding kind of resource-specific access that each app has.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to permissionGrants for teamwork
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/permissionGrants", m.BaseRequestBuilder.PathParameters)
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
