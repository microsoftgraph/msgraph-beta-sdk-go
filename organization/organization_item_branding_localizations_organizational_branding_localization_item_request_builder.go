package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
type OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters add different branding based on a locale.
type OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters
}
// OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackgroundImage provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) BackgroundImage()(*OrganizationItemBrandingLocalizationsItemBackgroundImageRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) BannerLogo()(*OrganizationItemBrandingLocalizationsItemBannerLogoRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    m := &OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder{
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
// NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property localizations for organization
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) CustomCSS()(*OrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemCustomCSSRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property localizations for organization
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Favicon()(*OrganizationItemBrandingLocalizationsItemFaviconRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get add different branding based on a locale.
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) HeaderLogo()(*OrganizationItemBrandingLocalizationsItemHeaderLogoRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemHeaderLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property localizations in organization
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, requestConfiguration *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingLocalizationable, error) {
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
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) SquareLogo()(*OrganizationItemBrandingLocalizationsItemSquareLogoRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) SquareLogoDark()(*OrganizationItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
