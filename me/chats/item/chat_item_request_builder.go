package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/operations"
    i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/permissiongrants"
    ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/installedapps"
    ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/members"
    idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/lastmessagepreview"
    if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/tabs"
    if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages"
    i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/tabs/item"
    i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item"
    i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/installedapps/item"
    i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/permissiongrants/item"
    ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/members/item"
    iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/operations/item"
)

// ChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatItemRequestBuilderDeleteOptions options for Delete
type ChatItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetOptions options for Get
type ChatItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetQueryParameters get chats from me
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatItemRequestBuilderPatchOptions options for Patch
type ChatItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property chats for me
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(options *ChatItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get chats from me
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(options *ChatItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property chats in me
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(options *ChatItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property chats for me
func (m *ChatItemRequestBuilder) Delete(options *ChatItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get chats from me
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable), nil
}
func (m *ChatItemRequestBuilder) InstalledApps()(*ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.InstalledAppsRequestBuilder) {
    return ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) LastMessagePreview()(*idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.LastMessagePreviewRequestBuilder) {
    return idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.MembersRequestBuilder) {
    return ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.MessagesRequestBuilder) {
    return if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Operations()(*i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.OperationsRequestBuilder) {
    return i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.operations.item collection
func (m *ChatItemRequestBuilder) OperationsById(id string)(*iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in me
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatItemRequestBuilder) PermissionGrants()(*i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.PermissionGrantsRequestBuilder) {
    return i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.permissionGrants.item collection
func (m *ChatItemRequestBuilder) PermissionGrantsById(id string)(*i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Tabs()(*if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.TabsRequestBuilder) {
    return if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
