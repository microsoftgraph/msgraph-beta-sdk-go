package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0c676084c053b74ed3a3133cb01b5e628a094e0413da8590be1ed0afd08d166f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/triggertypes"
    i0df2000278766f4e13bb6be9bde05e5eb0c186391a4a2cb9f634d92117307226 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/threatsubmission"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i29dbe885802105b8f24e18a3ad45a4baf47a032bed4a97cc90638ac720b35332 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests"
    i2d37e58be29c573dea772021cbc16fb5110ad11503c90334105692cd58482307 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescorecontrolprofiles"
    i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts"
    i3076976e1c2468886103a1402ff2835f889faba296f1c85ce51c99ee09e38cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cloudappsecurityprofiles"
    i69463578215a124f2797b45520f2b68cc0cbbf908f813f3f3f8585fd63973a93 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/incidents"
    i6b80d132c3193d55697cc8d6c1b7a29dd0535a152ced7eeb8aa3a99d27d63daa "github.com/microsoftgraph/msgraph-beta-sdk-go/security/filesecurityprofiles"
    i770070b4ce265ea6cee72f7d73d32b4c3ff094b8c0ccac5230bba695726f4bc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/domainsecurityprofiles"
    i7e095c516b5de7d1e6e8fa4aaf680dada81b5f596dc746bf98e54a2a345122ac "github.com/microsoftgraph/msgraph-beta-sdk-go/security/hostsecurityprofiles"
    i81ba5cba60c873845066ad89e2c6fe67a50b8f628aaa18e17de96216c62d64f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators"
    i81db0892066e847c456de63f4c19382f89016247d0403ce885a78520fdcd01b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/usersecurityprofiles"
    i8ae7bba33296af81c9eae761d1a04dfc084176f7cefbeb91b73b5b37bbb03740 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/providertenantsettings"
    i9d73acd74b99f8c28d8e6ba0f549aeb25b13d09f6e92903e09e541ed58de8218 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/triggers"
    ib0ce1c8104cd7a5487ab293f5f33f3c3b896da30437f31a1bb830fffeca99765 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securityactions"
    ibd30f474e0f2ed1ac20c040a2570fcc591ecadd864b59a598b03981d488e873b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/labels"
    ic9ba9daf256a77044250cba0d8c29ae1ae3e69cd91e81b4a7aba1988ececf405 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts_v2"
    icabca6351261bc62cf0a59e4be2a3d081990a9a025d76f7a2849d5c1be3089b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/runhuntingquery"
    id43f72338affca62b04ede5a8930eef2eacb7985a0d9ba5d084163b59793c1b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescores"
    id8c8c6e6d2248957fbc7631aef94353af5a458e67dcc26feaa6a1450a99a895c "github.com/microsoftgraph/msgraph-beta-sdk-go/security/ipsecurityprofiles"
    iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation"
    if9426fddcb85c4e3b0dbc38e5f8657c3e8823751cef07a81c4685904e8f6a39a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases"
    ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection"
    i02e0bbeb2df2eda7dbc98813d493b8564c05050583278d44a701c5e5b6ef886a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/hostsecurityprofiles/item"
    i0905673ef4d7fa54e4d3f55e349406fd6862ddc6c802fb42ff1e5afa87694e6e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/domainsecurityprofiles/item"
    i18f494f2b6028fa48209e2ed1901867809ed07f584269e2fa79310996bdc0c55 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts_v2/item"
    i5a3d2f037bd277d4a881c48a1993939f3df188957e325d6ae0bdf67636fcf03b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescorecontrolprofiles/item"
    i5c9089aa1bb65758fcfbff1d8a7aa9b5deeb5741c421c67322898649fd582848 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/item"
    i609377b6be386a3d85bc23c0d4098c9e56229a97a58f2b8b00bb39804bd3ec91 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/incidents/item"
    i6c9cc74f462fa9e89a76cab1d5a7489b11c2e26c4d3a81cda11945485ac87c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securescores/item"
    i8366108fde7cb9abb94de15d00c81b743720ced29144f928fd43323122437c0a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/ipsecurityprofiles/item"
    i956e12778a313e3750d20998bd9bc58b24ba11b94459928db9623590d22b6ddc "github.com/microsoftgraph/msgraph-beta-sdk-go/security/securityactions/item"
    i9b85a1e7bdfc97cbdd52f3319eecf44724fb90aed32a007e01bbf117230a1cbd "github.com/microsoftgraph/msgraph-beta-sdk-go/security/usersecurityprofiles/item"
    i9fe89cbebebc9e8421ea4e1df57c4ac03af0de51d9dea1f1d0ba1d6023c52df0 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item"
    ic633bb560cd35fd4a473b252391ace288264bc3e0245adcecf45bb4f003dcc2d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/filesecurityprofiles/item"
    icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/alerts/item"
    idfc7ee4e0b30b6fb2a96f1eed667daeae65a0816e9af7bc38a13ef77dac4c0d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/providertenantsettings/item"
    ie02fc175d3be34ac616ef3dda9f36f01be04b3ccf591093492385bc0ab62023a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cloudappsecurityprofiles/item"
)

