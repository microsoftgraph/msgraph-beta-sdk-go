package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/provisionemail"
    i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/members"
    i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/messages"
    i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/tabs"
    iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/filesfolder"
    ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/completemigration"
    ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/removeemail"
    i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/members/item"
    i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/messages/item"
    i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/tabs/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChannelItemRequestBuilderDeleteOptions options for Delete
type ChannelItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChannelItemRequestBuilderGetOptions options for Get
type ChannelItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChannelItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChannelItemRequestBuilderPatchOptions options for Patch
type ChannelItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChannelItemRequestBuilder) CompleteMigration()(*ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04.CompleteMigrationRequestBuilder) {
    return ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/channels/{channel_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for teams
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation(options *ChannelItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation(options *ChannelItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property channels in teams
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(options *ChannelItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property channels for teams
func (m *ChannelItemRequestBuilder) Delete(options *ChannelItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChannelItemRequestBuilder) FilesFolder()(*iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870.FilesFolderRequestBuilder) {
    return iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(options *ChannelItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable), nil
}
func (m *ChannelItemRequestBuilder) Members()(*i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c.MembersRequestBuilder) {
    return i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Messages()(*i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c.MessagesRequestBuilder) {
    return i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teams
func (m *ChannelItemRequestBuilder) Patch(options *ChannelItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59.ProvisionEmailRequestBuilder) {
    return i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be.RemoveEmailRequestBuilder) {
    return ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Tabs()(*i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa.TabsRequestBuilder) {
    return i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
