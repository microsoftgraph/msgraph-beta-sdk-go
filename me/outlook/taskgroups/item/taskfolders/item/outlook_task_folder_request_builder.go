package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties"
    iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties"
    ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks"
    i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties/item"
    ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties/item"
    id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item"
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
    m.urlTemplate = "https://graph.microsoft.com/beta/me/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}{?select,expand}";
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
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedProperties()(*iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae.MultiValueExtendedPropertiesRequestBuilder) {
    return iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedProperties()(*i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Tasks()(*ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f.TasksRequestBuilder) {
    return ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) TasksById(id string)(*id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
