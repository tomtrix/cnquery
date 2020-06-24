package statutil

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.io/mondoo/motor/motoros/types"
)

type CommandRunner interface {
	RunCommand(command string) (*types.Command, error)
}

var ESCAPEREGEX = regexp.MustCompile(`[^\w@%+=:,./-]`)

func ShellEscape(s string) string {
	if len(s) == 0 {
		return "''"
	}
	if ESCAPEREGEX.MatchString(s) {
		return "'" + strings.Replace(s, "'", "'\"'\"'", -1) + "'"
	}

	return s
}

func New(cmdRunner CommandRunner) *statHelper {
	return &statHelper{
		commandRunner: cmdRunner,
	}
}

// Stat helper implements the stat command for various unix systems
// since this helper is used by transports itself, we cannot rely on the
// platform detection mechanism (since it may rely on stat to determine the system)
// therefore we implement the minimum required to detect the right stat parser
type statHelper struct {
	commandRunner CommandRunner
	detected      bool
	isunix        bool
}

var bsdunix = map[string]bool{
	"openbsd":   true,
	"dragonfly": true,
	"freebsd":   true,
	"netbsd":    true,
}

func (s *statHelper) Stat(name string) (os.FileInfo, error) {
	// detect stat version
	if !s.detected {
		cmd, err := s.commandRunner.RunCommand("uname -s")
		if err != nil {
			log.Debug().Err(err).Str("file", name).Msg("could not stat the file")
		}

		data, err := ioutil.ReadAll(cmd.Stdout)
		if err != nil {
			return nil, err
		}

		// only switch to unix if we properly detected it, otherwise fallback to linux
		val := strings.ToLower(strings.TrimSpace(string(data)))

		isunix, ok := bsdunix[val]
		if ok && isunix {
			s.isunix = true
		}
		s.detected = true
	}

	if s.isunix {
		return s.unix(name)
	}
	return s.linux(name)
}

func (s *statHelper) linux(name string) (os.FileInfo, error) {
	lstat := "-L"
	format := "--printf"
	path := ShellEscape(name)

	var sb strings.Builder

	sb.WriteString("stat ")
	sb.WriteString(lstat)
	sb.WriteString(" ")
	sb.WriteString(path)
	sb.WriteString(" ")
	sb.WriteString(format)
	sb.WriteString(" '%s\n%f\n%u\n%g\n%X\n%Y\n%C'")

	// NOTE: handling the exit code here does not work for all cases
	// sometimes stat returns something like: failed to get security context of '/etc/ssh/sshd_config': No data available
	// Therefore we continue after this command and try to parse the result and focus on making the parsing more robust
	cmd, err := s.commandRunner.RunCommand(sb.String())
	if err != nil {
		log.Debug().Err(err).Str("file", name).Msg("could not stat the file")
	}

	data, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		return nil, err
	}

	statsData := strings.Split(string(data), "\n")
	if len(statsData) != 7 {
		log.Error().Str("name", name).Msg("could not stat the file")
		return nil, os.ErrNotExist
	}

	size, err := strconv.Atoi(statsData[0])
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	uid, err := strconv.ParseInt(statsData[2], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	gid, err := strconv.ParseInt(statsData[3], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	mask, err := strconv.ParseUint(statsData[1], 16, 32)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	mode := os.FileMode(uint32(mask) & 07777)

	mtime, err := strconv.ParseInt(statsData[4], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	return &types.FileInfo{
		FSize:    int64(size),
		FMode:    mode,
		FIsDir:   mode.IsDir(),
		FModTime: time.Unix(mtime, 0),
		Uid:      uid,
		Gid:      gid,
	}, nil
}

func (s *statHelper) unix(name string) (os.FileInfo, error) {
	lstat := "-L"
	format := "-f"
	path := ShellEscape(name)

	var sb strings.Builder
	sb.WriteString("stat ")
	sb.WriteString(lstat)
	sb.WriteString(" ")
	sb.WriteString(format)
	sb.WriteString(" '%z:%p:%u:%g:%a:%m'")
	sb.WriteString(" ")
	sb.WriteString(path)

	cmd, err := s.commandRunner.RunCommand(sb.String())
	if err != nil {
		log.Debug().Err(err).Str("file", name).Msg("could not stat the file")
	}

	data, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		return nil, err
	}

	statsData := strings.Split(string(data), ":")
	if len(statsData) != 6 {
		log.Error().Str("name", name).Msg("could parse stat response")
		return nil, os.ErrNotExist
	}

	size, err := strconv.Atoi(statsData[0])
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	uid, err := strconv.ParseInt(statsData[2], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	gid, err := strconv.ParseInt(statsData[3], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	// NOTE: the base is 8 instead of 16 on linux systems
	mask, err := strconv.ParseUint(statsData[1], 8, 32)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	mode := os.FileMode(uint32(mask) & 07777)

	mtime, err := strconv.ParseInt(statsData[4], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "could not stat "+name)
	}

	return &types.FileInfo{
		FSize:    int64(size),
		FMode:    mode,
		FIsDir:   mode.IsDir(),
		FModTime: time.Unix(mtime, 0),
		Uid:      uid,
		Gid:      gid,
	}, nil
}
