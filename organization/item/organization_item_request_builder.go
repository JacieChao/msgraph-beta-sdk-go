package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/activateservice"
    i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/getmembergroups"
    i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/settings"
    i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/restore"
    i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/checkmemberobjects"
    i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/extensions"
    i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/getmemberobjects"
    ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/setmobiledevicemanagementauthority"
    ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/certificatebasedauthconfiguration"
    ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/branding"
    if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/checkmembergroups"
    i37b836a8004e125a1b3420227b71278a7f91ba13727f5a526417e3f56babde57 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/certificatebasedauthconfiguration/item"
    i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64 "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item/extensions/item"
)

// OrganizationItemRequestBuilder provides operations to manage the collection of organization entities.
type OrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OrganizationItemRequestBuilderDeleteOptions options for Delete
type OrganizationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OrganizationItemRequestBuilderGetOptions options for Get
type OrganizationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OrganizationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OrganizationItemRequestBuilderGetQueryParameters get entity from organization by key
type OrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OrganizationItemRequestBuilderPatchOptions options for Patch
type OrganizationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ActivateService the activateService property
func (m *OrganizationItemRequestBuilder) ActivateService()(*i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f.ActivateServiceRequestBuilder) {
    return i01c2cf55bc9b20c5719617ab77e26031acc6603e9def1ad03f16311f618a097f.NewActivateServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Branding the branding property
func (m *OrganizationItemRequestBuilder) Branding()(*ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb.BrandingRequestBuilder) {
    return ie57eaa5f9a9f642f69c4a05abd9ccb4b0048bf0f967b3657df89bf183c04e2eb.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfiguration the certificateBasedAuthConfiguration property
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfiguration()(*ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b.CertificateBasedAuthConfigurationRequestBuilder) {
    return ie5657564522195dd5cec134caaa23e14e4064c2afac19838adbe58a7c8dcc69b.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfigurationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item.certificateBasedAuthConfiguration.item collection
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfigurationById(id string)(*i37b836a8004e125a1b3420227b71278a7f91ba13727f5a526417e3f56babde57.CertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration_id"] = id
    }
    return i37b836a8004e125a1b3420227b71278a7f91ba13727f5a526417e3f56babde57.NewCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *OrganizationItemRequestBuilder) CheckMemberGroups()(*if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283.CheckMemberGroupsRequestBuilder) {
    return if94213e38f10b90dde3c8b3b2ec63955951955ba89406c658f342939fdca7283.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrganizationItemRequestBuilder) CheckMemberObjects()(*i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35.CheckMemberObjectsRequestBuilder) {
    return i3086ab3ba30171d3e38d1e8eeabf2cbe86c2e2171b5a3e0fda13250935294c35.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationItemRequestBuilderInternal instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    m := &OrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemRequestBuilder instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from organization
func (m *OrganizationItemRequestBuilder) CreateDeleteRequestInformation(options *OrganizationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation get entity from organization by key
func (m *OrganizationItemRequestBuilder) CreateGetRequestInformation(options *OrganizationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in organization
func (m *OrganizationItemRequestBuilder) CreatePatchRequestInformation(options *OrganizationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from organization
func (m *OrganizationItemRequestBuilder) Delete(options *OrganizationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
// Extensions the extensions property
func (m *OrganizationItemRequestBuilder) Extensions()(*i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984.ExtensionsRequestBuilder) {
    return i50de44433a9e07319c7393b40edead1d7b173f7cf161d13659e880b1758a6984.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item.extensions.item collection
func (m *OrganizationItemRequestBuilder) ExtensionsById(id string)(*i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i5a9a1fe6f96ba4b698fe3cad99f9bf2e8e3994919ed21cb95a03fd584cae3a64.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from organization by key
func (m *OrganizationItemRequestBuilder) Get(options *OrganizationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *OrganizationItemRequestBuilder) GetMemberGroups()(*i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563.GetMemberGroupsRequestBuilder) {
    return i074e0fb14906f4213b401d21f0a458ad552430e15c71f5984106ab22d1aca563.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrganizationItemRequestBuilder) GetMemberObjects()(*i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67.GetMemberObjectsRequestBuilder) {
    return i95cbc854e27501fbe03075c88f756e8970d171baeebf5e9dffe945387c377d67.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in organization
func (m *OrganizationItemRequestBuilder) Patch(options *OrganizationItemRequestBuilderPatchOptions)(error) {
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
// Restore the restore property
func (m *OrganizationItemRequestBuilder) Restore()(*i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d.RestoreRequestBuilder) {
    return i2b874725bd7aa87adea08d05d3bbe22c3da38631563a9a7f79f171d18e88528d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetMobileDeviceManagementAuthority the setMobileDeviceManagementAuthority property
func (m *OrganizationItemRequestBuilder) SetMobileDeviceManagementAuthority()(*ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56.SetMobileDeviceManagementAuthorityRequestBuilder) {
    return ib11138bd8221a5f18c275978ce130521f1d65d12f9aeacb61a288cc17e61aa56.NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *OrganizationItemRequestBuilder) Settings()(*i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b.SettingsRequestBuilder) {
    return i14a549a86b04cb26e4212d1379ec190de95c8bfea1251853ad3119f900012c3b.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
