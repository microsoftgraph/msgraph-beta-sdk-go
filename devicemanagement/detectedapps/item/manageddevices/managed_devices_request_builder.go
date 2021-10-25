package manageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/executeaction"
    i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/bulkreprovisioncloudpc"
    idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/ref"
)

type ManagedDevicesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedDevicesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27.BulkReprovisionCloudPcRequestBuilder) {
    return i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/detectedApps/{detectedApp_id}/managedDevices{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformation(q func (value *ManagedDevicesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedDevicesRequestBuilderGetQueryParameters)
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
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951.ExecuteActionRequestBuilder) {
    return i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) Get(q func (value *ManagedDevicesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ManagedDevicesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevicesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedDevicesResponse), nil
}
func (m *ManagedDevicesRequestBuilder) Ref()(*idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8.RefRequestBuilder) {
    return idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
