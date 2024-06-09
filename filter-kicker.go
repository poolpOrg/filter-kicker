package main

/*
 * Copyright (c) 2024 Gilles Chehade <gilles@poolp.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

import (
	"net"
	"time"

	"github.com/poolpOrg/OpenSMTPD-framework/filter"
)

type SessionData struct {
	kickcount int
}

func kicker(session filter.Session) filter.Response {
	if session.Get().(*SessionData).kickcount >= 3 {
		return filter.Disconnect("421 4.7.0 Too many commands without progressing, goodbye.")
	} else {
		return filter.Proceed()
	}
}

/*
func protocolClientCb(timestamp time.Time, session filter.Session, command string) {
	localSession := session.Get().(*SessionData)
	fmt.Fprintf(os.Stderr, "%s: %s: kickercount: %d\n", timestamp, session, localSession.kickcount)
}
*/

func filterConnectCb(timestamp time.Time, session filter.Session, rdns string, src net.Addr) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterHeloCb(timestamp time.Time, session filter.Session, helo string) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterEhloCb(timestamp time.Time, session filter.Session, helo string) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterStartTLSCb(timestamp time.Time, session filter.Session, tls string) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterAuthCb(timestamp time.Time, session filter.Session, mechanism string) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterMailFromCb(timestamp time.Time, session filter.Session, from string) filter.Response {
	return filter.Proceed()
}

func filterRcptToCb(timestamp time.Time, session filter.Session, to string) filter.Response {
	return filter.Proceed()
}

func filterDataCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterCommitCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount = 0
	return filter.Proceed()
}

func filterNoopCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount += 1
	return kicker(session)
}

func filterRsetCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount += 1
	return kicker(session)
}

func filterHelpCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount += 1
	return kicker(session)
}

func filterWizCb(timestamp time.Time, session filter.Session) filter.Response {
	session.Get().(*SessionData).kickcount += 1
	return kicker(session)
}

func main() {
	filter.Init()

	filter.SMTP_IN.SessionAllocator(func() filter.SessionData {
		return &SessionData{}
	})

	//filter.SMTP_IN.OnProtocolClient(protocolClientCb)

	filter.SMTP_IN.ConnectRequest(filterConnectCb)
	filter.SMTP_IN.HeloRequest(filterHeloCb)
	filter.SMTP_IN.EhloRequest(filterEhloCb)
	filter.SMTP_IN.StartTLSRequest(filterStartTLSCb)
	filter.SMTP_IN.AuthRequest(filterAuthCb)
	filter.SMTP_IN.MailFromRequest(filterMailFromCb)
	filter.SMTP_IN.RcptToRequest(filterRcptToCb)
	filter.SMTP_IN.DataRequest(filterDataCb)
	filter.SMTP_IN.CommitRequest(filterCommitCb)
	filter.SMTP_IN.NoopRequest(filterNoopCb)
	filter.SMTP_IN.RsetRequest(filterRsetCb)
	filter.SMTP_IN.HelpRequest(filterHelpCb)
	filter.SMTP_IN.WizRequest(filterWizCb)

	filter.Dispatch()
}
