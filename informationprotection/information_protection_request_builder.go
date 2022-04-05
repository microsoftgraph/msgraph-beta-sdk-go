package informationprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/threatassessmentrequests"
    i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy"
    i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/datalosspreventionpolicies"
    i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitypolicysettings"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/encryptbuffer"
    i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitylabels"
    i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/bitlocker"
    ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/verifysignature"
    ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/signdigest"
    if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/decryptbuffer"
    i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/threatassessmentrequests/item"
    i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/datalosspreventionpolicies/item"
    ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitylabels/item"
)

// InformationProtectionRequestBuilder provides operations to manage the informationProtection singleton.
type InformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InformationProtectionRequestBuilderGetOptions options for Get
type InformationProtectionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *InformationProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// InformationProtectionRequestBuilderGetQueryParameters get informationProtection
type InformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InformationProtectionRequestBuilderPatchOptions options for Patch
type InformationProtectionRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Bitlocker the bitlocker property
func (m *InformationProtectionRequestBuilder) Bitlocker()(*i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad.BitlockerRequestBuilder) {
    return i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad.NewBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection{?select,expand}";
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
// CreateGetRequestInformation get informationProtection
func (m *InformationProtectionRequestBuilder) CreateGetRequestInformation(options *InformationProtectionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update informationProtection
func (m *InformationProtectionRequestBuilder) CreatePatchRequestInformation(options *InformationProtectionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// DataLossPreventionPolicies the dataLossPreventionPolicies property
func (m *InformationProtectionRequestBuilder) DataLossPreventionPolicies()(*i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524.DataLossPreventionPoliciesRequestBuilder) {
    return i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524.NewDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.dataLossPreventionPolicies.item collection
func (m *InformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0.DataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy_id"] = id
    }
    return i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0.NewDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DecryptBuffer the decryptBuffer property
func (m *InformationProtectionRequestBuilder) DecryptBuffer()(*if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e.DecryptBufferRequestBuilder) {
    return if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e.NewDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EncryptBuffer the encryptBuffer property
func (m *InformationProtectionRequestBuilder) EncryptBuffer()(*i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52.EncryptBufferRequestBuilder) {
    return i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52.NewEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection
func (m *InformationProtectionRequestBuilder) Get(options *InformationProtectionRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Patch update informationProtection
func (m *InformationProtectionRequestBuilder) Patch(options *InformationProtectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Policy the policy property
func (m *InformationProtectionRequestBuilder) Policy()(*i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6.PolicyRequestBuilder) {
    return i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels the sensitivityLabels property
func (m *InformationProtectionRequestBuilder) SensitivityLabels()(*i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1.SensitivityLabelsRequestBuilder) {
    return i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.sensitivityLabels.item collection
func (m *InformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd.SensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel_id"] = id
    }
    return ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd.NewSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings the sensitivityPolicySettings property
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4.SensitivityPolicySettingsRequestBuilder) {
    return i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4.NewSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignDigest the signDigest property
func (m *InformationProtectionRequestBuilder) SignDigest()(*ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43.SignDigestRequestBuilder) {
    return ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43.NewSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests the threatAssessmentRequests property
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequests()(*i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147.ThreatAssessmentRequestsRequestBuilder) {
    return i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147.NewThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.threatAssessmentRequests.item collection
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3.ThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest_id"] = id
    }
    return i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3.NewThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifySignature the verifySignature property
func (m *InformationProtectionRequestBuilder) VerifySignature()(*ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286.VerifySignatureRequestBuilder) {
    return ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286.NewVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
