package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i058780e515c8b4a770f21fe6800400e16271cc8d83fed056e2470fc6b85b93e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/customcss"
    i5d295c595a8dfd5c3199dc9f3162028bd78a99181260faabe9ff1148a75a9bbf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/headerlogo"
    i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/backgroundimage"
    i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/squarelogo"
    ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/favicon"
    ib06e3fb2ab0d52cd388a43e277edd31fc53e151e7c2cbb2c7729364e258b5717 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/squarelogodark"
    ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/bannerlogo"
)

// OrganizationalBrandingLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
type OrganizationalBrandingLocalizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters add different branding based on a locale.
type OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters
}
// OrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackgroundImage provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) BackgroundImage()(*i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d.BackgroundImageRequestBuilder) {
    return i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) BannerLogo()(*ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae.BannerLogoRequestBuilder) {
    return ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationalBrandingLocalizationItemRequestBuilderInternal instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationalBrandingLocalizationItemRequestBuilder) {
    m := &OrganizationalBrandingLocalizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationalBrandingLocalizationItemRequestBuilder instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationalBrandingLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property localizations for organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation add different branding based on a locale.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property localizations in organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CustomCSS provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CustomCSS()(*i058780e515c8b4a770f21fe6800400e16271cc8d83fed056e2470fc6b85b93e0.CustomCSSRequestBuilder) {
    return i058780e515c8b4a770f21fe6800400e16271cc8d83fed056e2470fc6b85b93e0.NewCustomCSSRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property localizations for organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Favicon provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Favicon()(*ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3.FaviconRequestBuilder) {
    return ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get add different branding based on a locale.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable), nil
}
// HeaderLogo provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) HeaderLogo()(*i5d295c595a8dfd5c3199dc9f3162028bd78a99181260faabe9ff1148a75a9bbf.HeaderLogoRequestBuilder) {
    return i5d295c595a8dfd5c3199dc9f3162028bd78a99181260faabe9ff1148a75a9bbf.NewHeaderLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property localizations in organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, requestConfiguration *OrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable), nil
}
// SquareLogo provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) SquareLogo()(*i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967.SquareLogoRequestBuilder) {
    return i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) SquareLogoDark()(*ib06e3fb2ab0d52cd388a43e277edd31fc53e151e7c2cbb2c7729364e258b5717.SquareLogoDarkRequestBuilder) {
    return ib06e3fb2ab0d52cd388a43e277edd31fc53e151e7c2cbb2c7729364e258b5717.NewSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
