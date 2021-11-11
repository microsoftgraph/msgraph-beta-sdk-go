package directory

import (
    i099820decfc9bea0a9207fe8eabbdbec38a02bc9f29b8c032c0f3dc2b1de2037 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits"
    i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains"
    i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems"
    i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies"
    iaf1d1cd519f0199547e1f554f5d9d5790e9b1bbf35a06ad1d200c0cbcf09c175 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions"
    ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/attributesets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/federationconfigurations"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item"
    i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/sharedemaildomains/item"
    i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item"
    i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions/item"
    ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item"
    ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/federationconfigurations/item"
    ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/attributesets/item"
)

// Builds and executes requests for operations under \directory
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
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
// Get directory
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type DirectoryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory;
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
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) AdministrativeUnitsById(id string)(*i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.AdministrativeUnitRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["administrativeUnit_id"] = id
    }
    return i6dbb852f4179b2099ef0d7cfea974aab29da28ff271825ed4facd1031635a55a.NewAdministrativeUnitRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) AttributeSets()(*ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1.AttributeSetsRequestBuilder) {
    return ib5dfa66383589c099255c501996228eba7a0bc49ce9c7f992a9607cce3a1b8b1.NewAttributeSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.attributeSets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) AttributeSetsById(id string)(*ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223.AttributeSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attributeSet_id"] = id
    }
    return ifebb2851fc7ede78464b250ca6ca9ee9431aaf2ea5e0cabbeec1de32f3d73223.NewAttributeSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DirectoryRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DirectoryRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get directory
// Parameters:
//  - options : Options for the request
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(options *DirectoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Update directory
// Parameters:
//  - options : Options for the request
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
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.customSecurityAttributeDefinitions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) CustomSecurityAttributeDefinitionsById(id string)(*i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6.CustomSecurityAttributeDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customSecurityAttributeDefinition_id"] = id
    }
    return i9b5ae363d6622db0767524c36dfc4944cb635f7495736d92423ec12e1c606ff6.NewCustomSecurityAttributeDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) DeletedItems()(*i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.DeletedItemsRequestBuilder) {
    return i3cd67c30f5b1dbe46ac1a4cc4e3febcb7347d9b6590ebcbb3a7c4ed486b7a33a.NewDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.deletedItems.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) DeletedItemsById(id string)(*ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.DirectoryObjectRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ia0a2720db6fe1c2002158844efe537fd50cb223ed65a824c2b17bccbc76d639b.NewDirectoryObjectRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) FeatureRolloutPolicies()(*i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.FeatureRolloutPoliciesRequestBuilder) {
    return i874ed4edfcf924c8bf52b0ce3fb2cbbe5980293405fe171ac739490cbbcdaabf.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.featureRolloutPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) FeatureRolloutPoliciesById(id string)(*i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.FeatureRolloutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy_id"] = id
    }
    return i5552436032dc99a204b9894d3cdec8c8fd49de8267b7be85641d407446e4b6b4.NewFeatureRolloutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) FederationConfigurations()(*iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a.FederationConfigurationsRequestBuilder) {
    return iee88df06b6c2f88b23f7c64f0bd54b255fc8bd6e4ed8805e6adf9aa81134997a.NewFederationConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.federationConfigurations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) FederationConfigurationsById(id string)(*ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b.IdentityProviderBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase_id"] = id
    }
    return ic56ab875d66f866a4e02c3e21e059765fd0587edf1d7176130c0d565ab17801b.NewIdentityProviderBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get directory
// Parameters:
//  - options : Options for the request
func (m *DirectoryRequestBuilder) Get(options *DirectoryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectory() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Directory), nil
}
// Update directory
// Parameters:
//  - options : Options for the request
func (m *DirectoryRequestBuilder) Patch(options *DirectoryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectoryRequestBuilder) SharedEmailDomains()(*i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.SharedEmailDomainsRequestBuilder) {
    return i24baa95180344241e5988e88abde66879965e1161c491971c4fef299d89ecd3d.NewSharedEmailDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.sharedEmailDomains.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DirectoryRequestBuilder) SharedEmailDomainsById(id string)(*i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.SharedEmailDomainRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedEmailDomain_id"] = id
    }
    return i596c69bb84519fa1bd095263c7f63372f6b77b2d592adcc064a8f58e4664f18a.NewSharedEmailDomainRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
