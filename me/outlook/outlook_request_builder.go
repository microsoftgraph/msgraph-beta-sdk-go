package outlook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezoneswithtimezonestandard"
    i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups"
    i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders"
    ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories"
    ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezones"
    ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks"
    ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedlanguages"
    i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories/item"
    i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item"
    ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item"
    id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item"
)

type OutlookRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OutlookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/outlook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOutlookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OutlookRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookRequestBuilder) CreateGetRequestInformation(q func (value *OutlookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OutlookRequestBuilderGetQueryParameters)
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
func (m *OutlookRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutlookRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookRequestBuilder) Get(q func (value *OutlookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser), nil
}
func (m *OutlookRequestBuilder) MasterCategories()(*ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.MasterCategoriesRequestBuilder) {
    return ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.NewMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.OutlookCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory_id"] = id
    }
    return i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.NewOutlookCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OutlookRequestBuilder) SupportedLanguages()(*ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.SupportedLanguagesRequestBuilder) {
    return ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.NewSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) SupportedTimeZones()(*ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.SupportedTimeZonesRequestBuilder) {
    return ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.NewSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
func (m *OutlookRequestBuilder) TaskFolders()(*i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.TaskFoldersRequestBuilder) {
    return i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.NewTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.OutlookTaskFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder_id"] = id
    }
    return ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.NewOutlookTaskFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskGroups()(*i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.TaskGroupsRequestBuilder) {
    return i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.NewTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.OutlookTaskGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup_id"] = id
    }
    return i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.NewOutlookTaskGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) Tasks()(*ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.TasksRequestBuilder) {
    return ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TasksById(id string)(*id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
