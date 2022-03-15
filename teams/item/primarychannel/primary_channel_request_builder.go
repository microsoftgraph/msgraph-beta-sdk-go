package primarychannel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0523c837e5aa043ea3780f6d384a657438fc499bdb768169ab97541970241ef4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/sharedwithteams"
    i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/tabs"
    i17b4ce191aaab30a5836a14057bb3fbc4cc34118922956129ebf9a726dbae62e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/removeemail"
    i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/completemigration"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages"
    i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/members"
    i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/filesfolder"
    ib49515d0c4c45b64bac89c4190d41ddf469ef5300189a69ab134543640527613 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/provisionemail"
    if7d3dc91aae7b81954f5794f652d60bf40bd040518a08ed894c8018b5b3871b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i01501d8bdaa88b3e325abf5899eb702455b1b03c3abf4845c994bfb2fe588af5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/sharedwithteams/item"
    i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/members/item"
    i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item"
    icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/tabs/item"
)

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrimaryChannelRequestBuilderDeleteOptions options for Delete
type PrimaryChannelRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrimaryChannelRequestBuilderGetOptions options for Get
type PrimaryChannelRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrimaryChannelRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrimaryChannelRequestBuilderPatchOptions options for Patch
type PrimaryChannelRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa.CompleteMigrationRequestBuilder) {
    return i1eb85305be35c87f82fc0651bbd70b3d14333c21a691f70f9e435fa28ff209aa.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/primaryChannel{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(options *PrimaryChannelRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(options *PrimaryChannelRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(options *PrimaryChannelRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property primaryChannel for teams
func (m *PrimaryChannelRequestBuilder) Delete(options *PrimaryChannelRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *PrimaryChannelRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName(tenantId *string, userId *string, userPrincipalName *string)(*if7d3dc91aae7b81954f5794f652d60bf40bd040518a08ed894c8018b5b3871b7.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return if7d3dc91aae7b81954f5794f652d60bf40bd040518a08ed894c8018b5b3871b7.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, tenantId, userId, userPrincipalName);
}
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e.FilesFolderRequestBuilder) {
    return i69bf49888498532e5d6fe758a7a9d13e8b2580d0ab19f737bd527d11303da79e.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get(options *PrimaryChannelRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Channelable), nil
}
func (m *PrimaryChannelRequestBuilder) Members()(*i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515.MembersRequestBuilder) {
    return i67923f345c8cd21dcca218335afaeba54ea28bf38a33354fadeaa4650de78515.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i1609b5720a26919bbe6f305c8e35b3e3d15fdf33acf3814e76e0b661e4b6ed44.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Messages()(*i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23.MessagesRequestBuilder) {
    return i5331dde414fe312a1eaee47c5565de7b4057d8a80eb8d45969b508953b35ac23.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i5475d1ec34efb9c05c75780100e516b117ec8c9908b6af485721aaad0050ef85.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in teams
func (m *PrimaryChannelRequestBuilder) Patch(options *PrimaryChannelRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
func (m *PrimaryChannelRequestBuilder) SharedWithTeams()(*i0523c837e5aa043ea3780f6d384a657438fc499bdb768169ab97541970241ef4.SharedWithTeamsRequestBuilder) {
    return i0523c837e5aa043ea3780f6d384a657438fc499bdb768169ab97541970241ef4.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.sharedWithTeams.item collection
func (m *PrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*i01501d8bdaa88b3e325abf5899eb702455b1b03c3abf4845c994bfb2fe588af5.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo_id"] = id
    }
    return i01501d8bdaa88b3e325abf5899eb702455b1b03c3abf4845c994bfb2fe588af5.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrimaryChannelRequestBuilder) Tabs()(*i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b.TabsRequestBuilder) {
    return i14f382464eb02055ca9dcbd98bbefcc5f5cde7c67fb548dc45519dbdabcd998b.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return icf6d8f24c127128ff064ffe65313e1ae4d46b1c21f5338abcde90ed64a944f80.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
