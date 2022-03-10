package informationprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i05a0f6d03a62687d4a0ab7a6623ce716d5ec495078059f7cb0b8e6cea75698b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitylabels"
    i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitypolicysettings"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/bitlocker"
    ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/policy"
    ic3c846ff56dc5e314d4c6bfa19cb00924d82997eea901fe1f9b84c1d80e78428 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/datalosspreventionpolicies"
    idbc0c28ddaa866046df733d92dac03d4c4f02bc523555cc4a6b74a509ee1b65c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/threatassessmentrequests"
    i191c8b3440dd1e8a6718ed78cfecae17bc1dbfd3c34dd568e3ee66f765bbfcef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/datalosspreventionpolicies/item"
    i881aabb372d33b8637b6e47f5cd810fa059621fecdca425af680154915bda58c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/sensitivitylabels/item"
    icaf3ff4e247245a4c08bfd8cebd5e2ed55e3a3ea59bac39141569435b536acf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/threatassessmentrequests/item"
)

// InformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.user entity.
type InformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InformationProtectionRequestBuilderDeleteOptions options for Delete
type InformationProtectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
// InformationProtectionRequestBuilderGetQueryParameters get informationProtection from users
type InformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InformationProtectionRequestBuilderPatchOptions options for Patch
type InformationProtectionRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *InformationProtectionRequestBuilder) Bitlocker()(*i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08.BitlockerRequestBuilder) {
    return i4d90195489ca45be2307fbe69c3159942127422787130dc6bf07c3af7c022a08.NewBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/informationProtection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for users
func (m *InformationProtectionRequestBuilder) CreateDeleteRequestInformation(options *InformationProtectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get informationProtection from users
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
// CreatePatchRequestInformation update the navigation property informationProtection in users
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
        urlTplParams["dataLossPreventionPolicy_id"] = id
    }
    return i191c8b3440dd1e8a6718ed78cfecae17bc1dbfd3c34dd568e3ee66f765bbfcef.NewDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property informationProtection for users
func (m *InformationProtectionRequestBuilder) Delete(options *InformationProtectionRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
// Get get informationProtection from users
func (m *InformationProtectionRequestBuilder) Get(options *InformationProtectionRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateInformationProtectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in users
func (m *InformationProtectionRequestBuilder) Patch(options *InformationProtectionRequestBuilderPatchOptions)(error) {
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
func (m *InformationProtectionRequestBuilder) Policy()(*ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389.PolicyRequestBuilder) {
    return ia6cbc7bc1ed8597b7b0e86b2d192d20f7941e47862fb78aec8df090b439c4389.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["sensitivityLabel_id"] = id
    }
    return i881aabb372d33b8637b6e47f5cd810fa059621fecdca425af680154915bda58c.NewSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c.SensitivityPolicySettingsRequestBuilder) {
    return i124183ab6498588a7b75532486ec809ed07330838d143a4f82efefe532baad7c.NewSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["threatAssessmentRequest_id"] = id
    }
    return icaf3ff4e247245a4c08bfd8cebd5e2ed55e3a3ea59bac39141569435b536acf5.NewThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
