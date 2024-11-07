package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamOwnersWithUserPrincipalNameRequestBuilder provides operations to manage the owners property of the microsoft.graph.team entity.
type ItemTeamOwnersWithUserPrincipalNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user-delegated permissions, no owner can be specified (the current user is the owner). The owner must be specified as an object ID (GUID), not a UPN.
type ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewItemTeamOwnersWithUserPrincipalNameRequestBuilderInternal instantiates a new ItemTeamOwnersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewItemTeamOwnersWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userPrincipalName *string)(*ItemTeamOwnersWithUserPrincipalNameRequestBuilder) {
    m := &ItemTeamOwnersWithUserPrincipalNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/owners(userPrincipalName='{userPrincipalName}'){?%24expand,%24select}", pathParameters),
    }
    if userPrincipalName != nil {
        m.BaseRequestBuilder.PathParameters["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewItemTeamOwnersWithUserPrincipalNameRequestBuilder instantiates a new ItemTeamOwnersWithUserPrincipalNameRequestBuilder and sets the default values.
func NewItemTeamOwnersWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamOwnersWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamOwnersWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user-delegated permissions, no owner can be specified (the current user is the owner). The owner must be specified as an object ID (GUID), not a UPN.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamOwnersWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation the list of this team's owners. Currently, when creating a team using application permissions, exactly one owner must be specified. When using user-delegated permissions, no owner can be specified (the current user is the owner). The owner must be specified as an object ID (GUID), not a UPN.
// returns a *RequestInformation when successful
func (m *ItemTeamOwnersWithUserPrincipalNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamOwnersWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamOwnersWithUserPrincipalNameRequestBuilder when successful
func (m *ItemTeamOwnersWithUserPrincipalNameRequestBuilder) WithUrl(rawUrl string)(*ItemTeamOwnersWithUserPrincipalNameRequestBuilder) {
    return NewItemTeamOwnersWithUserPrincipalNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
