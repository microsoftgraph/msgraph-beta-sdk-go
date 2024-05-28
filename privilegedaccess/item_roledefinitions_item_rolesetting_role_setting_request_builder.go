package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder provides operations to manage the roleSetting property of the microsoft.graph.governanceRoleDefinition entity.
type ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetQueryParameters the associated role setting for the role definition.
type ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetQueryParameters
}
// NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderInternal instantiates a new ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder and sets the default values.
func NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) {
    m := &ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/roleDefinitions/{governanceRoleDefinition%2Did}/roleSetting{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder instantiates a new ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder and sets the default values.
func NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the associated role setting for the role definition.
// returns a GovernanceRoleSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleSettingable), nil
}
// ToGetRequestInformation the associated role setting for the role definition.
// returns a *RequestInformation when successful
func (m *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder when successful
func (m *ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) WithUrl(rawUrl string)(*ItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder) {
    return NewItemRoledefinitionsItemRolesettingRoleSettingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
