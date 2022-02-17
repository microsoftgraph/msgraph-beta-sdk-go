package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/addkey"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/removekey"
    i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/addpassword"
    i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/federatedidentitycredentials"
    i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/connectorgroup"
    i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/restore"
    i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/homerealmdiscoverypolicies"
    i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/getmemberobjects"
    i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization"
    i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/getmembergroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/extensionproperties"
    i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies"
    i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/logo"
    i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/createdonbehalfof"
    i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenissuancepolicies"
    i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/checkmembergroups"
    ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenlifetimepolicies"
    ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/checkmemberobjects"
    ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/setverifiedpublisher"
    id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners"
    if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/unsetverifiedpublisher"
    ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/removepassword"
    i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/extensionproperties/item"
    icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/federatedidentitycredentials/item"
)

// ApplicationRequestBuilder builds and executes requests for operations under \applications\{application-id}
type ApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApplicationRequestBuilderDeleteOptions options for Delete
type ApplicationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplicationRequestBuilderGetOptions options for Get
type ApplicationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ApplicationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplicationRequestBuilderGetQueryParameters get entity from applications by key
type ApplicationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ApplicationRequestBuilderPatchOptions options for Patch
type ApplicationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ApplicationRequestBuilder) AddKey()(*i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a.AddKeyRequestBuilder) {
    return i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) AddPassword()(*i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa.AddPasswordRequestBuilder) {
    return i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) AppManagementPolicies()(*i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f.AppManagementPoliciesRequestBuilder) {
    return i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) CheckMemberGroups()(*i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86.CheckMemberGroupsRequestBuilder) {
    return i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) CheckMemberObjects()(*ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49.CheckMemberObjectsRequestBuilder) {
    return ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) ConnectorGroup()(*i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698.ConnectorGroupRequestBuilder) {
    return i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698.NewConnectorGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewApplicationRequestBuilderInternal instantiates a new ApplicationRequestBuilder and sets the default values.
func NewApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationRequestBuilder) {
    m := &ApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationRequestBuilder instantiates a new ApplicationRequestBuilder and sets the default values.
func NewApplicationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from applications
func (m *ApplicationRequestBuilder) CreateDeleteRequestInformation(options *ApplicationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ApplicationRequestBuilder) CreatedOnBehalfOf()(*i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4.CreatedOnBehalfOfRequestBuilder) {
    return i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entity from applications by key
func (m *ApplicationRequestBuilder) CreateGetRequestInformation(options *ApplicationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in applications
func (m *ApplicationRequestBuilder) CreatePatchRequestInformation(options *ApplicationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from applications
func (m *ApplicationRequestBuilder) Delete(options *ApplicationRequestBuilderDeleteOptions)(error) {
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
func (m *ApplicationRequestBuilder) ExtensionProperties()(*i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b.ExtensionPropertiesRequestBuilder) {
    return i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b.NewExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.extensionProperties.item collection
func (m *ApplicationRequestBuilder) ExtensionPropertiesById(id string)(*i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da.ExtensionPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extensionProperty_id"] = id
    }
    return i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da.NewExtensionPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) FederatedIdentityCredentials()(*i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9.FederatedIdentityCredentialsRequestBuilder) {
    return i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9.NewFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.federatedIdentityCredentials.item collection
func (m *ApplicationRequestBuilder) FederatedIdentityCredentialsById(id string)(*icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2.FederatedIdentityCredentialRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential_id"] = id
    }
    return icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2.NewFederatedIdentityCredentialRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from applications by key
func (m *ApplicationRequestBuilder) Get(options *ApplicationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApplication() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application), nil
}
func (m *ApplicationRequestBuilder) GetMemberGroups()(*i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991.GetMemberGroupsRequestBuilder) {
    return i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) GetMemberObjects()(*i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5.GetMemberObjectsRequestBuilder) {
    return i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) HomeRealmDiscoveryPolicies()(*i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Logo()(*i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c.LogoRequestBuilder) {
    return i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c.NewLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Owners()(*id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90.OwnersRequestBuilder) {
    return id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in applications
func (m *ApplicationRequestBuilder) Patch(options *ApplicationRequestBuilderPatchOptions)(error) {
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
func (m *ApplicationRequestBuilder) RemoveKey()(*i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b.RemoveKeyRequestBuilder) {
    return i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) RemovePassword()(*ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097.RemovePasswordRequestBuilder) {
    return ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Restore()(*i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9.RestoreRequestBuilder) {
    return i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) SetVerifiedPublisher()(*ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b.SetVerifiedPublisherRequestBuilder) {
    return ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b.NewSetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Synchronization()(*i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c.SynchronizationRequestBuilder) {
    return i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c.NewSynchronizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) TokenIssuancePolicies()(*i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b.TokenIssuancePoliciesRequestBuilder) {
    return i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) TokenLifetimePolicies()(*ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30.TokenLifetimePoliciesRequestBuilder) {
    return ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) UnsetVerifiedPublisher()(*if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0.UnsetVerifiedPublisherRequestBuilder) {
    return if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0.NewUnsetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
