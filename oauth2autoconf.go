package oauth2autoconf

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

/*
 * Copyright (c) 2020 Norwegian University of Science and Technology
 */

type Oauth2Conf struct {
	ResponseTypesSupported                     []string `json:"response_types_supported,omitempty"`
	RequestParameterSupported                  bool     `json:"request_parameter_supported,omitempty"`
	RequestUriParameterSupported               bool     `json:"request_uri_parameter_supported,omitempty"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported,omitempty"`
	UiLocalesSupported                         []string `json:"ui_locales_supported,omitempty"`
	EndSessionEndpoint                         string   `json:"end_session_endpoint,omitempty"`
	Issuer                                     string   `json:"issuer,omitempty"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint,omitempty"`
	UserinfoEndpoint                           string   `json:"userinfo_endpoint,omitempty"`
	ServiceDocumentation                       string   `json:"service_documentation,omitempty"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	ClaimsSupported                            []string `json:"claims_supported,omitempty"`
	CodeChallengeMethodsSupported              []string `json:"code_challenge_methods_supported,omitempty"`
	JwksUri                                    string   `json:"jwks_uri,omitempty"`
	SubjectTypesSupported                      []string `json:"subject_types_supported,omitempty"`
	IDTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported,omitempty"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	ResponseModesSupported                     []string `json:"response_modes_supported,omitempty"`
	TokenEndpoint                              string   `json:"token_endpoint,omitempty"`
}

func Get(ctx context.Context, endp string) (*Oauth2Conf, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endp, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	cfg := &Oauth2Conf{}
	if err = json.NewDecoder(resp.Body).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (oc *Oauth2Conf) String() string {
	build := &strings.Builder{}
	if err := json.NewEncoder(build).Encode(oc); err != nil {
		return ""
	}
	return build.String()
}
