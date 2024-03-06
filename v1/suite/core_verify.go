package suite

import (
	"context"
	"errors"
	"fmt"
	"github.com/zj-sh/meta/v1/base"
	"github.com/zj-sh/meta/v1/pipe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

func VerifyFile(file, cmd string) (*SuiteInfo, error) {
	if !strings.HasPrefix(file, "/") {
		return nil, errors.New("file must be absolute path")
	}
	if err := VerifyCmd(cmd); err != nil {
		return nil, err
	}
	clt, grpc, err := GRPCSuiteClient("", cmd, file)
	if err != nil {
		return nil, fmt.Errorf("start suite client error: %s", err.Error())
	}
	defer grpc.Kill()
	info, _ := clt.Inject(context.TODO(), &emptypb.Empty{})
	return VerifyInfo(info, clt)
}

func VerifyInfo(st *SuiteInfo, clt SuiteClient) (*SuiteInfo, error) {
	if st == nil {
		return st, errors.New("the suite info cannot be empty")
	}
	if _, err := clt.InjectConfig(context.TODO(), &base.JsonString{}); status.Code(err) == codes.Unimplemented {
		return nil, errors.New("the suite must implement InjectConfig")
	}
	if !strings.HasPrefix(st.GetSuiteCode(), "suite_") {
		return st, errors.New("the suite code must be prefix with suite_")
	}
	if err := VerifyCmd(st.GetSuiteCmd()); err != nil {
		return st, err
	}
	if err := verifyVersion(st.GetSuiteVersion()); err != nil {
		return st, err
	}
	if err := verifyJobEffects(st, clt); err != nil {
		return st, err
	}
	if err := verifyField(st, clt); err != nil {
		return st, err
	}
	return st, nil
}

func VerifyCmd(cmd string) error {
	nosafe := []string{"rm"}

	if cmd == "" {
		return errors.New("the suite cmd cannot be empty")
	}
	if !strings.Contains(cmd, "{file}") {
		return errors.New("the suite cmd must contains {file}")
	}
	for _, v := range nosafe {
		if strings.HasPrefix(cmd, fmt.Sprintf("%s ", v)) ||
			strings.Contains(cmd, fmt.Sprintf(" %s ", v)) ||
			strings.Contains(cmd, fmt.Sprintf(";%s ", v)) {
			return errors.New("the suite cmd not safe")
		}
	}
	return nil
}

func verifyJobEffects(st *SuiteInfo, clt SuiteClient) error {
	jes := st.GetJobEffects()
	if len(jes) == 0 {
		return errors.New("the job types that the suite can support cannot be empty")
	}
	ctx := context.TODO()
	for _, je := range jes {
		switch je {
		case pipe.JobEffect_EFFECT_UNKNOWN:
			return errors.New("the job effects that the suite can support cannot be JobEffect_EFFECT_UNKNOWN")
		case pipe.JobEffect_EFFECT_SOURCE_PULL:
			if _, err := clt.SourceClone(ctx, &SourcePullCloneParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_SOURCE_PULL, it must implement SourceClone")
			}
			if _, err := clt.SourceCreate(ctx, &SourcePullCreateParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_SOURCE_PULL, it must implement SourceCreate")
			}
			if _, err := clt.SourceMerge(ctx, &SourcePullMergeParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_SOURCE_PULL, it must implement SourceMerge")
			}
		case pipe.JobEffect_EFFECT_SOURCE_AUTH:
		case pipe.JobEffect_EFFECT_SOURCE_SCAN:
			if _, err := clt.SourceScan(ctx, &SourceScanParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_SOURCE_SCAN, it must implement SourceScan")
			}
		case pipe.JobEffect_EFFECT_CONFIG:
		case pipe.JobEffect_EFFECT_TEST:
			if _, err := clt.TestRun(ctx, &TestRunParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_TEST, it must implement TestRun")
			}
			if _, err := clt.TestCancel(ctx, &TestCancelParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_TEST, it must implement TestCancel")
			}
		case pipe.JobEffect_EFFECT_BUILD:
			if _, err := clt.Build(ctx, &BuildParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_BUILD, it must implement Build")
			}
		case pipe.JobEffect_EFFECT_IMAGE:
			if _, err := clt.ImageBuild(ctx, &ImageBuildParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_IMAGE, it must implement ImageBuild")
			}
		case pipe.JobEffect_EFFECT_DEPLOY:
		case pipe.JobEffect_EFFECT_DATA:
		case pipe.JobEffect_EFFECT_VERSION:
		case pipe.JobEffect_EFFECT_NOTIFY:
			if _, err := clt.Notify(ctx, &NotifyParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_NOTIFY, it must implement Notify")
			}
		case pipe.JobEffect_EFFECT_APPROVAL:
			if _, err := clt.ApprovalSend(ctx, &ApprovalSendParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_APPROVAL, it must implement ApprovalSend")
			}
			if _, err := clt.ApprovalCancel(ctx, &ApprovalCancelParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_APPROVAL, it must implement ApprovalCancel")
			}
		case pipe.JobEffect_EFFECT_STORAGE:
			if _, err := clt.StorageUploadFile(ctx, &StorageUploadFileParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_STORAGE, it must implement StorageUploadFile")
			}
			if _, err := clt.StorageUploadStream(ctx); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_STORAGE, it must implement StorageUploadStream")
			}
			if _, err := clt.StorageDelete(ctx, &StorageDeleteParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_STORAGE, it must implement StorageDelete")
			}
		case pipe.JobEffect_EFFECT_PRODUCT:
			if _, err := clt.ProductUploadFile(ctx, &ProductUploadFileParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_PRODUCT, it must implement ProductUploadFile")
			}
			if _, err := clt.ProductUploadStream(ctx); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_PRODUCT, it must implement ProductUploadStream")
			}
			if _, err := clt.ProductDelete(ctx, &ProductDeleteParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_PRODUCT, it must implement ProductDelete")
			}
		case pipe.JobEffect_EFFECT_SHELL:
			if _, err := clt.ShellExec(ctx); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_SHELL, it must implement ShellExec")
			}
		case pipe.JobEffect_EFFECT_ISSUES:
			if _, err := clt.Issues(ctx, &IssuesParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_ISSUES, it must implement Issues")
			}
		case pipe.JobEffect_EFFECT_ORG:
			if _, err := clt.OrgUserSync(ctx, &OrgUserParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_ORG, it must implement OrgUserSync")
			}
			if _, err := clt.OrgSync(ctx, &OrgParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_ORG, it must implement OrgSync")
			}
		case pipe.JobEffect_EFFECT_CUSTOM:
			if _, err := clt.Custom(ctx, &CustomParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_CUSTOM, it must implement Custom")
			}
		case pipe.JobEffect_EFFECT_PIPE:
			if _, err := clt.PipeRun(ctx, &PipelineRunParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_PIPE, it must implement PipeRun")
			}
			if _, err := clt.PipeCancel(ctx, &PipelineCancelParam{}); status.Code(err) == codes.Unimplemented {
				return errors.New("if the suite wants to support JobEffect_EFFECT_PIPE, it must implement PipeCancel")
			}
		default:
			return fmt.Errorf("the job effects that the suite can support cannot be %s", je.String())
		}
	}
	return nil
}
func verifyField(st *SuiteInfo, clt SuiteClient) error {
	return nil
}
func verifyVersion(version string) error {
	reg, _ := regexp.Compile("^[v|V]?\\d{1,4}(\\.\\d{1,4}){0,4}$")
	if !reg.MatchString(version) {
		return errors.New("版本号格式错误, 请满足：[v]x(.x)(.x)(.x), x区间0-9999")
	}
	return nil
}
