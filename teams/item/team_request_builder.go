package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/installedapps"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/operations"
    i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/clone"
    i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/completemigration"
    i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/archive"
    i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/tags"
    i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels"
    i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/owners"
    i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/group"
    i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/template"
    ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/members"
    ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/unarchive"
    ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/photo"
    icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel"
    iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/sendactivitynotification"
    iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants"
    i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/members/item"
    i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/installedapps/item"
    i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/tags/item"
    i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item"
    ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/operations/item"
    iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/item"
)

// TeamRequestBuilder builds and executes requests for operations under \teams\{team-id}
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
// TeamRequestBuilderGetQueryParameters get entity from teams by key
type TeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamRequestBuilderPatchOptions options for Patch
type TeamRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Team;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamRequestBuilder) Archive()(*i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518.ArchiveRequestBuilder) {
    return i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Channels()(*i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d.ChannelsRequestBuilder) {
    return i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8.ChannelRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel_id"] = id
    }
    return i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8.NewChannelRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Clone()(*i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7.CloneRequestBuilder) {
    return i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) CompleteMigration()(*i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c.CompleteMigrationRequestBuilder) {
    return i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from teams
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
// CreateGetRequestInformation get entity from teams by key
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
// CreatePatchRequestInformation update entity in teams
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
// Delete delete entity from teams
func (m *TeamRequestBuilder) Delete(options *TeamRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from teams by key
func (m *TeamRequestBuilder) Get(options *TeamRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Team, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTeam() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Team), nil
}
func (m *TeamRequestBuilder) Group()(*i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406.GroupRequestBuilder) {
    return i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) InstalledApps()(*i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca.InstalledAppsRequestBuilder) {
    return i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f.TeamsAppInstallationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f.NewTeamsAppInstallationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Members()(*ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd.MembersRequestBuilder) {
    return ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Operations()(*i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7.OperationsRequestBuilder) {
    return i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d.TeamsAsyncOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d.NewTeamsAsyncOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Owners()(*i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79.OwnersRequestBuilder) {
    return i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in teams
func (m *TeamRequestBuilder) Patch(options *TeamRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamRequestBuilder) PermissionGrants()(*iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a.PermissionGrantsRequestBuilder) {
    return iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.permissionGrants.item collection
func (m *TeamRequestBuilder) PermissionGrantsById(id string)(*iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4.ResourceSpecificPermissionGrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4.NewResourceSpecificPermissionGrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Photo()(*ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e.PhotoRequestBuilder) {
    return ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) PrimaryChannel()(*icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f.PrimaryChannelRequestBuilder) {
    return icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Schedule()(*i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e.ScheduleRequestBuilder) {
    return i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) SendActivityNotification()(*iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b.SendActivityNotificationRequestBuilder) {
    return iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Tags()(*i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c.TagsRequestBuilder) {
    return i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.tags.item collection
func (m *TeamRequestBuilder) TagsById(id string)(*i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708.TeamworkTagRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag_id"] = id
    }
    return i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708.NewTeamworkTagRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TeamRequestBuilder) Template()(*i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5.TemplateRequestBuilder) {
    return i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamRequestBuilder) Unarchive()(*ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d.UnarchiveRequestBuilder) {
    return ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
