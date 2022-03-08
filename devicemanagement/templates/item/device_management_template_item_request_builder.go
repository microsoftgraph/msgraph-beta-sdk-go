package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// DeviceManagementTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementTemplateItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateItemRequestBuilderGetOptions options for Get
type DeviceManagementTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateItemRequestBuilderGetQueryParameters the available templates
type DeviceManagementTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementTemplateItemRequestBuilderPatchOptions options for Patch
type DeviceManagementTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementTemplateItemRequestBuilder) Categories()(*i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.CategoriesRequestBuilder) {
    return i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.categories.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) CategoriesById(id string)(*i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.DeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory_id"] = id
    }
    return i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.NewDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementTemplateItemRequestBuilder) CompareWithTemplateId(templateId *string)(*ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.CompareWithTemplateIdRequestBuilder) {
    return ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementTemplateItemRequestBuilderInternal instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateItemRequestBuilder) {
    m := &DeviceManagementTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTemplateItemRequestBuilder instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property templates for deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementTemplateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementTemplateItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementTemplateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementTemplateItemRequestBuilder) CreateInstance()(*i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.CreateInstanceRequestBuilder) {
    return i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.NewCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementTemplateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property templates for deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) Delete(options *DeviceManagementTemplateItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the available templates
func (m *DeviceManagementTemplateItemRequestBuilder) Get(options *DeviceManagementTemplateItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceManagementTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateable), nil
}
func (m *DeviceManagementTemplateItemRequestBuilder) MigratableTo()(*i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.MigratableToRequestBuilder) {
    return i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.NewMigratableToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MigratableToById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.migratableTo.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) MigratableToById(id string)(*ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.DeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate_id1"] = id
    }
    return ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.NewDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) Patch(options *DeviceManagementTemplateItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementTemplateItemRequestBuilder) Settings()(*iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.SettingsRequestBuilder) {
    return iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.settings.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) SettingsById(id string)(*i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.DeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.NewDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
