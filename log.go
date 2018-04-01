package captainslog

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var announceOnce bool

// Debugf writes a debug level log entry.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	log(ctx, fmt.Sprintf("debug. "+format, args...))
}

// Errorf writes a debug level log entry.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	log(ctx, fmt.Sprintf("error. "+format, args...))
}

// Warningf writes a debug level log entry.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	log(ctx, fmt.Sprintf("warning. "+format, args...))
}

func log(ctx context.Context, s string) {
	if !announceOnce {
		announceOnce = true
		now := time.Now()
		stardate := speakn(now.Format("060201"))
		announce := fmt.Sprintf("Captains log star date %s point %s.", stardate, speakn(fmt.Sprintf("%d", now.Hour())))
		exec.CommandContext(ctx, "say", announce).Run()
	} else {
		exec.CommandContext(ctx, "say", "-v", "Daniel", "Captains log, supplemental.").Run()
	}
	exec.CommandContext(ctx, "say", "-v", "Daniel", s).Run()
}

func speakn(s string) string {
	var nos []string
	for i := range s {
		nos = append(nos, string(s[i]))
	}
	return strings.Join(nos, " ")
}
