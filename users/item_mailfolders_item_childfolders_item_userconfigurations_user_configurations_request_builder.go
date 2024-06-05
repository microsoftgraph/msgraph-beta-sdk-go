package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
type ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetQueryParameters get userConfigurations from users
type ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetQueryParameters struct {
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
// ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetQueryParameters
}
// ByUserConfigurationId provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
// returns a *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationItemRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) ByUserConfigurationId(userConfigurationId string)(*ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userConfigurationId != "" {
        urlTplParams["userConfiguration%2Did"] = userConfigurationId
    }
    return NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderInternal instantiates a new ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) {
    m := &ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/userConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder instantiates a new ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemMailfoldersItemChildfoldersItemUserconfigurationsCountRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) Count()(*ItemMailfoldersItemChildfoldersItemUserconfigurationsCountRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemUserconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get userConfigurations from users
// returns a UserConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConfigurationCollectionResponseable), nil
}
// ToGetRequestInformation get userConfigurations from users
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemUserconfigurationsUserConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
