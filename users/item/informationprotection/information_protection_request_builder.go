package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04cf05c5c8f994ff60f069fa17365461a81725d5b81546944542ea7ba50c33b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/encryptbuffer"
    i05a0f6d03a62687d4a0ab7a6623ce716d5ec495078059f7cb0b8e6cea75698b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitylabels"
    i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitypolicysettings"
    i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/bitlocker"
    i659c82714306f456ed196c36c3db55f7f532d452b83e200bc4c9fac27f759b58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/verifysignature"
    ia4c05b91b98e220a41cbdaa681d4415ce4d6cdbef9570b33a0dafc5be44f6bac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/decryptbuffer"
    ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy"
    ic3c846ff56dc5e314d4c6bfa19cb00924d82997eea901fe1f9b84c1d80e78428 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/datalosspreventionpolicies"
    idbc0c28ddaa866046df733d92dac03d4c4f02bc523555cc4a6b74a509ee1b65c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/threatassessmentrequests"
    ie9793034cd3f29fae723f1f5ff28ff39cb87eb37c66f6795183c08da0ede1ade "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/signdigest"
    i191c8b3440dd1e8a6718ed78cfecae17bc1dbfd3c34dd568e3ee66f765bbfcef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/datalosspreventionpolicies/item"
    i881aabb372d33b8637b6e47f5cd810fa059621fecdca425af680154915bda58c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitylabels/item"
    icaf3ff4e247245a4c08bfd8cebd5e2ed55e3a3ea59bac39141569435b536acf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/threatassessmentrequests/item"
)

// InformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.user entity.
type InformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InformationProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InformationProtectionRequestBuilderGetQueryParameters get informationProtection from users
type InformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InformationProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InformationProtectionRequestBuilderGetQueryParameters
}
// InformationProtectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bitlocker the bitlocker property
func (m *InformationProtectionRequestBuilder) Bitlocker()(*i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08.BitlockerRequestBuilder) {
    return i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08.NewBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for users
func (m *InformationProtectionRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property informationProtection for users
func (m *InformationProtectionRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *InformationProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get informationProtection from users
func (m *InformationProtectionRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get informationProtection from users
func (m *InformationProtectionRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InformationProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property informationProtection in users
func (m *InformationProtectionRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property informationProtection in users
func (m *InformationProtectionRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *InformationProtectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataLossPreventionPolicies the dataLossPreventionPolicies property
func (m *InformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ic3c846ff56dc5e314d4c6bfa19cb00924d82997eea901fe1f9b84c1d80e78428.DataLossPreventionPoliciesRequestBuilder) {
    return ic3c846ff56dc5e314d4c6bfa19cb00924d82997eea901fe1f9b84c1d80e78428.NewDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.informationProtection.dataLossPreventionPolicies.item collection
func (m *InformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*i191c8b3440dd1e8a6718ed78cfecae17bc1dbfd3c34dd568e3ee66f765bbfcef.DataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return i191c8b3440dd1e8a6718ed78cfecae17bc1dbfd3c34dd568e3ee66f765bbfcef.NewDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DecryptBuffer the decryptBuffer property
func (m *InformationProtectionRequestBuilder) DecryptBuffer()(*ia4c05b91b98e220a41cbdaa681d4415ce4d6cdbef9570b33a0dafc5be44f6bac.DecryptBufferRequestBuilder) {
    return ia4c05b91b98e220a41cbdaa681d4415ce4d6cdbef9570b33a0dafc5be44f6bac.NewDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property informationProtection for users
func (m *InformationProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *InformationProtectionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EncryptBuffer the encryptBuffer property
func (m *InformationProtectionRequestBuilder) EncryptBuffer()(*i04cf05c5c8f994ff60f069fa17365461a81725d5b81546944542ea7ba50c33b9.EncryptBufferRequestBuilder) {
    return i04cf05c5c8f994ff60f069fa17365461a81725d5b81546944542ea7ba50c33b9.NewEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection from users
func (m *InformationProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *InformationProtectionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in users
func (m *InformationProtectionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, requestConfiguration *InformationProtectionRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Policy the policy property
func (m *InformationProtectionRequestBuilder) Policy()(*ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389.PolicyRequestBuilder) {
    return ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels the sensitivityLabels property
func (m *InformationProtectionRequestBuilder) SensitivityLabels()(*i05a0f6d03a62687d4a0ab7a6623ce716d5ec495078059f7cb0b8e6cea75698b8.SensitivityLabelsRequestBuilder) {
    return i05a0f6d03a62687d4a0ab7a6623ce716d5ec495078059f7cb0b8e6cea75698b8.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.informationProtection.sensitivityLabels.item collection
func (m *InformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*i881aabb372d33b8637b6e47f5cd810fa059621fecdca425af680154915bda58c.SensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return i881aabb372d33b8637b6e47f5cd810fa059621fecdca425af680154915bda58c.NewSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings the sensitivityPolicySettings property
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c.SensitivityPolicySettingsRequestBuilder) {
    return i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c.NewSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignDigest the signDigest property
func (m *InformationProtectionRequestBuilder) SignDigest()(*ie9793034cd3f29fae723f1f5ff28ff39cb87eb37c66f6795183c08da0ede1ade.SignDigestRequestBuilder) {
    return ie9793034cd3f29fae723f1f5ff28ff39cb87eb37c66f6795183c08da0ede1ade.NewSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests the threatAssessmentRequests property
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequests()(*idbc0c28ddaa866046df733d92dac03d4c4f02bc523555cc4a6b74a509ee1b65c.ThreatAssessmentRequestsRequestBuilder) {
    return idbc0c28ddaa866046df733d92dac03d4c4f02bc523555cc4a6b74a509ee1b65c.NewThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.informationProtection.threatAssessmentRequests.item collection
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*icaf3ff4e247245a4c08bfd8cebd5e2ed55e3a3ea59bac39141569435b536acf5.ThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return icaf3ff4e247245a4c08bfd8cebd5e2ed55e3a3ea59bac39141569435b536acf5.NewThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifySignature the verifySignature property
func (m *InformationProtectionRequestBuilder) VerifySignature()(*i659c82714306f456ed196c36c3db55f7f532d452b83e200bc4c9fac27f759b58.VerifySignatureRequestBuilder) {
    return i659c82714306f456ed196c36c3db55f7f532d452b83e200bc4c9fac27f759b58.NewVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
