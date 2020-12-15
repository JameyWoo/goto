/**
 * @Author: JameyWoo
 * @Email: 2622075127wjh@gmail.com
 * @Date: 2020/12/14
 * @Desc: main
 * @Copyright (c) 2020, JameyWoo. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type MyFormatter struct{
	DisableColors bool
}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	s.DisableColors = false
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	fullPath := strings.Split(entry.Caller.File, "/")
	path := fullPath[len(fullPath) - 1] + ":" + strconv.Itoa(entry.Caller.Line)
	msg := fmt.Sprintf("%s %8s %15s\t %s\n",
		timestamp,
		"[" + strings.ToUpper(entry.Level.String()) + "]",
		path,
		entry.Message)
	return []byte(msg), nil
}

func init() {
	logrus.SetFormatter(new(MyFormatter))
	logrus.SetReportCaller(true)
}
