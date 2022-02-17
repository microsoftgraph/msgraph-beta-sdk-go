package settings

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/profilecardproperties"
    ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/peopleinsights"
    ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/iteminsights"
    i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings/profilecardproperties/item"
)

// SettingsRequestBuilder builds and executes requests for operations under \organization\{organization-id}\settings
type SettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SettingsRequestBuilderDeleteOptions options for Delete
type SettingsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SettingsRequestBuilderGetOptions options for Get
type SettingsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SettingsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SettingsRequestBuilderGetQueryParameters retrieve the properties and relationships of organizationSettings object. Nullable.
type SettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SettingsRequestBuilderPatchOptions options for Patch
type SettingsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationSettings;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSettingsRequestBuilderInternal instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SettingsRequestBuilder) {
    m := &SettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}/settings{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSettingsRequestBuilder instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) CreateDeleteRequestInformation(options *SettingsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) CreateGetRequestInformation(options *SettingsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) CreatePatchRequestInformation(options *SettingsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) Delete(options *SettingsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) Get(options *SettingsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationSettings, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOrganizationSettings() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrganizationSettings), nil
}
func (m *SettingsRequestBuilder) ItemInsights()(*ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7.ItemInsightsRequestBuilder) {
    return ic8dc6584c7fcf76e945a3a597a167477455672164e6ecef2a585705bb95f06f7.NewItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *SettingsRequestBuilder) Patch(options *SettingsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *SettingsRequestBuilder) PeopleInsights()(*ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa.PeopleInsightsRequestBuilder) {
    return ib5d9323a069790a40d8250ed97dce0595a704227bc01a0ff026952b59c07b6fa.NewPeopleInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SettingsRequestBuilder) ProfileCardProperties()(*i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf.ProfileCardPropertiesRequestBuilder) {
    return i604f8b04e00513fe80f3d85b5361d4a35dbdefc791356fd0e5d7cb3820a457bf.NewProfileCardPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProfileCardPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item.settings.profileCardProperties.item collection
func (m *SettingsRequestBuilder) ProfileCardPropertiesById(id string)(*i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df.ProfileCardPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profileCardProperty_id"] = id
    }
    return i164665736586dec4189f70a600626997bf7f7b32e1d5debbea1adff37d4860df.NewProfileCardPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
