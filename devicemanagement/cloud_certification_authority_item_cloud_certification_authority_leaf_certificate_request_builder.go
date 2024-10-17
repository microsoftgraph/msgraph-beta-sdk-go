package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.cloudCertificationAuthority entity.
type CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters required OData property to expose leaf certificate API.
type CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters struct {
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
// CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters
}
// ByCloudCertificationAuthorityLeafCertificateId provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.cloudCertificationAuthority entity.
// returns a *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder when successful
func (m *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) ByCloudCertificationAuthorityLeafCertificateId(cloudCertificationAuthorityLeafCertificateId string)(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudCertificationAuthorityLeafCertificateId != "" {
        urlTplParams["cloudCertificationAuthorityLeafCertificate%2Did"] = cloudCertificationAuthorityLeafCertificateId
    }
    return NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderInternal instantiates a new CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) {
    m := &CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthority/{cloudCertificationAuthority%2Did}/cloudCertificationAuthorityLeafCertificate{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder instantiates a new CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCountRequestBuilder when successful
func (m *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) Count()(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCountRequestBuilder) {
    return NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get required OData property to expose leaf certificate API.
// returns a CloudCertificationAuthorityLeafCertificateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityLeafCertificateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateCollectionResponseable), nil
}
// ToGetRequestInformation required OData property to expose leaf certificate API.
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityItemCloudCertificationAuthorityLeafCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
