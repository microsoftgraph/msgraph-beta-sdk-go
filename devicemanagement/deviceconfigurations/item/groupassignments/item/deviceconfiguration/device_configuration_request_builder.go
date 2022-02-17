package deviceconfiguration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i35c4a1a5c1df086cb97731936969720d785a08966501a6658b453edcd93f3f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/windowsprivacyaccesscontrols"
    i44855f5003b56d3b7820710073dec689deaa89cd719504f222558d1f1a396d40 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/assign"
    i9856a553b21b6538b5d4e8e6ce1c593a1b6b97ed315493ea55a9e87170f8701d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/windowsupdateforbusinessconfiguration"
    i99d9d96dd906104453f5fcac96b966cbac48c3ec85216223332576a71f1435e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/ref"
    ic909bb3a3fa6c5e1f4b987a7351d244d5621bc505d8c39cd4581c7b95d57f933 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/assignedaccessmultimodeprofiles"
    ie355c76def8c28a4b0abbe9e22a10db802a2a49b124aaf4edea9807f125627a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/getomasettingplaintextvaluewithsecretreferencevalueid"
)

// DeviceConfigurationRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\groupAssignments\{deviceConfigurationGroupAssignment-id}\deviceConfiguration
type DeviceConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationRequestBuilderGetOptions options for Get
type DeviceConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationRequestBuilderGetQueryParameters the navigation link to the Device Configuration being targeted.
type DeviceConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *DeviceConfigurationRequestBuilder) Assign()(*i44855f5003b56d3b7820710073dec689deaa89cd719504f222558d1f1a396d40.AssignRequestBuilder) {
    return i44855f5003b56d3b7820710073dec689deaa89cd719504f222558d1f1a396d40.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) AssignedAccessMultiModeProfiles()(*ic909bb3a3fa6c5e1f4b987a7351d244d5621bc505d8c39cd4581c7b95d57f933.AssignedAccessMultiModeProfilesRequestBuilder) {
    return ic909bb3a3fa6c5e1f4b987a7351d244d5621bc505d8c39cd4581c7b95d57f933.NewAssignedAccessMultiModeProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceConfigurationRequestBuilderInternal instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    m := &DeviceConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/groupAssignments/{deviceConfigurationGroupAssignment_id}/deviceConfiguration{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationRequestBuilder instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationRequestBuilder) Get(options *DeviceConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration), nil
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\groupAssignments\{deviceConfigurationGroupAssignment-id}\deviceConfiguration\microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')
func (m *DeviceConfigurationRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*ie355c76def8c28a4b0abbe9e22a10db802a2a49b124aaf4edea9807f125627a9.GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return ie355c76def8c28a4b0abbe9e22a10db802a2a49b124aaf4edea9807f125627a9.NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
func (m *DeviceConfigurationRequestBuilder) Ref()(*i99d9d96dd906104453f5fcac96b966cbac48c3ec85216223332576a71f1435e8.RefRequestBuilder) {
    return i99d9d96dd906104453f5fcac96b966cbac48c3ec85216223332576a71f1435e8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) WindowsPrivacyAccessControls()(*i35c4a1a5c1df086cb97731936969720d785a08966501a6658b453edcd93f3f0f.WindowsPrivacyAccessControlsRequestBuilder) {
    return i35c4a1a5c1df086cb97731936969720d785a08966501a6658b453edcd93f3f0f.NewWindowsPrivacyAccessControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationRequestBuilder) WindowsUpdateForBusinessConfiguration()(*i9856a553b21b6538b5d4e8e6ce1c593a1b6b97ed315493ea55a9e87170f8701d.WindowsUpdateForBusinessConfigurationRequestBuilder) {
    return i9856a553b21b6538b5d4e8e6ce1c593a1b6b97ed315493ea55a9e87170f8701d.NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
