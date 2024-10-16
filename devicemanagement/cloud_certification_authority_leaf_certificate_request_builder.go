package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudCertificationAuthorityLeafCertificateRequestBuilder provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.deviceManagement entity.
type CloudCertificationAuthorityLeafCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters collection of CloudCertificationAuthorityLeafCertificate records associated with account.
type CloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters struct {
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
// CloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudCertificationAuthorityLeafCertificateRequestBuilderGetQueryParameters
}
// CloudCertificationAuthorityLeafCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudCertificationAuthorityLeafCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudCertificationAuthorityLeafCertificateId provides operations to manage the cloudCertificationAuthorityLeafCertificate property of the microsoft.graph.deviceManagement entity.
// returns a *CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder when successful
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) ByCloudCertificationAuthorityLeafCertificateId(cloudCertificationAuthorityLeafCertificateId string)(*CloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudCertificationAuthorityLeafCertificateId != "" {
        urlTplParams["cloudCertificationAuthorityLeafCertificate%2Did"] = cloudCertificationAuthorityLeafCertificateId
    }
    return NewCloudCertificationAuthorityLeafCertificateCloudCertificationAuthorityLeafCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudCertificationAuthorityLeafCertificateRequestBuilderInternal instantiates a new CloudCertificationAuthorityLeafCertificateRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityLeafCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityLeafCertificateRequestBuilder) {
    m := &CloudCertificationAuthorityLeafCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudCertificationAuthorityLeafCertificate{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudCertificationAuthorityLeafCertificateRequestBuilder instantiates a new CloudCertificationAuthorityLeafCertificateRequestBuilder and sets the default values.
func NewCloudCertificationAuthorityLeafCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudCertificationAuthorityLeafCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudCertificationAuthorityLeafCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudCertificationAuthorityLeafCertificateCountRequestBuilder when successful
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) Count()(*CloudCertificationAuthorityLeafCertificateCountRequestBuilder) {
    return NewCloudCertificationAuthorityLeafCertificateCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of CloudCertificationAuthorityLeafCertificate records associated with account.
// returns a CloudCertificationAuthorityLeafCertificateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateCollectionResponseable, error) {
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
// Post create new navigation property to cloudCertificationAuthorityLeafCertificate for deviceManagement
// returns a CloudCertificationAuthorityLeafCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, requestConfiguration *CloudCertificationAuthorityLeafCertificateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable), nil
}
// ToGetRequestInformation collection of CloudCertificationAuthorityLeafCertificate records associated with account.
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudCertificationAuthorityLeafCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to cloudCertificationAuthorityLeafCertificate for deviceManagement
// returns a *RequestInformation when successful
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudCertificationAuthorityLeafCertificateable, requestConfiguration *CloudCertificationAuthorityLeafCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudCertificationAuthorityLeafCertificateRequestBuilder when successful
func (m *CloudCertificationAuthorityLeafCertificateRequestBuilder) WithUrl(rawUrl string)(*CloudCertificationAuthorityLeafCertificateRequestBuilder) {
    return NewCloudCertificationAuthorityLeafCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
