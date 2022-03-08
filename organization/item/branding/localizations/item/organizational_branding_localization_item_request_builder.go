package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/backgroundimage"
    i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/squarelogo"
    ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/favicon"
    ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding/localizations/item/bannerlogo"
)

// OrganizationalBrandingLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
type OrganizationalBrandingLocalizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OrganizationalBrandingLocalizationItemRequestBuilderDeleteOptions options for Delete
type OrganizationalBrandingLocalizationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrganizationalBrandingLocalizationItemRequestBuilderGetOptions options for Get
type OrganizationalBrandingLocalizationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters add different branding based on a locale.
type OrganizationalBrandingLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OrganizationalBrandingLocalizationItemRequestBuilderPatchOptions options for Patch
type OrganizationalBrandingLocalizationItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingLocalizationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) BackgroundImage()(*i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d.BackgroundImageRequestBuilder) {
    return i6a7478100780912a5be031e6c8d790dc49405d68dd9b8c7f08e4ce316b6b8e3d.NewBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) BannerLogo()(*ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae.BannerLogoRequestBuilder) {
    return ic2fbe7099579dba0316764b0f81d4a843f3004e003c2609224543e9897d0b2ae.NewBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationalBrandingLocalizationItemRequestBuilderInternal instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationalBrandingLocalizationItemRequestBuilder) {
    m := &OrganizationalBrandingLocalizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}/branding/localizations/{organizationalBrandingLocalization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationalBrandingLocalizationItemRequestBuilder instantiates a new OrganizationalBrandingLocalizationItemRequestBuilder and sets the default values.
func NewOrganizationalBrandingLocalizationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property localizations for organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreateDeleteRequestInformation(options *OrganizationalBrandingLocalizationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation add different branding based on a locale.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreateGetRequestInformation(options *OrganizationalBrandingLocalizationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property localizations in organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) CreatePatchRequestInformation(options *OrganizationalBrandingLocalizationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property localizations for organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Delete(options *OrganizationalBrandingLocalizationItemRequestBuilderDeleteOptions)(error) {
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
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Favicon()(*ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3.FaviconRequestBuilder) {
    return ia7baad9e196f596f211db6cea3f24b75276809ee9ba116821ef97852759ddad3.NewFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get add different branding based on a locale.
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Get(options *OrganizationalBrandingLocalizationItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingLocalizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOrganizationalBrandingLocalizationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationalBrandingLocalizationable), nil
}
// Patch update the navigation property localizations in organization
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) Patch(options *OrganizationalBrandingLocalizationItemRequestBuilderPatchOptions)(error) {
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
func (m *OrganizationalBrandingLocalizationItemRequestBuilder) SquareLogo()(*i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967.SquareLogoRequestBuilder) {
    return i9d66c98e71dc6ac8e5d6be7ba84577c585b7605c426a5121a43bdc47b7842967.NewSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
