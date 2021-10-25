package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/activateservice"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/getmembergroups"
    i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings"
    i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/restore"
    i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/checkmemberobjects"
    i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/extensions"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/getmemberobjects"
    ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/setmobiledevicemanagementauthority"
    ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/certificatebasedauthconfiguration"
    ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding"
    if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/checkmembergroups"
    i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/extensions/item"
)

type OrganizationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OrganizationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *OrganizationRequestBuilder) ActivateService()(*i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f.ActivateServiceRequestBuilder) {
    return i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f.NewActivateServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Branding()(*ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb.BrandingRequestBuilder) {
    return ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CertificateBasedAuthConfiguration()(*ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b.CertificateBasedAuthConfigurationRequestBuilder) {
    return ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CheckMemberGroups()(*if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283.CheckMemberGroupsRequestBuilder) {
    return if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CheckMemberObjects()(*i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35.CheckMemberObjectsRequestBuilder) {
    return i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    m := &OrganizationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/organization/{organization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOrganizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OrganizationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrganizationRequestBuilder) CreateGetRequestInformation(q func (value *OrganizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OrganizationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrganizationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Organization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrganizationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OrganizationRequestBuilder) Extensions()(*i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984.ExtensionsRequestBuilder) {
    return i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) ExtensionsById(id string)(*i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Get(q func (value *OrganizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Organization, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOrganization() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Organization), nil
}
func (m *OrganizationRequestBuilder) GetMemberGroups()(*i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563.GetMemberGroupsRequestBuilder) {
    return i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) GetMemberObjects()(*i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67.GetMemberObjectsRequestBuilder) {
    return i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Organization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OrganizationRequestBuilder) Restore()(*i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d.RestoreRequestBuilder) {
    return i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) SetMobileDeviceManagementAuthority()(*ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56.SetMobileDeviceManagementAuthorityRequestBuilder) {
    return ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56.NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Settings()(*i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b.SettingsRequestBuilder) {
    return i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
