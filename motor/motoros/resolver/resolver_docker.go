package resolver

import (
	"errors"

	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/rs/zerolog/log"
	motorcloud_docker "go.mondoo.io/mondoo/motor/discovery/docker"
	"go.mondoo.io/mondoo/motor/motoros/docker/docker_engine"
	"go.mondoo.io/mondoo/motor/motoros/docker/image"
	"go.mondoo.io/mondoo/motor/motoros/docker/snapshot"
	"go.mondoo.io/mondoo/motor/motoros/types"
)

type DockerInfo struct {
	Name       string
	Identifier string
	Labels     map[string]string
}

// When we talk about Docker, users think at leasst of 3 different things:
// - container runtime (e.g. docker engine)
// - container image (eg. from docker engine or registry)
// - container snapshot
//
// Docker made a very good job in abstracting the problem away from the user
// so that he normally does not think about the distinction. But we need to
// think about those aspects since all those need a different implementation and
// handling.
//
// The user wants and needs an easy way to point to those endpoints:
//
// # registries
// -t docker://gcr.io/project/image@sha256:label
// -t docker://index.docker.io/project/image:label
//
// # docker daemon
// -t docker://id -> image
// -t docker://id -> container
//
// # local directory
// -t docker:///path/link_to_image_archive.tar -> Docker Image
// -t docker:///path/link_to_image_archive2.tar -> OCI
// -t docker:///path/link_to_container.tar
//
// Therefore, this package will only implement the auto-discovery and
// redirect to specific implementations once the disovery is completed
func ResolveDockerTransport(endpoint *types.Endpoint) (types.Transport, DockerInfo, error) {
	// 0. check if we have a tar as input
	//    detect if the tar is a container image format -> container image
	//    or a container snapshot format -> container snapshot
	// 1. check if we have a container id
	//    check if the container is running -> docker engine
	//    check if the container is stopped -> container snapshot
	// 3. check if we have an image id -> container image
	// 4. check if we have a descriptor for a registry -> container image

	if endpoint == nil || len(endpoint.Host) == 0 {
		return nil, DockerInfo{}, errors.New("no endpoint provided")
	}

	// TODO: check if we are pointing to a local tar file
	log.Debug().Str("docker", endpoint.Host).Msg("try to resolve the container or image source")
	_, err := os.Stat(endpoint.Host)
	if err == nil {
		log.Debug().Msg("found local docker/image file")

		// try to load docker image tarball
		img, err := tarball.ImageFromPath(endpoint.Host, nil)
		if err == nil {
			log.Debug().Msg("detected docker image")
			var identifier string

			hash, err := img.Digest()
			if err == nil {
				identifier = motorcloud_docker.MondooContainerImageID(hash.String())
			} else {
				log.Warn().Err(err).Msg("could not determine referenceid")
			}

			rc := mutate.Extract(img)
			transport, err := image.New(rc)
			return transport, DockerInfo{
				Identifier: identifier,
			}, err
		} else {
			log.Debug().Msg("detected docker container snapshot")
			transport, err := snapshot.NewFromFile(endpoint.Host)
			return transport, DockerInfo{}, err
		}

		// TODO: detect file format
		return nil, DockerInfo{}, errors.New("could not find the container reference")
	}

	log.Debug().Msg("try to connect to docker engine")
	// could be an image id/name, container id/name or a short reference to an image in docker engine
	ded, err := NewDockerEngineDiscovery()
	if err == nil {
		ci, err := ded.ContainerInfo(endpoint.Host)
		if err == nil {
			if ci.Running {
				log.Debug().Msg("found running container " + ci.ID)
				transport, err := docker_engine.New(ci.ID)
				return transport, DockerInfo{
					Name:       motorcloud_docker.ShortContainerImageID(ci.ID),
					Identifier: motorcloud_docker.MondooContainerID(ci.ID),
					Labels:     ci.Labels,
				}, err
			} else {
				log.Debug().Msg("found stopped container " + ci.ID)
				transport, err := snapshot.NewFromDockerEngine(ci.ID)
				return transport, DockerInfo{
					Name:       motorcloud_docker.ShortContainerImageID(ci.ID),
					Identifier: motorcloud_docker.MondooContainerID(ci.ID),
					Labels:     ci.Labels,
				}, err
			}
		}

		ii, err := ded.ImageInfo(endpoint.Host)
		if err == nil {
			log.Debug().Msg("found docker engine image " + ii.ID)
			img, rc, err := image.LoadFromDockerEngine(ii.ID)
			if err != nil {
				return nil, DockerInfo{}, err
			}

			var identifier string
			hash, err := img.Digest()
			if err == nil {
				identifier = motorcloud_docker.MondooContainerImageID(hash.String())
			}

			transport, err := image.New(rc)
			return transport, DockerInfo{
				Name:       ii.Name,
				Identifier: identifier,
				Labels:     ii.Labels,
			}, err
		}
	}

	log.Debug().Msg("try to download the image from docker registry")
	// load container image from remote directoryload tar file into backend

	ref, err := name.ParseReference(endpoint.Host, name.WeakValidation)
	if err == nil {
		log.Debug().Str("ref", ref.Name()).Msg("found valid container registry reference")

		registryOpts := []image.Option{image.WithInsecure(endpoint.Insecure)}
		if len(endpoint.BearerToken) > 0 {
			log.Debug().Msg("enable bearer authentication for image")
			registryOpts = append(registryOpts, image.WithAuthenticator(&authn.Bearer{Token: endpoint.BearerToken}))
		}

		// image.WithAuthenticator()
		img, rc, err := image.LoadFromRegistry(ref, registryOpts...)
		if err != nil {
			return nil, DockerInfo{}, err
		}

		var identifier string
		hash, err := img.Digest()
		if err == nil {
			identifier = motorcloud_docker.MondooContainerImageID(hash.String())
		}

		transport, err := image.New(rc)
		return transport, DockerInfo{
			Name:       motorcloud_docker.ShortContainerImageID(hash.String()),
			Identifier: identifier,
		}, err
	} else {
		log.Debug().Str("image", endpoint.Host).Msg("Could not detect a valid repository url")
		return nil, DockerInfo{}, err
	}

	// if we reached here, we assume we have a name of an image or container from a registry
	return nil, DockerInfo{}, errors.New("could not find the container reference")
}
