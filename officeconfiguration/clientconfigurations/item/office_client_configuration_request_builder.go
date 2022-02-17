package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/policypayload"
    i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/userpreferencepayload"
    ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assign"
    ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments"
    ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments/item"
)

// OfficeClientConfigurationRequestBuilder builds and executes requests for operations under \officeConfiguration\clientConfigurations\{officeClientConfiguration-id}
type OfficeClientConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OfficeClientConfigurationRequestBuilderDeleteOptions options for Delete
type OfficeClientConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OfficeClientConfigurationRequestBuilderGetOptions options for Get
type OfficeClientConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OfficeClientConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OfficeClientConfigurationRequestBuilderGetQueryParameters list of office Client configuration.
type OfficeClientConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OfficeClientConfigurationRequestBuilderPatchOptions options for Patch
type OfficeClientConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OfficeClientConfigurationRequestBuilder) Assign()(*ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.AssignRequestBuilder) {
    return ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OfficeClientConfigurationRequestBuilder) Assignments()(*ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.AssignmentsRequestBuilder) {
    return ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.officeConfiguration.clientConfigurations.item.assignments.item collection
func (m *OfficeClientConfigurationRequestBuilder) AssignmentsById(id string)(*ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.OfficeClientConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfigurationAssignment_id"] = id
    }
    return ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.NewOfficeClientConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOfficeClientConfigurationRequestBuilderInternal instantiates a new OfficeClientConfigurationRequestBuilder and sets the default values.
func NewOfficeClientConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeClientConfigurationRequestBuilder) {
    m := &OfficeClientConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeClientConfigurationRequestBuilder instantiates a new OfficeClientConfigurationRequestBuilder and sets the default values.
func NewOfficeClientConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeClientConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeClientConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) CreateDeleteRequestInformation(options *OfficeClientConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) CreateGetRequestInformation(options *OfficeClientConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) CreatePatchRequestInformation(options *OfficeClientConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) Delete(options *OfficeClientConfigurationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) Get(options *OfficeClientConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOfficeClientConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration), nil
}
// Patch list of office Client configuration.
func (m *OfficeClientConfigurationRequestBuilder) Patch(options *OfficeClientConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *OfficeClientConfigurationRequestBuilder) PolicyPayload()(*i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.PolicyPayloadRequestBuilder) {
    return i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.NewPolicyPayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OfficeClientConfigurationRequestBuilder) UserPreferencePayload()(*i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.UserPreferencePayloadRequestBuilder) {
    return i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.NewUserPreferencePayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
