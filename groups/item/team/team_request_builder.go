package team

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/owners"
    i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/template"
    i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/tags"
    i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/group"
    i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/operations"
    i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/installedapps"
    i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/photo"
    ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel"
    ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule"
    id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels"
    ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/members"
    ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants"
    i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/members/item"
    i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/installedapps/item"
    i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/operations/item"
    ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item"
    ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/tags/item"
    id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/owners/item"
    idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item"
)

// TeamRequestBuilder provides operations to manage the team property of the microsoft.graph.group entity.
type TeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamRequestBuilderDeleteOptions options for Delete
type TeamRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamRequestBuilderGetOptions options for Get
type TeamRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamRequestBuilderGetQueryParameters get team from groups
type TeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamRequestBuilderPatchOptions options for Patch
type TeamRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Teamable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamRequestBuilder) Channels()(*id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550.ChannelsRequestBuilder) {
    return id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel_id"] = id
    }
    return idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property team for groups
func (m *TeamRequestBuilder) CreateDeleteRequestInformation(options *TeamRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get team from groups
func (m *TeamRequestBuilder) CreateGetRequestInformation(options *TeamRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property team in groups
func (m *TeamRequestBuilder) CreatePatchRequestInformation(options *TeamRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property team for groups
func (m *TeamRequestBuilder) Delete(options *TeamRequestBuilderDeleteOptions)(error) {
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
// Get get team from groups
func (m *TeamRequestBuilder) Get(options *TeamRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTeamFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Teamable), nil
}
func (m *TeamRequestBuilder) Group()(*i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac.GroupRequestBuilder) {
    return i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) InstalledApps()(*i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58.InstalledAppsRequestBuilder) {
    return i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Members()(*ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356.MembersRequestBuilder) {
    return ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Operations()(*i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53.OperationsRequestBuilder) {
    return i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Owners()(*i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26.OwnersRequestBuilder) {
    return i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.owners.item collection
func (m *TeamRequestBuilder) OwnersById(id string)(*id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user_id"] = id
    }
    return id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property team in groups
func (m *TeamRequestBuilder) Patch(options *TeamRequestBuilderPatchOptions)(error) {
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
func (m *TeamRequestBuilder) PermissionGrants()(*ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b.PermissionGrantsRequestBuilder) {
    return ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.permissionGrants.item collection
func (m *TeamRequestBuilder) PermissionGrantsById(id string)(*ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Photo()(*i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0.PhotoRequestBuilder) {
    return i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) PrimaryChannel()(*ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b.PrimaryChannelRequestBuilder) {
    return ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Schedule()(*ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e.ScheduleRequestBuilder) {
    return ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Tags()(*i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136.TagsRequestBuilder) {
    return i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.tags.item collection
func (m *TeamRequestBuilder) TagsById(id string)(*ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag_id"] = id
    }
    return ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Template()(*i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b.TemplateRequestBuilder) {
    return i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
