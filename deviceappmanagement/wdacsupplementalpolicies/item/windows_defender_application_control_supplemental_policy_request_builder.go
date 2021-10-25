package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses"
    i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/deploysummary"
    i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assign"
    idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments"
    ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments/item"
    if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses/item"
)

type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assign()(*i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.AssignRequestBuilder) {
    return i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assignments()(*idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.AssignmentsRequestBuilder) {
    return idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) AssignmentsById(id string)(*ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.WindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyAssignment_id"] = id
    }
    return ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy_id}{?select,expand}";
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
func NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreateGetRequestInformation(q func (value *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters)
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeploySummary()(*i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.DeploySummaryRequestBuilder) {
    return i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.NewDeploySummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeviceStatuses()(*i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.DeviceStatusesRequestBuilder) {
    return i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeviceStatusesById(id string)(*if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus_id"] = id
    }
    return if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Get(q func (value *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsDefenderApplicationControlSupplementalPolicy() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy), nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
