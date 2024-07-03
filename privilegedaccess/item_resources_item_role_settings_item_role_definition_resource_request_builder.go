package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder provides operations to manage the resource property of the microsoft.graph.governanceRoleDefinition entity.
type ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters read-only. The associated resource for the role definition.
type ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters
}
// NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal instantiates a new ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder and sets the default values.
func NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    m := &ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleSettings/{governanceRoleSetting%2Did}/roleDefinition/resource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder instantiates a new ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder and sets the default values.
func NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only. The associated resource for the role definition.
// returns a GovernanceResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// ToGetRequestInformation read-only. The associated resource for the role definition.
// returns a *RequestInformation when successful
func (m *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder when successful
func (m *ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    return NewItemResourcesItemRoleSettingsItemRoleDefinitionResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
