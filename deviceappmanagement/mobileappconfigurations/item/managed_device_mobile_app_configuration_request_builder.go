package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
    i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
)

type ManagedDeviceMobileAppConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assign()(*ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.AssignRequestBuilder) {
    return ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assignments()(*ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.AssignmentsRequestBuilder) {
    return ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) AssignmentsById(id string)(*i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.NewManagedDeviceMobileAppConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedDeviceMobileAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters)
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatuses()(*if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.DeviceStatusesRequestBuilder) {
    return if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusesById(id string)(*i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusSummary()(*i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.DeviceStatusSummaryRequestBuilder) {
    return i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Get(q func (value *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDeviceMobileAppConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration), nil
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatuses()(*i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.UserStatusesRequestBuilder) {
    return i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusesById(id string)(*i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusSummary()(*if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.UserStatusSummaryRequestBuilder) {
    return if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
