package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
type ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetQueryParameters get applications from onPremisesPublishingProfiles
type ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetQueryParameters
}
// NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderInternal instantiates a new ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, uniqueName *string)(*ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) {
    m := &ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/applications(uniqueName='{uniqueName}'){?%24expand,%24select}", pathParameters),
    }
    if uniqueName != nil {
        m.BaseRequestBuilder.PathParameters["uniqueName"] = *uniqueName
    }
    return m
}
// NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder instantiates a new ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get applications from onPremisesPublishingProfiles
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get applications from onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder when successful
func (m *ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationswithuniquenameApplicationsWithUniqueNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
