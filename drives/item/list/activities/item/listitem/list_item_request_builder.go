package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1368978dd349500afb32e29589514d44843c7f8140a279254e03f3cb5fdd5ea8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/fields"
    i2957c7bbc447e9341b55f852f9bb5e40967c0a5ce8884866fb564e23e10ed0b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/driveitem"
    ia0fcfa007eb1b316366eea0b84b2869af1acf6e3df295d66f83750542f364baa "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/activities"
    ia5393902a0c161f5f9284aa604caafb1bc15944cd9fc4878749a15d6b499313c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ib3765913f4f529452e97d0c999737a92248a030362442ed1ffd4dd24bbaf323f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/analytics"
    ibdfd1b435bdb7ada027c836c926ab5c8b32fb9b2b0a0af91492be49dfda1647f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/createlink"
    ifcc48aea51f73e0a7917d7e199317a4d216a8b7b7507a2b2dd25e86d20331e06 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/versions"
    i1a31aeb1242533ab4c4c2fb7a45f2d4fe00e9f42efc4b707288b35c01020aac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/versions/item"
    i5bb1e1afea04aa6755f9952a6a2faf2b72e8e596d71b85b7437800f6920d6488 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item/listitem/activities/item"
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
func (m *ListItemRequestBuilder) Activities()(*ia0fcfa007eb1b316366eea0b84b2869af1acf6e3df295d66f83750542f364baa.ActivitiesRequestBuilder) {
    return ia0fcfa007eb1b316366eea0b84b2869af1acf6e3df295d66f83750542f364baa.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i5bb1e1afea04aa6755f9952a6a2faf2b72e8e596d71b85b7437800f6920d6488.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return i5bb1e1afea04aa6755f9952a6a2faf2b72e8e596d71b85b7437800f6920d6488.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ib3765913f4f529452e97d0c999737a92248a030362442ed1ffd4dd24bbaf323f.AnalyticsRequestBuilder) {
    return ib3765913f4f529452e97d0c999737a92248a030362442ed1ffd4dd24bbaf323f.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drives/{drive_id}/list/activities/{itemActivityOLD_id}/listItem{?select,expand}";
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
func (m *ListItemRequestBuilder) CreateLink()(*ibdfd1b435bdb7ada027c836c926ab5c8b32fb9b2b0a0af91492be49dfda1647f.CreateLinkRequestBuilder) {
    return ibdfd1b435bdb7ada027c836c926ab5c8b32fb9b2b0a0af91492be49dfda1647f.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) DriveItem()(*i2957c7bbc447e9341b55f852f9bb5e40967c0a5ce8884866fb564e23e10ed0b0.DriveItemRequestBuilder) {
    return i2957c7bbc447e9341b55f852f9bb5e40967c0a5ce8884866fb564e23e10ed0b0.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i1368978dd349500afb32e29589514d44843c7f8140a279254e03f3cb5fdd5ea8.FieldsRequestBuilder) {
    return i1368978dd349500afb32e29589514d44843c7f8140a279254e03f3cb5fdd5ea8.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ia5393902a0c161f5f9284aa604caafb1bc15944cd9fc4878749a15d6b499313c.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ia5393902a0c161f5f9284aa604caafb1bc15944cd9fc4878749a15d6b499313c.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
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
func (m *ListItemRequestBuilder) Versions()(*ifcc48aea51f73e0a7917d7e199317a4d216a8b7b7507a2b2dd25e86d20331e06.VersionsRequestBuilder) {
    return ifcc48aea51f73e0a7917d7e199317a4d216a8b7b7507a2b2dd25e86d20331e06.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) VersionsById(id string)(*i1a31aeb1242533ab4c4c2fb7a45f2d4fe00e9f42efc4b707288b35c01020aac6.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i1a31aeb1242533ab4c4c2fb7a45f2d4fe00e9f42efc4b707288b35c01020aac6.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
