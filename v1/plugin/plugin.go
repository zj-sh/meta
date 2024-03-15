package plugin

import (
	"fmt"
	"github.com/zohu/reg"
	"os"
	"path/filepath"
	"reflect"
)

const (
	DefaultExecMainFilePattern  = `^main(\.\w+)?$`
	DefaultExecPrescriptPattern = `(^@$)|(^@\x20.*)|(.*\x20@$)|(.*\x20@\x20.*)`
)

func (p Type) PathName() string {
	switch p {
	case Type_KIT:
		return "kit"
	case Type_EXECFILE:
		return "exec"
	case Type_IMAGE:
		return "image"
	case Type_HELM:
		return "helm"
	default:
		return "unknown"
	}
}
func (p *Plugin) Verify(dir string) error {
	if reg.String(p.GetName()).IsTruthAlphanumericUnderline().MaxLen(128).NotB() {
		return fmt.Errorf("name can only contain letters, numbers, underscores, the first can not be a number, maximum length of 128 digits")
	}
	if reg.String(p.GetFullName()).MaxLen(128).NotB() {
		return fmt.Errorf("full name maximum length of 128 digits")
	}
	if reg.Version(p.GetVersion()).IsVersion().NotB() {
		return fmt.Errorf("incorrect version format, only supports x.x.x standard formats")
	}
	if len(p.GetEffects()) == 0 {
		return fmt.Errorf("effects is required")
	}
	if reg.Version(p.GetSupport().GetRuntimeVersion()).IsVersionSupport().NotB() {
		return fmt.Errorf("runtime version must be specified, supports prefixes: (~) patch, (^) minor, (>=) greater than or equal to, (<=) less than or equal to, default *")
	}
	if len(p.GetSupport().GetOs()) == 0 {
		return fmt.Errorf("support os is required")
	}
	if len(p.GetSupport().GetArch()) == 0 {
		return fmt.Errorf("support arch is required")
	}
	switch p.GetType() {
	case Type_KIT:
		if !existed(filepath.Join(dir, "kit.sh")) {
			return fmt.Errorf("%s/kit.sh is required", dir)
		}
	case Type_EXECFILE:
		p.Prescript = firstTruthValue(p.GetPrescript(), "@")
		if reg.String(p.GetPrescript()).Match(DefaultExecPrescriptPattern).NotB() {
			return fmt.Errorf("prescript must have the @ symbol")
		}
		var hasMain bool
		fis, _ := os.ReadDir(dir)
		for _, fi := range fis {
			if !fi.IsDir() && reg.String(fi.Name()).Match(DefaultExecMainFilePattern).B() {
				hasMain = true
				break
			}
		}
		if !hasMain {
			return fmt.Errorf("%s/main(.*)? is required", dir)
		}
	case Type_IMAGE:
		if p.GetImage() != nil {
			if p.GetImage().GetName() == "" {
				return fmt.Errorf("image.name is required")
			}
			p.Image.Tag = firstTruthValue(p.GetImage().GetTag(), "latest")
		} else {
			if !existed(filepath.Join(dir, "Dockerfile")) &&
				!existed(filepath.Join(dir, "dockerfile")) {
				return fmt.Errorf("type=image, must have Plugin.image or ./Dockerfile")
			}
		}
	case Type_HELM:
		if p.GetHelm() == nil {
			return fmt.Errorf("type=helm, Plugin.helm is required")
		}
		if reg.String(p.GetHelm().GetRepo()).IsUrl().B() {
			if p.GetHelm().GetChart() == "" {
				return fmt.Errorf("remote helm repo Plugin.helm.chart is required")
			}
		} else {
			if !existed(filepath.Join(dir, p.GetHelm().GetRepo(), "Chart.yaml")) ||
				!existed(filepath.Join(dir, p.GetHelm().GetRepo(), "values.yaml")) {
				return fmt.Errorf("local chart repo the %s dir must is a helm Chart repo", p.GetHelm().GetRepo())
			}
			p.Helm.Chart = filepath.Base(firstTruthValue(p.GetHelm().GetChart(), p.GetHelm().GetRepo()))
		}
	default:
		return fmt.Errorf("plugin type is required")
	}
	return nil
}
func existed(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
func firstTruthValue[T any](args ...T) T {
	for _, item := range args {
		if !reflect.ValueOf(item).IsZero() {
			return item
		}
	}
	return args[0]
}
