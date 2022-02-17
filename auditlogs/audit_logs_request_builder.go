package auditlogs

import (
    i61884affea02999b55c34273d0b26c7d1ae48580ccad62b3a4ecefd4ecca9be5 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/restrictedsignins"
    i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryprovisioning"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryaudits"
    ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/signins"
    ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/provisioning"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/signins/item"
    i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryprovisioning/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6cc8224410e5f63fa35653adcdb9e8272fa04027bd7951a66a001c4821dbc387 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/restrictedsignins/item"
    ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/provisioning/item"
    iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs/directoryaudits/item"
)

// AuditLogsRequestBuilder builds and executes requests for operations under \auditLogs
type AuditLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuditLogsRequestBuilderGetOptions options for Get
type AuditLogsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuditLogsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuditLogsRequestBuilderGetQueryParameters get auditLogs
type AuditLogsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuditLogsRequestBuilderPatchOptions options for Patch
type AuditLogsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditLogRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuditLogsRequestBuilderInternal instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/auditLogs{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get auditLogs
func (m *AuditLogsRequestBuilder) CreateGetRequestInformation(options *AuditLogsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update auditLogs
func (m *AuditLogsRequestBuilder) CreatePatchRequestInformation(options *AuditLogsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3.DirectoryAuditsRequestBuilder) {
    return iddb9d350ed2f2204a816a8b2c9737202cbbb609b2b77635734c77c461c7b11b3.NewDirectoryAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryAuditsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.auditLogs.directoryAudits.item collection
func (m *AuditLogsRequestBuilder) DirectoryAuditsById(id string)(*iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3.DirectoryAuditRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryAudit_id"] = id
    }
    return iecdf88a9efb8a7e905d5b4313dc1bc508fc1f53e48ec460b752ecce4944a19d3.NewDirectoryAuditRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuditLogsRequestBuilder) DirectoryProvisioning()(*i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48.DirectoryProvisioningRequestBuilder) {
    return i897c3994bb157d73a049b5f8e090a75f6296cc505d281e096f9c2f5015269d48.NewDirectoryProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryProvisioningById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.auditLogs.directoryProvisioning.item collection
func (m *AuditLogsRequestBuilder) DirectoryProvisioningById(id string)(*i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5.ProvisioningObjectSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary_id"] = id
    }
    return i43b93442c076ce1053480c2aaba8f405e6875171980c0abd4c67e9a1a25b53e5.NewProvisioningObjectSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get auditLogs
func (m *AuditLogsRequestBuilder) Get(options *AuditLogsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditLogRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuditLogRoot() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuditLogRoot), nil
}
// Patch update auditLogs
func (m *AuditLogsRequestBuilder) Patch(options *AuditLogsRequestBuilderPatchOptions)(error) {
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
func (m *AuditLogsRequestBuilder) Provisioning()(*ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852.ProvisioningRequestBuilder) {
    return ie7cd14714146bead7b23538aa08e8ec73dbbd6fdac83c25da131496b868f0852.NewProvisioningRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.auditLogs.provisioning.item collection
func (m *AuditLogsRequestBuilder) ProvisioningById(id string)(*ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5.ProvisioningObjectSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["provisioningObjectSummary_id"] = id
    }
    return ida9b6a06c8d64bad7537948374efa593835455cc0fac39900b364fe69288a5b5.NewProvisioningObjectSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuditLogsRequestBuilder) RestrictedSignIns()(*i61884affea02999b55c34273d0b26c7d1ae48580ccad62b3a4ecefd4ecca9be5.RestrictedSignInsRequestBuilder) {
    return i61884affea02999b55c34273d0b26c7d1ae48580ccad62b3a4ecefd4ecca9be5.NewRestrictedSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestrictedSignInsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.auditLogs.restrictedSignIns.item collection
func (m *AuditLogsRequestBuilder) RestrictedSignInsById(id string)(*i6cc8224410e5f63fa35653adcdb9e8272fa04027bd7951a66a001c4821dbc387.RestrictedSignInRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["restrictedSignIn_id"] = id
    }
    return i6cc8224410e5f63fa35653adcdb9e8272fa04027bd7951a66a001c4821dbc387.NewRestrictedSignInRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuditLogsRequestBuilder) SignIns()(*ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7.SignInsRequestBuilder) {
    return ie1f17e2e55a9a96b4c8eaf8d9efd0982978d50135d920222ccacbe442d0074b7.NewSignInsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SignInsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.auditLogs.signIns.item collection
func (m *AuditLogsRequestBuilder) SignInsById(id string)(*i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187.SignInRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["signIn_id"] = id
    }
    return i18c1f0dab1391e55e1b65be523640813575f4930aeeb8f410364c82046e73187.NewSignInRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
