package branding

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/backgroundimage"
    i1b40b605aeb59b28ba3fa6abed7afbe4b8f7d36ecb4da595918a6e40e86b5489 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/squarelogodark"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/squarelogo"
    i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/favicon"
    ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/bannerlogo"
    icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/localizations"
    i71fbeabc44f0c477a20e1c3a1515ecf3df8929b67e5f5aeb33b83cf6335140ad "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/localizations/item"
)

// BrandingRequestBuilder provides operations to manage the organizationalBranding singleton.
type BrandingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BrandingRequestBuilderGetQueryParameters get branding
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
// BackgroundImage provides operations to manage the media for the organizationalBranding entity.
func (m *BrandingRequestBuilder) BackgroundImage()(*i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.BackgroundImageRequestBuilder) {
    return i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo provides operations to manage the media for the organizationalBranding entity.
func (m *BrandingRequestBuilder) BannerLogo()(*ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973.BannerLogoRequestBuilder) {
    return ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrandingRequestBuilder) {
    m := &BrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/branding{?%24select,%24expand}";
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
// CreateGetRequestInformation get branding
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
// CreatePatchRequestInformation update branding
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
// Favicon provides operations to manage the media for the organizationalBranding entity.
func (m *BrandingRequestBuilder) Favicon()(*i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.FaviconRequestBuilder) {
    return i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get branding
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
// Localizations provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *BrandingRequestBuilder) Localizations()(*icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87.LocalizationsRequestBuilder) {
    return icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *BrandingRequestBuilder) LocalizationsById(id string)(*i71fbeabc44f0c477a20e1c3a1515ecf3df8929b67e5f5aeb33b83cf6335140ad.OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization%2Did"] = id
    }
    return i71fbeabc44f0c477a20e1c3a1515ecf3df8929b67e5f5aeb33b83cf6335140ad.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update branding
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
// SquareLogo provides operations to manage the media for the organizationalBranding entity.
func (m *BrandingRequestBuilder) SquareLogo()(*i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.SquareLogoRequestBuilder) {
    return i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark provides operations to manage the media for the organizationalBranding entity.
func (m *BrandingRequestBuilder) SquareLogoDark()(*i1b40b605aeb59b28ba3fa6abed7afbe4b8f7d36ecb4da595918a6e40e86b5489.SquareLogoDarkRequestBuilder) {
    return i1b40b605aeb59b28ba3fa6abed7afbe4b8f7d36ecb4da595918a6e40e86b5489.NewSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
