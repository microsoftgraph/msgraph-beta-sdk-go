package identityprotection

import (
    i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskdetections"
    i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyusers"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskdetections/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyusers/item"
)

// Builds and executes requests for operations under \identityProtection
type IdentityProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type IdentityProtectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IdentityProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get identityProtection
type IdentityProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type IdentityProtectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityProtection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new IdentityProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIdentityProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    m := &IdentityProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new IdentityProtectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIdentityProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get identityProtection
// Parameters:
//  - options : Options for the request
func (m *IdentityProtectionRequestBuilder) CreateGetRequestInformation(options *IdentityProtectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update identityProtection
// Parameters:
//  - options : Options for the request
func (m *IdentityProtectionRequestBuilder) CreatePatchRequestInformation(options *IdentityProtectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get identityProtection
// Parameters:
//  - options : Options for the request
func (m *IdentityProtectionRequestBuilder) Get(options *IdentityProtectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityProtectionRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIdentityProtectionRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityProtectionRoot), nil
}
// Update identityProtection
// Parameters:
//  - options : Options for the request
func (m *IdentityProtectionRequestBuilder) Patch(options *IdentityProtectionRequestBuilderPatchOptions)(error) {
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
func (m *IdentityProtectionRequestBuilder) RiskDetections()(*i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4.RiskDetectionsRequestBuilder) {
    return i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4.NewRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskDetections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *IdentityProtectionRequestBuilder) RiskDetectionsById(id string)(*i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83.RiskDetectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskDetection_id"] = id
    }
    return i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83.NewRiskDetectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityProtectionRequestBuilder) RiskyUsers()(*i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989.RiskyUsersRequestBuilder) {
    return i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskyUsers.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *IdentityProtectionRequestBuilder) RiskyUsersById(id string)(*i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05.RiskyUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUser_id"] = id
    }
    return i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05.NewRiskyUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
