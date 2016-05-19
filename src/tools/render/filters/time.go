package filters

import (
	"github.com/flosch/pongo2"

	"fmt"
	"time"
)

func init() {
	pongo2.RegisterFilter("duration", durationFilter)
	pongo2.RegisterFilter("naturaltime", naturaltimeFilter)
}

func durationFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {

	diff, err := timeDiff(in, param)

	if err != nil {

		return nil, &pongo2.Error{
			Sender:   "filter:duration",
			ErrorMsg: err.Error(),
		}
	}

	return pongo2.AsValue(duration(diff)), nil
}

func naturaltimeFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {

	diff, err := timeDiff(in, param)

	if err != nil {

		return nil, &pongo2.Error{
			Sender:   "filter:naturaltime",
			ErrorMsg: err.Error(),
		}
	}

	return pongo2.AsValue(naturaltime(diff)), nil
}

func timeDiff(in *pongo2.Value, param *pongo2.Value) (time.Duration, error) {

	var (
		basetime  time.Time
		paramtime time.Time
	)

	switch in.Interface().(type) {
	case time.Time:
		basetime = in.Interface().(time.Time)
	case *time.Time:
		t := in.Interface().(*time.Time)
		basetime = *t
	default:
		return 0, fmt.Errorf("value is not a time")
	}

	if !param.IsNil() {

		switch param.Interface().(type) {
		case time.Time:
			paramtime = param.Interface().(time.Time)
		case *time.Time:
			t := param.Interface().(*time.Time)
			paramtime = *t
		default:
			return 0, fmt.Errorf("parameter is not a time")
		}

	} else {

		paramtime = time.Now()
	}

	return basetime.Sub(paramtime), nil
}
