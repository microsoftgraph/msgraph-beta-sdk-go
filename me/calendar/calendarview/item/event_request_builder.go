package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/dismissreminder"
    i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences"
    i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/forward"
    i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/accept"
    i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments"
    i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties"
    i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/calendar"
    i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances"
    ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/tentativelyaccept"
    ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/decline"
    icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/cancel"
    id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions"
    if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/snoozereminder"
    ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties"
    ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item"
    iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions/item"
    ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments/item"
    ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties/item"
    if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item"
    ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    EndDateTime *string;
    Expand []string;
    Select_escaped []string;
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.AcceptRequestBuilder) {
    return i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.AttachmentsRequestBuilder) {
    return i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.CalendarRequestBuilder) {
    return i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.CancelRequestBuilder) {
    return icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
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
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.DeclineRequestBuilder) {
    return ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.DismissReminderRequestBuilder) {
    return i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.ExceptionOccurrencesRequestBuilder) {
    return i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.ExtensionsRequestBuilder) {
    return id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.ForwardRequestBuilder) {
    return i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.InstancesRequestBuilder) {
    return i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.MultiValueExtendedPropertiesRequestBuilder) {
    return i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.SingleValueExtendedPropertiesRequestBuilder) {
    return ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.SnoozeReminderRequestBuilder) {
    return if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.TentativelyAcceptRequestBuilder) {
    return ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
