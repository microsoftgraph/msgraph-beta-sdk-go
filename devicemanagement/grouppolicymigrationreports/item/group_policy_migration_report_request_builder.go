package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings"
    ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions"
    i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings/item"
    i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions/item"
)

type GroupPolicyMigrationReportRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GroupPolicyMigrationReportRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewGroupPolicyMigrationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportRequestBuilder) {
    m := &GroupPolicyMigrationReportRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGroupPolicyMigrationReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupPolicyMigrationReportRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyMigrationReportRequestBuilder) CreateGetRequestInformation(q func (value *GroupPolicyMigrationReportRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GroupPolicyMigrationReportRequestBuilderGetQueryParameters)
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
func (m *GroupPolicyMigrationReportRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyMigrationReportRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupPolicyMigrationReportRequestBuilder) Get(q func (value *GroupPolicyMigrationReportRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyMigrationReport() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport), nil
}
func (m *GroupPolicyMigrationReportRequestBuilder) GroupPolicySettingMappings()(*i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.GroupPolicySettingMappingsRequestBuilder) {
    return i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.NewGroupPolicySettingMappingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyMigrationReportRequestBuilder) GroupPolicySettingMappingsById(id string)(*i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.GroupPolicySettingMappingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicySettingMapping_id"] = id
    }
    return i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.NewGroupPolicySettingMappingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupPolicyMigrationReportRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupPolicyMigrationReportRequestBuilder) UnsupportedGroupPolicyExtensions()(*ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.UnsupportedGroupPolicyExtensionsRequestBuilder) {
    return ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.NewUnsupportedGroupPolicyExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyMigrationReportRequestBuilder) UnsupportedGroupPolicyExtensionsById(id string)(*i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.UnsupportedGroupPolicyExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unsupportedGroupPolicyExtension_id"] = id
    }
    return i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.NewUnsupportedGroupPolicyExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
