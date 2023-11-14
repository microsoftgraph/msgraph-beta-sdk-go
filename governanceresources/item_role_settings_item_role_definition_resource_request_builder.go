package governanceresources

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder provides operations to manage the resource property of the microsoft.graph.governanceRoleDefinition entity.
type ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters read-only. The associated resource for the role definition.
type ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetQueryParameters
}
// NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal instantiates a new ResourceRequestBuilder and sets the default values.
func NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    m := &ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/governanceResources/{governanceResource%2Did}/roleSettings/{governanceRoleSetting%2Did}/roleDefinition/resource{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilder instantiates a new ResourceRequestBuilder and sets the default values.
func NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only. The associated resource for the role definition.
func (m *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) WithUrl(rawUrl string)(*ItemRoleSettingsItemRoleDefinitionResourceRequestBuilder) {
    return NewItemRoleSettingsItemRoleDefinitionResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
