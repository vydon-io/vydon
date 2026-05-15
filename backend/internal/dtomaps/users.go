package dtomaps

import (
	"github.com/jackc/pgx/v5/pgtype"
	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAccountTypeDto(aType vydondb.AccountType) mgmtv1alpha1.UserAccountType {
	switch aType {
	case 0:
		return mgmtv1alpha1.UserAccountType_USER_ACCOUNT_TYPE_PERSONAL
	case 1:
		return mgmtv1alpha1.UserAccountType_USER_ACCOUNT_TYPE_TEAM
	case 2:
		return mgmtv1alpha1.UserAccountType_USER_ACCOUNT_TYPE_ENTERPRISE
	default:
		return mgmtv1alpha1.UserAccountType_USER_ACCOUNT_TYPE_UNSPECIFIED
	}
}

func ToAccountInviteDto(input *db_queries.NeosyncApiAccountInvite) *mgmtv1alpha1.AccountInvite {
	return &mgmtv1alpha1.AccountInvite{
		Id:           vydondb.UUIDString(input.ID),
		AccountId:    vydondb.UUIDString(input.AccountID),
		SenderUserId: vydondb.UUIDString(input.SenderUserID),
		Email:        input.Email,
		Token:        input.Token,
		Accepted:     input.Accepted.Bool,
		CreatedAt:    timestamppb.New(input.CreatedAt.Time),
		UpdatedAt:    timestamppb.New(input.UpdatedAt.Time),
		ExpiresAt:    timestamppb.New(input.ExpiresAt.Time),
		Role:         toRoleDto(input.Role),
	}
}

func toRoleDto(role pgtype.Int4) mgmtv1alpha1.AccountRole {
	if !role.Valid {
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_UNSPECIFIED
	}
	return mgmtv1alpha1.AccountRole(role.Int32)
}
func ToUserAccount(input *db_queries.NeosyncApiAccount) *mgmtv1alpha1.UserAccount {
	return &mgmtv1alpha1.UserAccount{
		Id:                  vydondb.UUIDString(input.ID),
		Name:                input.AccountSlug,
		Type:                ToAccountTypeDto(vydondb.AccountType(input.AccountType)),
		HasStripeCustomerId: hasStripeCustomerId(input.StripeCustomerID),
	}
}

func hasStripeCustomerId(customerId pgtype.Text) bool {
	return customerId.Valid && customerId.String != ""
}
