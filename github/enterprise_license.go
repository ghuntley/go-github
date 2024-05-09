// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// EnterpriseConsumedLicense represents license consumption information for all users, including those from connected servers, associated with an enterprise.
type EnterpriseConsumedLicense struct {
	TotalSeatsConsumed  *int                             `json:"total_seats_consumed"`
	TotalSeatsPurchased *int                             `json:"total_seats_purchased"`
	Users               *[]EnterpriseConsumedLicenseUser `json:"users"`
}

// EnterpriseConsumedLicenseUser represents license consumption information users, including those from connected servers, associated with an enterprise.

type EnterpriseConsumedLicenseUser struct {
	EnterpriseServerEmails          *[]string `json:"enterprise_server_emails"`
	EnterpriseServerUser            *bool     `json:"enterprise_server_user"`
	EnterpriseServerUserIds         *[]string `json:"enterprise_server_user_ids"`
	GithubComEnterpriseRole         *string   `json:"github_com_enterprise_role"`
	GithubComEnterpriseRoles        *[]string `json:"github_com_enterprise_roles"`
	GithubComLogin                  *string   `json:"github_com_login"`
	GithubComMemberRoles            *[]string `json:"github_com_member_roles"`
	GithubComName                   *string   `json:"github_com_name"`
	GithubComOrgsWithPendingInvites *[]string `json:"github_com_orgs_with_pending_invites"`
	GithubComProfile                *string   `json:"github_com_profile"`
	GithubComSamlNameID             *string   `json:"github_com_saml_name_id"`
	GithubComTwoFactorAuth          *bool     `json:"github_com_two_factor_auth"`
	GithubComUser                   *bool     `json:"github_com_user"`
	GithubComVerifiedDomainEmails   *[]string `json:"github_com_verified_domain_emails"`
	LicenseType                     *string   `json:"license_type"`
	TotalUserAccounts               *int      `json:"total_user_accounts"`
	VisualStudioLicenseStatus       *string   `json:"visual_studio_license_status"`
	VisualStudioSubscriptionEmail   *string   `json:"visual_studio_subscription_email"`
	VisualStudioSubscriptionUser    *bool     `json:"visual_studio_subscription_user"`
}

// ListConsumedLicenses lists the license consumption information for all users, including those from connected servers, associated with an enterprise.
//
// GitHub API docs: https://docs.github.com/en/enterprise-cloud@latest/rest/enterprise-admin/license?apiVersion=2022-11-28#list-enterprise-consumed-licenses
//
//meta:operation GET /enterprises/{enterprise}/consumed-licenses
func (s *EnterpriseService) ListConsumedLicenses(ctx context.Context, enterprise string, opts *ListOptions) (*EnterpriseConsumedLicense, *Response, error) {
	u := fmt.Sprintf("enterprises/%v/consumed-licenses", enterprise)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var consumedLicenses *EnterpriseConsumedLicense
	resp, err := s.client.Do(ctx, req, &consumedLicenses)
	if err != nil {
		return nil, resp, err
	}

	return consumedLicenses, resp, nil

}
