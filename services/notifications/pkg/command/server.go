package command

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/events/stream"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/go-micro/plugins/v4/events/natsjs"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/crypto"
	"github.com/owncloud/ocis/v2/ocis-pkg/service/grpc"
	settingssvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/settings/v0"
	"github.com/owncloud/ocis/v2/services/notifications/pkg/channels"
	"github.com/owncloud/ocis/v2/services/notifications/pkg/config"
	"github.com/owncloud/ocis/v2/services/notifications/pkg/config/parser"
	"github.com/owncloud/ocis/v2/services/notifications/pkg/logging"
	"github.com/owncloud/ocis/v2/services/notifications/pkg/service"
	"github.com/urfave/cli/v2"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", cfg.Service.Name),
		Category: "server",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)

			// evs defines a list of events to subscribe to
			evs := []events.Unmarshaller{
				events.ShareCreated{},
				events.ShareExpired{},
				events.SpaceShared{},
				events.SpaceUnshared{},
				events.SpaceMembershipExpired{},
			}

			evtsCfg := cfg.Notifications.Events

			var tlsConf *tls.Config
			if evtsCfg.EnableTLS {
				var rootCAPool *x509.CertPool
				if evtsCfg.TLSRootCACertificate != "" {
					rootCrtFile, err := os.Open(evtsCfg.TLSRootCACertificate)
					if err != nil {
						return err
					}

					rootCAPool, err = crypto.NewCertPoolFromPEM(rootCrtFile)
					if err != nil {
						return err
					}
					evtsCfg.TLSInsecure = false
				}

				tlsConf = &tls.Config{
					MinVersion:         tls.VersionTLS12,
					InsecureSkipVerify: evtsCfg.TLSInsecure, //nolint:gosec
					RootCAs:            rootCAPool,
				}
			}
			client, err := stream.Nats(
				natsjs.TLSConfig(tlsConf),
				natsjs.Address(evtsCfg.Endpoint),
				natsjs.ClusterID(evtsCfg.Cluster),
			)
			if err != nil {
				return err
			}
			evts, err := events.Consume(client, evtsCfg.ConsumerGroup, evs...)
			if err != nil {
				return err
			}
			channel, err := channels.NewMailChannel(*cfg, logger)
			if err != nil {
				return err
			}
			tm, err := pool.StringToTLSMode(cfg.Notifications.GRPCClientTLS.Mode)
			if err != nil {
				return err
			}
			gwclient, err := pool.GetGatewayServiceClient(
				cfg.Notifications.RevaGateway,
				pool.WithTLSCACert(cfg.Notifications.GRPCClientTLS.CACert),
				pool.WithTLSMode(tm),
			)
			if err != nil {
				logger.Fatal().Err(err).Str("addr", cfg.Notifications.RevaGateway).Msg("could not get reva client")
			}
			valueService := settingssvc.NewValueService("com.owncloud.api.settings", grpc.DefaultClient())
			svc := service.NewEventsNotifier(evts, channel, logger, gwclient, valueService, cfg.Notifications.MachineAuthAPIKey, cfg.Notifications.EmailTemplatePath, cfg.WebUIURL)
			return svc.Run()
		},
	}
}
