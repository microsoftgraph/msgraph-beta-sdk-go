package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ic1d9737fc1c0a18008b6fd41b32eaa54f6eda77bc7590be0af2005ec52282707 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item/assignments"
    icaa99ae9f7101879792bea499febb85724e7e9b797da108c4987daca0e192244 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item/assign"
    ie9df73c2d706e7c3d9032c2ca1099f0d8f8ef487b1e09cfeee97d1301d4de667 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item/setpriority"
    i2743b28ceb0dd313f262858ca8786962b9d8aea0a4285b4dd4cb11ca6ed541da "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item/assignments/item"
)

// DeviceEnrollmentConfigurationRequestBuilder builds and executes requests for operations under \users\{user-id}\deviceEnrollmentConfigurations\{deviceEnrollmentConfiguration-id}
type DeviceEnrollmentConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceEnrollmentConfigurationRequestBuilderDeleteOptions options for Delete
type DeviceEnrollmentConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceEnrollmentConfigurationRequestBuilderGetOptions options for Get
type DeviceEnrollmentConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceEnrollmentConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceEnrollmentConfigurationRequestBuilderGetQueryParameters get enrollment configurations targeted to the user
type DeviceEnrollmentConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceEnrollmentConfigurationRequestBuilderPatchOptions options for Patch
type DeviceEnrollmentConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceEnrollmentConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceEnrollmentConfigurationRequestBuilder) Assign()(*icaa99ae9f7101879792bea499febb85724e7e9b797da108c4987daca0e192244.AssignRequestBuilder) {
    return icaa99ae9f7101879792bea499febb85724e7e9b797da108c4987daca0e192244.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceEnrollmentConfigurationRequestBuilder) Assignments()(*ic1d9737fc1c0a18008b6fd41b32eaa54f6eda77bc7590be0af2005ec52282707.AssignmentsRequestBuilder) {
    return ic1d9737fc1c0a18008b6fd41b32eaa54f6eda77bc7590be0af2005ec52282707.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.deviceEnrollmentConfigurations.item.assignments.item collection
func (m *DeviceEnrollmentConfigurationRequestBuilder) AssignmentsById(id string)(*i2743b28ceb0dd313f262858ca8786962b9d8aea0a4285b4dd4cb11ca6ed541da.EnrollmentConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentConfigurationAssignment_id"] = id
    }
    return i2743b28ceb0dd313f262858ca8786962b9d8aea0a4285b4dd4cb11ca6ed541da.NewEnrollmentConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceEnrollmentConfigurationRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationRequestBuilder) {
    m := &DeviceEnrollmentConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceEnrollmentConfigurationRequestBuilder instantiates a new DeviceEnrollmentConfigurationRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceEnrollmentConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreateDeleteRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreateGetRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) CreatePatchRequestInformation(options *DeviceEnrollmentConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) Delete(options *DeviceEnrollmentConfigurationRequestBuilderDeleteOptions)(error) {
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
// Get get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) Get(options *DeviceEnrollmentConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceEnrollmentConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceEnrollmentConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceEnrollmentConfiguration), nil
}
// Patch get enrollment configurations targeted to the user
func (m *DeviceEnrollmentConfigurationRequestBuilder) Patch(options *DeviceEnrollmentConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *DeviceEnrollmentConfigurationRequestBuilder) SetPriority()(*ie9df73c2d706e7c3d9032c2ca1099f0d8f8ef487b1e09cfeee97d1301d4de667.SetPriorityRequestBuilder) {
    return ie9df73c2d706e7c3d9032c2ca1099f0d8f8ef487b1e09cfeee97d1301d4de667.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
