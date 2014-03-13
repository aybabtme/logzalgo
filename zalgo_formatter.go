package logzalgo

import (
	"bytes"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/kortschak/zalgo"
)

// New corrupts a default logrus logger, making it much better at logging stuff.
func New() *logrus.Logger {
	thatThing := logrus.New()
	thatThing.Formatter = NewZalgoFormatterrrrrr()
	return thatThing
}

// ZalgoFormatter, just look at it, don't touch it.
type ZalgoFormatter struct {
	victim *logrus.TextFormatter
	pain   *bytes.Buffer
	z      *zalgo.Corrupter
}

// NewZalgoFormatterrrrrr gives you a new Zalgo formatterrrrrr...!
func NewZalgoFormatterrrrrr() *ZalgoFormatter {
	pain := bytes.NewBuffer(nil)
	z := zalgo.NewCorrupter(pain)

	z.Zalgo = func(n int, z *zalgo.Corrupter) {
		z.Up += 0.1
		z.Middle += complex(0.01, 0.01)
		z.Down += complex(real(z.Down)*0.1, 0)
	}

	return &ZalgoFormatter{
		victim: &logrus.TextFormatter{},
		pain:   pain,
		z:      z,
	}
}

// Format Formats Stuff™.
func (zal *ZalgoFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	zal.pain.Reset()

	zal.z.Up = complex(0, 0.2)
	zal.z.Middle = complex(0, 0.2)
	zal.z.Down = complex(0.001, 0.3)

	victimsWish := entry.Data["msg"].(string)

	_, _ = fmt.Fprint(zal.z, victimsWish)

	victimsReality := zal.pain.String()

	entry.Data["msg"] = victimsReality
	return zal.victim.Format(entry)
}
