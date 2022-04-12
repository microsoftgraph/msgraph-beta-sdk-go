package branding

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/squarelogo"
    i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/bannerlogo"
    i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/favicon"
    ibd06dae6052c9c44018f5200f072f15bc200a78f5b996495a11369beb8cb87c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations"
    id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/backgroundimage"
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
// BrandingRequestBuilderDeleteOptions options for Delete
type BrandingRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// BrandingRequestBuilderGetOptions options for Get
type BrandingRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BrandingRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// BrandingRequestBuilderGetQueryParameters get branding from organization
type BrandingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BrandingRequestBuilderPatchOptions options for Patch
type BrandingRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *BrandingRequestBuilder) CreateDeleteRequestInformation(options *BrandingRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get branding from organization
func (m *BrandingRequestBuilder) CreateGetRequestInformation(options *BrandingRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property branding in organization
func (m *BrandingRequestBuilder) CreatePatchRequestInformation(options *BrandingRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property branding for organization
func (m *BrandingRequestBuilder) Delete(options *BrandingRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Favicon the favicon property
func (m *BrandingRequestBuilder) Favicon()(*i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.FaviconRequestBuilder) {
    return i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get branding from organization
func (m *BrandingRequestBuilder) Get(options *BrandingRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
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
func (m *BrandingRequestBuilder) Patch(options *BrandingRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SquareLogo the squareLogo property
func (m *BrandingRequestBuilder) SquareLogo()(*i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.SquareLogoRequestBuilder) {
    return i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
