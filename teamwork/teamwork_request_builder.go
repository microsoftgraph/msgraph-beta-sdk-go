package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3fce2babf7b4f0d1cb31c054814424b76c87dc90eadacea48f98134cc0ffe6e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices"
    i5f853021bcd09dc31bcfe1b81013b0b4476658eb54310db69fef8d76e06b714a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams"
    ib14bcbfd4e3bccc76a4fe3dff4da7ebc6577392679e6ba6cd1710cd5c03eeba0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates"
    ib9c577680442ee756cdbf50e3c2f2d61fc1877ecb12e13bf9ec9802c878f1b45 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/workforceintegrations"
    id1ff05524188cf798ff28c962060e88b92f5555111be336cede7088f8719c6a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/sendactivitynotificationtorecipients"
    idad6e2713a756a61b26caca3cc1da42b705a499834ed5b89ca8f4eade11dac1d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamsappsettings"
    i17ffe63fafac127289aa66b154f468d1a6a27dee2b4568c67bf179cbf15811a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item"
    i75138276537be7f359112029ca90eee1631bb436b9f9db869b58bf15fd0cd092 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/workforceintegrations/item"
    ic86fd621ddc51f4cb681ecef66a6c6bbfecda42d24db8e141175f82fdabf74d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item"
    id6e8a1d6759febda5acd2ed7c5930f2db6a83ef4fdb4eddde327904716515370 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item"
)

// TeamworkRequestBuilder provides operations to manage the teamwork singleton.
type TeamworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkRequestBuilderGetQueryParameters get teamwork
type TeamworkRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkRequestBuilderGetQueryParameters
}
// TeamworkRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamworkRequestBuilderInternal instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkRequestBuilder) {
    m := &TeamworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkRequestBuilder instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get teamwork
func (m *TeamworkRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get teamwork
func (m *TeamworkRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update teamwork
func (m *TeamworkRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update teamwork
func (m *TeamworkRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeletedTeams the deletedTeams property
func (m *TeamworkRequestBuilder) DeletedTeams()(*i5f853021bcd09dc31bcfe1b81013b0b4476658eb54310db69fef8d76e06b714a.DeletedTeamsRequestBuilder) {
    return i5f853021bcd09dc31bcfe1b81013b0b4476658eb54310db69fef8d76e06b714a.NewDeletedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletedTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.deletedTeams.item collection
func (m *TeamworkRequestBuilder) DeletedTeamsById(id string)(*ic86fd621ddc51f4cb681ecef66a6c6bbfecda42d24db8e141175f82fdabf74d2.DeletedTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deletedTeam%2Did"] = id
    }
    return ic86fd621ddc51f4cb681ecef66a6c6bbfecda42d24db8e141175f82fdabf74d2.NewDeletedTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices the devices property
func (m *TeamworkRequestBuilder) Devices()(*i3fce2babf7b4f0d1cb31c054814424b76c87dc90eadacea48f98134cc0ffe6e7.DevicesRequestBuilder) {
    return i3fce2babf7b4f0d1cb31c054814424b76c87dc90eadacea48f98134cc0ffe6e7.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.devices.item collection
func (m *TeamworkRequestBuilder) DevicesById(id string)(*i17ffe63fafac127289aa66b154f468d1a6a27dee2b4568c67bf179cbf15811a8.TeamworkDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkDevice%2Did"] = id
    }
    return i17ffe63fafac127289aa66b154f468d1a6a27dee2b4568c67bf179cbf15811a8.NewTeamworkDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get teamwork
func (m *TeamworkRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable), nil
}
// Patch update teamwork
func (m *TeamworkRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable), nil
}
// SendActivityNotificationToRecipients the sendActivityNotificationToRecipients property
func (m *TeamworkRequestBuilder) SendActivityNotificationToRecipients()(*id1ff05524188cf798ff28c962060e88b92f5555111be336cede7088f8719c6a6.SendActivityNotificationToRecipientsRequestBuilder) {
    return id1ff05524188cf798ff28c962060e88b92f5555111be336cede7088f8719c6a6.NewSendActivityNotificationToRecipientsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsAppSettings the teamsAppSettings property
func (m *TeamworkRequestBuilder) TeamsAppSettings()(*idad6e2713a756a61b26caca3cc1da42b705a499834ed5b89ca8f4eade11dac1d.TeamsAppSettingsRequestBuilder) {
    return idad6e2713a756a61b26caca3cc1da42b705a499834ed5b89ca8f4eade11dac1d.NewTeamsAppSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamTemplates the teamTemplates property
func (m *TeamworkRequestBuilder) TeamTemplates()(*ib14bcbfd4e3bccc76a4fe3dff4da7ebc6577392679e6ba6cd1710cd5c03eeba0.TeamTemplatesRequestBuilder) {
    return ib14bcbfd4e3bccc76a4fe3dff4da7ebc6577392679e6ba6cd1710cd5c03eeba0.NewTeamTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item collection
func (m *TeamworkRequestBuilder) TeamTemplatesById(id string)(*id6e8a1d6759febda5acd2ed7c5930f2db6a83ef4fdb4eddde327904716515370.TeamTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamTemplate%2Did"] = id
    }
    return id6e8a1d6759febda5acd2ed7c5930f2db6a83ef4fdb4eddde327904716515370.NewTeamTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WorkforceIntegrations the workforceIntegrations property
func (m *TeamworkRequestBuilder) WorkforceIntegrations()(*ib9c577680442ee756cdbf50e3c2f2d61fc1877ecb12e13bf9ec9802c878f1b45.WorkforceIntegrationsRequestBuilder) {
    return ib9c577680442ee756cdbf50e3c2f2d61fc1877ecb12e13bf9ec9802c878f1b45.NewWorkforceIntegrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkforceIntegrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.workforceIntegrations.item collection
func (m *TeamworkRequestBuilder) WorkforceIntegrationsById(id string)(*i75138276537be7f359112029ca90eee1631bb436b9f9db869b58bf15fd0cd092.WorkforceIntegrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workforceIntegration%2Did"] = id
    }
    return i75138276537be7f359112029ca90eee1631bb436b9f9db869b58bf15fd0cd092.NewWorkforceIntegrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
