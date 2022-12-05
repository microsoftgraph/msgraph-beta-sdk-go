package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder provides operations to manage the connectorGroups property of the microsoft.graph.onPremisesPublishingProfile entity.
type OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
type OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetQueryParameters
}
// OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Applications provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) Applications()(*OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationsById provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) ApplicationsById(id string)(*OnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["application%2Did"] = id
    }
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderInternal instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connectorGroups for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property connectorGroups in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property connectorGroups for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.connectorGroup entity.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) Members()(*OnPremisesPublishingProfilesItemConnectorGroupsItemMembersRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectorGroups.item.members.item collection
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) MembersById(id string)(*OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector%2Did"] = id
    }
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connectorGroups in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsConnectorGroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupable), nil
}
