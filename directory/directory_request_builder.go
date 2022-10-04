package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains"
    i3cbe87545047edd523f75c012306a35d5b10521745adafe10e986c1fc5101fda "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources"
    i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems"
    i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles"
    i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies"
    iaf1d1cd519f0199547e1f554f5d9d5790e9b1bbf35a06ad1d200c0cbcf09c175 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions"
    ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles"
    ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/attributesets"
    ide623d002dd8885a2ce7a958e38ce0272f6cc33213ccafda55701bab0956e95f "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations"
    iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/federationconfigurations"
    i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles/item"
    i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item"
    i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains/item"
    i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item"
    i9134fdf6cb6de8e59b19b19500831f6af1d8bb62e86076ee3571e1840ede40fb "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles/item"
    i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions/item"
    ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item"
    ia636abbcfb3919c66696d3c05179fcf8f8119643975a04700a611861aa160b5c "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item"
    ic0e4d6d4cf5e413a63e73490ea86e089e106e5a7b365a7cbc55fdf6d36d2bbf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item"
    ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/federationconfigurations/item"
    ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/attributesets/item"
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
// AdministrativeUnits the administrativeUnits property
func (m *DirectoryRequestBuilder) AdministrativeUnits()(*i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037.AdministrativeUnitsRequestBuilder) {
    return i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037.NewAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdministrativeUnitsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item collection
func (m *DirectoryRequestBuilder) AdministrativeUnitsById(id string)(*i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.AdministrativeUnitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["administrativeUnit%2Did"] = id
    }
    return i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.NewAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttributeSets the attributeSets property
func (m *DirectoryRequestBuilder) AttributeSets()(*ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1.AttributeSetsRequestBuilder) {
    return ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1.NewAttributeSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttributeSetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.attributeSets.item collection
func (m *DirectoryRequestBuilder) AttributeSetsById(id string)(*ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223.AttributeSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attributeSet%2Did"] = id
    }
    return ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223.NewAttributeSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// CustomSecurityAttributeDefinitions the customSecurityAttributeDefinitions property
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitions()(*iaf1d1cd519f0199547e1f554f5d9d5790e9b1bbf35a06ad1d200c0cbcf09c175.CustomSecurityAttributeDefinitionsRequestBuilder) {
    return iaf1d1cd519f0199547e1f554f5d9d5790e9b1bbf35a06ad1d200c0cbcf09c175.NewCustomSecurityAttributeDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomSecurityAttributeDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.customSecurityAttributeDefinitions.item collection
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitionsById(id string)(*i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6.CustomSecurityAttributeDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customSecurityAttributeDefinition%2Did"] = id
    }
    return i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6.NewCustomSecurityAttributeDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeletedItems the deletedItems property
func (m *DirectoryRequestBuilder) DeletedItems()(*i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.DeletedItemsRequestBuilder) {
    return i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.NewDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletedItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.deletedItems.item collection
func (m *DirectoryRequestBuilder) DeletedItemsById(id string)(*ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FeatureRolloutPolicies the featureRolloutPolicies property
func (m *DirectoryRequestBuilder) FeatureRolloutPolicies()(*i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.FeatureRolloutPoliciesRequestBuilder) {
    return i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.featureRolloutPolicies.item collection
func (m *DirectoryRequestBuilder) FeatureRolloutPoliciesById(id string)(*i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.FeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = id
    }
    return i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.NewFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederationConfigurations the federationConfigurations property
func (m *DirectoryRequestBuilder) FederationConfigurations()(*iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a.FederationConfigurationsRequestBuilder) {
    return iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a.NewFederationConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederationConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.federationConfigurations.item collection
func (m *DirectoryRequestBuilder) FederationConfigurationsById(id string)(*ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// ImpactedResources the impactedResources property
func (m *DirectoryRequestBuilder) ImpactedResources()(*i3cbe87545047edd523f75c012306a35d5b10521745adafe10e986c1fc5101fda.ImpactedResourcesRequestBuilder) {
    return i3cbe87545047edd523f75c012306a35d5b10521745adafe10e986c1fc5101fda.NewImpactedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImpactedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.impactedResources.item collection
func (m *DirectoryRequestBuilder) ImpactedResourcesById(id string)(*ia636abbcfb3919c66696d3c05179fcf8f8119643975a04700a611861aa160b5c.RecommendationResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendationResource%2Did"] = id
    }
    return ia636abbcfb3919c66696d3c05179fcf8f8119643975a04700a611861aa160b5c.NewRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InboundSharedUserProfiles the inboundSharedUserProfiles property
func (m *DirectoryRequestBuilder) InboundSharedUserProfiles()(*i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689.InboundSharedUserProfilesRequestBuilder) {
    return i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689.NewInboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InboundSharedUserProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.inboundSharedUserProfiles.item collection
func (m *DirectoryRequestBuilder) InboundSharedUserProfilesById(id string)(*i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273.InboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["inboundSharedUserProfile%2DuserId"] = id
    }
    return i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273.NewInboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OutboundSharedUserProfiles the outboundSharedUserProfiles property
func (m *DirectoryRequestBuilder) OutboundSharedUserProfiles()(*ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca.OutboundSharedUserProfilesRequestBuilder) {
    return ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca.NewOutboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutboundSharedUserProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.outboundSharedUserProfiles.item collection
func (m *DirectoryRequestBuilder) OutboundSharedUserProfilesById(id string)(*i9134fdf6cb6de8e59b19b19500831f6af1d8bb62e86076ee3571e1840ede40fb.OutboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outboundSharedUserProfile%2DuserId"] = id
    }
    return i9134fdf6cb6de8e59b19b19500831f6af1d8bb62e86076ee3571e1840ede40fb.NewOutboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// Recommendations the recommendations property
func (m *DirectoryRequestBuilder) Recommendations()(*ide623d002dd8885a2ce7a958e38ce0272f6cc33213ccafda55701bab0956e95f.RecommendationsRequestBuilder) {
    return ide623d002dd8885a2ce7a958e38ce0272f6cc33213ccafda55701bab0956e95f.NewRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecommendationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.recommendations.item collection
func (m *DirectoryRequestBuilder) RecommendationsById(id string)(*ic0e4d6d4cf5e413a63e73490ea86e089e106e5a7b365a7cbc55fdf6d36d2bbf7.RecommendationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendation%2Did"] = id
    }
    return ic0e4d6d4cf5e413a63e73490ea86e089e106e5a7b365a7cbc55fdf6d36d2bbf7.NewRecommendationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedEmailDomains the sharedEmailDomains property
func (m *DirectoryRequestBuilder) SharedEmailDomains()(*i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.SharedEmailDomainsRequestBuilder) {
    return i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.NewSharedEmailDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedEmailDomainsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.sharedEmailDomains.item collection
func (m *DirectoryRequestBuilder) SharedEmailDomainsById(id string)(*i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.SharedEmailDomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedEmailDomain%2Did"] = id
    }
    return i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.NewSharedEmailDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
