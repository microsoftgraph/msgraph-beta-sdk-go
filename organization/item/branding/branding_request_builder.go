package branding

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BrandingRequestBuilderDeleteOptions options for Delete
type BrandingRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BrandingRequestBuilderGetOptions options for Get
type BrandingRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BrandingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BrandingRequestBuilderGetQueryParameters get branding from organization
type BrandingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BrandingRequestBuilderPatchOptions options for Patch
type BrandingRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BrandingRequestBuilder) BackgroundImage()(*id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.BackgroundImageRequestBuilder) {
    return id1e3e547ca552ddc7d060d1dce939722dd18581fb9a5143610155bad7476a574.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BrandingRequestBuilder) BannerLogo()(*i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25.BannerLogoRequestBuilder) {
    return i7bdb509589c083f5cd687ccb30e62736aa8c72f641f5bcf6328e4a655e072b25.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BrandingRequestBuilder) {
    m := &BrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}/branding{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBrandingRequestBuilder instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BrandingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBrandingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property branding for organization
func (m *BrandingRequestBuilder) CreateDeleteRequestInformation(options *BrandingRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get branding from organization
func (m *BrandingRequestBuilder) CreateGetRequestInformation(options *BrandingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property branding in organization
func (m *BrandingRequestBuilder) CreatePatchRequestInformation(options *BrandingRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *BrandingRequestBuilder) Favicon()(*i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.FaviconRequestBuilder) {
    return i9b7f20e9ef38f4aefd99cbe38e2c0b1a9b0e0df96bc3c1f8e0e5588c555411b5.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get branding from organization
func (m *BrandingRequestBuilder) Get(options *BrandingRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOrganizationalBrandingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingable), nil
}
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
        urlTplParams["organizationalBrandingLocalization_id"] = id
    }
    return i43dcd1c55663be982ab2636ede2d587f63bd26aacf6715f4d6076aad7e3256ff.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property branding in organization
func (m *BrandingRequestBuilder) Patch(options *BrandingRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *BrandingRequestBuilder) SquareLogo()(*i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.SquareLogoRequestBuilder) {
    return i250ae21e56996eaf8b3a69882264a2c7de924d1e4f067694db380361586a9e3d.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
