package directory

import (
    i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits"
    i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains"
    i3cbe87545047edd523f75c012306a35d5b10521745adafe10e986c1fc5101fda "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources"
    i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems"
    i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles"
    i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies"
    iaf1d1cd519f0199547e1f554f5d9d5790e9b1bbf35a06ad1d200c0cbcf09c175 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions"
    ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles"
    ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/attributesets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ide623d002dd8885a2ce7a958e38ce0272f6cc33213ccafda55701bab0956e95f "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations"
    iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/federationconfigurations"
    i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRequestBuilderGetOptions options for Get
type DirectoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRequestBuilderGetQueryParameters get directory
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRequestBuilderPatchOptions options for Patch
type DirectoryRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directoryable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["administrativeUnit_id"] = id
    }
    return i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.NewAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["attributeSet_id"] = id
    }
    return ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223.NewAttributeSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get directory
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(options *DirectoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update directory
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(options *DirectoryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
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
        urlTplParams["customSecurityAttributeDefinition_id"] = id
    }
    return i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6.NewCustomSecurityAttributeDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["featureRolloutPolicy_id"] = id
    }
    return i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.NewFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["identityProviderBase_id"] = id
    }
    return ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get directory
func (m *DirectoryRequestBuilder) Get(options *DirectoryRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDirectoryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directoryable), nil
}
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
        urlTplParams["recommendationResource_id"] = id
    }
    return ia636abbcfb3919c66696d3c05179fcf8f8119643975a04700a611861aa160b5c.NewRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) InboundSharedUserProfiles()(*i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689.InboundSharedUserProfilesRequestBuilder) {
    return i77c0ae00685e8d9a81a4aecf1b394178c22d70b82eef43c63dc29f705d913689.NewInboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InboundSharedUserProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.inboundSharedUserProfiles.item collection
func (m *DirectoryRequestBuilder) InboundSharedUserProfilesById(id string)(*i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273.InboundSharedUserProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["inboundSharedUserProfile_userId"] = id
    }
    return i223f1150f57c5b2623641984f128295eead7dbc629507266bd840afb65dde273.NewInboundSharedUserProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) OutboundSharedUserProfiles()(*ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca.OutboundSharedUserProfilesRequestBuilder) {
    return ib4c7c90d94588a870946e877d26cfcd42ca3d90d8f00320757f06ef38d0c08ca.NewOutboundSharedUserProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutboundSharedUserProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.outboundSharedUserProfiles.item collection
func (m *DirectoryRequestBuilder) OutboundSharedUserProfilesById(id string)(*i9134fdf6cb6de8e59b19b19500831f6af1d8bb62e86076ee3571e1840ede40fb.OutboundSharedUserProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outboundSharedUserProfile_userId"] = id
    }
    return i9134fdf6cb6de8e59b19b19500831f6af1d8bb62e86076ee3571e1840ede40fb.NewOutboundSharedUserProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update directory
func (m *DirectoryRequestBuilder) Patch(options *DirectoryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["recommendation_id"] = id
    }
    return ic0e4d6d4cf5e413a63e73490ea86e089e106e5a7b365a7cbc55fdf6d36d2bbf7.NewRecommendationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["sharedEmailDomain_id"] = id
    }
    return i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.NewSharedEmailDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
