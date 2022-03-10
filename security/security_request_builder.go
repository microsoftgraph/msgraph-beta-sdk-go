package security

import (
    i2d37e58be29c573dea772021cbc16fb5110ad11503c90334105692cd58482307 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescorecontrolprofiles"
    i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts"
    i3076976e1c2468886103a1402ff2835f889faba296f1c85ce51c99ee09e38cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cloudappsecurityprofiles"
    i6b80d132c3193d55697cc8d6c1b7a29dd0535a152ced7eeb8aa3a99d27d63daa "github.com/microsoftgraph/msgraph-beta-sdk-go/security/filesecurityprofiles"
    i770070b4ce265ea6cee72f7d73d32b4c3ff094b8c0ccac5230bba695726f4bc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/domainsecurityprofiles"
    i7e095c516b5de7d1e6e8fa4aaf680dada81b5f596dc746bf98e54a2a345122ac "github.com/microsoftgraph/msgraph-beta-sdk-go/security/hostsecurityprofiles"
    i81ba5cba60c873845066ad89e2c6fe67a50b8f628aaa18e17de96216c62d64f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators"
    i81db0892066e847c456de63f4c19382f89016247d0403ce885a78520fdcd01b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/usersecurityprofiles"
    i8ae7bba33296af81c9eae761d1a04dfc084176f7cefbeb91b73b5b37bbb03740 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/providertenantsettings"
    ib0ce1c8104cd7a5487ab293f5f33f3c3b896da30437f31a1bb830fffeca99765 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securityactions"
    id43f72338affca62b04ede5a8930eef2eacb7985a0d9ba5d084163b59793c1b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescores"
    id8c8c6e6d2248957fbc7631aef94353af5a458e67dcc26feaa6a1450a99a895c "github.com/microsoftgraph/msgraph-beta-sdk-go/security/ipsecurityprofiles"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation"
    ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection"
    i02e0bbeb2df2eda7dbc98813d493b8564c05050583278d44a701c5e5b6ef886a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/hostsecurityprofiles/item"
    i0905673ef4d7fa54e4d3f55e349406fd6862ddc6c802fb42ff1e5afa87694e6e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/domainsecurityprofiles/item"
    i5a3d2f037bd277d4a881c48a1993939f3df188957e325d6ae0bdf67636fcf03b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescorecontrolprofiles/item"
    i5c9089aa1bb65758fcfbff1d8a7aa9b5deeb5741c421c67322898649fd582848 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/item"
    i6c9cc74f462fa9e89a76cab1d5a7489b11c2e26c4d3a81cda11945485ac87c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescores/item"
    i8366108fde7cb9abb94de15d00c81b743720ced29144f928fd43323122437c0a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/ipsecurityprofiles/item"
    i956e12778a313e3750d20998bd9bc58b24ba11b94459928db9623590d22b6ddc "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securityactions/item"
    i9b85a1e7bdfc97cbdd52f3319eecf44724fb90aed32a007e01bbf117230a1cbd "github.com/microsoftgraph/msgraph-beta-sdk-go/security/usersecurityprofiles/item"
    ic633bb560cd35fd4a473b252391ace288264bc3e0245adcecf45bb4f003dcc2d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/filesecurityprofiles/item"
    icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts/item"
    idfc7ee4e0b30b6fb2a96f1eed667daeae65a0816e9af7bc38a13ef77dac4c0d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/providertenantsettings/item"
    ie02fc175d3be34ac616ef3dda9f36f01be04b3ccf591093492385bc0ab62023a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cloudappsecurityprofiles/item"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// SecurityRequestBuilder provides operations to manage the security singleton.
type SecurityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SecurityRequestBuilderGetOptions options for Get
type SecurityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SecurityRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SecurityRequestBuilderGetQueryParameters get security
type SecurityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SecurityRequestBuilderPatchOptions options for Patch
type SecurityRequestBuilderPatchOptions struct {
    // 
    Body i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.Securityable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SecurityRequestBuilder) Alerts()(*i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9.AlertsRequestBuilder) {
    return i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9.NewAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.alerts.item collection
func (m *SecurityRequestBuilder) AlertsById(id string)(*icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7.AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert_id"] = id
    }
    return icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7.NewAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) AttackSimulation()(*iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6.AttackSimulationRequestBuilder) {
    return iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6.NewAttackSimulationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SecurityRequestBuilder) CloudAppSecurityProfiles()(*i3076976e1c2468886103a1402ff2835f889faba296f1c85ce51c99ee09e38cdc.CloudAppSecurityProfilesRequestBuilder) {
    return i3076976e1c2468886103a1402ff2835f889faba296f1c85ce51c99ee09e38cdc.NewCloudAppSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudAppSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cloudAppSecurityProfiles.item collection
func (m *SecurityRequestBuilder) CloudAppSecurityProfilesById(id string)(*ie02fc175d3be34ac616ef3dda9f36f01be04b3ccf591093492385bc0ab62023a.CloudAppSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudAppSecurityProfile_id"] = id
    }
    return ie02fc175d3be34ac616ef3dda9f36f01be04b3ccf591093492385bc0ab62023a.NewCloudAppSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSecurityRequestBuilderInternal instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityRequestBuilder instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get security
func (m *SecurityRequestBuilder) CreateGetRequestInformation(options *SecurityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update security
func (m *SecurityRequestBuilder) CreatePatchRequestInformation(options *SecurityRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SecurityRequestBuilder) DomainSecurityProfiles()(*i770070b4ce265ea6cee72f7d73d32b4c3ff094b8c0ccac5230bba695726f4bc2.DomainSecurityProfilesRequestBuilder) {
    return i770070b4ce265ea6cee72f7d73d32b4c3ff094b8c0ccac5230bba695726f4bc2.NewDomainSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.domainSecurityProfiles.item collection
func (m *SecurityRequestBuilder) DomainSecurityProfilesById(id string)(*i0905673ef4d7fa54e4d3f55e349406fd6862ddc6c802fb42ff1e5afa87694e6e.DomainSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainSecurityProfile_id"] = id
    }
    return i0905673ef4d7fa54e4d3f55e349406fd6862ddc6c802fb42ff1e5afa87694e6e.NewDomainSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) FileSecurityProfiles()(*i6b80d132c3193d55697cc8d6c1b7a29dd0535a152ced7eeb8aa3a99d27d63daa.FileSecurityProfilesRequestBuilder) {
    return i6b80d132c3193d55697cc8d6c1b7a29dd0535a152ced7eeb8aa3a99d27d63daa.NewFileSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FileSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.fileSecurityProfiles.item collection
func (m *SecurityRequestBuilder) FileSecurityProfilesById(id string)(*ic633bb560cd35fd4a473b252391ace288264bc3e0245adcecf45bb4f003dcc2d.FileSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileSecurityProfile_id"] = id
    }
    return ic633bb560cd35fd4a473b252391ace288264bc3e0245adcecf45bb4f003dcc2d.NewFileSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get security
func (m *SecurityRequestBuilder) Get(options *SecurityRequestBuilderGetOptions)(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.Securityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateSecurityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.Securityable), nil
}
func (m *SecurityRequestBuilder) HostSecurityProfiles()(*i7e095c516b5de7d1e6e8fa4aaf680dada81b5f596dc746bf98e54a2a345122ac.HostSecurityProfilesRequestBuilder) {
    return i7e095c516b5de7d1e6e8fa4aaf680dada81b5f596dc746bf98e54a2a345122ac.NewHostSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.hostSecurityProfiles.item collection
func (m *SecurityRequestBuilder) HostSecurityProfilesById(id string)(*i02e0bbeb2df2eda7dbc98813d493b8564c05050583278d44a701c5e5b6ef886a.HostSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostSecurityProfile_id"] = id
    }
    return i02e0bbeb2df2eda7dbc98813d493b8564c05050583278d44a701c5e5b6ef886a.NewHostSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) InformationProtection()(*ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f.InformationProtectionRequestBuilder) {
    return ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SecurityRequestBuilder) IpSecurityProfiles()(*id8c8c6e6d2248957fbc7631aef94353af5a458e67dcc26feaa6a1450a99a895c.IpSecurityProfilesRequestBuilder) {
    return id8c8c6e6d2248957fbc7631aef94353af5a458e67dcc26feaa6a1450a99a895c.NewIpSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IpSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.ipSecurityProfiles.item collection
func (m *SecurityRequestBuilder) IpSecurityProfilesById(id string)(*i8366108fde7cb9abb94de15d00c81b743720ced29144f928fd43323122437c0a.IpSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ipSecurityProfile_id"] = id
    }
    return i8366108fde7cb9abb94de15d00c81b743720ced29144f928fd43323122437c0a.NewIpSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update security
func (m *SecurityRequestBuilder) Patch(options *SecurityRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SecurityRequestBuilder) ProviderTenantSettings()(*i8ae7bba33296af81c9eae761d1a04dfc084176f7cefbeb91b73b5b37bbb03740.ProviderTenantSettingsRequestBuilder) {
    return i8ae7bba33296af81c9eae761d1a04dfc084176f7cefbeb91b73b5b37bbb03740.NewProviderTenantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProviderTenantSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.providerTenantSettings.item collection
func (m *SecurityRequestBuilder) ProviderTenantSettingsById(id string)(*idfc7ee4e0b30b6fb2a96f1eed667daeae65a0816e9af7bc38a13ef77dac4c0d5.ProviderTenantSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["providerTenantSetting_id"] = id
    }
    return idfc7ee4e0b30b6fb2a96f1eed667daeae65a0816e9af7bc38a13ef77dac4c0d5.NewProviderTenantSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*i2d37e58be29c573dea772021cbc16fb5110ad11503c90334105692cd58482307.SecureScoreControlProfilesRequestBuilder) {
    return i2d37e58be29c573dea772021cbc16fb5110ad11503c90334105692cd58482307.NewSecureScoreControlProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoreControlProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.secureScoreControlProfiles.item collection
func (m *SecurityRequestBuilder) SecureScoreControlProfilesById(id string)(*i5a3d2f037bd277d4a881c48a1993939f3df188957e325d6ae0bdf67636fcf03b.SecureScoreControlProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScoreControlProfile_id"] = id
    }
    return i5a3d2f037bd277d4a881c48a1993939f3df188957e325d6ae0bdf67636fcf03b.NewSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScores()(*id43f72338affca62b04ede5a8930eef2eacb7985a0d9ba5d084163b59793c1b0.SecureScoresRequestBuilder) {
    return id43f72338affca62b04ede5a8930eef2eacb7985a0d9ba5d084163b59793c1b0.NewSecureScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoresById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.secureScores.item collection
func (m *SecurityRequestBuilder) SecureScoresById(id string)(*i6c9cc74f462fa9e89a76cab1d5a7489b11c2e26c4d3a81cda11945485ac87c86.SecureScoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScore_id"] = id
    }
    return i6c9cc74f462fa9e89a76cab1d5a7489b11c2e26c4d3a81cda11945485ac87c86.NewSecureScoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecurityActions()(*ib0ce1c8104cd7a5487ab293f5f33f3c3b896da30437f31a1bb830fffeca99765.SecurityActionsRequestBuilder) {
    return ib0ce1c8104cd7a5487ab293f5f33f3c3b896da30437f31a1bb830fffeca99765.NewSecurityActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityActionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.securityActions.item collection
func (m *SecurityRequestBuilder) SecurityActionsById(id string)(*i956e12778a313e3750d20998bd9bc58b24ba11b94459928db9623590d22b6ddc.SecurityActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityAction_id"] = id
    }
    return i956e12778a313e3750d20998bd9bc58b24ba11b94459928db9623590d22b6ddc.NewSecurityActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) TiIndicators()(*i81ba5cba60c873845066ad89e2c6fe67a50b8f628aaa18e17de96216c62d64f4.TiIndicatorsRequestBuilder) {
    return i81ba5cba60c873845066ad89e2c6fe67a50b8f628aaa18e17de96216c62d64f4.NewTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TiIndicatorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.tiIndicators.item collection
func (m *SecurityRequestBuilder) TiIndicatorsById(id string)(*i5c9089aa1bb65758fcfbff1d8a7aa9b5deeb5741c421c67322898649fd582848.TiIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tiIndicator_id"] = id
    }
    return i5c9089aa1bb65758fcfbff1d8a7aa9b5deeb5741c421c67322898649fd582848.NewTiIndicatorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) UserSecurityProfiles()(*i81db0892066e847c456de63f4c19382f89016247d0403ce885a78520fdcd01b2.UserSecurityProfilesRequestBuilder) {
    return i81db0892066e847c456de63f4c19382f89016247d0403ce885a78520fdcd01b2.NewUserSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSecurityProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.userSecurityProfiles.item collection
func (m *SecurityRequestBuilder) UserSecurityProfilesById(id string)(*i9b85a1e7bdfc97cbdd52f3319eecf44724fb90aed32a007e01bbf117230a1cbd.UserSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSecurityProfile_id"] = id
    }
    return i9b85a1e7bdfc97cbdd52f3319eecf44724fb90aed32a007e01bbf117230a1cbd.NewUserSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
