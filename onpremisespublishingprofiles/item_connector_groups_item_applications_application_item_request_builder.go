package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
type ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters get applications from onPremisesPublishingProfiles
type ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters
}
// NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) {
    m := &ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/applications/{application%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get applications from onPremisesPublishingProfiles
func (m *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// Logo provides operations to manage the media for the onPremisesPublishingProfile entity.
func (m *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) Logo()(*ItemConnectorGroupsItemApplicationsItemLogoRequestBuilder) {
    return NewItemConnectorGroupsItemApplicationsItemLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get applications from onPremisesPublishingProfiles
func (m *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder) {
    return NewItemConnectorGroupsItemApplicationsApplicationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
