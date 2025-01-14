// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db_queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"
)

type Querier interface {
	CreateAccountInvite(ctx context.Context, db DBTX, arg CreateAccountInviteParams) (NeosyncApiAccountInvite, error)
	CreateAccountUserAssociation(ctx context.Context, db DBTX, arg CreateAccountUserAssociationParams) (NeosyncApiAccountUserAssociation, error)
	CreateAuth0IdentityProviderAssociation(ctx context.Context, db DBTX, arg CreateAuth0IdentityProviderAssociationParams) (NeosyncApiUserIdentityProviderAssociation, error)
	CreateConnection(ctx context.Context, db DBTX, arg CreateConnectionParams) (NeosyncApiConnection, error)
	CreateCustomTransformer(ctx context.Context, db DBTX, arg CreateCustomTransformerParams) (NeosyncApiTransformer, error)
	CreateJob(ctx context.Context, db DBTX, arg CreateJobParams) (NeosyncApiJob, error)
	CreateJobConnectionDestination(ctx context.Context, db DBTX, arg CreateJobConnectionDestinationParams) (NeosyncApiJobDestinationConnectionAssociation, error)
	CreateJobConnectionDestinations(ctx context.Context, db DBTX, arg []CreateJobConnectionDestinationsParams) (int64, error)
	CreatePersonalAccount(ctx context.Context, db DBTX, accountSlug string) (NeosyncApiAccount, error)
	CreateTeamAccount(ctx context.Context, db DBTX, accountSlug string) (NeosyncApiAccount, error)
	CreateUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	DeleteCustomTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) error
	DeleteJob(ctx context.Context, db DBTX, id pgtype.UUID) error
	GetAccount(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccount, error)
	GetAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountInvite, error)
	GetAccountUserAssociation(ctx context.Context, db DBTX, arg GetAccountUserAssociationParams) (NeosyncApiAccountUserAssociation, error)
	GetAccountsByUser(ctx context.Context, db DBTX, id pgtype.UUID) ([]NeosyncApiAccount, error)
	GetActiveAccountInvites(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiAccountInvite, error)
	GetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	GetConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiConnection, error)
	GetConnectionByNameAndAccount(ctx context.Context, db DBTX, arg GetConnectionByNameAndAccountParams) (NeosyncApiConnection, error)
	GetConnectionsByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiConnection, error)
	GetCustomTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiTransformer, error)
	GetCustomTransformersByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiTransformer, error)
	GetJobById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiJob, error)
	GetJobByNameAndAccount(ctx context.Context, db DBTX, arg GetJobByNameAndAccountParams) (NeosyncApiJob, error)
	GetJobConnectionDestination(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobConnectionDestinations(ctx context.Context, db DBTX, id pgtype.UUID) ([]NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobConnectionDestinationsByJobIds(ctx context.Context, db DBTX, jobids []pgtype.UUID) ([]NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobsByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiJob, error)
	GetPersonalAccountByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) (NeosyncApiAccount, error)
	GetTeamAccountsByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) ([]NeosyncApiAccount, error)
	GetTemporalConfigByAccount(ctx context.Context, db DBTX, id pgtype.UUID) (*pg_models.TemporalConfig, error)
	GetTemporalConfigByUserAccount(ctx context.Context, db DBTX, arg GetTemporalConfigByUserAccountParams) (*pg_models.TemporalConfig, error)
	GetUser(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiUser, error)
	GetUserAssociationByAuth0Id(ctx context.Context, db DBTX, auth0ProviderID string) (NeosyncApiUserIdentityProviderAssociation, error)
	GetUserByAuth0Id(ctx context.Context, db DBTX, auth0ProviderID string) (NeosyncApiUser, error)
	GetUsersByTeamAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiUser, error)
	IsConnectionInAccount(ctx context.Context, db DBTX, arg IsConnectionInAccountParams) (int64, error)
	IsConnectionNameAvailable(ctx context.Context, db DBTX, arg IsConnectionNameAvailableParams) (int64, error)
	IsJobNameAvailable(ctx context.Context, db DBTX, arg IsJobNameAvailableParams) (int64, error)
	IsTransformerNameAvailable(ctx context.Context, db DBTX, arg IsTransformerNameAvailableParams) (int64, error)
	IsUserInAccount(ctx context.Context, db DBTX, arg IsUserInAccountParams) (int64, error)
	RemoveAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveAccountUser(ctx context.Context, db DBTX, arg RemoveAccountUserParams) error
	RemoveConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveConnectionByNameAndAccount(ctx context.Context, db DBTX, arg RemoveConnectionByNameAndAccountParams) error
	RemoveJobById(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveJobConnectionDestination(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveJobConnectionDestinations(ctx context.Context, db DBTX, jobids []pgtype.UUID) error
	SetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	UpdateActiveAccountInvitesToExpired(ctx context.Context, db DBTX, arg UpdateActiveAccountInvitesToExpiredParams) (NeosyncApiAccountInvite, error)
	UpdateConnection(ctx context.Context, db DBTX, arg UpdateConnectionParams) (NeosyncApiConnection, error)
	UpdateCustomTransformer(ctx context.Context, db DBTX, arg UpdateCustomTransformerParams) (NeosyncApiTransformer, error)
	UpdateJobConnectionDestination(ctx context.Context, db DBTX, arg UpdateJobConnectionDestinationParams) (NeosyncApiJobDestinationConnectionAssociation, error)
	UpdateJobMappings(ctx context.Context, db DBTX, arg UpdateJobMappingsParams) (NeosyncApiJob, error)
	UpdateJobSchedule(ctx context.Context, db DBTX, arg UpdateJobScheduleParams) (NeosyncApiJob, error)
	UpdateJobSource(ctx context.Context, db DBTX, arg UpdateJobSourceParams) (NeosyncApiJob, error)
	UpdateTemporalConfigByAccount(ctx context.Context, db DBTX, arg UpdateTemporalConfigByAccountParams) (NeosyncApiAccount, error)
}

var _ Querier = (*Queries)(nil)
