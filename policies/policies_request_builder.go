package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
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
// AccessReviewPolicy the accessReviewPolicy property
func (m *PoliciesRequestBuilder) AccessReviewPolicy()(*i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.AccessReviewPolicyRequestBuilder) {
    return i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.NewAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPolicies the activityBasedTimeoutPolicies property
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.ActivityBasedTimeoutPoliciesRequestBuilder) {
    return ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.NewActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.activityBasedTimeoutPolicies.item collection
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
// AdminConsentRequestPolicy the adminConsentRequestPolicy property
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.AdminConsentRequestPolicyRequestBuilder) {
    return i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies the appManagementPolicies property
func (m *PoliciesRequestBuilder) AppManagementPolicies()(*icb9925789218a9f15cfe348750b0959961d5079e5d10f9ce6723285bcf81edfe.AppManagementPoliciesRequestBuilder) {
    return icb9925789218a9f15cfe348750b0959961d5079e5d10f9ce6723285bcf81edfe.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.appManagementPolicies.item collection
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
// AuthenticationFlowsPolicy the authenticationFlowsPolicy property
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.AuthenticationFlowsPolicyRequestBuilder) {
    return i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.NewAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodsPolicy the authenticationMethodsPolicy property
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.AuthenticationMethodsPolicyRequestBuilder) {
    return i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicy the authorizationPolicy property
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.AuthorizationPolicyRequestBuilder) {
    return idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicyById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.authorizationPolicy.item collection
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
// B2cAuthenticationMethodsPolicy the b2cAuthenticationMethodsPolicy property
func (m *PoliciesRequestBuilder) B2cAuthenticationMethodsPolicy()(*i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.B2cAuthenticationMethodsPolicyRequestBuilder) {
    return i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.NewB2cAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies the claimsMappingPolicies property
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.ClaimsMappingPoliciesRequestBuilder) {
    return i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.claimsMappingPolicies.item collection
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
// ConditionalAccessPolicies the conditionalAccessPolicies property
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.ConditionalAccessPoliciesRequestBuilder) {
    return ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.conditionalAccessPolicies.item collection
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
// CreateGetRequestInformationWithRequestConfiguration get policies
func (m *PoliciesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get policies
func (m *PoliciesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update policies
func (m *PoliciesRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update policies
func (m *PoliciesRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CrossTenantAccessPolicy the crossTenantAccessPolicy property
func (m *PoliciesRequestBuilder) CrossTenantAccessPolicy()(*if3614f47055e200240fcbfc6301166529ebd3a0ae749e2a305e4126a286e6c20.CrossTenantAccessPolicyRequestBuilder) {
    return if3614f47055e200240fcbfc6301166529ebd3a0ae749e2a305e4126a286e6c20.NewCrossTenantAccessPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultAppManagementPolicy the defaultAppManagementPolicy property
func (m *PoliciesRequestBuilder) DefaultAppManagementPolicy()(*ia66b26489abda0f16a4552e41f842faa0920a32a210488eb7684d2c0a3ac678e.DefaultAppManagementPolicyRequestBuilder) {
    return ia66b26489abda0f16a4552e41f842faa0920a32a210488eb7684d2c0a3ac678e.NewDefaultAppManagementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleAccessReviewPolicy the directoryRoleAccessReviewPolicy property
func (m *PoliciesRequestBuilder) DirectoryRoleAccessReviewPolicy()(*i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.DirectoryRoleAccessReviewPolicyRequestBuilder) {
    return i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.NewDirectoryRoleAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalIdentitiesPolicy the externalIdentitiesPolicy property
func (m *PoliciesRequestBuilder) ExternalIdentitiesPolicy()(*iabb1075117ee91d4cff52c295dea39fce30d32c8e7045b3aba6ab87f3e634321.ExternalIdentitiesPolicyRequestBuilder) {
    return iabb1075117ee91d4cff52c295dea39fce30d32c8e7045b3aba6ab87f3e634321.NewExternalIdentitiesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPolicies the featureRolloutPolicies property
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.FeatureRolloutPoliciesRequestBuilder) {
    return i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.featureRolloutPolicies.item collection
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
// GetWithResponseHandler get policies
func (m *PoliciesRequestBuilder) GetWithResponseHandler(requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get policies
func (m *PoliciesRequestBuilder) GetWithResponseHandler(requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// HomeRealmDiscoveryPolicies the homeRealmDiscoveryPolicies property
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.homeRealmDiscoveryPolicies.item collection
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
// IdentitySecurityDefaultsEnforcementPolicy the identitySecurityDefaultsEnforcementPolicy property
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPolicies the mobileAppManagementPolicies property
func (m *PoliciesRequestBuilder) MobileAppManagementPolicies()(*ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.MobileAppManagementPoliciesRequestBuilder) {
    return ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.NewMobileAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.mobileAppManagementPolicies.item collection
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
// MobileDeviceManagementPolicies the mobileDeviceManagementPolicies property
func (m *PoliciesRequestBuilder) MobileDeviceManagementPolicies()(*ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.MobileDeviceManagementPoliciesRequestBuilder) {
    return ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.NewMobileDeviceManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileDeviceManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.mobileDeviceManagementPolicies.item collection
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
// PatchWithResponseHandler update policies
func (m *PoliciesRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update policies
func (m *PoliciesRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PermissionGrantPolicies the permissionGrantPolicies property
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.PermissionGrantPoliciesRequestBuilder) {
    return i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.permissionGrantPolicies.item collection
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
// RoleManagementPolicies the roleManagementPolicies property
func (m *PoliciesRequestBuilder) RoleManagementPolicies()(*i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.RoleManagementPoliciesRequestBuilder) {
    return i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.NewRoleManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicies.item collection
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
// RoleManagementPolicyAssignments the roleManagementPolicyAssignments property
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignments()(*i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.RoleManagementPolicyAssignmentsRequestBuilder) {
    return i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.NewRoleManagementPolicyAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPolicyAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicyAssignments.item collection
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
// ServicePrincipalCreationPolicies the servicePrincipalCreationPolicies property
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPolicies()(*i63cc74da591a22ec8c6bde2e4173c23db89b2ae47fef9d9e7c5634f87419136f.ServicePrincipalCreationPoliciesRequestBuilder) {
    return i63cc74da591a22ec8c6bde2e4173c23db89b2ae47fef9d9e7c5634f87419136f.NewServicePrincipalCreationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalCreationPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.servicePrincipalCreationPolicies.item collection
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
// TokenIssuancePolicies the tokenIssuancePolicies property
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.TokenIssuancePoliciesRequestBuilder) {
    return i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.tokenIssuancePolicies.item collection
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
// TokenLifetimePolicies the tokenLifetimePolicies property
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.TokenLifetimePoliciesRequestBuilder) {
    return ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.tokenLifetimePolicies.item collection
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
