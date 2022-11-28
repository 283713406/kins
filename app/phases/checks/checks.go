package checks

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/283713406/kins/app/constants"
	"github.com/pkg/errors"
	"k8s.io/klog"
)

type Error struct {
	Msg string
}

// Error implements the standard error interface
func (e *Error) Error() string {
	return fmt.Sprintf("[check] Some fatal errors occurred:\n%s", e.Msg)
}

type Checker interface {
	Check() (warnings, errorList []error)
	Name() string
}

// gateway verifies the node gateway.
type GatewayCheck struct {
	gatewayName string
}

// Name returns label for GatewayCheck.
func (GatewayCheck) Name() string {
	return "default"
}

func (gw GatewayCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("checking node default gateway")
	errorList = []error{}
	warnings = []error{}
	cmd := exec.Command("ip route show", gw.gatewayName)
	err := cmd.Run()
	if err != nil {
		errorList = append(errorList, errors.Wrapf(err, "Failed to found route: \"%s\"", gw.gatewayName))
	}
	return warnings, errorList
}

// gateway verifies the node gateway.
type DnsCheck struct {
}

// Name returns label for GatewayCheck.
func (DnsCheck) Name() string {
	return "114.114.114.114"
}

func (dns DnsCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("checking node default gateway")
	errorList = []error{}
	warnings = []error{}
	dnsOutputBytes, err := exec.Command("cat", constants.DnsFileDir).Output()
	if err != nil {
		errorList = append(errorList, errors.Wrapf(err, "Failed find dns file: \"%s\"", constants.DnsFileDir))
	}
	if ok := strings.Contains(string(dnsOutputBytes), constants.Dns144); ok{
		klog.V(1).Infoln("checking node default dns OK")
	}else {
		error := fmt.Sprintf("find dns Ok, But does not include: \"%s\"", constants.Dns144)
		errorList = append(errorList, errors.New(error))
	}
	return warnings, errorList
}

func RunPreChecks() error {

	checks := []Checker{
		GatewayCheck{gatewayName: constants.DefaultGateway},
		DnsCheck{},
	}

	return RunChecks(checks, os.Stderr)
}

func RunChecks(checks []Checker, ww io.Writer) error {
	var errsBuffer bytes.Buffer

	for _, c := range checks {
		name := c.Name()
		warnings, errs := c.Check()

		for _, w := range warnings {
			io.WriteString(ww, fmt.Sprintf("\t[WARNING %s]: %v\n", name, w))
		}
		for _, i := range errs {
			errsBuffer.WriteString(fmt.Sprintf("\t[ERROR %s]: %v\n", name, i.Error()))
		}
	}

	if errsBuffer.Len() > 0 {
		return &Error{Msg: errsBuffer.String()}
	}
	return nil
}