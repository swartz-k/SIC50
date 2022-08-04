package tensor

import (
	"encoding/json"
	"fmt"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/utils/log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func getDirEntryLen(dir string) int {
	var minLen = 9999
	if entry, err := os.ReadDir(dir); err == nil {
		for _, e := range entry {
			tmp := e
			if tmp.IsDir() {
				if es, err := os.ReadDir(path.Join(dir, tmp.Name())); err == nil {
					tmpEsLen := len(es)
					if tmpEsLen < minLen {
						minLen = tmpEsLen
					}
				}
			}
		}
	}
	return minLen
}

func getEpochStep(dir string, count int) (int, int) {
	if count == 0 {
		count = getDirEntryLen(dir)
	}
	epoch := count
	stepsPreEpoch := 1
	if count%5 == 0 {
		epoch = count / 5
		stepsPreEpoch = 5
	}
	if count%10 == 0 {
		epoch = count / 10
		stepsPreEpoch = 10
	}
	return epoch, stepsPreEpoch
}

func Train(dir string, minDirTotal int) (*float32, error) {
	// ~/miniconda3/envs/myenv/bin/python3.7 server/train.py
	trainScript := path.Join(config.Cfg.WorkDir, "train.py")
	epoch, step := getEpochStep(dir, minDirTotal)
	cmd := []string{
		trainScript, dir, fmt.Sprintf("%d", epoch), fmt.Sprintf("%d", step),
	}
	log.Info("exec cmd %s %+v", config.Cfg.PythonPath, cmd)
	output, err := exec.Command(config.Cfg.PythonPath, cmd...).Output()
	if err != nil {
		log.Info("train dir %s output %s err %s", dir, output, err.Error())
		return nil, err
	}
	results := strings.Split(string(output), "\n")
	var tryR string
	for _, r := range results {
		if strings.Contains(r, "[") && strings.Contains(r, "]") {
			tryR = r
		}
	}
	var float32R []float32
	log.Info("json parse %s", tryR)
	err = json.Unmarshal([]byte(tryR), &float32R)
	if err != nil {
		return nil, err
	}
	return &float32R[0], nil
}
