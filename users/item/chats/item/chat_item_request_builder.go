package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ca2a9cae7818d61f3ef187fe1f74ff11c7c24021f1563fde2c359d2926bf597 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/lastmessagepreview"
    i54e0ee28261ed5d0a6d4455159e895f2499fa33bc64b49a236251a1e6d997e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/operations"
    i9a5a07ec6132990ff8f9ab04a7e9bfaafe2225776ed9ae987b9cb2a4c724571e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/installedapps"
    ia99138ca96fe463c15270f256dd28f43a8debd37e2257bdb7f17c8b3c4932c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/members"
    ib4cae878dd83473762ab431e83f66914a9a52daee2f4a232de4f43fc71f6df34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/tabs"
    ib54911005485cdcc79f8d37f7e9f76efcb97cc6e45037bf3a31ae82d0ba25b4c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/messages"
    if4d7355d033005aaed0211ee8e0857a76c6a17ae379b409d56c7f00eabe0ced8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/permissiongrants"
    i32dcaf418f207c7408cdb1084a37cffe7983989294c800c84f726e1e3165eabb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/messages/item"
    i90fcffd68aa029c4102834a67aea89332247a5f005ef3371cc76730ff3224bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/members/item"
    iad400fe8235cbf891298d9a0a89793e66c7ddf0fbebabd6bb395fe8e34a8eeef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/permissiongrants/item"
    ic9f83acf4b742966208ff504012d0fb43ab9a221c9e4bb1bf31beb79e7d45390 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/tabs/item"
    ice239d24ff0b0732d7cf1b28ea933776f0a04e61cd76f1c961bbf19740466dff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/operations/item"
    id41aec000e052a36a9d7d55d2d725c278965c65d45c20ecf0599d3497821d07c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item/installedapps/item"
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
// ChatItemRequestBuilderGetQueryParameters get chats from users
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/chats/{chat_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property chats for users
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
// CreateGetRequestInformation get chats from users
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
// CreatePatchRequestInformation update the navigation property chats in users
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
// Delete delete navigation property chats for users
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
// Get get chats from users
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
func (m *ChatItemRequestBuilder) InstalledApps()(*i9a5a07ec6132990ff8f9ab04a7e9bfaafe2225776ed9ae987b9cb2a4c724571e.InstalledAppsRequestBuilder) {
    return i9a5a07ec6132990ff8f9ab04a7e9bfaafe2225776ed9ae987b9cb2a4c724571e.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*id41aec000e052a36a9d7d55d2d725c278965c65d45c20ecf0599d3497821d07c.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return id41aec000e052a36a9d7d55d2d725c278965c65d45c20ecf0599d3497821d07c.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) LastMessagePreview()(*i1ca2a9cae7818d61f3ef187fe1f74ff11c7c24021f1563fde2c359d2926bf597.LastMessagePreviewRequestBuilder) {
    return i1ca2a9cae7818d61f3ef187fe1f74ff11c7c24021f1563fde2c359d2926bf597.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*ia99138ca96fe463c15270f256dd28f43a8debd37e2257bdb7f17c8b3c4932c80.MembersRequestBuilder) {
    return ia99138ca96fe463c15270f256dd28f43a8debd37e2257bdb7f17c8b3c4932c80.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*i90fcffd68aa029c4102834a67aea89332247a5f005ef3371cc76730ff3224bfd.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i90fcffd68aa029c4102834a67aea89332247a5f005ef3371cc76730ff3224bfd.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*ib54911005485cdcc79f8d37f7e9f76efcb97cc6e45037bf3a31ae82d0ba25b4c.MessagesRequestBuilder) {
    return ib54911005485cdcc79f8d37f7e9f76efcb97cc6e45037bf3a31ae82d0ba25b4c.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*i32dcaf418f207c7408cdb1084a37cffe7983989294c800c84f726e1e3165eabb.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i32dcaf418f207c7408cdb1084a37cffe7983989294c800c84f726e1e3165eabb.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Operations()(*i54e0ee28261ed5d0a6d4455159e895f2499fa33bc64b49a236251a1e6d997e29.OperationsRequestBuilder) {
    return i54e0ee28261ed5d0a6d4455159e895f2499fa33bc64b49a236251a1e6d997e29.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.operations.item collection
func (m *ChatItemRequestBuilder) OperationsById(id string)(*ice239d24ff0b0732d7cf1b28ea933776f0a04e61cd76f1c961bbf19740466dff.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return ice239d24ff0b0732d7cf1b28ea933776f0a04e61cd76f1c961bbf19740466dff.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in users
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
func (m *ChatItemRequestBuilder) PermissionGrants()(*if4d7355d033005aaed0211ee8e0857a76c6a17ae379b409d56c7f00eabe0ced8.PermissionGrantsRequestBuilder) {
    return if4d7355d033005aaed0211ee8e0857a76c6a17ae379b409d56c7f00eabe0ced8.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.permissionGrants.item collection
func (m *ChatItemRequestBuilder) PermissionGrantsById(id string)(*iad400fe8235cbf891298d9a0a89793e66c7ddf0fbebabd6bb395fe8e34a8eeef.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return iad400fe8235cbf891298d9a0a89793e66c7ddf0fbebabd6bb395fe8e34a8eeef.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Tabs()(*ib4cae878dd83473762ab431e83f66914a9a52daee2f4a232de4f43fc71f6df34.TabsRequestBuilder) {
    return ib4cae878dd83473762ab431e83f66914a9a52daee2f4a232de4f43fc71f6df34.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*ic9f83acf4b742966208ff504012d0fb43ab9a221c9e4bb1bf31beb79e7d45390.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return ic9f83acf4b742966208ff504012d0fb43ab9a221c9e4bb1bf31beb79e7d45390.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