// SecurityRequestBuilder provides operations to manage the security singleton.
type SecurityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityRequestBuilderGetQueryParameters get security
type SecurityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityRequestBuilderGetQueryParameters
}
// SecurityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alerts the alerts property
func (m *SecurityRequestBuilder) Alerts()(*i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9.AlertsRequestBuilder) {
    return i30697897f04d53b804bcad1f0153ad552b622f468be9ca14198e001717fcbdd9.NewAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Alerts_v2 the alerts_v2 property
func (m *SecurityRequestBuilder) Alerts_v2()(*ic9ba9daf256a77044250cba0d8c29ae1ae3e69cd91e81b4a7aba1988ececf405.Alerts_v2RequestBuilder) {
    return ic9ba9daf256a77044250cba0d8c29ae1ae3e69cd91e81b4a7aba1988ececf405.NewAlerts_v2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Alerts_v2ById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.alerts_v2.item collection
func (m *SecurityRequestBuilder) Alerts_v2ById(id string)(*i18f494f2b6028fa48209e2ed1901867809ed07f584269e2fa79310996bdc0c55.AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return i18f494f2b6028fa48209e2ed1901867809ed07f584269e2fa79310996bdc0c55.NewAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.alerts.item collection
func (m *SecurityRequestBuilder) AlertsById(id string)(*icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7.AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return icf7abdec996315842957a23a296d4a13dfa457345ce67dfb1fcf777acecb7ef7.NewAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttackSimulation the attackSimulation property
func (m *SecurityRequestBuilder) AttackSimulation()(*iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6.AttackSimulationRequestBuilder) {
    return iee2d4af12376b55381e42df8f22c5ffc39043076828fbebb0a459384ee44c9e6.NewAttackSimulationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cases the cases property
func (m *SecurityRequestBuilder) Cases()(*if9426fddcb85c4e3b0dbc38e5f8657c3e8823751cef07a81c4685904e8f6a39a.CasesRequestBuilder) {
    return if9426fddcb85c4e3b0dbc38e5f8657c3e8823751cef07a81c4685904e8f6a39a.NewCasesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudAppSecurityProfiles the cloudAppSecurityProfiles property
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
        urlTplParams["cloudAppSecurityProfile%2Did"] = id
    }
    return ie02fc175d3be34ac616ef3dda9f36f01be04b3ccf591093492385bc0ab62023a.NewCloudAppSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSecurityRequestBuilderInternal instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityRequestBuilder instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get security
func (m *SecurityRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get security
func (m *SecurityRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update security
func (m *SecurityRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update security
func (m *SecurityRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DomainSecurityProfiles the domainSecurityProfiles property
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
        urlTplParams["domainSecurityProfile%2Did"] = id
    }
    return i0905673ef4d7fa54e4d3f55e349406fd6862ddc6c802fb42ff1e5afa87694e6e.NewDomainSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FileSecurityProfiles the fileSecurityProfiles property
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
        urlTplParams["fileSecurityProfile%2Did"] = id
    }
    return ic633bb560cd35fd4a473b252391ace288264bc3e0245adcecf45bb4f003dcc2d.NewFileSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get security
func (m *SecurityRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get security
func (m *SecurityRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SecurityRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable), nil
}
// HostSecurityProfiles the hostSecurityProfiles property
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
        urlTplParams["hostSecurityProfile%2Did"] = id
    }
    return i02e0bbeb2df2eda7dbc98813d493b8564c05050583278d44a701c5e5b6ef886a.NewHostSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Incidents the incidents property
func (m *SecurityRequestBuilder) Incidents()(*i69463578215a124f2797b45520f2b68cc0cbbf908f813f3f3f8585fd63973a93.IncidentsRequestBuilder) {
    return i69463578215a124f2797b45520f2b68cc0cbbf908f813f3f3f8585fd63973a93.NewIncidentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncidentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.incidents.item collection
func (m *SecurityRequestBuilder) IncidentsById(id string)(*i609377b6be386a3d85bc23c0d4098c9e56229a97a58f2b8b00bb39804bd3ec91.IncidentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["incident%2Did"] = id
    }
    return i609377b6be386a3d85bc23c0d4098c9e56229a97a58f2b8b00bb39804bd3ec91.NewIncidentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InformationProtection the informationProtection property
func (m *SecurityRequestBuilder) InformationProtection()(*ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f.InformationProtectionRequestBuilder) {
    return ife41332ee3198396d805ce9faa05ed3df4e05203fb4e2de40b902e920b0e388f.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IpSecurityProfiles the ipSecurityProfiles property
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
        urlTplParams["ipSecurityProfile%2Did"] = id
    }
    return i8366108fde7cb9abb94de15d00c81b743720ced29144f928fd43323122437c0a.NewIpSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Labels the labels property
func (m *SecurityRequestBuilder) Labels()(*ibd30f474e0f2ed1ac20c040a2570fcc591ecadd864b59a598b03981d488e873b.LabelsRequestBuilder) {
    return ibd30f474e0f2ed1ac20c040a2570fcc591ecadd864b59a598b03981d488e873b.NewLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update security
func (m *SecurityRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update security
func (m *SecurityRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProviderTenantSettings the providerTenantSettings property
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
        urlTplParams["providerTenantSetting%2Did"] = id
    }
    return idfc7ee4e0b30b6fb2a96f1eed667daeae65a0816e9af7bc38a13ef77dac4c0d5.NewProviderTenantSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RunHuntingQuery the runHuntingQuery property
func (m *SecurityRequestBuilder) RunHuntingQuery()(*icabca6351261bc62cf0a59e4be2a3d081990a9a025d76f7a2849d5c1be3089b7.RunHuntingQueryRequestBuilder) {
    return icabca6351261bc62cf0a59e4be2a3d081990a9a025d76f7a2849d5c1be3089b7.NewRunHuntingQueryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoreControlProfiles the secureScoreControlProfiles property
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
        urlTplParams["secureScoreControlProfile%2Did"] = id
    }
    return i5a3d2f037bd277d4a881c48a1993939f3df188957e325d6ae0bdf67636fcf03b.NewSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SecureScores the secureScores property
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
        urlTplParams["secureScore%2Did"] = id
    }
    return i6c9cc74f462fa9e89a76cab1d5a7489b11c2e26c4d3a81cda11945485ac87c86.NewSecureScoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SecurityActions the securityActions property
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
        urlTplParams["securityAction%2Did"] = id
    }
    return i956e12778a313e3750d20998bd9bc58b24ba11b94459928db9623590d22b6ddc.NewSecurityActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SubjectRightsRequests the subjectRightsRequests property
func (m *SecurityRequestBuilder) SubjectRightsRequests()(*i29dbe885802105b8f24e18a3ad45a4baf47a032bed4a97cc90638ac720b35332.SubjectRightsRequestsRequestBuilder) {
    return i29dbe885802105b8f24e18a3ad45a4baf47a032bed4a97cc90638ac720b35332.NewSubjectRightsRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubjectRightsRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.subjectRightsRequests.item collection
func (m *SecurityRequestBuilder) SubjectRightsRequestsById(id string)(*i9fe89cbebebc9e8421ea4e1df57c4ac03af0de51d9dea1f1d0ba1d6023c52df0.SubjectRightsRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subjectRightsRequest%2Did"] = id
    }
    return i9fe89cbebebc9e8421ea4e1df57c4ac03af0de51d9dea1f1d0ba1d6023c52df0.NewSubjectRightsRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ThreatSubmission the threatSubmission property
func (m *SecurityRequestBuilder) ThreatSubmission()(*i0df2000278766f4e13bb6be9bde05e5eb0c186391a4a2cb9f634d92117307226.ThreatSubmissionRequestBuilder) {
    return i0df2000278766f4e13bb6be9bde05e5eb0c186391a4a2cb9f634d92117307226.NewThreatSubmissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TiIndicators the tiIndicators property
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
        urlTplParams["tiIndicator%2Did"] = id
    }
    return i5c9089aa1bb65758fcfbff1d8a7aa9b5deeb5741c421c67322898649fd582848.NewTiIndicatorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Triggers the triggers property
func (m *SecurityRequestBuilder) Triggers()(*i9d73acd74b99f8c28d8e6ba0f549aeb25b13d09f6e92903e09e541ed58de8218.TriggersRequestBuilder) {
    return i9d73acd74b99f8c28d8e6ba0f549aeb25b13d09f6e92903e09e541ed58de8218.NewTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerTypes the triggerTypes property
func (m *SecurityRequestBuilder) TriggerTypes()(*i0c676084c053b74ed3a3133cb01b5e628a094e0413da8590be1ed0afd08d166f.TriggerTypesRequestBuilder) {
    return i0c676084c053b74ed3a3133cb01b5e628a094e0413da8590be1ed0afd08d166f.NewTriggerTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSecurityProfiles the userSecurityProfiles property
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
        urlTplParams["userSecurityProfile%2Did"] = id
    }
    return i9b85a1e7bdfc97cbdd52f3319eecf44724fb90aed32a007e01bbf117230a1cbd.NewUserSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
