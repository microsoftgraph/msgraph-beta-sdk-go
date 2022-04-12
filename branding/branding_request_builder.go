package branding

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/backgroundimage"
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
// BrandingRequestBuilderGetQueryParameters get branding
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
func (m *BrandingRequestBuilder) BackgroundImage()(*i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.BackgroundImageRequestBuilder) {
    return i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo the bannerLogo property
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
// CreatePatchRequestInformation update branding
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
// Favicon the favicon property
func (m *BrandingRequestBuilder) Favicon()(*i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.FaviconRequestBuilder) {
    return i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get branding
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
func (m *BrandingRequestBuilder) Localizations()(*icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87.LocalizationsRequestBuilder) {
    return icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.branding.localizations.item collection
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
func (m *BrandingRequestBuilder) SquareLogo()(*i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.SquareLogoRequestBuilder) {
    return i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
