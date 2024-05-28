package governanceroledefinitions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRolesettingRoleSettingRequestBuilder provides operations to manage the roleSetting property of the microsoft.graph.governanceRoleDefinition entity.
type ItemRolesettingRoleSettingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRolesettingRoleSettingRequestBuilderGetQueryParameters the associated role setting for the role definition.
type ItemRolesettingRoleSettingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRolesettingRoleSettingRequestBuilderGetQueryParameters
}
// NewItemRolesettingRoleSettingRequestBuilderInternal instantiates a new ItemRolesettingRoleSettingRequestBuilder and sets the default values.
func NewItemRolesettingRoleSettingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRolesettingRoleSettingRequestBuilder) {
    m := &ItemRolesettingRoleSettingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/governanceRoleDefinitions/{governanceRoleDefinition%2Did}/roleSetting{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRolesettingRoleSettingRequestBuilder instantiates a new ItemRolesettingRoleSettingRequestBuilder and sets the default values.
func NewItemRolesettingRoleSettingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRolesettingRoleSettingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRolesettingRoleSettingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the associated role setting for the role definition.
// returns a GovernanceRoleSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRolesettingRoleSettingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleSettingable, error) {
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
func (m *ItemRolesettingRoleSettingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRolesettingRoleSettingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRolesettingRoleSettingRequestBuilder when successful
func (m *ItemRolesettingRoleSettingRequestBuilder) WithUrl(rawUrl string)(*ItemRolesettingRoleSettingRequestBuilder) {
    return NewItemRolesettingRoleSettingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
