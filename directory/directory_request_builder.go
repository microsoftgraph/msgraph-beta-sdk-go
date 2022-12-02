package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRequestBuilder provides operations to manage the directory singleton.
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRequestBuilderGetQueryParameters get directory
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters
}
// DirectoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdministrativeUnits provides operations to manage the administrativeUnits property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) AdministrativeUnits()(*DirectoryAdministrativeUnitsRequestBuilder) {
    return NewDirectoryAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdministrativeUnitsById provides operations to manage the administrativeUnits property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) AdministrativeUnitsById(id string)(*DirectoryAdministrativeUnitsAdministrativeUnitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["administrativeUnit%2Did"] = id
    }
    return NewDirectoryAdministrativeUnitsAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttributeSets provides operations to manage the attributeSets property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) AttributeSets()(*DirectoryAttributeSetsRequestBuilder) {
    return NewDirectoryAttributeSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttributeSetsById provides operations to manage the attributeSets property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) AttributeSetsById(id string)(*DirectoryAttributeSetsAttributeSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attributeSet%2Did"] = id
    }
    return NewDirectoryAttributeSetsAttributeSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get directory
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update directory
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustomSecurityAttributeDefinitions provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitions()(*DirectoryCustomSecurityAttributeDefinitionsRequestBuilder) {
    return NewDirectoryCustomSecurityAttributeDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomSecurityAttributeDefinitionsById provides operations to manage the customSecurityAttributeDefinitions property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitionsById(id string)(*DirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customSecurityAttributeDefinition%2Did"] = id
    }
    return NewDirectoryCustomSecurityAttributeDefinitionsCustomSecurityAttributeDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeletedItems provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) DeletedItems()(*DirectoryDeletedItemsRequestBuilder) {
    return NewDirectoryDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletedItemsById provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) DeletedItemsById(id string)(*DirectoryDeletedItemsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewDirectoryDeletedItemsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FeatureRolloutPolicies provides operations to manage the featureRolloutPolicies property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) FeatureRolloutPolicies()(*DirectoryFeatureRolloutPoliciesRequestBuilder) {
    return NewDirectoryFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById provides operations to manage the featureRolloutPolicies property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) FeatureRolloutPoliciesById(id string)(*DirectoryFeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = id
    }
    return NewDirectoryFeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederationConfigurations provides operations to manage the federationConfigurations property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) FederationConfigurations()(*DirectoryFederationConfigurationsRequestBuilder) {
    return NewDirectoryFederationConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederationConfigurationsById provides operations to manage the federationConfigurations property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) FederationConfigurationsById(id string)(*DirectoryFederationConfigurationsIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return NewDirectoryFederationConfigurationsIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get directory
func (m *DirectoryRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable), nil
}
// ImpactedResources provides operations to manage the impactedResources property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) ImpactedResources()(*DirectoryImpactedResourcesRequestBuilder) {
    return NewDirectoryImpactedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImpactedResourcesById provides operations to manage the impactedResources property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) ImpactedResourcesById(id string)(*DirectoryImpactedResourcesRecommendationResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendationResource%2Did"] = id
    }
    return NewDirectoryImpactedResourcesRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InboundSharedUserProfiles provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) InboundSharedUserProfiles()(*DirectoryInboundSharedUserProfilesRequestBuilder) {
    return NewDirectoryInboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InboundSharedUserProfilesById provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) InboundSharedUserProfilesById(id string)(*DirectoryInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["inboundSharedUserProfile%2DuserId"] = id
    }
    return NewDirectoryInboundSharedUserProfilesInboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OnPremisesSynchronization provides operations to manage the onPremisesSynchronization property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) OnPremisesSynchronization()(*DirectoryOnPremisesSynchronizationRequestBuilder) {
    return NewDirectoryOnPremisesSynchronizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesSynchronizationById provides operations to manage the onPremisesSynchronization property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) OnPremisesSynchronizationById(id string)(*DirectoryOnPremisesSynchronizationOnPremisesDirectorySynchronizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesDirectorySynchronization%2Did"] = id
    }
    return NewDirectoryOnPremisesSynchronizationOnPremisesDirectorySynchronizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OutboundSharedUserProfiles provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) OutboundSharedUserProfiles()(*DirectoryOutboundSharedUserProfilesRequestBuilder) {
    return NewDirectoryOutboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutboundSharedUserProfilesById provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) OutboundSharedUserProfilesById(id string)(*DirectoryOutboundSharedUserProfilesOutboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outboundSharedUserProfile%2DuserId"] = id
    }
    return NewDirectoryOutboundSharedUserProfilesOutboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update directory
func (m *DirectoryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Directoryable), nil
}
// Recommendations provides operations to manage the recommendations property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) Recommendations()(*DirectoryRecommendationsRequestBuilder) {
    return NewDirectoryRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecommendationsById provides operations to manage the recommendations property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) RecommendationsById(id string)(*DirectoryRecommendationsRecommendationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendation%2Did"] = id
    }
    return NewDirectoryRecommendationsRecommendationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedEmailDomains provides operations to manage the sharedEmailDomains property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) SharedEmailDomains()(*DirectorySharedEmailDomainsRequestBuilder) {
    return NewDirectorySharedEmailDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedEmailDomainsById provides operations to manage the sharedEmailDomains property of the microsoft.graph.directory entity.
func (m *DirectoryRequestBuilder) SharedEmailDomainsById(id string)(*DirectorySharedEmailDomainsSharedEmailDomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedEmailDomain%2Did"] = id
    }
    return NewDirectorySharedEmailDomainsSharedEmailDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
