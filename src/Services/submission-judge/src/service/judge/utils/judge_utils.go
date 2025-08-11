package judgeutils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/utils"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/judge"
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
)

func ReturnIsolateIfFail(pService *poolservice.PoolService, i *domain.Isolate, err error) {
	if err != nil {
		i.Logger.Warn().Msgf("Return the isolate because something went wrong: %v", err)
		(*pService).Put(i)
	}
}

func ParseMetaFile(data []byte) (*judge.RunVerdict, error) {
	verdict := &judge.RunVerdict{}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		log := config.GetLogger()
		log.Debug().Msgf("Key: %s - Val: %s", key, val)

		switch key {
		case "status":
			verdict.Status = val
		case "exitcode":
			if code, err := strconv.Atoi(val); err == nil {
				verdict.ExitCode = code
			}
		case "message":
			verdict.Message = val
		case "time":
			if t, err := strconv.ParseFloat(val, 64); err == nil {
				verdict.Time = t
			}
		case "time-wall":
			if t, err := strconv.ParseFloat(val, 64); err == nil {
				verdict.TimeWall = t
			}
		case "cg-mem":
			if mem, err := strconv.Atoi(val); err == nil {
				verdict.CgMem = memory.KiB * memory.Memory(mem)
			}
		case "cg-mem-sw":
			if mem, err := strconv.Atoi(val); err == nil {
				verdict.CgMemSw = memory.KiB * memory.Memory(mem)
			}
		case "max-rss":
			if rss, err := strconv.Atoi(val); err == nil {
				verdict.MaxRss = memory.KiB * memory.Memory(rss)
			}
		case "csw":
			if csw, err := strconv.Atoi(val); err == nil {
				verdict.Csw = csw
			}
		case "csw-forced":
			if csw, err := strconv.Atoi(val); err == nil {
				verdict.CswForced = csw
			}
		case "cg-oom-killed":
			if killed, err := strconv.Atoi(val); err == nil {
				verdict.CgOomKilled = killed
			}
		case "exited-normally":
			verdict.ExitedNormally = val == "1" || strings.ToLower(val) == "true"
		case "killed":
			if sig, err := strconv.Atoi(val); err == nil {
				verdict.KilledBySignal = sig
			}
		}
	}
	return verdict, scanner.Err()
}

func CheckRunStatus(i *domain.Isolate, submissionId string) (*judge.RunVerdict, error) {
	metaAddr, err := utils.GetMetaFilePath(i, submissionId)
	if err != nil {
		return nil, err
	}

	meta, err := os.ReadFile(metaAddr)
	if err != nil {
		return nil, fmt.Errorf("Can't open meta file: %v", err)
	}

	return ParseMetaFile(meta)
}
