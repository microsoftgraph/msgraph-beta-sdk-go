package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
type EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetQueryParameters the Windows Enterprise Code Signing Certificate.
type EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetQueryParameters struct {
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
// EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetQueryParameters
}
// EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEnterpriseCodeSigningCertificateId provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
// returns a *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificateItemRequestBuilder when successful
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) ByEnterpriseCodeSigningCertificateId(enterpriseCodeSigningCertificateId string)(*EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if enterpriseCodeSigningCertificateId != "" {
        urlTplParams["enterpriseCodeSigningCertificate%2Did"] = enterpriseCodeSigningCertificateId
    }
    return NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderInternal instantiates a new EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder and sets the default values.
func NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) {
    m := &EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/enterpriseCodeSigningCertificates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder instantiates a new EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder and sets the default values.
func NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EnterprisecodesigningcertificatesCountRequestBuilder when successful
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) Count()(*EnterprisecodesigningcertificatesCountRequestBuilder) {
    return NewEnterprisecodesigningcertificatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Windows Enterprise Code Signing Certificate.
// returns a EnterpriseCodeSigningCertificateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnterpriseCodeSigningCertificateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateCollectionResponseable), nil
}
// Post create new navigation property to enterpriseCodeSigningCertificates for deviceAppManagement
// returns a EnterpriseCodeSigningCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateable, requestConfiguration *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateable), nil
}
// ToGetRequestInformation the Windows Enterprise Code Signing Certificate.
// returns a *RequestInformation when successful
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to enterpriseCodeSigningCertificates for deviceAppManagement
// returns a *RequestInformation when successful
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnterpriseCodeSigningCertificateable, requestConfiguration *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder when successful
func (m *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) WithUrl(rawUrl string)(*EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) {
    return NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
