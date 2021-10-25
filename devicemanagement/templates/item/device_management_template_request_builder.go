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

type DeviceManagementTemplateRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceManagementTemplateRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *DeviceManagementTemplateRequestBuilder) Categories()(*i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.CategoriesRequestBuilder) {
    return i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) CategoriesById(id string)(*i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.DeviceManagementTemplateSettingCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory_id"] = id
    }
    return i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.NewDeviceManagementTemplateSettingCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) CompareWithTemplateId(templateId *string)(*ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.CompareWithTemplateIdRequestBuilder) {
    return ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
func NewDeviceManagementTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateRequestBuilder) {
    m := &DeviceManagementTemplateRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/templates/{deviceManagementTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDeviceManagementTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementTemplateRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceManagementTemplateRequestBuilder) CreateGetRequestInformation(q func (value *DeviceManagementTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceManagementTemplateRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceManagementTemplateRequestBuilder) CreateInstance()(*i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.CreateInstanceRequestBuilder) {
    return i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.NewCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DeviceManagementTemplateRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementTemplateRequestBuilder) Get(q func (value *DeviceManagementTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate), nil
}
func (m *DeviceManagementTemplateRequestBuilder) MigratableTo()(*i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.MigratableToRequestBuilder) {
    return i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.NewMigratableToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) MigratableToById(id string)(*ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.DeviceManagementTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate_id1"] = id
    }
    return ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.NewDeviceManagementTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementTemplateRequestBuilder) Settings()(*iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.SettingsRequestBuilder) {
    return iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementTemplateRequestBuilder) SettingsById(id string)(*i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.DeviceManagementSettingInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.NewDeviceManagementSettingInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
