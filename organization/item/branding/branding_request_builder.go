package branding

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/squarelogo"
    i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/bannerlogo"
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
// BrandingRequestBuilderGetQueryParameters resource to manage the default branding for the organization. Nullable.
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
// BackgroundImage the backgroundImage property
func (m *BrandingRequestBuilder) BackgroundImage()(*id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.BackgroundImageRequestBuilder) {
    return id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo the bannerLogo property
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
// CreateDeleteRequestInformation delete navigation property branding for organization
func (m *BrandingRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property branding for organization
func (m *BrandingRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *BrandingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation resource to manage the default branding for the organization. Nullable.
func (m *BrandingRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration resource to manage the default branding for the organization. Nullable.
func (m *BrandingRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BrandingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property branding in organization
func (m *BrandingRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property branding in organization
func (m *BrandingRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *BrandingRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property branding for organization
func (m *BrandingRequestBuilder) Delete(ctx context.Context, requestConfiguration *BrandingRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Favicon the favicon property
func (m *BrandingRequestBuilder) Favicon()(*i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.FaviconRequestBuilder) {
    return i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get resource to manage the default branding for the organization. Nullable.
func (m *BrandingRequestBuilder) Get(ctx context.Context, requestConfiguration *BrandingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Localizations the localizations property
func (m *BrandingRequestBuilder) Localizations()(*ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9.LocalizationsRequestBuilder) {
    return ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item.branding.localizations.item collection
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
// Patch update the navigation property branding in organization
func (m *BrandingRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *BrandingRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SquareLogo the squareLogo property
func (m *BrandingRequestBuilder) SquareLogo()(*i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.SquareLogoRequestBuilder) {
    return i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark the squareLogoDark property
func (m *BrandingRequestBuilder) SquareLogoDark()(*idc11574055e00ea67da03c794a8998ec5c5b8ce2eb6c7c68aeed2f0563e1f802.SquareLogoDarkRequestBuilder) {
    return idc11574055e00ea67da03c794a8998ec5c5b8ce2eb6c7c68aeed2f0563e1f802.NewSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
