package managedtenants

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates"
    i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends"
    i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages"
    i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices"
    i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation"
    i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates"
    i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries"
    i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses"
    i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates"
    i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation"
    i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions"
    ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries"
    iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags"
    ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview"
    icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections"
    icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups"
    icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents"
    idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances"
    idd936945fe349ae55442678f4cc9336c26385c7d63fad7702cdd3ed7f48759b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/riskyusers"
    ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances"
    i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates/item"
    i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents/item"
    i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation/item"
    i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices/item"
    i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates/item"
    i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries/item"
    i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends/item"
    i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags/item"
    i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances/item"
    i69ca8f00785bb577702d827861686c081c74eb4d6cdd80b73f488dab59d3a13b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/riskyusers/item"
    i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups/item"
    i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants/item"
    ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/item"
    ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates/item"
    ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries/item"
    ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages/item"
    ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation/item"
    idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances/item"
    ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions/item"
    ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections/item"
    if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview/item"
)

type ManagedTenantsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedTenantsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.AggregatedPolicyCompliancesRequestBuilder) {
    return idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.NewAggregatedPolicyCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliancesById(id string)(*i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.AggregatedPolicyComplianceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["aggregatedPolicyCompliance_id"] = id
    }
    return i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.NewAggregatedPolicyComplianceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcConnections()(*icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.CloudPcConnectionsRequestBuilder) {
    return icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.NewCloudPcConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcConnectionsById(id string)(*ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.CloudPcConnectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcConnection_id"] = id
    }
    return ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.NewCloudPcConnectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcDevices()(*i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.CloudPcDevicesRequestBuilder) {
    return i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.NewCloudPcDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcDevicesById(id string)(*i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.CloudPcDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDevice_id"] = id
    }
    return i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.NewCloudPcDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcsOverview()(*ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.CloudPcsOverviewRequestBuilder) {
    return ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.NewCloudPcsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcsOverviewById(id string)(*if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.CloudPcOverviewRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOverview_tenantId"] = id
    }
    return if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.NewCloudPcOverviewRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.ConditionalAccessPolicyCoveragesRequestBuilder) {
    return i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.NewConditionalAccessPolicyCoveragesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoveragesById(id string)(*ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.ConditionalAccessPolicyCoverageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicyCoverage_id"] = id
    }
    return ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.NewConditionalAccessPolicyCoverageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    m := &ManagedTenantsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/tenantRelationships/managedTenants{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedTenantsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedTenantsRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedTenantsRequestBuilder) CreateGetRequestInformation(q func (value *ManagedTenantsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedTenantsRequestBuilderGetQueryParameters)
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
func (m *ManagedTenantsRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.CredentialUserRegistrationsSummariesRequestBuilder) {
    return ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.NewCredentialUserRegistrationsSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummariesById(id string)(*ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.CredentialUserRegistrationsSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationsSummary_id"] = id
    }
    return ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.NewCredentialUserRegistrationsSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.DeviceCompliancePolicySettingStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary_id"] = id
    }
    return i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) Get(q func (value *ManagedTenantsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedTenant() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant), nil
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.ManagedDeviceCompliancesRequestBuilder) {
    return ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.NewManagedDeviceCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliancesById(id string)(*idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.ManagedDeviceComplianceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceCompliance_id"] = id
    }
    return idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.NewManagedDeviceComplianceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.ManagedDeviceComplianceTrendsRequestBuilder) {
    return i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.NewManagedDeviceComplianceTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrendsById(id string)(*i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.ManagedDeviceComplianceTrendRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceComplianceTrend_id"] = id
    }
    return i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.NewManagedDeviceComplianceTrendRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActions()(*i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.ManagementActionsRequestBuilder) {
    return i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.NewManagementActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActionsById(id string)(*ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.ManagementActionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementAction_id"] = id
    }
    return ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.NewManagementActionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.ManagementActionTenantDeploymentStatusesRequestBuilder) {
    return i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatusesById(id string)(*ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.ManagementActionTenantDeploymentStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementActionTenantDeploymentStatus_id"] = id
    }
    return ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.NewManagementActionTenantDeploymentStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementIntents()(*icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.ManagementIntentsRequestBuilder) {
    return icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.NewManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementIntentsById(id string)(*i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.ManagementIntentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementIntent_id"] = id
    }
    return i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.NewManagementIntentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplates()(*i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.ManagementTemplatesRequestBuilder) {
    return i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.NewManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplatesById(id string)(*i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.ManagementTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplate_id"] = id
    }
    return i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.NewManagementTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ManagedTenantsRequestBuilder) RiskyUsers()(*idd936945fe349ae55442678f4cc9336c26385c7d63fad7702cdd3ed7f48759b2.RiskyUsersRequestBuilder) {
    return idd936945fe349ae55442678f4cc9336c26385c7d63fad7702cdd3ed7f48759b2.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) RiskyUsersById(id string)(*i69ca8f00785bb577702d827861686c081c74eb4d6cdd80b73f488dab59d3a13b.RiskyUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUser_id"] = id
    }
    return i69ca8f00785bb577702d827861686c081c74eb4d6cdd80b73f488dab59d3a13b.NewRiskyUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantGroups()(*icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.TenantGroupsRequestBuilder) {
    return icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.NewTenantGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantGroupsById(id string)(*i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.TenantGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantGroup_id"] = id
    }
    return i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.NewTenantGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) Tenants()(*i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.TenantsRequestBuilder) {
    return i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.NewTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsById(id string)(*i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.TenantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenant_id"] = id
    }
    return i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.NewTenantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.TenantsCustomizedInformationRequestBuilder) {
    return i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.NewTenantsCustomizedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformationById(id string)(*i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.TenantCustomizedInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantCustomizedInformation_id"] = id
    }
    return i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.NewTenantCustomizedInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformation()(*i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.TenantsDetailedInformationRequestBuilder) {
    return i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.NewTenantsDetailedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformationById(id string)(*ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.TenantDetailedInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantDetailedInformation_id"] = id
    }
    return ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.NewTenantDetailedInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantTags()(*iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.TenantTagsRequestBuilder) {
    return iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.NewTenantTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantTagsById(id string)(*i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.TenantTagRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantTag_id"] = id
    }
    return i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.NewTenantTagRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.WindowsDeviceMalwareStatesRequestBuilder) {
    return i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.NewWindowsDeviceMalwareStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStatesById(id string)(*i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.WindowsDeviceMalwareStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState_id"] = id
    }
    return i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.NewWindowsDeviceMalwareStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStates()(*i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.WindowsProtectionStatesRequestBuilder) {
    return i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.NewWindowsProtectionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStatesById(id string)(*ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.WindowsProtectionStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsProtectionState_id"] = id
    }
    return ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.NewWindowsProtectionStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}