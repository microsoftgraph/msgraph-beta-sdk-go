package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/fields"
    i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/driveitem"
    i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities"
    ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/analytics"
    ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions"
    icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/createlink"
    id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities/item"
    ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions/item"
)

type ListItemRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ListItemRequestBuilder) Activities()(*i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.ActivitiesRequestBuilder) {
    return i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.AnalyticsRequestBuilder) {
    return ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drive/list/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemRequestBuilder) CreateGetRequestInformation(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListItemRequestBuilderGetQueryParameters)
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
func (m *ListItemRequestBuilder) CreateLink()(*icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.CreateLinkRequestBuilder) {
    return icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) DriveItem()(*i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.DriveItemRequestBuilder) {
    return i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.FieldsRequestBuilder) {
    return i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Get(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
func (m *ListItemRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.VersionsRequestBuilder) {
    return ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) VersionsById(id string)(*ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
