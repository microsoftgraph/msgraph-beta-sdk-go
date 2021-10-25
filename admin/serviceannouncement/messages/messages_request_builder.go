package messages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unfavorite"
    i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/favorite"
    i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markunread"
    ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markread"
    icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/archive"
    icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unarchive"
)

type MessagesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MessagesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func (m *MessagesRequestBuilder) Archive()(*icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.ArchiveRequestBuilder) {
    return icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    m := &MessagesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/admin/serviceAnnouncement/messages{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewMessagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessagesRequestBuilder) CreateGetRequestInformation(q func (value *MessagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MessagesRequestBuilderGetQueryParameters)
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
func (m *MessagesRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *MessagesRequestBuilder) Favorite()(*i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.FavoriteRequestBuilder) {
    return i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.NewFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) Get(q func (value *MessagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*MessagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessagesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*MessagesResponse), nil
}
func (m *MessagesRequestBuilder) MarkRead()(*ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.MarkReadRequestBuilder) {
    return ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.NewMarkReadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) MarkUnread()(*i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.MarkUnreadRequestBuilder) {
    return i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.NewMarkUnreadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServiceUpdateMessage() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceUpdateMessage), nil
}
func (m *MessagesRequestBuilder) Unarchive()(*icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.UnarchiveRequestBuilder) {
    return icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessagesRequestBuilder) Unfavorite()(*i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.UnfavoriteRequestBuilder) {
    return i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.NewUnfavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
