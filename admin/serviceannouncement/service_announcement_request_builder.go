package serviceannouncement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/issues"
    iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages"
    ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/issues/item"
    ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/item"
    ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews/item"
)

type ServiceAnnouncementRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ServiceAnnouncementRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    m := &ServiceAnnouncementRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/admin/serviceAnnouncement{?select,expand}";
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
func NewServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ServiceAnnouncementRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ServiceAnnouncementRequestBuilder) CreateGetRequestInformation(q func (value *ServiceAnnouncementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ServiceAnnouncementRequestBuilderGetQueryParameters)
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
func (m *ServiceAnnouncementRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceAnnouncement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ServiceAnnouncementRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ServiceAnnouncementRequestBuilder) Get(q func (value *ServiceAnnouncementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceAnnouncement, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServiceAnnouncement() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceAnnouncement), nil
}
func (m *ServiceAnnouncementRequestBuilder) HealthOverviews()(*i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4.HealthOverviewsRequestBuilder) {
    return i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4.NewHealthOverviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) HealthOverviewsById(id string)(*ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98.ServiceHealthRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["serviceHealth_id"] = id
    }
    return ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98.NewServiceHealthRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Issues()(*i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948.IssuesRequestBuilder) {
    return i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) IssuesById(id string)(*ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7.ServiceHealthIssueRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["serviceHealthIssue_id"] = id
    }
    return ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7.NewServiceHealthIssueRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Messages()(*iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa.MessagesRequestBuilder) {
    return iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) MessagesById(id string)(*ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04.ServiceUpdateMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["serviceUpdateMessage_id"] = id
    }
    return ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04.NewServiceUpdateMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServiceAnnouncementRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceAnnouncement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
