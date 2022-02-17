package informationprotection

import (
    i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/threatassessmentrequests"
    i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy"
    i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/datalosspreventionpolicies"
    i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitypolicysettings"
    i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/encryptbuffer"
    i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitylabels"
    i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/bitlocker"
    ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/verifysignature"
    ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/signdigest"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/decryptbuffer"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/threatassessmentrequests/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/datalosspreventionpolicies/item"
    ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/sensitivitylabels/item"
)

// InformationProtectionRequestBuilder builds and executes requests for operations under \informationProtection
type InformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InformationProtectionRequestBuilderGetOptions options for Get
type InformationProtectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InformationProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *InformationProtectionRequestBuilder) Bitlocker()(*i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad.BitlockerRequestBuilder) {
    return i5474609105224b94733e494f0a76d1ee2b2f78e3da0f7269cb4c4de7cb4835ad.NewBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get informationProtection
func (m *InformationProtectionRequestBuilder) CreateGetRequestInformation(options *InformationProtectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update informationProtection
func (m *InformationProtectionRequestBuilder) CreatePatchRequestInformation(options *InformationProtectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *InformationProtectionRequestBuilder) DataLossPreventionPolicies()(*i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524.DataLossPreventionPoliciesRequestBuilder) {
    return i1a95523e1b7c794d87cb658c728777e101fb0ee597c029847e1a346429f38524.NewDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.dataLossPreventionPolicies.item collection
func (m *InformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0.DataLossPreventionPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy_id"] = id
    }
    return i81855f40af60427eaa02288c5b3a8c7910b71b92f7590558464a7545621900e0.NewDataLossPreventionPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) DecryptBuffer()(*if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e.DecryptBufferRequestBuilder) {
    return if1d4de22a4e73e8600fbc436c47981b98858d08128f85fdfca3d0206acce2e0e.NewDecryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) EncryptBuffer()(*i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52.EncryptBufferRequestBuilder) {
    return i2661c5b59d8f82198e7fabe71500b75071f0700081cbe2409e33a8769eea2a52.NewEncryptBufferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get informationProtection
func (m *InformationProtectionRequestBuilder) Get(options *InformationProtectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInformationProtection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtection), nil
}
// Patch update informationProtection
func (m *InformationProtectionRequestBuilder) Patch(options *InformationProtectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *InformationProtectionRequestBuilder) Policy()(*i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6.PolicyRequestBuilder) {
    return i0bba7088fde5fe36bfa9328661fff812736b327f39cbf0cf119f8e1adeb31fb6.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) SensitivityLabels()(*i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1.SensitivityLabelsRequestBuilder) {
    return i4a2863a2873ddcee341eca5db8b2392a687082091015ea021687fd47d5cc62b1.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.sensitivityLabels.item collection
func (m *InformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd.SensitivityLabelRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel_id"] = id
    }
    return ia08232ff2c92a1e8e8fe0d1ca164e23faf909caa06499c80937301423e6eeecd.NewSensitivityLabelRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4.SensitivityPolicySettingsRequestBuilder) {
    return i1f1fc4efa4f7beefe6f28e0e47da69e7db2e007893565f6a9709f0136e0664b4.NewSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) SignDigest()(*ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43.SignDigestRequestBuilder) {
    return ice41f225667c8f03ae582d3f84fd885e36e168dd2e411baeff49b2bed2923f43.NewSignDigestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequests()(*i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147.ThreatAssessmentRequestsRequestBuilder) {
    return i01a4da13cab38ce326543975818120437c502ed331f722fb657d431f4668f147.NewThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.informationProtection.threatAssessmentRequests.item collection
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3.ThreatAssessmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest_id"] = id
    }
    return i327f9371a6a981bfca40891bd4712a18ae93dba5046bc3d5eb9091886aa46aa3.NewThreatAssessmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) VerifySignature()(*ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286.VerifySignatureRequestBuilder) {
    return ib1e5ef178341a9997924bbcf59d55402eb0c8d72a3fe1ae61ffc28668ff4e286.NewVerifySignatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
