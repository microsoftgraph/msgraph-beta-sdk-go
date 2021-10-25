package primarychannel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/tabs"
    i17b4ce191aaab30a5836a14057bb3fbc4cc34118922956129ebf9a726dbae62e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/removeemail"
    i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/completemigration"
    i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages"
    i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/members"
    i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/filesfolder"
    ib49515d0c4c45b64bac89c4190d41ddf469ef5300189a69ab134543640527613 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/provisionemail"
    i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/members/item"
    i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item"
    icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/tabs/item"
)

type PrimaryChannelRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa.CompleteMigrationRequestBuilder) {
    return i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/teams/{team_id}/primaryChannel{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(q func (value *PrimaryChannelRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrimaryChannelRequestBuilderGetQueryParameters)
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
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channel, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrimaryChannelRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e.FilesFolderRequestBuilder) {
    return i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Get(q func (value *PrimaryChannelRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channel, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewChannel() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channel), nil
}
func (m *PrimaryChannelRequestBuilder) Members()(*i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515.MembersRequestBuilder) {
    return i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Messages()(*i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23.MessagesRequestBuilder) {
    return i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channel, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*ib49515d0c4c45b64bac89c4190d41ddf469ef5300189a69ab134543640527613.ProvisionEmailRequestBuilder) {
    return ib49515d0c4c45b64bac89c4190d41ddf469ef5300189a69ab134543640527613.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i17b4ce191aaab30a5836a14057bb3fbc4cc34118922956129ebf9a726dbae62e.RemoveEmailRequestBuilder) {
    return i17b4ce191aaab30a5836a14057bb3fbc4cc34118922956129ebf9a726dbae62e.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Tabs()(*i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b.TabsRequestBuilder) {
    return i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
