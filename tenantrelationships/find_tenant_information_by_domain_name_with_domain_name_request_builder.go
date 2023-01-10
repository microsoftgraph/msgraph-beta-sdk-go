package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FindTenantInformationByDomainNameWithDomainNameRequestBuilder provides operations to call the findTenantInformationByDomainName method.
type FindTenantInformationByDomainNameWithDomainNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal instantiates a new FindTenantInformationByDomainNameWithDomainNameRequestBuilder and sets the default values.
func NewFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*FindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    m := &FindTenantInformationByDomainNameWithDomainNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/microsoft.graph.findTenantInformationByDomainName(domainName='{domainName}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if domainName != nil {
        urlTplParams["domainName"] = *domainName
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFindTenantInformationByDomainNameWithDomainNameRequestBuilder instantiates a new FindTenantInformationByDomainNameWithDomainNameRequestBuilder and sets the default values.
func NewFindTenantInformationByDomainNameWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findTenantInformationByDomainName
func (m *FindTenantInformationByDomainNameWithDomainNameRequestBuilder) Get(ctx context.Context, requestConfiguration *FindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantInformationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantInformationable), nil
}
// ToGetRequestInformation invoke function findTenantInformationByDomainName
func (m *FindTenantInformationByDomainNameWithDomainNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
