package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
func (m *TeamworkRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeletedTeams provides operations to manage the deletedTeams property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) DeletedTeams()(*TeamworkDeletedTeamsRequestBuilder) {
    return NewTeamworkDeletedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletedTeamsById provides operations to manage the deletedTeams property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) DeletedTeamsById(id string)(*TeamworkDeletedTeamsDeletedTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deletedTeam%2Did"] = id
    }
    return NewTeamworkDeletedTeamsDeletedTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices provides operations to manage the devices property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) Devices()(*TeamworkDevicesRequestBuilder) {
    return NewTeamworkDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById provides operations to manage the devices property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) DevicesById(id string)(*TeamworkDevicesTeamworkDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkDevice%2Did"] = id
    }
    return NewTeamworkDevicesTeamworkDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get teamwork
func (m *TeamworkRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamworkable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// SendActivityNotificationToRecipients provides operations to call the sendActivityNotificationToRecipients method.
func (m *TeamworkRequestBuilder) SendActivityNotificationToRecipients()(*TeamworkSendActivityNotificationToRecipientsRequestBuilder) {
    return NewTeamworkSendActivityNotificationToRecipientsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsAppSettings provides operations to manage the teamsAppSettings property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) TeamsAppSettings()(*TeamworkTeamsAppSettingsRequestBuilder) {
    return NewTeamworkTeamsAppSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamTemplates provides operations to manage the teamTemplates property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) TeamTemplates()(*TeamworkTeamTemplatesRequestBuilder) {
    return NewTeamworkTeamTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamTemplatesById provides operations to manage the teamTemplates property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) TeamTemplatesById(id string)(*TeamworkTeamTemplatesTeamTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamTemplate%2Did"] = id
    }
    return NewTeamworkTeamTemplatesTeamTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WorkforceIntegrations provides operations to manage the workforceIntegrations property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) WorkforceIntegrations()(*TeamworkWorkforceIntegrationsRequestBuilder) {
    return NewTeamworkWorkforceIntegrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkforceIntegrationsById provides operations to manage the workforceIntegrations property of the microsoft.graph.teamwork entity.
func (m *TeamworkRequestBuilder) WorkforceIntegrationsById(id string)(*TeamworkWorkforceIntegrationsWorkforceIntegrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workforceIntegration%2Did"] = id
    }
    return NewTeamworkWorkforceIntegrationsWorkforceIntegrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
