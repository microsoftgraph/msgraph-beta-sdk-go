package branding

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/squarelogo"
    i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/bannerlogo"
    i8ab8f774a5a97f1d672f15ad971f6404b3a1261eba32549ee9704bbf936b05f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/customcss"
    i957b7effb04db0e2b1c5cea3a7dc2d5e14018afff7d8e350d23de1228772c818 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/headerlogo"
    i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/favicon"
    ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations"
    id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/backgroundimage"
    idc11574055e00ea67da03c794a8998ec5c5b8ce2eb6c7c68aeed2f0563e1f802 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/squarelogodark"
    i43dcd1c55663be982ab2636ede2d587f63bd26aacf6715f4d6076aad7e3256ff "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item"
)

// BrandingRequestBuilder provides operations to manage the branding property of the microsoft.graph.organization entity.
type BrandingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BrandingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrandingRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BrandingRequestBuilderGetQueryParameters retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
type BrandingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BrandingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrandingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BrandingRequestBuilderGetQueryParameters
}
// BrandingRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrandingRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackgroundImage provides operations to manage the media for the organization entity.
func (m *BrandingRequestBuilder) BackgroundImage()(*id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.BackgroundImageRequestBuilder) {
    return id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo provides operations to manage the media for the organization entity.
func (m *BrandingRequestBuilder) BannerLogo()(*i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25.BannerLogoRequestBuilder) {
    return i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrandingRequestBuilder) {
    m := &BrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/branding{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBrandingRequestBuilder instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrandingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBrandingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
func (m *BrandingRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *BrandingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
func (m *BrandingRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BrandingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of the default branding object specified by the organizationalBranding resource.
func (m *BrandingRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *BrandingRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *BrandingRequestBuilder) CustomCSS()(*i8ab8f774a5a97f1d672f15ad971f6404b3a1261eba32549ee9704bbf936b05f4.CustomCSSRequestBuilder) {
    return i8ab8f774a5a97f1d672f15ad971f6404b3a1261eba32549ee9704bbf936b05f4.NewCustomCSSRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
func (m *BrandingRequestBuilder) Delete(ctx context.Context, requestConfiguration *BrandingRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *BrandingRequestBuilder) Favicon()(*i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.FaviconRequestBuilder) {
    return i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
func (m *BrandingRequestBuilder) Get(ctx context.Context, requestConfiguration *BrandingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable), nil
}
// HeaderLogo provides operations to manage the media for the organization entity.
func (m *BrandingRequestBuilder) HeaderLogo()(*i957b7effb04db0e2b1c5cea3a7dc2d5e14018afff7d8e350d23de1228772c818.HeaderLogoRequestBuilder) {
    return i957b7effb04db0e2b1c5cea3a7dc2d5e14018afff7d8e350d23de1228772c818.NewHeaderLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Localizations provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *BrandingRequestBuilder) Localizations()(*ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9.LocalizationsRequestBuilder) {
    return ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *BrandingRequestBuilder) LocalizationsById(id string)(*i43dcd1c55663be982ab2636ede2d587f63bd26aacf6715f4d6076aad7e3256ff.OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization%2Did"] = id
    }
    return i43dcd1c55663be982ab2636ede2d587f63bd26aacf6715f4d6076aad7e3256ff.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of the default branding object specified by the organizationalBranding resource.
func (m *BrandingRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *BrandingRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable), nil
}
// SquareLogo provides operations to manage the media for the organization entity.
func (m *BrandingRequestBuilder) SquareLogo()(*i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.SquareLogoRequestBuilder) {
    return i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *BrandingRequestBuilder) SquareLogoDark()(*idc11574055e00ea67da03c794a8998ec5c5b8ce2eb6c7c68aeed2f0563e1f802.SquareLogoDarkRequestBuilder) {
    return idc11574055e00ea67da03c794a8998ec5c5b8ce2eb6c7c68aeed2f0563e1f802.NewSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
