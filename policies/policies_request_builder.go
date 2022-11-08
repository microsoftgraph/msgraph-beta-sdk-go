package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i01d7860ea198dbe947e5a97722c9ca4edec8fd26cab7760478431192c3e470ea "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/deviceregistrationpolicy"
    i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/featurerolloutpolicies"
    i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/b2cauthenticationmethodspolicy"
    i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/identitysecuritydefaultsenforcementpolicy"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicyassignments"
    i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationflowspolicy"
    i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationmethodspolicy"
    i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/adminconsentrequestpolicy"
    i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies"
    i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies"
    i6390b41fee35d38ad58432cb872f7b0981fd1b58b2e222bdf8c0d929c212d726 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationstrengthpolicies"
    i63cc74da591a22ec8c6bde2e4173c23db89b2ae47fef9d9e7c5634f87419136f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies"
    i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/claimsmappingpolicies"
    i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/accessreviewpolicy"
    i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenissuancepolicies"
    i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/directoryroleaccessreviewpolicy"
    ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/conditionalaccesspolicies"
    ia66b26489abda0f16a4552e41f842faa0920a32a210488eb7684d2c0a3ac678e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/defaultappmanagementpolicy"
    ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobiledevicemanagementpolicies"
    iabb1075117ee91d4cff52c295dea39fce30d32c8e7045b3aba6ab87f3e634321 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/externalidentitiespolicy"
    ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/activitybasedtimeoutpolicies"
    ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenlifetimepolicies"
    icb9925789218a9f15cfe348750b0959961d5079e5d10f9ce6723285bcf81edfe "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/appmanagementpolicies"
    idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy"
    iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/homerealmdiscoverypolicies"
    if3614f47055e200240fcbfc6301166529ebd3a0ae749e2a305e4126a286e6c20 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy"
    ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies"
    i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenlifetimepolicies/item"
    i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item"
    i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicyassignments/item"
    i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/claimsmappingpolicies/item"
    i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenissuancepolicies/item"
    i48ae16125bc24acea72c4577c552f952c9f32fec09f609757df4b612154ccbef "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationstrengthpolicies/item"
    i5059ef490cc51ecd77364347b98e03ea1c34501de1dede44e9b212932f620cbc "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy/item"
    i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/featurerolloutpolicies/item"
    i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobiledevicemanagementpolicies/item"
    iadb7d5b1d89c4738d0e9ddc6026962b7c676342a12ca09eb194e68257d3c7eae "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/appmanagementpolicies/item"
    ib3f60717aa9a4ac09e9bb8a247a048202cb7a0c0e5abad3e0036a87c95f673ae "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item"
    ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/homerealmdiscoverypolicies/item"
    ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies/item"
    ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/activitybasedtimeoutpolicies/item"
    if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/conditionalaccesspolicies/item"
    if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item"
)

