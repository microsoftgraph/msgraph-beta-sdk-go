package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assign"
    i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/categories"
    ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/devicestates"
    iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assignments"
    iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary"
    ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/installsummary"
    i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/devicestates/item"
    i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary/item"
    i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assignments/item"
)

type ManagedEBookRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedEBookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ManagedEBookRequestBuilder) Assign()(*i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1.AssignRequestBuilder) {
    return i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) Assignments()(*iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006.AssignmentsRequestBuilder) {
    return iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) AssignmentsById(id string)(*i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b.ManagedEBookAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment_id"] = id
    }
    return i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b.NewManagedEBookAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) Categories()(*i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738.CategoriesRequestBuilder) {
    return i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewManagedEBookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookRequestBuilder) {
    m := &ManagedEBookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/managedEBooks/{managedEBook_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedEBookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedEBookRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedEBookRequestBuilder) CreateGetRequestInformation(q func (value *ManagedEBookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedEBookRequestBuilderGetQueryParameters)
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
func (m *ManagedEBookRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedEBook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedEBookRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedEBookRequestBuilder) DeviceStates()(*ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6.DeviceStatesRequestBuilder) {
    return ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) DeviceStatesById(id string)(*i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256.DeviceInstallStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState_id"] = id
    }
    return i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256.NewDeviceInstallStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) Get(q func (value *ManagedEBookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedEBook, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedEBook() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedEBook), nil
}
func (m *ManagedEBookRequestBuilder) InstallSummary()(*ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403.InstallSummaryRequestBuilder) {
    return ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedEBook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedEBookRequestBuilder) UserStateSummary()(*iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314.UserStateSummaryRequestBuilder) {
    return iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) UserStateSummaryById(id string)(*i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe.UserInstallStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary_id"] = id
    }
    return i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe.NewUserInstallStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
