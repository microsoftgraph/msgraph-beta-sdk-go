package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetQueryParameters get a list of the customAccessPackageWorkflowExtension objects and their properties. The resulting list includes all the customAccessPackageWorkflowExtension objects for the catalog that the caller has access to read.
type EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomAccessPackageWorkflowExtensionId provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) ByCustomAccessPackageWorkflowExtensionId(customAccessPackageWorkflowExtensionId string)(*EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customAccessPackageWorkflowExtensionId != "" {
        urlTplParams["customAccessPackageWorkflowExtension%2Did"] = customAccessPackageWorkflowExtensionId
    }
    return NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    m := &EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/customAccessPackageWorkflowExtensions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder instantiates a new EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) Count()(*EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the customAccessPackageWorkflowExtension objects and their properties. The resulting list includes all the customAccessPackageWorkflowExtension objects for the catalog that the caller has access to read.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomAccessPackageWorkflowExtensionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-list-customaccesspackageworkflowextensions?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAccessPackageWorkflowExtensionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionCollectionResponseable), nil
}
// Post create a new customAccessPackageWorkflowExtension object and add it to an existing accessPackageCatalog object.  
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomAccessPackageWorkflowExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-post-customaccesspackageworkflowextensions?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable), nil
}
// ToGetRequestInformation get a list of the customAccessPackageWorkflowExtension objects and their properties. The resulting list includes all the customAccessPackageWorkflowExtension objects for the catalog that the caller has access to read.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new customAccessPackageWorkflowExtension object and add it to an existing accessPackageCatalog object.  
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemCustomaccesspackageworkflowextensionsCustomAccessPackageWorkflowExtensionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
