// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/teamsadministration"
)

// TeamsUserConfigurationsRequestBuilder provides operations to manage the userConfigurations property of the microsoft.graph.teamsAdministration.teamsAdminRoot entity.
type TeamsUserConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsUserConfigurationsRequestBuilderGetQueryParameters get user configurations for all Teams users who belong to a tenant.
type TeamsUserConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TeamsUserConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsUserConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsUserConfigurationsRequestBuilderGetQueryParameters
}
// TeamsUserConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsUserConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTeamsUserConfigurationId provides operations to manage the userConfigurations property of the microsoft.graph.teamsAdministration.teamsAdminRoot entity.
// returns a *TeamsUserConfigurationsTeamsUserConfigurationItemRequestBuilder when successful
func (m *TeamsUserConfigurationsRequestBuilder) ByTeamsUserConfigurationId(teamsUserConfigurationId string)(*TeamsUserConfigurationsTeamsUserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if teamsUserConfigurationId != "" {
        urlTplParams["teamsUserConfiguration%2Did"] = teamsUserConfigurationId
    }
    return NewTeamsUserConfigurationsTeamsUserConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamsUserConfigurationsRequestBuilderInternal instantiates a new TeamsUserConfigurationsRequestBuilder and sets the default values.
func NewTeamsUserConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsUserConfigurationsRequestBuilder) {
    m := &TeamsUserConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/teams/userConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamsUserConfigurationsRequestBuilder instantiates a new TeamsUserConfigurationsRequestBuilder and sets the default values.
func NewTeamsUserConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsUserConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsUserConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamsUserConfigurationsCountRequestBuilder when successful
func (m *TeamsUserConfigurationsRequestBuilder) Count()(*TeamsUserConfigurationsCountRequestBuilder) {
    return NewTeamsUserConfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get user configurations for all Teams users who belong to a tenant.
// returns a TeamsUserConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamsadministration-teamsadminroot-list-userconfigurations?view=graph-rest-beta
func (m *TeamsUserConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsUserConfigurationsRequestBuilderGetRequestConfiguration)(ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.CreateTeamsUserConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationCollectionResponseable), nil
}
// Post create new navigation property to userConfigurations for admin
// returns a TeamsUserConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsUserConfigurationsRequestBuilder) Post(ctx context.Context, body ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationable, requestConfiguration *TeamsUserConfigurationsRequestBuilderPostRequestConfiguration)(ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.CreateTeamsUserConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationable), nil
}
// ToGetRequestInformation get user configurations for all Teams users who belong to a tenant.
// returns a *RequestInformation when successful
func (m *TeamsUserConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsUserConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userConfigurations for admin
// returns a *RequestInformation when successful
func (m *TeamsUserConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ida00aeab33baad931d9ea7ddda72fd1a74178962fdc58f5709ddf52b9f0fcc0f.TeamsUserConfigurationable, requestConfiguration *TeamsUserConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamsUserConfigurationsRequestBuilder when successful
func (m *TeamsUserConfigurationsRequestBuilder) WithUrl(rawUrl string)(*TeamsUserConfigurationsRequestBuilder) {
    return NewTeamsUserConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
