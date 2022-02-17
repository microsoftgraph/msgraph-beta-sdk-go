package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assignments"
    i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/settings"
    iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/createcopy"
    ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assign"
    i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/assignments/item"
    iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicies/item/settings/item"
)

// DeviceManagementConfigurationPolicyRequestBuilder builds and executes requests for operations under \deviceManagement\configurationPolicies\{deviceManagementConfigurationPolicy-id}
type DeviceManagementConfigurationPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters list of all Configuration policies
type DeviceManagementConfigurationPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementConfigurationPolicyRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Assign()(*ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170.AssignRequestBuilder) {
    return ibbfd5cc1d108b0a4813146b6ecc9e4823ba67ae5bf93fdd3fb64f99ca4e2d170.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Assignments()(*i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e.AssignmentsRequestBuilder) {
    return i69f80c23199c4402937b20220cf1e92e26db31e6a9495384f6071feff5fb666e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicies.item.assignments.item collection
func (m *DeviceManagementConfigurationPolicyRequestBuilder) AssignmentsById(id string)(*i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153.DeviceManagementConfigurationPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment_id"] = id
    }
    return i5e2155a9211be69588396fffd88611feb903891dc97905be90cf3ece71e3c153.NewDeviceManagementConfigurationPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementConfigurationPolicyRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyRequestBuilder instantiates a new DeviceManagementConfigurationPolicyRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateCopy()(*iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968.CreateCopyRequestBuilder) {
    return iab3edfc9e4992848bcd442122584eb8dcfbf0066b96fb9aee67506b106fea968.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Delete(options *DeviceManagementConfigurationPolicyRequestBuilderDeleteOptions)(error) {
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
// Get list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Get(options *DeviceManagementConfigurationPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementConfigurationPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicy), nil
}
// Patch list of all Configuration policies
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Patch(options *DeviceManagementConfigurationPolicyRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementConfigurationPolicyRequestBuilder) Settings()(*i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877.SettingsRequestBuilder) {
    return i9e07f7775b89cd0dc8237114af526cce0bc28e5356bbb8234a9c06ded2bd1877.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicies.item.settings.item collection
func (m *DeviceManagementConfigurationPolicyRequestBuilder) SettingsById(id string)(*iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba.DeviceManagementConfigurationSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting_id"] = id
    }
    return iedbbb6f68a2c0078d0b3190bbd5f0cd4f2d9a2f8875eead0f7a4e33156ffe5ba.NewDeviceManagementConfigurationSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
