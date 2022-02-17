package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates"
    ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates/item"
)

// DeviceManagementConfigurationPolicyTemplateRequestBuilder builds and executes requests for operations under \deviceManagement\configurationPolicyTemplates\{deviceManagementConfigurationPolicyTemplate-id}
type DeviceManagementConfigurationPolicyTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementConfigurationPolicyTemplateRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationPolicyTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyTemplateRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationPolicyTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementConfigurationPolicyTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationPolicyTemplateRequestBuilderGetQueryParameters list of all templates
type DeviceManagementConfigurationPolicyTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementConfigurationPolicyTemplateRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationPolicyTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementConfigurationPolicyTemplateRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyTemplateRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplateRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyTemplateRequestBuilder instantiates a new DeviceManagementConfigurationPolicyTemplateRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) Delete(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderDeleteOptions)(error) {
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
// Get list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) Get(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementConfigurationPolicyTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationPolicyTemplate), nil
}
// Patch list of all templates
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) Patch(options *DeviceManagementConfigurationPolicyTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) SettingTemplates()(*i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c.SettingTemplatesRequestBuilder) {
    return i3126bfb1a915f0c85e8910f85b7fbdc19b25d34462fa8ede5a13c297a190602c.NewSettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicyTemplates.item.settingTemplates.item collection
func (m *DeviceManagementConfigurationPolicyTemplateRequestBuilder) SettingTemplatesById(id string)(*ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1.DeviceManagementConfigurationSettingTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate_id"] = id
    }
    return ifeb77407d3839d95c54629e5156854062d7a5a2e76728f24033dce5b2d483dd1.NewDeviceManagementConfigurationSettingTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
