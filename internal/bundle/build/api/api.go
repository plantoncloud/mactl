package api

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/app/build/api/graphql/client/pg"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/client/bloomrpc"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/client/grpcurl"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/client/wombat"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/compiler/buf"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/compiler/protobuf"
	"github.com/plantoncloud/mactl/internal/app/build/api/rest/client/postman"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Infof("ensuring rest-api client - postman")
	if err := postman.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure rest-api client - postman")
	}
	log.Infof("ensured rest-api client - postman")
	log.Infof("ensuring graphql api client - graphql-playground")
	if err := pg.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure graphql api client - graphql-playground")
	}
	log.Infof("ensured graphql api client - graphql-playground")
	log.Infof("ensuring grpc api compilers")
	log.Infof("ensuring grpc api compiler - prtobuf")
	if err := protobuf.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api compiler - prtobuf")
	}
	log.Infof("ensured grpc api compiler - prtobuf")
	log.Infof("ensuring grpc api compiler - buf")
	if err := buf.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api compiler - buf")
	}
	log.Infof("ensured grpc api compiler - buf")
	log.Infof("ensured grpc api compilers")
	log.Infof("ensuring grpc api clients")
	log.Infof("ensuring grpc api client - bloomrpc")
	if err := bloomrpc.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api client - bloomrpc")
	}
	log.Infof("ensured grpc api client - bloomrpc")
	log.Infof("ensuring grpc api client - wombat")
	if err := wombat.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api client - wombat")
	}
	log.Infof("ensured grpc api client - wombat")
	log.Infof("ensuring grpc api client - grpcurl")
	if err := grpcurl.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api client - grpcurl")
	}
	log.Infof("ensured grpc api client - grpcurl")
	log.Infof("ensured grpc api clients")
	return nil
}

func Upgrade() error {
	log.Infof("upgrading rest-api client - postman")
	if err := postman.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade rest-api client - postman")
	}
	log.Infof("upgraded rest-api client - postman")

	log.Infof("upgrading graphql api client - graphql-playground")
	if err := pg.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade graphql api client - graphql-playground")
	}
	log.Infof("upgraded graphql api client - graphql-playground")

	log.Infof("upgrading grpc api compilers")
	log.Infof("upgrading grpc api compiler - protoc")
	if err := protobuf.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api compiler - protoc")
	}
	log.Infof("upgraded grpc api compiler - protoc")

	log.Infof("upgrading grpc api compiler - buf")
	if err := buf.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api compiler - buf")
	}
	log.Infof("upgraded grpc api compiler - buf")
	log.Infof("upgraded grpc api compilers")

	log.Infof("upgrading grpc api clients")
	log.Infof("upgrading grpc api client - bloomrpc")
	if err := bloomrpc.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api client - bloomrpc")
	}
	log.Infof("upgraded grpc api client - bloomrpc")

	log.Infof("upgrading grpc api client - wombat")
	if err := wombat.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api client - wombat")
	}
	log.Infof("upgraded grpc api client - wombat")

	log.Infof("upgrading grpc api client - grpcurl")
	if err := grpcurl.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api client - grpcurl")
	}
	log.Infof("upgraded grpc api client - grpcurl")
	log.Infof("upgraded grpc api clients")
	return nil
}
