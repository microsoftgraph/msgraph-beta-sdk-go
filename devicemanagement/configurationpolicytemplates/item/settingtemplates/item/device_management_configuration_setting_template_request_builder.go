package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4ef26cac8f2c79ce988999738305c17f201665e71cef01186dae6f9c2b7ccbb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates/item/settingdefinitions"
    i17e29a57a5f9a404c2017596b5ae58278006f69df8d598ba8b023689b417d3b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/configurationpolicytemplates/item/settingtemplates/item/settingdefinitions/item"
)

// DeviceManagementConfigurationSettingTemplateRequestBuilder builds and executes requests for operations under \deviceManagement\configurationPolicyTemplates\{deviceManagementConfigurationPolicyTemplate-id}\settingTemplates\{deviceManagementConfigurationSettingTemplate-id}
type DeviceManagementConfigurationSettingTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementConfigurationSettingTemplateRequestBuilderDeleteOptions options for Delete
type DeviceManagementConfigurationSettingTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationSettingTemplateRequestBuilderGetOptions options for Get
type DeviceManagementConfigurationSettingTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementConfigurationSettingTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementConfigurationSettingTemplateRequestBuilderGetQueryParameters setting templates
type DeviceManagementConfigurationSettingTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementConfigurationSettingTemplateRequestBuilderPatchOptions options for Patch
type DeviceManagementConfigurationSettingTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationSettingTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementConfigurationSettingTemplateRequestBuilderInternal instantiates a new DeviceManagementConfigurationSettingTemplateRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationSettingTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationSettingTemplateRequestBuilder) {
    m := &DeviceManagementConfigurationSettingTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate_id}/settingTemplates/{deviceManagementConfigurationSettingTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationSettingTemplateRequestBuilder instantiates a new DeviceManagementConfigurationSettingTemplateRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationSettingTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementConfigurationSettingTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationSettingTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementConfigurationSettingTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) CreateGetRequestInformation(options *DeviceManagementConfigurationSettingTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementConfigurationSettingTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) Delete(options *DeviceManagementConfigurationSettingTemplateRequestBuilderDeleteOptions)(error) {
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
// Get setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) Get(options *DeviceManagementConfigurationSettingTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationSettingTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementConfigurationSettingTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementConfigurationSettingTemplate), nil
}
// Patch setting templates
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) Patch(options *DeviceManagementConfigurationSettingTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) SettingDefinitions()(*i4ef26cac8f2c79ce988999738305c17f201665e71cef01186dae6f9c2b7ccbb7.SettingDefinitionsRequestBuilder) {
    return i4ef26cac8f2c79ce988999738305c17f201665e71cef01186dae6f9c2b7ccbb7.NewSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.configurationPolicyTemplates.item.settingTemplates.item.settingDefinitions.item collection
func (m *DeviceManagementConfigurationSettingTemplateRequestBuilder) SettingDefinitionsById(id string)(*i17e29a57a5f9a404c2017596b5ae58278006f69df8d598ba8b023689b417d3b1.DeviceManagementConfigurationSettingDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition_id"] = id
    }
    return i17e29a57a5f9a404c2017596b5ae58278006f69df8d598ba8b023689b417d3b1.NewDeviceManagementConfigurationSettingDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
