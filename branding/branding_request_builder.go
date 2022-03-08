package branding

import (
    i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/backgroundimage"
    i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/squarelogo"
    i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/favicon"
    ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/bannerlogo"
    icb26067e77492c83cb4e87ec1e3f5164c5a2b6eccb9e23f4a03609a175e65c87 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/localizations"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i71fbeabc44f0c477a20e1c3a1515ecf3df8929b67e5f5aeb33b83cf6335140ad "github.com/microsoftgraph/msgraph-beta-sdk-go/branding/localizations/item"
)

// BrandingRequestBuilder provides operations to manage the organizationalBranding singleton.
type BrandingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
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
// BrandingRequestBuilderGetQueryParameters get branding
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
func (m *BrandingRequestBuilder) BackgroundImage()(*i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.BackgroundImageRequestBuilder) {
    return i0cc7ec2bb77cfb998ba8a1ad527a89fb9a51a4881030337bde331a8e4f476070.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BrandingRequestBuilder) BannerLogo()(*ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973.BannerLogoRequestBuilder) {
    return ib0aa6f52ba8b83b6b9f06a69897f4792f8858acad7d73aaf3dbe9c3583eb1973.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BrandingRequestBuilder) {
    m := &BrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/branding{?select,expand}";
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
// CreateGetRequestInformation get branding
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
// CreatePatchRequestInformation update branding
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
func (m *BrandingRequestBuilder) Favicon()(*i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.FaviconRequestBuilder) {
    return i6dbe8712bcbf565967e359a0b5474825946d283c9a8f7154520173aba185bbdb.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get branding
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
        urlTplParams["organizationalBrandingLocalization_id"] = id
    }
    return i71fbeabc44f0c477a20e1c3a1515ecf3df8929b67e5f5aeb33b83cf6335140ad.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update branding
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
func (m *BrandingRequestBuilder) SquareLogo()(*i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.SquareLogoRequestBuilder) {
    return i2dc84c458003c3284caa395048d6882b017ff8d4d2f1da317816287aff8f011c.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