// PoliciesRequestBuilder provides operations to manage the policyRoot singleton.
type PoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesRequestBuilderGetQueryParameters get policies
type PoliciesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesRequestBuilderGetQueryParameters
}
// PoliciesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewPolicy provides operations to manage the accessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AccessReviewPolicy()(*i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.AccessReviewPolicyRequestBuilder) {
    return i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.NewAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPolicies provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.ActivityBasedTimeoutPoliciesRequestBuilder) {
    return ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.NewActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPoliciesById provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPoliciesById(id string)(*ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87.ActivityBasedTimeoutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["activityBasedTimeoutPolicy%2Did"] = id
    }
    return ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87.NewActivityBasedTimeoutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AdminConsentRequestPolicy provides operations to manage the adminConsentRequestPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.AdminConsentRequestPolicyRequestBuilder) {
    return i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPolicies()(*icb9925789218a9f15cfe348750b0959961d5079e5d10f9ce6723285bcf81edfe.AppManagementPoliciesRequestBuilder) {
    return icb9925789218a9f15cfe348750b0959961d5079e5d10f9ce6723285bcf81edfe.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPoliciesById(id string)(*iadb7d5b1d89c4738d0e9ddc6026962b7c676342a12ca09eb194e68257d3c7eae.AppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appManagementPolicy%2Did"] = id
    }
    return iadb7d5b1d89c4738d0e9ddc6026962b7c676342a12ca09eb194e68257d3c7eae.NewAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationFlowsPolicy provides operations to manage the authenticationFlowsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.AuthenticationFlowsPolicyRequestBuilder) {
    return i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.NewAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.AuthenticationMethodsPolicyRequestBuilder) {
    return i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPolicies provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPolicies()(*i6390b41fee35d38ad58432cb872f7b0981fd1b58b2e222bdf8c0d929c212d726.AuthenticationStrengthPoliciesRequestBuilder) {
    return i6390b41fee35d38ad58432cb872f7b0981fd1b58b2e222bdf8c0d929c212d726.NewAuthenticationStrengthPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPoliciesById provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPoliciesById(id string)(*i48ae16125bc24acea72c4577c552f952c9f32fec09f609757df4b612154ccbef.AuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = id
    }
    return i48ae16125bc24acea72c4577c552f952c9f32fec09f609757df4b612154ccbef.NewAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthorizationPolicy provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.AuthorizationPolicyRequestBuilder) {
    return idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicyById provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicyById(id string)(*i5059ef490cc51ecd77364347b98e03ea1c34501de1dede44e9b212932f620cbc.AuthorizationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authorizationPolicy%2Did"] = id
    }
    return i5059ef490cc51ecd77364347b98e03ea1c34501de1dede44e9b212932f620cbc.NewAuthorizationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2cAuthenticationMethodsPolicy provides operations to manage the b2cAuthenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) B2cAuthenticationMethodsPolicy()(*i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.B2cAuthenticationMethodsPolicyRequestBuilder) {
    return i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.NewB2cAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.ClaimsMappingPoliciesRequestBuilder) {
    return i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPoliciesById(id string)(*i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100.ClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100.NewClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicies provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.ConditionalAccessPoliciesRequestBuilder) {
    return ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPoliciesById provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPoliciesById(id string)(*if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb.ConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb.NewConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPoliciesRequestBuilderInternal instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    m := &PoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesRequestBuilder instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get policies
func (m *PoliciesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update policies
func (m *PoliciesRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CrossTenantAccessPolicy provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) CrossTenantAccessPolicy()(*if3614f47055e200240fcbfc6301166529ebd3a0ae749e2a305e4126a286e6c20.CrossTenantAccessPolicyRequestBuilder) {
    return if3614f47055e200240fcbfc6301166529ebd3a0ae749e2a305e4126a286e6c20.NewCrossTenantAccessPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultAppManagementPolicy provides operations to manage the defaultAppManagementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DefaultAppManagementPolicy()(*ia66b26489abda0f16a4552e41f842faa0920a32a210488eb7684d2c0a3ac678e.DefaultAppManagementPolicyRequestBuilder) {
    return ia66b26489abda0f16a4552e41f842faa0920a32a210488eb7684d2c0a3ac678e.NewDefaultAppManagementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRegistrationPolicy provides operations to manage the deviceRegistrationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DeviceRegistrationPolicy()(*i01d7860ea198dbe947e5a97722c9ca4edec8fd26cab7760478431192c3e470ea.DeviceRegistrationPolicyRequestBuilder) {
    return i01d7860ea198dbe947e5a97722c9ca4edec8fd26cab7760478431192c3e470ea.NewDeviceRegistrationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleAccessReviewPolicy provides operations to manage the directoryRoleAccessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DirectoryRoleAccessReviewPolicy()(*i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.DirectoryRoleAccessReviewPolicyRequestBuilder) {
    return i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.NewDirectoryRoleAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalIdentitiesPolicy provides operations to manage the externalIdentitiesPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ExternalIdentitiesPolicy()(*iabb1075117ee91d4cff52c295dea39fce30d32c8e7045b3aba6ab87f3e634321.ExternalIdentitiesPolicyRequestBuilder) {
    return iabb1075117ee91d4cff52c295dea39fce30d32c8e7045b3aba6ab87f3e634321.NewExternalIdentitiesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPolicies provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.FeatureRolloutPoliciesRequestBuilder) {
    return i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPoliciesById(id string)(*i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28.FeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = id
    }
    return i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28.NewFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get policies
func (m *PoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945.HomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945.NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IdentitySecurityDefaultsEnforcementPolicy provides operations to manage the identitySecurityDefaultsEnforcementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPolicies provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPolicies()(*ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.MobileAppManagementPoliciesRequestBuilder) {
    return ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.NewMobileAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPoliciesById provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPoliciesById(id string)(*ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011.MobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011.NewMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileDeviceManagementPolicies provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPolicies()(*ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.MobileDeviceManagementPoliciesRequestBuilder) {
    return ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.NewMobileDeviceManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileDeviceManagementPoliciesById provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPoliciesById(id string)(*i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302.MobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302.NewMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update policies
func (m *PoliciesRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// PermissionGrantPolicies provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.PermissionGrantPoliciesRequestBuilder) {
    return i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPoliciesById provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPoliciesById(id string)(*i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa.PermissionGrantPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantPolicy%2Did"] = id
    }
    return i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa.NewPermissionGrantPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicies provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicies()(*i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.RoleManagementPoliciesRequestBuilder) {
    return i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.NewRoleManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPoliciesById provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPoliciesById(id string)(*if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54.UnifiedRoleManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicy%2Did"] = id
    }
    return if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54.NewUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicyAssignments provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignments()(*i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.RoleManagementPolicyAssignmentsRequestBuilder) {
    return i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.NewRoleManagementPolicyAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPolicyAssignmentsById provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignmentsById(id string)(*i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585.UnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyAssignment%2Did"] = id
    }
    return i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585.NewUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ServicePrincipalCreationPolicies provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPolicies()(*i63cc74da591a22ec8c6bde2e4173c23db89b2ae47fef9d9e7c5634f87419136f.ServicePrincipalCreationPoliciesRequestBuilder) {
    return i63cc74da591a22ec8c6bde2e4173c23db89b2ae47fef9d9e7c5634f87419136f.NewServicePrincipalCreationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalCreationPoliciesById provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPoliciesById(id string)(*ib3f60717aa9a4ac09e9bb8a247a048202cb7a0c0e5abad3e0036a87c95f673ae.ServicePrincipalCreationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationPolicy%2Did"] = id
    }
    return ib3f60717aa9a4ac09e9bb8a247a048202cb7a0c0e5abad3e0036a87c95f673ae.NewServicePrincipalCreationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.TokenIssuancePoliciesRequestBuilder) {
    return i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePoliciesById(id string)(*i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9.TokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9.NewTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.TokenLifetimePoliciesRequestBuilder) {
    return ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePoliciesById(id string)(*i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e.TokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e.NewTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
