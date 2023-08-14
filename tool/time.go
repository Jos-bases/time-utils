package tool

import (
	"time"
)

type TimeTool struct {
	timeStr string
	format TimeFormat
	loc *time.Location
	tm *time.Time
}

func (t *TimeTool) NewTimeTool(timeStr string, format TimeFormat) *TimeTool  {
	if timeStr!= "" {
		t.timeStr = timeStr
	} else {
		if format != "" {
			t.timeStr = time.Now().Format(string(format))
		} else {
			t.timeStr = time.Now().Format(string(Y_M_D_H_I_S))
		}
	}

	t.format = format
	return t
}

func (t *TimeTool) LoadLocation(loc TimeLocation) *TimeTool  {

	t.loc, _ = time.LoadLocation(string(loc))

	return t
}

func (t *TimeTool) AddDay(day int) *TimeTool  {

	if t.tm == nil {
		if t.loc != nil {
			tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
			tm = tm.AddDate(0,0, day)
			t.tm = &tm

		} else {
			tm, _ := time.Parse(string(t.format), t.timeStr)
			tm = tm.AddDate(0,0, day)
			t.tm = &tm

		}
	} else {
		tm := t.tm.AddDate(0,0, day)
		t.tm = &tm
	}
	


	return t
}

func (t *TimeTool) AddMonth(month int) *TimeTool {

	if t.tm == nil {
		if t.loc != nil {
			tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
			tm = tm.AddDate(0,month, 0)
			t.tm = &tm

		} else {
			tm, _ := time.Parse(string(t.format), t.timeStr)
			tm = tm.AddDate(0,month, 0)
			t.tm = &tm
		}
	} else {
		tm := t.tm.AddDate(0,month, 0)
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) AddYear(year int) *TimeTool  {

	if t.tm == nil {
		if t.loc != nil {
			tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
			tm = tm.AddDate(year,0, 0)
			t.tm = &tm

		} else {
			tm, _ := time.Parse(string(t.format), t.timeStr)
			tm = tm.AddDate(year,0, 0)
			t.tm = &tm
		}
	} else {
		tm := t.tm.AddDate(year,0, 0)
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) AddSecond(second int) *TimeTool  {

	if t.tm == nil {
		if t.loc != nil {
			tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
			tm = tm.Add(time.Duration(second))
			t.tm = &tm

		} else {
			tm, _ := time.Parse(string(t.format), t.timeStr)
			tm = tm.Add(time.Duration(second))
			t.tm = &tm
		}
	} else {
		tm := t.tm.Add(time.Duration(second))
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) GetTime() time.Time  {

	var tm time.Time
	if t.loc != nil {
		tm = t.tm.In(t.loc)
	} else {
		tm = *t.tm
	}

	return tm
}

func (t *TimeTool) GetString() string  {

	var tm string
	if t.loc != nil {
		tm = t.tm.In(t.loc).Format(string(t.format))
	} else {
		tm = t.tm.Format(string(t.format))
	}
	return tm
}