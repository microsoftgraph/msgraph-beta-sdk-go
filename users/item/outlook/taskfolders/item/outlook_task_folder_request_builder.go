package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i70753d531cd2d03ce05924ac80789572a170a34271c5f1ff498ac69103e9ce6c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/multivalueextendedproperties"
    i86d3aaafac363883fd22a677c5fdcc68ca4a7ef0f722b9829c10a6200d62350f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks"
    if335ec0373b6fcbd74f5eb099d4683fc01fe9cf664574729fbb4fa9afe750ea6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/singlevalueextendedproperties"
    i12842b485e0c49d2f190adb9c34384e7ed531197aa90be3121b9db498d9492f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/singlevalueextendedproperties/item"
    i5fee4ecfccc04b9486e0e1107b1775092d66b5e99c175918cbc8476b8d1c0cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/multivalueextendedproperties/item"
    i70a06e760505c690b521534fb4682338b80fb764f653aa241eb813f4b9e367c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item"
)

type OutlookTaskFolderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OutlookTaskFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewOutlookTaskFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    m := &OutlookTaskFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/outlook/taskFolders/{outlookTaskFolder_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOutlookTaskFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OutlookTaskFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookTaskFolderRequestBuilder) CreateGetRequestInformation(q func (value *OutlookTaskFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OutlookTaskFolderRequestBuilderGetQueryParameters)
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
func (m *OutlookTaskFolderRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookTaskFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookTaskFolderRequestBuilder) Get(q func (value *OutlookTaskFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTaskFolder() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder), nil
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedProperties()(*i70753d531cd2d03ce05924ac80789572a170a34271c5f1ff498ac69103e9ce6c.MultiValueExtendedPropertiesRequestBuilder) {
    return i70753d531cd2d03ce05924ac80789572a170a34271c5f1ff498ac69103e9ce6c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5fee4ecfccc04b9486e0e1107b1775092d66b5e99c175918cbc8476b8d1c0cd7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5fee4ecfccc04b9486e0e1107b1775092d66b5e99c175918cbc8476b8d1c0cd7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedProperties()(*if335ec0373b6fcbd74f5eb099d4683fc01fe9cf664574729fbb4fa9afe750ea6.SingleValueExtendedPropertiesRequestBuilder) {
    return if335ec0373b6fcbd74f5eb099d4683fc01fe9cf664574729fbb4fa9afe750ea6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i12842b485e0c49d2f190adb9c34384e7ed531197aa90be3121b9db498d9492f9.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i12842b485e0c49d2f190adb9c34384e7ed531197aa90be3121b9db498d9492f9.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Tasks()(*i86d3aaafac363883fd22a677c5fdcc68ca4a7ef0f722b9829c10a6200d62350f.TasksRequestBuilder) {
    return i86d3aaafac363883fd22a677c5fdcc68ca4a7ef0f722b9829c10a6200d62350f.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) TasksById(id string)(*i70a06e760505c690b521534fb4682338b80fb764f653aa241eb813f4b9e367c5.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i70a06e760505c690b521534fb4682338b80fb764f653aa241eb813f4b9e367c5.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
