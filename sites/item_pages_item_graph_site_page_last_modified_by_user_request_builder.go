package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
type ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetQueryParameters get lastModifiedByUser from sites
type ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetQueryParameters
}
// NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderInternal instantiates a new ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) {
    m := &ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/lastModifiedByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder instantiates a new ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get lastModifiedByUser from sites
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
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
// MailboxSettings the mailboxSettings property
// returns a *ItemPagesItemGraphSitePageLastModifiedByUserMailboxSettingsRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) MailboxSettings()(*ItemPagesItemGraphSitePageLastModifiedByUserMailboxSettingsRequestBuilder) {
    return NewItemPagesItemGraphSitePageLastModifiedByUserMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *ItemPagesItemGraphSitePageLastModifiedByUserServiceProvisioningErrorsRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) ServiceProvisioningErrors()(*ItemPagesItemGraphSitePageLastModifiedByUserServiceProvisioningErrorsRequestBuilder) {
    return NewItemPagesItemGraphSitePageLastModifiedByUserServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get lastModifiedByUser from sites
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) {
    return NewItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
