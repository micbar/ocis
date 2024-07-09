package command

import (
	"errors"
	"fmt"
	"path/filepath"

	ocisbs "github.com/cs3org/reva/v2/pkg/storage/fs/ocis/blobstore"
	"github.com/cs3org/reva/v2/pkg/storage/fs/posix/lookup"
	s3bs "github.com/cs3org/reva/v2/pkg/storage/fs/s3ng/blobstore"
	"github.com/cs3org/reva/v2/pkg/storagespace"
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/ocis/pkg/revisions"
	"github.com/urfave/cli/v2"
)

var (
	// _nodesGlobPattern is the glob pattern to find all nodes
	_nodesGlobPattern = "spaces/*/*/*/*/*/*/*/*"
)

// RevisionsCommand is the entrypoint for the revisions command.
func RevisionsCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "revisions",
		Usage: "ocis revisions functionality",
		Subcommands: []*cli.Command{
			PurgeRevisionsCommand(cfg),
		},
		Before: func(c *cli.Context) error {
			return configlog.ReturnError(parser.ParseConfig(cfg, true))
		},
		Action: func(_ *cli.Context) error {
			fmt.Println("Read the docs")
			return nil
		},
	}
}

// PurgeRevisionsCommand allows removing all revisions from a storage provider.
func PurgeRevisionsCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "purge",
		Usage: "purge revisions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "basepath",
				Aliases:  []string{"p"},
				Usage:    "the basepath of the decomposedfs (e.g. /var/tmp/ocis/storage/metadata)",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "blobstore",
				Aliases: []string{"b"},
				Usage:   "the blobstore type. Can be (none, ocis, s3ng). Default ocis. Note: When using s3ng this needs same configuration as the storage-users service",
				Value:   "ocis",
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "do not delete anything, just print what would be deleted",
				Value: true,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "print verbose output",
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "resource-id",
				Aliases: []string{"r"},
				Usage:   "purge all revisions of this file/space. If not set, all revisions will be purged",
			},
		},
		Action: func(c *cli.Context) error {
			basePath := c.String("basepath")
			if basePath == "" {
				fmt.Println("basepath is required")
				return cli.ShowCommandHelp(c, "revisions")
			}

			var (
				bs  revisions.DelBlobstore
				err error
			)
			switch c.String("blobstore") {
			case "s3ng":
				bs, err = s3bs.New(
					cfg.StorageUsers.Drivers.S3NG.Endpoint,
					cfg.StorageUsers.Drivers.S3NG.Region,
					cfg.StorageUsers.Drivers.S3NG.Bucket,
					cfg.StorageUsers.Drivers.S3NG.AccessKey,
					cfg.StorageUsers.Drivers.S3NG.SecretKey,
					s3bs.Options{},
				)
			case "ocis":
				bs, err = ocisbs.New(basePath)
			case "none":
				bs = nil
			default:
				err = errors.New("blobstore type not supported")
			}
			if err != nil {
				fmt.Println(err)
				return err
			}

			p, err := generatePath(basePath, c.String("resource-id"))
			if err != nil {
				fmt.Printf("❌ Error parsing resourceID: %s", err)
				return err
			}

			if err := revisions.PurgeRevisions(p, bs, c.Bool("dry-run"), c.Bool("verbose")); err != nil {
				fmt.Printf("❌ Error purging revisions: %s", err)
				return err
			}

			return nil
		},
	}
}

func generatePath(basePath string, resourceID string) (string, error) {
	if resourceID == "" {
		return filepath.Join(basePath, _nodesGlobPattern), nil
	}

	rid, err := storagespace.ParseID(resourceID)
	if err != nil {
		return "", err
	}

	sid := lookup.Pathify(rid.GetSpaceId(), 1, 2)
	if sid == "" {
		sid = "*/*"
	}

	nid := lookup.Pathify(rid.GetOpaqueId(), 4, 2)
	if nid == "" {
		nid = "*/*/*/*/"
	}

	return filepath.Join(basePath, "spaces", sid, "nodes", nid+"*"), nil
}

func init() {
	register.AddCommand(RevisionsCommand)
}
