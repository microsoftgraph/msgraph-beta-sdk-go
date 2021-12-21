package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i745325089c8daf027097f621c96e1b7d845792fe5700c21d6cdb1fab8faebd98 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/runsummary"
    i76a58dc0ecd0613c93853cb6cc1cd3de1ec88851e809f9a201945659759ec324 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/assignments"
    i7f4b9bc8a9cb843e23253015b7a3880d3bc1299adceaf4be79a4dbd5eea94019 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/devicerunstates"
    ibbefa6569e41c8f73a9271dcdf516ffcb2ea4c4f1276219d1625ec944670bda5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/userrunstates"
    ic885200d8e720d935c1869a84197bf7854e2399b0092c0faf376ae616d483f73 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/assign"
    i5f96a372e4933548465bcb1bc6b021a80b1e97ec654f47f8a9a10940aef94e71 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/devicerunstates/item"
    ia3919e3a46ca0475695453048d8c26f33eefdae67b5ff84f9da11a5cd3938813 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/userrunstates/item"
    ib666025a5395af48f5ed203a2e914cb459b6384e11b373ad07a664bf9f3c0bfc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/hardwareconfigurations/item/assignments/item"
)

// HardwareConfigurationRequestBuilder builds and executes requests for operations under \deviceManagement\hardwareConfigurations\{hardwareConfiguration-id}
type HardwareConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// HardwareConfigurationRequestBuilderDeleteOptions options for Delete
type HardwareConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// HardwareConfigurationRequestBuilderGetOptions options for Get
type HardwareConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *HardwareConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// HardwareConfigurationRequestBuilderGetQueryParameters the hardware configurations for this account.
type HardwareConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// HardwareConfigurationRequestBuilderPatchOptions options for Patch
type HardwareConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HardwareConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *HardwareConfigurationRequestBuilder) Assign()(*ic885200d8e720d935c1869a84197bf7854e2399b0092c0faf376ae616d483f73.AssignRequestBuilder) {
    return ic885200d8e720d935c1869a84197bf7854e2399b0092c0faf376ae616d483f73.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *HardwareConfigurationRequestBuilder) Assignments()(*i76a58dc0ecd0613c93853cb6cc1cd3de1ec88851e809f9a201945659759ec324.AssignmentsRequestBuilder) {
    return i76a58dc0ecd0613c93853cb6cc1cd3de1ec88851e809f9a201945659759ec324.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.hardwareConfigurations.item.assignments.item collection
func (m *HardwareConfigurationRequestBuilder) AssignmentsById(id string)(*ib666025a5395af48f5ed203a2e914cb459b6384e11b373ad07a664bf9f3c0bfc.HardwareConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hardwareConfigurationAssignment_id"] = id
    }
    return ib666025a5395af48f5ed203a2e914cb459b6384e11b373ad07a664bf9f3c0bfc.NewHardwareConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewHardwareConfigurationRequestBuilderInternal instantiates a new HardwareConfigurationRequestBuilder and sets the default values.
func NewHardwareConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HardwareConfigurationRequestBuilder) {
    m := &HardwareConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHardwareConfigurationRequestBuilder instantiates a new HardwareConfigurationRequestBuilder and sets the default values.
func NewHardwareConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HardwareConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) CreateDeleteRequestInformation(options *HardwareConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) CreateGetRequestInformation(options *HardwareConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) CreatePatchRequestInformation(options *HardwareConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) Delete(options *HardwareConfigurationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *HardwareConfigurationRequestBuilder) DeviceRunStates()(*i7f4b9bc8a9cb843e23253015b7a3880d3bc1299adceaf4be79a4dbd5eea94019.DeviceRunStatesRequestBuilder) {
    return i7f4b9bc8a9cb843e23253015b7a3880d3bc1299adceaf4be79a4dbd5eea94019.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.hardwareConfigurations.item.deviceRunStates.item collection
func (m *HardwareConfigurationRequestBuilder) DeviceRunStatesById(id string)(*i5f96a372e4933548465bcb1bc6b021a80b1e97ec654f47f8a9a10940aef94e71.HardwareConfigurationDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hardwareConfigurationDeviceState_id"] = id
    }
    return i5f96a372e4933548465bcb1bc6b021a80b1e97ec654f47f8a9a10940aef94e71.NewHardwareConfigurationDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) Get(options *HardwareConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HardwareConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewHardwareConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HardwareConfiguration), nil
}
// Patch the hardware configurations for this account.
func (m *HardwareConfigurationRequestBuilder) Patch(options *HardwareConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *HardwareConfigurationRequestBuilder) RunSummary()(*i745325089c8daf027097f621c96e1b7d845792fe5700c21d6cdb1fab8faebd98.RunSummaryRequestBuilder) {
    return i745325089c8daf027097f621c96e1b7d845792fe5700c21d6cdb1fab8faebd98.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *HardwareConfigurationRequestBuilder) UserRunStates()(*ibbefa6569e41c8f73a9271dcdf516ffcb2ea4c4f1276219d1625ec944670bda5.UserRunStatesRequestBuilder) {
    return ibbefa6569e41c8f73a9271dcdf516ffcb2ea4c4f1276219d1625ec944670bda5.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.hardwareConfigurations.item.userRunStates.item collection
func (m *HardwareConfigurationRequestBuilder) UserRunStatesById(id string)(*ia3919e3a46ca0475695453048d8c26f33eefdae67b5ff84f9da11a5cd3938813.HardwareConfigurationUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hardwareConfigurationUserState_id"] = id
    }
    return ia3919e3a46ca0475695453048d8c26f33eefdae67b5ff84f9da11a5cd3938813.NewHardwareConfigurationUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
