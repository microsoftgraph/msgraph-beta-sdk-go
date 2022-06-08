package teamwork

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/sendactivitynotification"
    i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/associatedteams"
    i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps"
    i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/installedapps/item"
    i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork/associatedteams/item"
)

// TeamworkRequestBuilder provides operations to manage the teamwork property of the microsoft.graph.user entity.
type TeamworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkRequestBuilderGetQueryParameters a container for Microsoft Teams features available for the user. Read-only. Nullable.
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
// AssociatedTeams the associatedTeams property
func (m *TeamworkRequestBuilder) AssociatedTeams()(*i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df.AssociatedTeamsRequestBuilder) {
    return i200af14bc8a01d318341d7ffd4bd69adb3b9b97703c97f7fd8c69b8f1dc415df.NewAssociatedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssociatedTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.teamwork.associatedTeams.item collection
func (m *TeamworkRequestBuilder) AssociatedTeamsById(id string)(*i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545.AssociatedTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["associatedTeamInfo%2Did"] = id
    }
    return i884fd4ffb36d199e7dec37be072607e2769510674770c99847363a2f72615545.NewAssociatedTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamworkRequestBuilderInternal instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkRequestBuilder) {
    m := &TeamworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/teamwork{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property teamwork for me
func (m *TeamworkRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property teamwork for me
func (m *TeamworkRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamworkRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a container for Microsoft Teams features available for the user. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property teamwork in me
func (m *TeamworkRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property teamwork in me
func (m *TeamworkRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property teamwork for me
func (m *TeamworkRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property teamwork for me
func (m *TeamworkRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamworkRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserTeamworkFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable), nil
}
// InstalledApps the installedApps property
func (m *TeamworkRequestBuilder) InstalledApps()(*i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.InstalledAppsRequestBuilder) {
    return i5352384a816a8116471b29a72d69ff38c387d8edbbac8d15cc0e43ef5d88cb0e.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.teamwork.installedApps.item collection
func (m *TeamworkRequestBuilder) InstalledAppsById(id string)(*i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.UserScopeTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userScopeTeamsAppInstallation%2Did"] = id
    }
    return i2f8fc19d0ffa378974d134d26c9bd4ed6b3083770f50add7d052132db4ad8e8b.NewUserScopeTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamwork in me
func (m *TeamworkRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property teamwork in me
func (m *TeamworkRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamworkRequestBuilder) SendActivityNotification()(*i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.SendActivityNotificationRequestBuilder) {
    return i1621e3eac6e3236524f354f9b5e3dcec4b2ac8b61604c4a81f7db4b934cd90c1.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
