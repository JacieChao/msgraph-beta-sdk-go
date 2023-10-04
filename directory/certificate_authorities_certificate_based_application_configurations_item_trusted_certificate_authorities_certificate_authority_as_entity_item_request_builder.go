package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
type CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters get trustedCertificateAuthorities from directory
type CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters
}
// CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal instantiates a new CertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    m := &CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}/trustedCertificateAuthorities/{certificateAuthorityAsEntity%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder instantiates a new CertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property trustedCertificateAuthorities for directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get trustedCertificateAuthorities from directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityAsEntityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable), nil
}
// Patch update the navigation property trustedCertificateAuthorities in directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityAsEntityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable), nil
}
// ToDeleteRequestInformation delete navigation property trustedCertificateAuthorities for directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get trustedCertificateAuthorities from directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property trustedCertificateAuthorities in directory
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) WithUrl(rawUrl string)(*CertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    return NewCertificateAuthoritiesCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
