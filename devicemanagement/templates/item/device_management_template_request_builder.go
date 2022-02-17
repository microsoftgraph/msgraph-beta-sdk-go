package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories"
    i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/migratableto"
    i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/createinstance"
    iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/settings"
    ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/comparewithtemplateid"
    i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/settings/item"
    i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories/item"
    ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/migratableto/item"
)

// DeviceManagementTemplateRequestBuilder builds and executes requests for operations under \deviceManagement\templates\{deviceManagementTemplate-id}
type DeviceManagementTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementTemplateRequestBuilderDeleteOptions options for Delete
type DeviceManagementTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateRequestBuilderGetOptions options for Get
type DeviceManagementTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateRequestBuilderGetQueryParameters the available templates
type DeviceManagementTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementTemplateRequestBuilderPatchOptions options for Patch
type DeviceManagementTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementTemplateRequestBuilder) Categories()(*i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.CategoriesRequestBuilder) {
    return i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.categories.item collection
func (m *DeviceManagementTemplateRequestBuilder) CategoriesById(id string)(*i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.DeviceManagementTemplateSettingCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory_id"] = id
    }
    return i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.NewDeviceManagementTemplateSettingCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId builds and executes requests for operations under \deviceManagement\templates\{deviceManagementTemplate-id}\microsoft.graph.compare(templateId='{templateId}')
func (m *DeviceManagementTemplateRequestBuilder) CompareWithTemplateId(templateId *string)(*ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.CompareWithTemplateIdRequestBuilder) {
    return ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementTemplateRequestBuilderInternal instantiates a new DeviceManagementTemplateRequestBuilder and sets the default values.
func NewDeviceManagementTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateRequestBuilder) {
    m := &DeviceManagementTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTemplateRequestBuilder instantiates a new DeviceManagementTemplateRequestBuilder and sets the default values.
func NewDeviceManagementTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the available templates
func (m *DeviceManagementTemplateRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the available templates
func (m *DeviceManagementTemplateRequestBuilder) CreateGetRequestInformation(options *DeviceManagementTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementTemplateRequestBuilder) CreateInstance()(*i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.CreateInstanceRequestBuilder) {
    return i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.NewCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation the available templates
func (m *DeviceManagementTemplateRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the available templates
func (m *DeviceManagementTemplateRequestBuilder) Delete(options *DeviceManagementTemplateRequestBuilderDeleteOptions)(error) {
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
// Get the available templates
func (m *DeviceManagementTemplateRequestBuilder) Get(options *DeviceManagementTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate), nil
}
func (m *DeviceManagementTemplateRequestBuilder) MigratableTo()(*i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.MigratableToRequestBuilder) {
    return i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.NewMigratableToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MigratableToById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.migratableTo.item collection
func (m *DeviceManagementTemplateRequestBuilder) MigratableToById(id string)(*ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.DeviceManagementTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate_id1"] = id
    }
    return ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.NewDeviceManagementTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the available templates
func (m *DeviceManagementTemplateRequestBuilder) Patch(options *DeviceManagementTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementTemplateRequestBuilder) Settings()(*iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.SettingsRequestBuilder) {
    return iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.settings.item collection
func (m *DeviceManagementTemplateRequestBuilder) SettingsById(id string)(*i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.DeviceManagementSettingInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.NewDeviceManagementSettingInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
