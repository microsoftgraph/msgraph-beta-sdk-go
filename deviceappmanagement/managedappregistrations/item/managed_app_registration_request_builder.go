package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies"
    id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations"
    if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies"
    i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item"
    i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations/item"
    id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item"
)

type ManagedAppRegistrationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedAppRegistrationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ManagedAppRegistrationRequestBuilder) AppliedPolicies()(*i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.AppliedPoliciesRequestBuilder) {
    return i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.NewAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) AppliedPoliciesById(id string)(*i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewManagedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    m := &ManagedAppRegistrationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedAppRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedAppRegistrationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedAppRegistrationRequestBuilder) CreateGetRequestInformation(q func (value *ManagedAppRegistrationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedAppRegistrationRequestBuilderGetQueryParameters)
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
func (m *ManagedAppRegistrationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedAppRegistrationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedAppRegistrationRequestBuilder) Get(q func (value *ManagedAppRegistrationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedAppRegistration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration), nil
}
func (m *ManagedAppRegistrationRequestBuilder) IntendedPolicies()(*if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.IntendedPoliciesRequestBuilder) {
    return if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.NewIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) IntendedPoliciesById(id string)(*id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) Operations()(*id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.OperationsRequestBuilder) {
    return id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) OperationsById(id string)(*i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.ManagedAppOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation_id"] = id
    }
    return i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.NewManagedAppOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
