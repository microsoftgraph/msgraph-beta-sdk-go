package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder provides operations to manage the owners property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetQueryParameters the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user delegated permissions, no owner can be specified (the current user is the owner). Owner must be specified as an object ID (GUID), not a UPN.
type TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetQueryParameters
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/owners/{user%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user delegated permissions, no owner can be specified (the current user is the owner). Owner must be specified as an object ID (GUID), not a UPN.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) MailboxSettings()(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersItemMailboxSettingsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersItemMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) ServiceProvisioningErrors()(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersItemServiceProvisioningErrorsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersItemServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user delegated permissions, no owner can be specified (the current user is the owner). Owner must be specified as an object ID (GUID), not a UPN.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
